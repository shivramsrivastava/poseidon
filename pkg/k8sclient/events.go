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
	sync.Mutex
}

var poseidonEvents *PoseidonEvents
var poseidonEventsLock sync.Mutex

func NewPoseidonEvents() *PoseidonEvents {

	poseidonEventsLock.Lock()
	if poseidonEvents == nil {
		poseidonEvents = &PoseidonEvents{
			podEvents: NewPodEvents(),
			podInfo:   make(map[PodIdentifier]int64),
		}
	}
	poseidonEventsLock.Unlock()

	return poseidonEvents
}

func NewPodEvents() *PodEvents {

	broadCaster := record.NewBroadcaster()
	recorder := broadCaster.NewRecorder(legacyscheme.Scheme, corev1.EventSource{Component: "Poseidon"})

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
		case podIdentifier := <-PodPendingChan:
			{
				//no need to lock this
				posiedonEvents.Lock()
				posiedonEvents.podInfo[podIdentifier]++
				posiedonEvents.Unlock()
			}
		case podIdentifier := <-PodSheduledChan:
			{
				//if we receive from this channel it means \
				// we got schedule delta

				// check the map and delete
				// we don't need this step just to sure if the pod exist
				posiedonEvents.Lock()
				if _, ok := posiedonEvents.podInfo[podIdentifier]; !ok {
					glog.Info(podIdentifier.Name, " Not found ")
				} else {
					//remove the pod after the message has been broadcasted
					delete(posiedonEvents.podInfo, podIdentifier)
					PodToK8sPodLock.Lock()
					delete(PodToK8sPod, podIdentifier)
					PodToK8sPodLock.Unlock()
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

		if val > 2 {
			PodToK8sPodLock.Lock()
			// Note: accessing the below map PodToK8sPod without checking if the value exists
			posiedonEvents.podEvents.Recorder.Eventf(PodToK8sPod[podIdentifier], corev1.EventTypeWarning, "FailedScheduling", "Firmament failed to schedule the pod %s in %s namespace", podIdentifier.Name, podIdentifier.Namespace)
			delete(PodToK8sPod, podIdentifier)
			delete(posiedonEvents.podInfo, podIdentifier)
			PodToK8sPodLock.Unlock()
		}
		posiedonEvents.podInfo[podIdentifier]++
	}
	posiedonEvents.Unlock()
}
