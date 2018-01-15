// Poseidon
// Copyright (c) The Poseidon Authors.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// THIS CODE IS PROVIDED ON AN *AS IS* BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING WITHOUT
// LIMITATION ANY IMPLIED WARRANTIES OR CONDITIONS OF TITLE, FITNESS FOR
// A PARTICULAR PURPOSE, MERCHANTABLITY OR NON-INFRINGEMENT.
//
// See the Apache Version 2.0 License for specific language governing
// permissions and limitations under the License.

package k8sclient

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/poseidon/pkg/firmament"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type NodeSelectors map[string]string

func SortNodeSelectors(nodeSelector NodeSelectors) NodeSelectors {
	newSortedNodeSelectors := make(NodeSelectors)
	var keyArray []string
	for key := range nodeSelector {
		// Skip key with networkRequirement
		if key == "networkRequirement" {
			continue
		}
		keyArray = append(keyArray, key)
	}
	sort.Strings(keyArray)
	for _, key := range keyArray {
		newSortedNodeSelectors[key] = nodeSelector[key]
	}
	return newSortedNodeSelectors
}

func NewPodWatcher(kubeVerMajor, kubeVerMinor int, schedulerName string, client kubernetes.Interface, fc firmament.FirmamentSchedulerClient) *PodWatcher {
	glog.Info("Starting PodWatcher...")
	PodsCond = &sync.Mutex{}
	//PodsCond = sync.NewCond(&sync.Mutex{})
	PodToTD = make(map[PodIdentifier]*firmament.TaskDescriptor)
	TaskIDToPod = make(map[uint64]PodIdentifier)
	jobIDToJD = make(map[string]*firmament.JobDescriptor)
	jobNumTasksToRemove = make(map[string]int)
	podWatcher := &PodWatcher{
		clientset: client,
		fc:        fc,
	}
	schedulerSelector := labels.Everything()
	var err error
	if kubeVerMajor >= 1 && kubeVerMinor >= 6 {
		// TODO:(shiv)
		// Please refer issue : https://github.com/kubernetes/kubernetes/issues/49190
		// schedulerName support which is available in Kubernetes >= 1.6 is broken.
		// The current workaround to this issue is to assign a mandatory label as mentioned below to all the object
		// ( like pods/deployments/job etc) that are to be scheduled by poseidon.
		// label:
		//  schedulerName: poseidon
		schedulerSelector, err = labels.Parse("schedulerName==" + schedulerName)
		if err != nil {
			glog.Fatal("Failed to parse scheduler label selector")
		}
	} else {
		var err error
		schedulerSelector, err = labels.Parse("scheduler in (" + schedulerName + ")")
		if err != nil {
			glog.Fatal("Failed to parse scheduler label selector")
		}
	}
	_, controller := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(alo metav1.ListOptions) (runtime.Object, error) {
				alo.LabelSelector = schedulerSelector.String()
				return client.CoreV1().Pods("").List(alo)
			},
			WatchFunc: func(alo metav1.ListOptions) (watch.Interface, error) {
				alo.LabelSelector = schedulerSelector.String()
				return client.CoreV1().Pods("").Watch(alo)
			},
		},
		&v1.Pod{},
		time.Millisecond, //(time.Second * 10) ,//(1 * time.Millisecond),
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				key, err := cache.MetaNamespaceKeyFunc(obj)
				if err != nil {
					glog.Errorf("AddFunc: error getting key %v", err)
				}
				podWatcher.enqueuePodAddition(key, obj)
			},
			UpdateFunc: func(old, new interface{}) {
				key, err := cache.MetaNamespaceKeyFunc(new)
				if err != nil {
					glog.Errorf("UpdateFunc: error getting key %v", err)
				}
				podWatcher.enqueuePodUpdate(key, old, new)
			},
			DeleteFunc: func(obj interface{}) {
				key, err := cache.MetaNamespaceKeyFunc(obj)
				if err != nil {
					glog.Errorf("DeleteFunc: error getting key %v", err)
				}
				podWatcher.enqueuePodDeletion(key, obj)
			},
		},
	)
	podWatcher.controller = controller
	podWatcher.podWorkQueue = NewKeyedQueue()
	return podWatcher
}

func (this *PodWatcher) getCpuMemRequest(pod *v1.Pod) (int64, int64) {
	cpuReq := int64(0)
	memReq := int64(0)
	for _, container := range pod.Spec.Containers {
		request := container.Resources.Requests
		cpuReqQuantity := request["cpu"]
		cpuReq += cpuReqQuantity.MilliValue()
		memReqQuantity := request["memory"]
		memReqCont, _ := memReqQuantity.AsInt64()
		memReq += memReqCont
	}
	return cpuReq, memReq
}

