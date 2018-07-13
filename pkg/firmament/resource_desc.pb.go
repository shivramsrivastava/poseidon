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

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource_desc.proto

package firmament

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ResourceDescriptor_ResourceState int32

const (
	ResourceDescriptor_RESOURCE_UNKNOWN ResourceDescriptor_ResourceState = 0
	ResourceDescriptor_RESOURCE_IDLE    ResourceDescriptor_ResourceState = 1
	ResourceDescriptor_RESOURCE_BUSY    ResourceDescriptor_ResourceState = 2
	ResourceDescriptor_RESOURCE_LOST    ResourceDescriptor_ResourceState = 3
)

var ResourceDescriptor_ResourceState_name = map[int32]string{
	0: "RESOURCE_UNKNOWN",
	1: "RESOURCE_IDLE",
	2: "RESOURCE_BUSY",
	3: "RESOURCE_LOST",
}
var ResourceDescriptor_ResourceState_value = map[string]int32{
	"RESOURCE_UNKNOWN": 0,
	"RESOURCE_IDLE":    1,
	"RESOURCE_BUSY":    2,
	"RESOURCE_LOST":    3,
}

func (x ResourceDescriptor_ResourceState) String() string {
	return proto.EnumName(ResourceDescriptor_ResourceState_name, int32(x))
}
func (ResourceDescriptor_ResourceState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor10, []int{0, 0}
}

type ResourceDescriptor_ResourceType int32

const (
	ResourceDescriptor_RESOURCE_PU          ResourceDescriptor_ResourceType = 0
	ResourceDescriptor_RESOURCE_CORE        ResourceDescriptor_ResourceType = 1
	ResourceDescriptor_RESOURCE_CACHE       ResourceDescriptor_ResourceType = 2
	ResourceDescriptor_RESOURCE_NIC         ResourceDescriptor_ResourceType = 3
	ResourceDescriptor_RESOURCE_DISK        ResourceDescriptor_ResourceType = 4
	ResourceDescriptor_RESOURCE_SSD         ResourceDescriptor_ResourceType = 5
	ResourceDescriptor_RESOURCE_MACHINE     ResourceDescriptor_ResourceType = 6
	ResourceDescriptor_RESOURCE_LOGICAL     ResourceDescriptor_ResourceType = 7
	ResourceDescriptor_RESOURCE_NUMA_NODE   ResourceDescriptor_ResourceType = 8
	ResourceDescriptor_RESOURCE_SOCKET      ResourceDescriptor_ResourceType = 9
	ResourceDescriptor_RESOURCE_COORDINATOR ResourceDescriptor_ResourceType = 10
)

var ResourceDescriptor_ResourceType_name = map[int32]string{
	0:  "RESOURCE_PU",
	1:  "RESOURCE_CORE",
	2:  "RESOURCE_CACHE",
	3:  "RESOURCE_NIC",
	4:  "RESOURCE_DISK",
	5:  "RESOURCE_SSD",
	6:  "RESOURCE_MACHINE",
	7:  "RESOURCE_LOGICAL",
	8:  "RESOURCE_NUMA_NODE",
	9:  "RESOURCE_SOCKET",
	10: "RESOURCE_COORDINATOR",
}
var ResourceDescriptor_ResourceType_value = map[string]int32{
	"RESOURCE_PU":          0,
	"RESOURCE_CORE":        1,
	"RESOURCE_CACHE":       2,
	"RESOURCE_NIC":         3,
	"RESOURCE_DISK":        4,
	"RESOURCE_SSD":         5,
	"RESOURCE_MACHINE":     6,
	"RESOURCE_LOGICAL":     7,
	"RESOURCE_NUMA_NODE":   8,
	"RESOURCE_SOCKET":      9,
	"RESOURCE_COORDINATOR": 10,
}

func (x ResourceDescriptor_ResourceType) String() string {
	return proto.EnumName(ResourceDescriptor_ResourceType_name, int32(x))
}
func (ResourceDescriptor_ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor10, []int{0, 1}
}

