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
	"github.com/kubernetes-sigs/poseidon/pkg/firmament"
	"sync"

	"github.com/golang/glog"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

// NewPodWatcher initialize a PodWatcher.
func NewK8sPodWatcher(kubeVerMajor, kubeVerMinor int, schedulerName string, client kubernetes.Interface, fc firmament.FirmamentSchedulerClient) *K8sPodWatcher {
	glog.V(2).Info("Starting K8sPodWatcher...")
	PodMux = new(sync.RWMutex)
	PodToTD = make(map[PodIdentifier]*firmament.TaskDescriptor)
	TaskIDToPod = make(map[uint64]PodIdentifier)
	jobIDToJD = make(map[string]*firmament.JobDescriptor)
	jobNumTasksToRemove = make(map[string]int)
	podWatcher := &K8sPodWatcher{
		clientset: client,
		fc:        fc,
	}
	schedulerSelector := fields.Everything()
	podSelector := labels.Everything()
	if kubeVerMajor >= 1 && kubeVerMinor >= 6 {
		// schedulerName is only available in Kubernetes >= 1.6.
		schedulerSelector = fields.ParseSelectorOrDie("spec.schedulerName!=" + schedulerName)
	} else {
		var err error
		podSelector, err = labels.Parse("scheduler notin (" + schedulerName + ")")
		if err != nil {
			glog.Fatal("Failed to parse scheduler label selector")
		}
	}
	_, controller := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(alo metav1.ListOptions) (runtime.Object, error) {
				alo.FieldSelector = schedulerSelector.String()
				alo.LabelSelector = podSelector.String()
				return client.CoreV1().Pods("").List(alo)
			},
			WatchFunc: func(alo metav1.ListOptions) (watch.Interface, error) {
				alo.FieldSelector = schedulerSelector.String()
				alo.LabelSelector = podSelector.String()
				return client.CoreV1().Pods("").Watch(alo)
			},
		},
		&v1.Pod{},
		0,
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
	stop := make(chan struct{})
	go podWatcher.controller.Run(stop)
	return podWatcher
}

func (pw *K8sPodWatcher) getCPUMemRequest(pod *v1.Pod) (int64, int64, int64) {
	cpuReq := int64(0)
	memReq := int64(0)
	ephemeralReq := int64(0)
	for _, container := range pod.Spec.Containers {
		request := container.Resources.Requests
		cpuReqQuantity := request[v1.ResourceCPU]
		cpuReq += cpuReqQuantity.MilliValue()
		memReqQuantity := request[v1.ResourceMemory]
		memReqCont, _ := memReqQuantity.AsInt64()
		memReq += memReqCont
		ephemeralReqQuantity := request[v1.ResourceEphemeralStorage]
		ephemeralReqCont, _ := ephemeralReqQuantity.AsInt64()
		ephemeralReq += ephemeralReqCont

	}
	return cpuReq, memReq, ephemeralReq
}

func (pw *K8sPodWatcher) parsePod(pod *v1.Pod) *firmament.TaskInfo {
	var resourceID string
	cpuReq, memReq, ephemeralReq := pw.getCPUMemRequest(pod)
	podPhase := PodUnknown
	opType := firmament.TaskInfoType_TASKINFO_ADD
	switch pod.Status.Phase {
	case v1.PodPending:
		podPhase = PodPending
	case v1.PodRunning:
		podPhase = PodRunning
		opType = firmament.TaskInfoType_TASKINFO_ADD // send add task info
	case v1.PodSucceeded:
		podPhase = PodSucceeded
		opType = firmament.TaskInfoType_TASKINFO_REMOVE // send remove task info
	case v1.PodFailed:
		podPhase = PodFailed
		opType = firmament.TaskInfoType_TASKINFO_REMOVE //when a pod fails send remove taskinfo
	}

	if podPhase == PodPending {
		glog.V(2).Info("for Pending ignore AddTaskInfo", pod.Name+"/"+pod.Namespace)
		return nil
	}
	// check if the node name is updated in the pod spec
	if pod.Spec.NodeName == "" {
		glog.V(2).Info("for pod in ", pod.Status.Phase, " state node-name not set so ignoring AddTaskInfo", pod.Name+"/"+pod.Namespace)
		return nil
	} else {
		NodeMux.Lock()
		if rtnd, ok := NodeToRTND[pod.Spec.NodeName]; ok {
			resourceID = rtnd.GetResourceDesc().GetUuid()
		} else {
			glog.Errorf("Node ", pod.Spec.NodeName, " doesn't exist")
			return nil
		}
	}
	return &firmament.TaskInfo{
		TaskName:                    pod.Name + "/" + pod.Namespace,
		ResourceId:                  resourceID,
		CpuUtilization:              cpuReq,
		MemUtilization:              memReq,
		EphemeralStorageUtilization: ephemeralReq,
		Type: opType,
	}
}