func (this *PodWatcher) parsePod(pod *v1.Pod) *Pod {
	cpuReq, memReq := this.getCpuMemRequest(pod)
	podPhase := PodPhase("Unknown")
	switch pod.Status.Phase {
	case "Pending":
		podPhase = "Pending"
	case "Running":
		podPhase = "Running"
	case "Succeeded":
		podPhase = "Succeeded"
	case "Failed":
		podPhase = "Failed"
	}
	return &Pod{
		Identifier: PodIdentifier{
			Name:      pod.Name,
			Namespace: pod.Namespace,
		},
		State:        podPhase,
		CpuRequest:   cpuReq,
		MemRequestKb: memReq / bytesToKb,
		Labels:       pod.Labels,
		Annotations:  pod.Annotations,
		NodeSelector: pod.Spec.NodeSelector,
		OwnerRef:     GetOwnerReference(pod),
	}
}

func (this *PodWatcher) enqueuePodAddition(key interface{}, obj interface{}) {
	pod := obj.(*v1.Pod)
	addedPod := this.parsePod(pod)
	this.podWorkQueue.Add(key, addedPod)
	glog.V(2).Info("enqueuePodAddition: Added pod ", addedPod.Identifier)

}

func (this *PodWatcher) enqueuePodDeletion(key interface{}, obj interface{}) {
	pod := obj.(*v1.Pod)
	if pod.DeletionTimestamp != nil {
		// Only delete pods if they have a DeletionTimestamp.
		deletedPod := &Pod{
			Identifier: PodIdentifier{
				Name:      pod.Name,
				Namespace: pod.Namespace,
			},
			State:    PodDeleted,
			OwnerRef: GetOwnerReference(pod),
		}
		this.podWorkQueue.Add(key, deletedPod)
		glog.V(2).Info("enqueuePodDeletion: Added pod ", deletedPod.Identifier)
	}
}

func (this *PodWatcher) enqueuePodUpdate(key, oldObj, newObj interface{}) {
	oldPod := oldObj.(*v1.Pod)
	newPod := newObj.(*v1.Pod)
	if oldPod.Status.Phase != newPod.Status.Phase {
		if newPod.Status.Phase == "Running" {
			return
		}
		// TODO(ionel): This code assumes that if other fields changed as well then Firmament will automatically update them upon state transition. This is currently not true.
		updatedPod := this.parsePod(newPod)
		this.podWorkQueue.Add(key, updatedPod)
		glog.V(2).Infof("enqueuePodUpdate: Updated pod state change %v %s", updatedPod.Identifier, updatedPod.State)
		return
	}
	oldCpuReq, oldMemReq := this.getCpuMemRequest(oldPod)
	newCpuReq, newMemReq := this.getCpuMemRequest(newPod)
	if oldCpuReq != newCpuReq || oldMemReq != newMemReq ||
		!reflect.DeepEqual(oldPod.Labels, newPod.Labels) ||
		!reflect.DeepEqual(oldPod.Annotations, newPod.Annotations) ||
		!reflect.DeepEqual(oldPod.Spec.NodeSelector, newPod.Spec.NodeSelector) {
		updatedPod := this.parsePod(newPod)
		this.podWorkQueue.Add(key, updatedPod)
		glog.V(2).Info("enqueuePodUpdate: Updated pod ", updatedPod.Identifier)
		return
	}
}

func (this *PodWatcher) Run(stopCh <-chan struct{}, nWorkers int) {
	defer utilruntime.HandleCrash()

	// The workers can stop when we are done.
	defer this.podWorkQueue.ShutDown()
	defer glog.Info("Shutting down PodWatcher")
	glog.Info("Getting pod updates...")

	go this.controller.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, this.controller.HasSynced) {
		utilruntime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	glog.Info("Starting pod watching workers")
	for i := 0; i < nWorkers; i++ { //nWorkers; i++ {
		go wait.Until(this.podWorker, (time.Millisecond), stopCh)
	}

	<-stopCh
	glog.Info("Stopping pod watcher")
}

