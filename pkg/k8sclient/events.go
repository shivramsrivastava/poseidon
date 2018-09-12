/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package k8sclient

import (
	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"sync"
)

type PodEvents struct {
	Recorder    record.EventRecorder
	BroadCaster record.EventBroadcaster
}

type PoseidonEvents struct {
	podEvents *PodEvents
	podInfo   map[PodIdentifier]int64
	k8sClient kubernetes.Interface
	sync.Mutex
}

var poseidonEvents *PoseidonEvents
var poseidonEventsLock sync.Mutex

func NewPoseidonEvents(coreEvent kubernetes.Interface) *PoseidonEvents {

	poseidonEventsLock.Lock()
	if poseidonEvents == nil {
		poseidonEvents = &PoseidonEvents{
			podEvents: NewPodEvents(coreEvent),
			podInfo:   make(map[PodIdentifier]int64),
			k8sClient: coreEvent,
		}
	}
	poseidonEventsLock.Unlock()

	return poseidonEvents
}

func NewPodEvents(coreEvent kubernetes.Interface) *PodEvents {

	sch := legacyscheme.Scheme

	sch.AddKnownTypes(schema.GroupVersion{Group: "Poseidon", Version: "v1"}, &corev1.Pod{})
	objarr, ok, err := sch.ObjectKinds(&corev1.Pod{})

	glog.Info(objarr, ok, err)
	if err != nil {
		glog.Fatalf("legacysheme err %v", err)
	}

	err = corev1.AddToScheme(legacyscheme.Scheme)
	//err=api.AddToScheme(legacyscheme.Scheme)
	if err != nil {
		glog.Fatalf("could not register schemes", err)
	}
	broadCaster := record.NewBroadcaster()
	recorder := broadCaster.NewRecorder(sch, corev1.EventSource{Component: "Poseidon"})
	broadCaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: coreEvent.CoreV1().Events("")})
	//recorder.Event()

	return &PodEvents{
		BroadCaster: broadCaster,
		Recorder:    recorder,
	}
}

// This will reveive when a pod is added
// and also when a schdule delta i called
func (posiedonEvents *PoseidonEvents) ReceivePodInfo() {

	for {
		select {
		case pendingPod := <-PodPendingChan:
			{
				glog.Info("got a pod in pending state", pendingPod)
				//no need to lock this
				posiedonEvents.Lock()
				posiedonEvents.podInfo[pendingPod]++
				posiedonEvents.Unlock()
			}
		case podIdentifier := <-PodSheduledChan:
			{
				glog.Info("got a pod schedules by firmament", podIdentifier)
				//if we receive from this channel it means \
				// we got schedule delta

				// check the map and delete
				// we don't need this step just to sure if the pod exist
				posiedonEvents.Lock()
				//temporary hack
				mapAccesskey := podIdentifier
				mapAccesskey.NodeName = ""
				if _, ok := posiedonEvents.podInfo[mapAccesskey]; !ok {
					// this means the first event for a failiure
					glog.Info(podIdentifier.Name, " Not found ")
					glog.Info("What we have is ", posiedonEvents.podInfo)
					glog.Info("What we have in k8s pods", PodToK8sPod)
					PodToK8sPodLock.Lock()
					// Note: accessing the below map PodToK8sPod without checking if the value exists
					// check if the pod exists
					poseidonPod, ok := PodToK8sPod[mapAccesskey]
					PodToK8sPodLock.Unlock()
					if !ok {
						glog.Info(podIdentifier, " Dose not exit but got a schedule call")
						posiedonEvents.Unlock()
						continue
					}
					posiedonEvents.podEvents.Recorder.Eventf(poseidonPod, corev1.EventTypeNormal, "Scheduled", "Successfully assigned %v/%v to %v", podIdentifier.Name, podIdentifier.Namespace, podIdentifier.NodeName)
				} else {
					//remove the pod after the message has been broadcasted
					posiedonEvents.podEvents.Recorder.Eventf(PodToK8sPod[mapAccesskey], corev1.EventTypeNormal, "Scheduled", "Successfully assigned %v/%v to %v", podIdentifier.Name, podIdentifier.Namespace, podIdentifier.NodeName)
					glog.Info("Deleting after sending success event", podIdentifier, " from posiedonEvents map in ReceivePodInfo")
					delete(posiedonEvents.podInfo, mapAccesskey)
					//PodToK8sPodLock.Lock()
					//delete(PodToK8sPod, podIdentifier)
					//PodToK8sPodLock.Unlock()
				}
				posiedonEvents.Unlock()
			}
		}
	}

}

// RecalulateMissingPods called after every schedule call
func (posiedonEvents *PoseidonEvents) RecalculateMissingPods() {

	posiedonEvents.Lock()
	for podIdentifier, val := range posiedonEvents.podInfo {
		glog.Info("RecalculateMissingPods called ", podIdentifier, val)
		if val > 2 {
			glog.Info(podIdentifier, " cannot be scheduled anymore it pass 2 scheduling runs")
			PodToK8sPodLock.Lock()
			// Note: accessing the below map PodToK8sPod without checking if the value exists
			// check if the pod exists
			poseidonPod, ok := PodToK8sPod[podIdentifier]
			if !ok {
				glog.Info(podIdentifier, " Dose not exit")
				PodToK8sPodLock.Unlock()
				continue
			}
			PodToK8sPodLock.Unlock()
			posiedonEvents.podEvents.Recorder.Eventf(poseidonPod, corev1.EventTypeWarning, "FailedScheduling", "Firmament failed to schedule the pod %s in %s namespace", podIdentifier.Name, podIdentifier.Namespace)
			Update(posiedonEvents.k8sClient, poseidonPod, &corev1.PodCondition{
				Type:    corev1.PodScheduled,
				Status:  corev1.ConditionFalse,
				Reason:  corev1.PodReasonUnschedulable,
				Message: "Firmament unable to schedule the pod",
			})
			glog.Info("Sent FaileScheduling Events deleting from the events map")
			//delete(PodToK8sPod, podIdentifier)
			glog.Info("Deleting ", podIdentifier, " from posiedonEvents map in RecalculateMissingPods")
			delete(posiedonEvents.podInfo, podIdentifier)
		} else {
			posiedonEvents.podInfo[podIdentifier]++
		}
	}
	posiedonEvents.Unlock()
}