// CheckAndUpdateK8sPodMap will return true if a new object added else will return false
func (pw *K8sPodWatcher) CheckAndUpdateK8sPodMap(taskinfo *firmament.TaskInfo) bool {
	ok := false
	pw.Lock()
	if _, ok := pw.K8sPods[taskinfo.GetTaskName()]; !ok {
		pw.K8sPods[taskinfo.GetTaskName()] = taskinfo
		ok = true
	}
	pw.Unlock()
	return ok
}

// RemoveTaskfromK8sPodMap return true if remove was successful
func (pw *K8sPodWatcher) RemoveTaskfromK8sPodMap(taskinfo *firmament.TaskInfo) bool {
	ok := false
	pw.Lock()
	if _, ok := pw.K8sPods[taskinfo.GetTaskName()]; ok {
		delete(pw.K8sPods, taskinfo.GetTaskName())
		ok = true
	}
	pw.Unlock()
	return ok

}

// CheckOpType return true if the optype match
func (pw *K8sPodWatcher) CheckOpType(taskinfo *firmament.TaskInfo, optype firmament.TaskInfoType) bool {
	if taskinfo.GetType() == optype {
		return true
	}
	return false
}

func (pw *K8sPodWatcher) enqueuePodAddition(key interface{}, obj interface{}) {
	pod := obj.(*v1.Pod)
	if addedPod := pw.parsePod(pod); addedPod != nil {
		if pw.CheckAndUpdateK8sPodMap(addedPod) {
			//can send the info
			// this can be for a pod already running/succeded or newly added and in pending state
			firmament.AddTaskInfo(pw.fc, addedPod)
			glog.V(2).Info("AddTaskInfo sent for pod", addedPod.GetTaskName())
		} else {
			glog.V(2).Info("igoring the AddTaskInfo for already existing task", addedPod.GetTaskName())
		}
	}
}

func (pw *K8sPodWatcher) enqueuePodDeletion(key interface{}, obj interface{}) {
	pod := obj.(*v1.Pod)
	if pod.DeletionTimestamp != nil {
		if deletePod := pw.parsePod(pod); deletePod != nil {
			if _, ok := pw.K8sPods[deletePod.GetTaskName()]; ok {
				//check the opType and send it to firmament
				if pw.CheckOpType(deletePod, firmament.TaskInfoType_TASKINFO_REMOVE) {
					firmament.AddTaskInfo(pw.fc, deletePod)
					_ = pw.RemoveTaskfromK8sPodMap(deletePod)
				} else {
					glog.V(2).Info("OpType for deleting pod is different", deletePod.GetType())
					deletePod.Type = firmament.TaskInfoType_TASKINFO_REMOVE
					firmament.AddTaskInfo(pw.fc, deletePod)
					_ = pw.RemoveTaskfromK8sPodMap(deletePod)
				}
			} else {
				glog.V(2).Info("Deleting Pod doesn't exit in K8sPods map", deletePod.GetTaskName())
			}
		} else {
			glog.V(2).Info("Pending pod getting deleted no AddTaskInfo needed", pod.Name+"/"+pod.Namespace)
		}
	}
}

func (pw *K8sPodWatcher) enqueuePodUpdate(key, oldObj, newObj interface{}) {
	oldPod := oldObj.(*v1.Pod)
	newPod := newObj.(*v1.Pod)
	if oldPod.Status.Phase != newPod.Status.Phase {

		if oldPod.Status.Phase == v1.PodPending && newPod.Status.Phase == v1.PodRunning {
			if addedPod := pw.parsePod(newPod); addedPod != nil {
				if pw.CheckAndUpdateK8sPodMap(addedPod) {
					//can send the info
					// this can be for a pod already running/succeded or newly added and in pending state
					firmament.AddTaskInfo(pw.fc, addedPod)
					glog.V(2).Info("AddTaskInfo sent for pod", addedPod.GetTaskName())
				} else {
					glog.V(2).Info("igoring the AddTaskInfo for already existing task", addedPod.GetTaskName())
				}
			}
		} else {
			// should be a remove op
			if deletePod := pw.parsePod(newPod); deletePod != nil {
				if _, ok := pw.K8sPods[deletePod.GetTaskName()]; ok {
					//check the opType and send it to firmament
					if pw.CheckOpType(deletePod, firmament.TaskInfoType_TASKINFO_REMOVE) {
						firmament.AddTaskInfo(pw.fc, deletePod)
						_ = pw.RemoveTaskfromK8sPodMap(deletePod)
					} else {
						glog.V(2).Info("OpType for deleting pod is different", deletePod.GetType())
						deletePod.Type = firmament.TaskInfoType_TASKINFO_REMOVE
						firmament.AddTaskInfo(pw.fc, deletePod)
						_ = pw.RemoveTaskfromK8sPodMap(deletePod)
					}
				} else {
					glog.V(2).Info("Deleting Pod doesn't exit in K8sPods map", deletePod.GetTaskName())
				}
			} else {
				glog.V(2).Info("Pending pod getting deleted no AddTaskInfo needed", newPod.Name+"/"+newPod.Namespace)
			}

		}
	}
}