type ResourceDescriptor struct {
	Uuid            string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	FriendlyName    string `protobuf:"bytes,2,opt,name=friendly_name,json=friendlyName" json:"friendly_name,omitempty"`
	DescriptiveName string `protobuf:"bytes,3,opt,name=descriptive_name,json=descriptiveName" json:"descriptive_name,omitempty"`
	// state is the state of resource.
	State ResourceDescriptor_ResourceState `protobuf:"varint,4,opt,name=state,enum=firmament.ResourceDescriptor_ResourceState" json:"state,omitempty"`
	// task_capacity is the max number of tasks that can be scheduled on the resource.
	TaskCapacity uint64 `protobuf:"varint,5,opt,name=task_capacity,json=taskCapacity" json:"task_capacity,omitempty"`
	// last heartbeat of the resource, e.g. node.
	LastHeartbeat uint64 `protobuf:"varint,6,opt,name=last_heartbeat,json=lastHeartbeat" json:"last_heartbeat,omitempty"`
	// Type of the resource
	Type ResourceDescriptor_ResourceType `protobuf:"varint,7,opt,name=type,enum=firmament.ResourceDescriptor_ResourceType" json:"type,omitempty"`
	// schedulable indicates if the resource, e.g. node is able to schedule tasks.
	Schedulable bool `protobuf:"varint,8,opt,name=schedulable" json:"schedulable,omitempty"`
	// current_running_tasks stores all the running tasks on the resource, e.g. node.
	CurrentRunningTasks []uint64 `protobuf:"varint,9,rep,packed,name=current_running_tasks,json=currentRunningTasks" json:"current_running_tasks,omitempty"`
	// num_running_tasks_below stores the number of running tasks on the resources that are below this node.
	NumRunningTasksBelow uint64 `protobuf:"varint,10,opt,name=num_running_tasks_below,json=numRunningTasksBelow" json:"num_running_tasks_below,omitempty"`
	NumSlotsBelow        uint64 `protobuf:"varint,11,opt,name=num_slots_below,json=numSlotsBelow" json:"num_slots_below,omitempty"`
	// Resource capacity and load tracking
	AvailableResources          *ResourceVector `protobuf:"bytes,12,opt,name=available_resources,json=availableResources" json:"available_resources,omitempty"`
	ReservedResources           *ResourceVector `protobuf:"bytes,13,opt,name=reserved_resources,json=reservedResources" json:"reserved_resources,omitempty"`
	MinAvailableResourcesBelow  *ResourceVector `protobuf:"bytes,14,opt,name=min_available_resources_below,json=minAvailableResourcesBelow" json:"min_available_resources_below,omitempty"`
	MaxAvailableResourcesBelow  *ResourceVector `protobuf:"bytes,15,opt,name=max_available_resources_below,json=maxAvailableResourcesBelow" json:"max_available_resources_below,omitempty"`
	MinUnreservedResourcesBelow *ResourceVector `protobuf:"bytes,16,opt,name=min_unreserved_resources_below,json=minUnreservedResourcesBelow" json:"min_unreserved_resources_below,omitempty"`
	MaxUnreservedResourcesBelow *ResourceVector `protobuf:"bytes,17,opt,name=max_unreserved_resources_below,json=maxUnreservedResourcesBelow" json:"max_unreserved_resources_below,omitempty"`
	ResourceCapacity            *ResourceVector `protobuf:"bytes,18,opt,name=resource_capacity,json=resourceCapacity" json:"resource_capacity,omitempty"`
	// Cost-model-specific statistics
	WhareMapStats          *WhareMapStats          `protobuf:"bytes,19,opt,name=whare_map_stats,json=whareMapStats" json:"whare_map_stats,omitempty"`
	CocoInterferenceScores *CoCoInterferenceScores `protobuf:"bytes,20,opt,name=coco_interference_scores,json=cocoInterferenceScores" json:"coco_interference_scores,omitempty"`
	// Simulation related fields
	TraceMachineId uint64 `protobuf:"varint,21,opt,name=trace_machine_id,json=traceMachineId" json:"trace_machine_id,omitempty"`
	// Resource labels
	Labels []*Label `protobuf:"bytes,32,rep,name=labels" json:"labels,omitempty"`
	// Taints
	Taints []*Taint `protobuf:"bytes,33,rep,name=taints" json:"taints,omitempty"`
}