func (this *PodWatcher) podWorker() {

	for {
		func() {
			key, items, quit := this.podWorkQueue.Get()
			if quit {
				glog.V(2).Info(" quit called from Queue")
				return
			}
			for _, item := range items {
				pod := item.(*Pod)
				switch pod.State {
				case PodPending:
					glog.V(2).Info("PodPending ", pod.Identifier)
					jobId := this.generateJobID(pod.OwnerRef)
					PodsCond.Lock()
					glog.V(2).Info("Jod ID for this pod: ", jobId)
					jd, ok := jobIDToJD[jobId]
					if !ok {
						jd = this.createNewJob(pod.OwnerRef, jobId)
						jobIDToJD[jobId] = jd
						// Changed below value from 0 to 1 because 1st task added.
						td := this.addTaskToJob(pod, jd.Uuid, jd.Name, 0)
						jobNumTasksToRemove[jobId] = 1
						jd.RootTask = td
						PodToTD[pod.Identifier] = td
						TaskIDToPod[td.GetUid()] = pod.Identifier
						taskDescription := &firmament.TaskDescription{
							TaskDescriptor: td,
							JobDescriptor:  jd,
						}
						firmament.TaskSubmitted(this.fc, taskDescription)
						PodsCond.Unlock()
						continue
					}
					jobNumTasksToRemove[jobId]++
					taskCount := jobNumTasksToRemove[jobId]
					PodsCond.Unlock()
					td := this.addTaskToJob(pod, jd.Uuid, jd.Name, (taskCount))
					PodsCond.Lock()
					PodToTD[pod.Identifier] = td
					TaskIDToPod[td.GetUid()] = pod.Identifier
					taskDescription := &firmament.TaskDescription{
						TaskDescriptor: td,
						JobDescriptor:  jd,
					}
					PodsCond.Unlock()
					go firmament.TaskSubmitted(this.fc, taskDescription)
				case PodSucceeded:
					glog.V(2).Info("PodSucceeded ", pod.Identifier)
					PodsCond.Lock()
					td, ok := PodToTD[pod.Identifier]
					PodsCond.Unlock()
					if !ok {
						glog.Fatalf("Pod %v does not exist", pod.Identifier)
					}
					firmament.TaskCompleted(this.fc, &firmament.TaskUID{TaskUid: td.Uid})
				case PodDeleted:
					glog.Info("PodDeleted ", pod.Identifier)
					PodsCond.Lock()
					td, ok := PodToTD[pod.Identifier]
					PodsCond.Unlock()
					if !ok {
						glog.Fatalf("Pod %s does not exist", pod.Identifier)
					}
					firmament.TaskRemoved(this.fc, &firmament.TaskUID{TaskUid: td.Uid})
					PodsCond.Lock()
					delete(PodToTD, pod.Identifier)
					delete(TaskIDToPod, td.GetUid())
					// TODO(ionel): Should we delete the task from JD's spawned field?
					jobId := this.generateJobID(pod.OwnerRef)
					jobNumTasksToRemove[jobId]--
					if jobNumTasksToRemove[jobId] == 0 {
						// Clean state because the job doesn't have any tasks left.
						delete(jobNumTasksToRemove, jobId)
						delete(jobIDToJD, jobId)
					}
					PodsCond.Unlock()
				case PodFailed:
					glog.V(2).Info("PodFailed ", pod.Identifier)
					PodsCond.Lock()
					td, ok := PodToTD[pod.Identifier]
					PodsCond.Unlock()
					if !ok {
						glog.Fatalf("Pod %s does not exist", pod.Identifier)
					}
					firmament.TaskFailed(this.fc, &firmament.TaskUID{TaskUid: td.Uid})
				case PodRunning:
					glog.V(2).Info("PodRunning ", pod.Identifier)
					// We don't have to do anything.
				case PodUnknown:
					glog.Errorf("Pod %s in unknown state", pod.Identifier)
					// TODO(ionel): Handle Unknown case.
				case PodUpdated:
					glog.V(2).Info("PodUpdated ", pod.Identifier)
					PodsCond.Lock()
					jobId := this.generateJobID(pod.OwnerRef)
					jd, okJob := jobIDToJD[jobId]
					td, okPod := PodToTD[pod.Identifier]
					PodsCond.Unlock()
					if !okJob {
						glog.Fatalf("Pod's %v job does not exist", pod.Identifier)
					}
					if !okPod {
						glog.Fatalf("Pod %v does not exist", pod.Identifier)
					}
					this.updateTask(pod, td)
					taskDescription := &firmament.TaskDescription{
						TaskDescriptor: td,
						JobDescriptor:  jd,
					}
					firmament.TaskUpdated(this.fc, taskDescription)
				default:
					glog.Fatalf("Pod %v in unexpected state %v", pod.Identifier, pod.State)
				}
			}
			defer this.podWorkQueue.Done(key)
		}()
	}
}

func (this *PodWatcher) createNewJob(jobName string, jobID string) *firmament.JobDescriptor {
	jobDesc := &firmament.JobDescriptor{
		Uuid:  jobID,
		Name:  jobName,
		State: firmament.JobDescriptor_CREATED,
	}
	return jobDesc
}

