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
	"github.com/camsas/poseidon/pkg/firmament"
	"github.com/golang/glog"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"time"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

var clientSet, clientSet1, clientSet2 kubernetes.Interface
var totalElapsedTime time.Duration // this time should give the over all schedul time
var PushRateLimit flowcontrol.RateLimiter

func BindPodToNode(podName string, namespace string, nodeName string) {
	//singleBindcallTime := time.Now()
	PushRateLimit.Accept()
	err := clientSet.CoreV1().Pods(namespace).Bind(&v1.Binding{
		meta_v1.TypeMeta{},
		meta_v1.ObjectMeta{
			Name: podName,
		},
		v1.ObjectReference{
			Namespace: namespace,
			Name:      nodeName,
		}})
	//sinceTime := time.Now()
	//glog.Info("Current Bind call time", sinceTime.Sub(singleBindcallTime), " Start:", singleBindcallTime, " End:", sinceTime)
	//totalElapsedTime = totalElapsedTime + sinceTime.Sub(singleBindcallTime)
	//glog.Info(" Total Bind call summzation ", totalElapsedTime)
	if err != nil {
		glog.Info("Could not bind %v", err)
		time.Sleep(time.Millisecond * 5)
		BindPodToNode(podName,namespace,nodeName)
	}
}

func DeletePod(podName string, namespace string) {
	clientSet.CoreV1().Pods(namespace).Delete(podName, &meta_v1.DeleteOptions{})
}

func GetClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}

func New(schedulerName string, kubeConfig string, kubeVersionMajor, kubeVersionMinor int, firmamentAddress string) {

	PushRateLimit=flowcontrol.NewTokenBucketRateLimiter(500,500)
	config, err := GetClientConfig(kubeConfig)
	if err != nil {
		glog.Fatalf("Failed to load client config: %v", err)
	}
	config.QPS = config.QPS + 1500
	config.Burst = config.Burst + 1000
	clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		glog.Fatalf("Failed to create connection: %v", err)
	}

	clientSet1, err = kubernetes.NewForConfig(config)
	if err != nil {
		glog.Fatalf("Failed to create connection: %v", err)
	}

	clientSet2, err = kubernetes.NewForConfig(config)
	if err != nil {
		glog.Fatalf("Failed to create connection: %v", err)
	}

	//var arrayFC []firmament.FirmamentSchedulerClient
	/*arrayFC:=make([]firmament.FirmamentSchedulerClient,500)
	for i:=0;i<500;i++{
	fc1, conn1, err := firmament.New(firmamentAddress)
	if err != nil {
		glog.Fatalf("Failed to connect to Firmament: %v", err)
	}
		arrayFC[i] = fc1
		defer conn1.Close()
	}*/

	fc1, conn1, err1 := firmament.New(firmamentAddress)
	if err1 != nil {
		glog.Fatalf("Failed to connect to Firmament: %v", err)
	}
	defer conn1.Close()

	fc2, conn2, err := firmament.New(firmamentAddress)
	if err != nil {
		glog.Fatalf("Failed to connect to Firmament: %v", err)
	}
	defer conn2.Close()

	glog.Info("k8s newclient called")
	stopCh := make(chan struct{})
	go NewPodWatcher(kubeVersionMajor, kubeVersionMinor, schedulerName, clientSet1, fc1).Run(stopCh, 100)
	go NewNodeWatcher(clientSet2, fc2).Run(stopCh, 10)

	// We block here.
	<-stopCh
}