func (m *ResourceDescriptor) Reset()                    { *m = ResourceDescriptor{} }
func (m *ResourceDescriptor) String() string            { return proto.CompactTextString(m) }
func (*ResourceDescriptor) ProtoMessage()               {}
func (*ResourceDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *ResourceDescriptor) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ResourceDescriptor) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func (m *ResourceDescriptor) GetDescriptiveName() string {
	if m != nil {
		return m.DescriptiveName
	}
	return ""
}

func (m *ResourceDescriptor) GetState() ResourceDescriptor_ResourceState {
	if m != nil {
		return m.State
	}
	return ResourceDescriptor_RESOURCE_UNKNOWN
}

func (m *ResourceDescriptor) GetTaskCapacity() uint64 {
	if m != nil {
		return m.TaskCapacity
	}
	return 0
}

func (m *ResourceDescriptor) GetLastHeartbeat() uint64 {
	if m != nil {
		return m.LastHeartbeat
	}
	return 0
}

func (m *ResourceDescriptor) GetType() ResourceDescriptor_ResourceType {
	if m != nil {
		return m.Type
	}
	return ResourceDescriptor_RESOURCE_PU
}

func (m *ResourceDescriptor) GetSchedulable() bool {
	if m != nil {
		return m.Schedulable
	}
	return false
}

func (m *ResourceDescriptor) GetCurrentRunningTasks() []uint64 {
	if m != nil {
		return m.CurrentRunningTasks
	}
	return nil
}

func (m *ResourceDescriptor) GetNumRunningTasksBelow() uint64 {
	if m != nil {
		return m.NumRunningTasksBelow
	}
	return 0
}

func (m *ResourceDescriptor) GetNumSlotsBelow() uint64 {
	if m != nil {
		return m.NumSlotsBelow
	}
	return 0
}

func (m *ResourceDescriptor) GetAvailableResources() *ResourceVector {
	if m != nil {
		return m.AvailableResources
	}
	return nil
}

func (m *ResourceDescriptor) GetReservedResources() *ResourceVector {
	if m != nil {
		return m.ReservedResources
	}
	return nil
}

func (m *ResourceDescriptor) GetMinAvailableResourcesBelow() *ResourceVector {
	if m != nil {
		return m.MinAvailableResourcesBelow
	}
	return nil
}

func (m *ResourceDescriptor) GetMaxAvailableResourcesBelow() *ResourceVector {
	if m != nil {
		return m.MaxAvailableResourcesBelow
	}
	return nil
}

func (m *ResourceDescriptor) GetMinUnreservedResourcesBelow() *ResourceVector {
	if m != nil {
		return m.MinUnreservedResourcesBelow
	}
	return nil
}

func (m *ResourceDescriptor) GetMaxUnreservedResourcesBelow() *ResourceVector {
	if m != nil {
		return m.MaxUnreservedResourcesBelow
	}
	return nil
}

func (m *ResourceDescriptor) GetResourceCapacity() *ResourceVector {
	if m != nil {
		return m.ResourceCapacity
	}
	return nil
}

func (m *ResourceDescriptor) GetWhareMapStats() *WhareMapStats {
	if m != nil {
		return m.WhareMapStats
	}
	return nil
}

func (m *ResourceDescriptor) GetCocoInterferenceScores() *CoCoInterferenceScores {
	if m != nil {
		return m.CocoInterferenceScores
	}
	return nil
}

func (m *ResourceDescriptor) GetTraceMachineId() uint64 {
	if m != nil {
		return m.TraceMachineId
	}
	return 0
}