func (this *PodWatcher) updateTask(pod *Pod, td *firmament.TaskDescriptor) {
	// TODO(ionel): Update LabelSelector!
	td.ResourceRequest.CpuCores = float32(pod.CpuRequest)
	td.ResourceRequest.RamCap = uint64(pod.MemRequestKb)
	// Update labels.
	td.Labels = nil
	for label, value := range pod.Labels {
		td.Labels = append(td.Labels,
			&firmament.Label{
				Key:   label,
				Value: value,
			})
	}
}

func (this *PodWatcher) addTaskToJob(pod *Pod, jdUid string, jdName string, tdID int) *firmament.TaskDescriptor {
	task := &firmament.TaskDescriptor{
		Name:  pod.Identifier.UniqueName(),
		State: firmament.TaskDescriptor_CREATED,
		JobId: jdUid,
		ResourceRequest: &firmament.ResourceVector{
			// TODO(ionel): Update types so no cast is required.
			CpuCores: float32(pod.CpuRequest),
			RamCap:   uint64(pod.MemRequestKb),
		},
	}

	// Add labels.
	for label, value := range pod.Labels {
		task.Labels = append(task.Labels,
			&firmament.Label{
				Key:   label,
				Value: value,
			})
	}
	// Get the network requirement from pods label, and set it in ResourceRequest of the TaskDescriptor
	setTaskNetworkRequirement(task, pod.Labels)
	task.LabelSelectors = this.getFirmamentLabelSelectorFromNodeSelectorMap(SortNodeSelectors(pod.NodeSelector))
	setTaskType(task)
	task.Uid = this.generateTaskID(jdName, tdID)
	return task
}

func (this *PodWatcher) generateJobID(seed string) string {
	if seed == "" {
		glog.Fatal("Seed value is nil")
	}

	return GenerateUUID(seed)
}

func (this *PodWatcher) generateTaskID(jdUid string, taskNum int) uint64 {
	return HashCombine(jdUid, taskNum)
}

// GetOwnerReference to get the parent object reference
func GetOwnerReference(pod *v1.Pod) string {

	// Return if owner reference exists.
	ownerRefs := pod.GetObjectMeta().GetOwnerReferences()
	if len(ownerRefs) != 0 {
		for x := range ownerRefs {
			ref := &ownerRefs[x]
			if ref.Controller != nil && *ref.Controller {
				return string(ref.UID)
			}
		}
	}

	// Return the controller-uid label if it exists.
	if controllerID := pod.GetObjectMeta().GetLabels()["controller-uid"]; controllerID != "" {
		return controllerID
	}

	// Return 'kubernetes.io/created-by' if it exists.
	if createdByAnnotation, ok := pod.GetObjectMeta().GetAnnotations()["kubernetes.io/created-by"]; ok {
		var serialCreatedBy v1.SerializedReference
		err := json.Unmarshal([]byte(createdByAnnotation), &serialCreatedBy)
		if err == nil {
			return string(serialCreatedBy.Reference.UID)
		}
	}

	// Return the uid of the ObjectMeta if none from the above is present.
	return string(pod.GetObjectMeta().GetUID())
}

func (this *PodWatcher) getFirmamentLabelSelectorFromNodeSelectorMap(nodeSelector NodeSelectors) []*firmament.LabelSelector {
	var firmamentLabelSelector []*firmament.LabelSelector
	for key, value := range nodeSelector {
		firmamentLabelSelector = append(firmamentLabelSelector, &firmament.LabelSelector{
			Type:   firmament.LabelSelector_IN_SET,
			Values: []string{value},
			Key:    key,
		})
	}
	return firmamentLabelSelector
}

func setTaskNetworkRequirement(td *firmament.TaskDescriptor, nodeSelectors NodeSelectors) {
	if val, ok := nodeSelectors["networkRequirement"]; ok {
		res, err := strconv.ParseUint(val, 10, 64)
		if err == nil {
			td.ResourceRequest.NetRxBw = res
		} else {
			glog.Errorf("Failed to parse networkRequirement %v", err)
		}
	}
}

func setTaskType(td *firmament.TaskDescriptor) {
	for _, label := range td.Labels {
		if label.Key == "taskType" {
			switch label.Value {
			case "Sheep":
				td.TaskType = firmament.TaskDescriptor_SHEEP
			case "Rabbit":
				td.TaskType = firmament.TaskDescriptor_RABBIT
			case "Devil":
				td.TaskType = firmament.TaskDescriptor_DEVIL
			case "Turtle":
				td.TaskType = firmament.TaskDescriptor_TURTLE
			default:
				glog.Errorf("Unexpected task type %s for task %s", label.Value, td.Name)
			}
		}
	}
}