func (m *ResourceDescriptor) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ResourceDescriptor) GetTaints() []*Taint {
	if m != nil {
		return m.Taints
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceDescriptor)(nil), "firmament.ResourceDescriptor")
	proto.RegisterEnum("firmament.ResourceDescriptor_ResourceState", ResourceDescriptor_ResourceState_name, ResourceDescriptor_ResourceState_value)
	proto.RegisterEnum("firmament.ResourceDescriptor_ResourceType", ResourceDescriptor_ResourceType_name, ResourceDescriptor_ResourceType_value)
}

func init() { proto.RegisterFile("resource_desc.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 806 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xed, 0x6e, 0xe3, 0xc4,
	0x17, 0xc6, 0xff, 0x69, 0xd2, 0x6e, 0x7b, 0xf2, 0x36, 0x9d, 0xb4, 0xfb, 0x1f, 0x8a, 0x58, 0x79,
	0x8b, 0x40, 0x06, 0xa4, 0x7e, 0x28, 0xe2, 0x2b, 0x22, 0xeb, 0x04, 0x1a, 0x9a, 0xc6, 0x68, 0x9c,
	0xb0, 0x42, 0xa0, 0xb5, 0x26, 0xce, 0x94, 0x5a, 0xc4, 0xe3, 0x68, 0x3c, 0xee, 0xcb, 0x0d, 0x72,
	0x3d, 0x5c, 0x02, 0x9a, 0xf1, 0x4b, 0x9c, 0xcd, 0x76, 0xc3, 0xb7, 0xf8, 0x39, 0xcf, 0xf9, 0xf9,
	0x39, 0x23, 0xcf, 0x09, 0xf4, 0x24, 0x4f, 0xe2, 0x54, 0x06, 0xdc, 0x5f, 0xf0, 0x24, 0xb8, 0x58,
	0xc9, 0x58, 0xc5, 0xf8, 0xe8, 0x36, 0x94, 0x11, 0x8b, 0xb8, 0x50, 0x67, 0xaf, 0x82, 0x38, 0x88,
	0xfd, 0x50, 0x28, 0x2e, 0x6f, 0xb9, 0xe4, 0x22, 0xe0, 0x7e, 0x12, 0xc4, 0x92, 0x27, 0x99, 0xf5,
	0xac, 0xb9, 0x64, 0x73, 0xbe, 0xcc, 0x1f, 0x4e, 0x4b, 0xd8, 0x3d, 0x0f, 0x54, 0x2c, 0x0b, 0xf9,
	0xe1, 0x8e, 0x49, 0xee, 0x47, 0x6c, 0xe5, 0x27, 0x8a, 0xa9, 0xa2, 0xb5, 0xa5, 0x58, 0x28, 0x8a,
	0xa7, 0xf3, 0xbf, 0xdb, 0x80, 0x69, 0xde, 0x3e, 0xe0, 0x49, 0x20, 0xc3, 0x95, 0x8a, 0x25, 0xc6,
	0xd0, 0x48, 0xd3, 0x70, 0x41, 0x6a, 0x56, 0xcd, 0x3e, 0xa2, 0xe6, 0x37, 0xfe, 0x1c, 0xda, 0xb7,
	0x32, 0xe4, 0x62, 0xb1, 0x7c, 0xf2, 0x05, 0x8b, 0x38, 0xd9, 0x33, 0xc5, 0x56, 0x21, 0x4e, 0x58,
	0xc4, 0xf1, 0x57, 0x80, 0x16, 0x39, 0x26, 0xbc, 0xe7, 0x99, 0xaf, 0x6e, 0x7c, 0xdd, 0x8a, 0x6e,
	0xac, 0x7d, 0xd8, 0xd7, 0xb9, 0x38, 0x69, 0x58, 0x35, 0xbb, 0x73, 0xf9, 0xcd, 0x45, 0x39, 0xfe,
	0xc5, 0x76, 0xa2, 0x52, 0xf2, 0x74, 0x0b, 0xcd, 0x3a, 0x75, 0x24, 0xc5, 0x92, 0xbf, 0xfc, 0x80,
	0xad, 0x58, 0x10, 0xaa, 0x27, 0xb2, 0x6f, 0xd5, 0xec, 0x06, 0x6d, 0x69, 0xd1, 0xc9, 0x35, 0xfc,
	0x05, 0x74, 0x96, 0x2c, 0x51, 0xfe, 0x1d, 0x67, 0x52, 0xcd, 0x39, 0x53, 0xe4, 0xc0, 0xb8, 0xda,
	0x5a, 0xbd, 0x2a, 0x44, 0xfc, 0x3d, 0x34, 0xd4, 0xd3, 0x8a, 0x93, 0x17, 0x26, 0xcd, 0xd7, 0xff,
	0x2d, 0xcd, 0xf4, 0x69, 0xc5, 0xa9, 0xe9, 0xc3, 0x16, 0x34, 0x93, 0xe0, 0x8e, 0x2f, 0xd2, 0x25,
	0x9b, 0x2f, 0x39, 0x39, 0xb4, 0x6a, 0xf6, 0x21, 0xad, 0x4a, 0xf8, 0x12, 0x4e, 0x83, 0x54, 0x4a,
	0x2e, 0x94, 0x2f, 0x53, 0x21, 0x42, 0xf1, 0xa7, 0xaf, 0x83, 0x26, 0xe4, 0xc8, 0xaa, 0xdb, 0x0d,
	0xda, 0xcb, 0x8b, 0x34, 0xab, 0x4d, 0x75, 0x09, 0x7f, 0x07, 0xff, 0x17, 0x69, 0xb4, 0xe9, 0xf7,
	0xe7, 0x7c, 0x19, 0x3f, 0x10, 0x30, 0x53, 0x9c, 0x88, 0x34, 0xaa, 0x76, 0xbc, 0xd1, 0x35, 0xfc,
	0x25, 0x74, 0x75, 0x5b, 0xb2, 0x8c, 0x55, 0x61, 0x6f, 0x66, 0x43, 0x8b, 0x34, 0xf2, 0xb4, 0x9a,
	0xf9, 0x7e, 0x86, 0x1e, 0xbb, 0x67, 0xa1, 0xc9, 0xe7, 0x17, 0x9f, 0x51, 0x42, 0x5a, 0x56, 0xcd,
	0x6e, 0x5e, 0x7e, 0xf2, 0x81, 0x33, 0xf8, 0xd5, 0x7c, 0x61, 0x14, 0x97, 0x5d, 0x45, 0x21, 0xc1,
	0x57, 0x80, 0x25, 0x4f, 0xb8, 0xbc, 0xe7, 0x8b, 0x0a, 0xaa, 0xbd, 0x0b, 0x75, 0x5c, 0x34, 0xad,
	0x49, 0x7f, 0xc0, 0x67, 0x51, 0x28, 0xfc, 0x0f, 0x24, 0xcb, 0x67, 0xe9, 0xec, 0x82, 0x9e, 0x45,
	0xa1, 0xe8, 0x6f, 0x45, 0xcc, 0x66, 0xd6, 0x74, 0xf6, 0xf8, 0x11, 0x7a, 0x77, 0x37, 0x9d, 0x3d,
	0x3e, 0x47, 0x7f, 0x07, 0xaf, 0x74, 0xf6, 0x54, 0x6c, 0x9f, 0x45, 0x8e, 0x47, 0xbb, 0xf0, 0x9f,
	0x46, 0xa1, 0x98, 0x89, 0xad, 0x63, 0x59, 0xf3, 0xd9, 0xe3, 0xc7, 0xf8, 0xc7, 0xbb, 0xf9, 0xec,
	0xf1, 0x59, 0xfe, 0x8f, 0x70, 0x5c, 0xae, 0x93, 0xf2, 0x5a, 0xe1, 0x5d, 0x48, 0x54, 0xf4, 0x94,
	0xb7, 0xee, 0x07, 0xe8, 0xbe, 0xb7, 0x7f, 0x48, 0xcf, 0x50, 0x48, 0x85, 0xf2, 0x56, 0x3b, 0x6e,
	0xd8, 0x4a, 0x5f, 0xea, 0x84, 0xb6, 0x1f, 0xaa, 0x8f, 0xf8, 0x77, 0x20, 0xcf, 0x6d, 0x41, 0x72,
	0x62, 0x50, 0xaf, 0x2b, 0x28, 0x27, 0x76, 0xe2, 0x51, 0xc5, 0xe9, 0x19, 0x23, 0x7d, 0xa9, 0x11,
	0xdb, 0x3a, 0xb6, 0x01, 0x29, 0xc9, 0x02, 0x1d, 0x2f, 0xb8, 0x0b, 0x05, 0xf7, 0xc3, 0x05, 0x39,
	0x35, 0x37, 0xa4, 0x63, 0xf4, 0x9b, 0x4c, 0x1e, 0x2d, 0xb0, 0x0d, 0x07, 0x66, 0xd9, 0x26, 0xc4,
	0xb2, 0xea, 0x76, 0xf3, 0x12, 0x55, 0x5e, 0x3a, 0xd6, 0x05, 0x9a, 0xd7, 0xb5, 0x33, 0xdb, 0xad,
	0xe4, 0xf5, 0x96, 0x73, 0xaa, 0x0b, 0x34, 0xaf, 0x9f, 0xbf, 0x83, 0xf6, 0xc6, 0x3e, 0xc3, 0x27,
	0x80, 0xe8, 0xd0, 0x73, 0x67, 0xd4, 0x19, 0xfa, 0xb3, 0xc9, 0xf5, 0xc4, 0x7d, 0x3b, 0x41, 0xff,
	0xc3, 0xc7, 0xd0, 0x2e, 0xd5, 0xd1, 0x60, 0x3c, 0x44, 0xb5, 0x0d, 0xe9, 0xcd, 0xcc, 0xfb, 0x0d,
	0xed, 0x6d, 0x48, 0x63, 0xd7, 0x9b, 0xa2, 0xfa, 0xf9, 0x3f, 0x35, 0x68, 0x55, 0x57, 0x14, 0xee,
	0x42, 0xb3, 0xf4, 0xfc, 0x32, 0x7b, 0x0f, 0xed, 0xb8, 0x54, 0xa3, 0x31, 0x74, 0xd6, 0x52, 0xdf,
	0xb9, 0x1a, 0xa2, 0x3d, 0x8c, 0xa0, 0x55, 0x6a, 0x93, 0x91, 0x83, 0xea, 0x1b, 0x8d, 0x83, 0x91,
	0x77, 0x8d, 0x1a, 0x1b, 0x26, 0xcf, 0x1b, 0xa0, 0xfd, 0x8d, 0x71, 0x6e, 0xfa, 0xce, 0xd5, 0x68,
	0x32, 0x44, 0x07, 0x1b, 0xea, 0xd8, 0xfd, 0x69, 0xe4, 0xf4, 0xc7, 0xe8, 0x05, 0x7e, 0x09, 0x78,
	0xfd, 0x8a, 0xd9, 0x4d, 0xdf, 0x9f, 0xb8, 0x83, 0x21, 0x3a, 0xc4, 0x3d, 0xe8, 0xae, 0xa9, 0xae,
	0x73, 0x3d, 0x9c, 0xa2, 0x23, 0x4c, 0xe0, 0xa4, 0x12, 0xdb, 0xa5, 0x83, 0xd1, 0xa4, 0x3f, 0x75,
	0x29, 0x82, 0xf9, 0x81, 0xf9, 0x3f, 0xfb, 0xf6, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x67,
	0xc2, 0x36, 0x5a, 0x07, 0x00, 0x00,
}
