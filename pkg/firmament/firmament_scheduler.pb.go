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
// source: firmament_scheduler.proto

package firmament

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TaskReplyType int32

const (
	TaskReplyType_TASK_COMPLETED_OK      TaskReplyType = 0
	TaskReplyType_TASK_SUBMITTED_OK      TaskReplyType = 1
	TaskReplyType_TASK_REMOVED_OK        TaskReplyType = 2
	TaskReplyType_TASK_FAILED_OK         TaskReplyType = 3
	TaskReplyType_TASK_UPDATED_OK        TaskReplyType = 4
	TaskReplyType_TASK_NOT_FOUND         TaskReplyType = 5
	TaskReplyType_TASK_JOB_NOT_FOUND     TaskReplyType = 6
	TaskReplyType_TASK_ALREADY_SUBMITTED TaskReplyType = 7
	TaskReplyType_TASK_STATE_NOT_CREATED TaskReplyType = 8
)

var TaskReplyType_name = map[int32]string{
	0: "TASK_COMPLETED_OK",
	1: "TASK_SUBMITTED_OK",
	2: "TASK_REMOVED_OK",
	3: "TASK_FAILED_OK",
	4: "TASK_UPDATED_OK",
	5: "TASK_NOT_FOUND",
	6: "TASK_JOB_NOT_FOUND",
	7: "TASK_ALREADY_SUBMITTED",
	8: "TASK_STATE_NOT_CREATED",
}
var TaskReplyType_value = map[string]int32{
	"TASK_COMPLETED_OK":      0,
	"TASK_SUBMITTED_OK":      1,
	"TASK_REMOVED_OK":        2,
	"TASK_FAILED_OK":         3,
	"TASK_UPDATED_OK":        4,
	"TASK_NOT_FOUND":         5,
	"TASK_JOB_NOT_FOUND":     6,
	"TASK_ALREADY_SUBMITTED": 7,
	"TASK_STATE_NOT_CREATED": 8,
}

func (x TaskReplyType) String() string {
	return proto.EnumName(TaskReplyType_name, int32(x))
}
func (TaskReplyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type NodeReplyType int32

const (
	NodeReplyType_NODE_ADDED_OK       NodeReplyType = 0
	NodeReplyType_NODE_FAILED_OK      NodeReplyType = 1
	NodeReplyType_NODE_REMOVED_OK     NodeReplyType = 2
	NodeReplyType_NODE_UPDATED_OK     NodeReplyType = 3
	NodeReplyType_NODE_NOT_FOUND      NodeReplyType = 4
	NodeReplyType_NODE_ALREADY_EXISTS NodeReplyType = 5
)

var NodeReplyType_name = map[int32]string{
	0: "NODE_ADDED_OK",
	1: "NODE_FAILED_OK",
	2: "NODE_REMOVED_OK",
	3: "NODE_UPDATED_OK",
	4: "NODE_NOT_FOUND",
	5: "NODE_ALREADY_EXISTS",
}
var NodeReplyType_value = map[string]int32{
	"NODE_ADDED_OK":       0,
	"NODE_FAILED_OK":      1,
	"NODE_REMOVED_OK":     2,
	"NODE_UPDATED_OK":     3,
	"NODE_NOT_FOUND":      4,
	"NODE_ALREADY_EXISTS": 5,
}

func (x NodeReplyType) String() string {
	return proto.EnumName(NodeReplyType_name, int32(x))
}
func (NodeReplyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type ServingStatus int32

const (
	ServingStatus_UNKNOWN     ServingStatus = 0
	ServingStatus_SERVING     ServingStatus = 1
	ServingStatus_NOT_SERVING ServingStatus = 2
)

var ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x ServingStatus) String() string {
	return proto.EnumName(ServingStatus_name, int32(x))
}
func (ServingStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type ScheduleRequest struct {
}

func (m *ScheduleRequest) Reset()                    { *m = ScheduleRequest{} }
func (m *ScheduleRequest) String() string            { return proto.CompactTextString(m) }
func (*ScheduleRequest) ProtoMessage()               {}
func (*ScheduleRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type SchedulingDeltas struct {
	Deltas []*SchedulingDelta `protobuf:"bytes,1,rep,name=deltas" json:"deltas,omitempty"`
}

func (m *SchedulingDeltas) Reset()                    { *m = SchedulingDeltas{} }
func (m *SchedulingDeltas) String() string            { return proto.CompactTextString(m) }
func (*SchedulingDeltas) ProtoMessage()               {}
func (*SchedulingDeltas) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SchedulingDeltas) GetDeltas() []*SchedulingDelta {
	if m != nil {
		return m.Deltas
	}
	return nil
}

type TaskCompletedResponse struct {
	Type TaskReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.TaskReplyType" json:"type,omitempty"`
}

func (m *TaskCompletedResponse) Reset()                    { *m = TaskCompletedResponse{} }
func (m *TaskCompletedResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskCompletedResponse) ProtoMessage()               {}
func (*TaskCompletedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *TaskCompletedResponse) GetType() TaskReplyType {
	if m != nil {
		return m.Type
	}
	return TaskReplyType_TASK_COMPLETED_OK
}

type TaskDescription struct {
	TaskDescriptor *TaskDescriptor `protobuf:"bytes,1,opt,name=task_descriptor,json=taskDescriptor" json:"task_descriptor,omitempty"`
	JobDescriptor  *JobDescriptor  `protobuf:"bytes,2,opt,name=job_descriptor,json=jobDescriptor" json:"job_descriptor,omitempty"`
}

func (m *TaskDescription) Reset()                    { *m = TaskDescription{} }
func (m *TaskDescription) String() string            { return proto.CompactTextString(m) }
func (*TaskDescription) ProtoMessage()               {}
func (*TaskDescription) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *TaskDescription) GetTaskDescriptor() *TaskDescriptor {
	if m != nil {
		return m.TaskDescriptor
	}
	return nil
}

func (m *TaskDescription) GetJobDescriptor() *JobDescriptor {
	if m != nil {
		return m.JobDescriptor
	}
	return nil
}

type TaskSubmittedResponse struct {
	Type TaskReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.TaskReplyType" json:"type,omitempty"`
}

func (m *TaskSubmittedResponse) Reset()                    { *m = TaskSubmittedResponse{} }
func (m *TaskSubmittedResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskSubmittedResponse) ProtoMessage()               {}
func (*TaskSubmittedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *TaskSubmittedResponse) GetType() TaskReplyType {
	if m != nil {
		return m.Type
	}
	return TaskReplyType_TASK_COMPLETED_OK
}

type TaskRemovedResponse struct {
	Type TaskReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.TaskReplyType" json:"type,omitempty"`
}

func (m *TaskRemovedResponse) Reset()                    { *m = TaskRemovedResponse{} }
func (m *TaskRemovedResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskRemovedResponse) ProtoMessage()               {}
func (*TaskRemovedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *TaskRemovedResponse) GetType() TaskReplyType {
	if m != nil {
		return m.Type
	}
	return TaskReplyType_TASK_COMPLETED_OK
}

type TaskFailedResponse struct {
	Type TaskReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.TaskReplyType" json:"type,omitempty"`
}

func (m *TaskFailedResponse) Reset()                    { *m = TaskFailedResponse{} }
func (m *TaskFailedResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskFailedResponse) ProtoMessage()               {}
func (*TaskFailedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *TaskFailedResponse) GetType() TaskReplyType {
	if m != nil {
		return m.Type
	}
	return TaskReplyType_TASK_COMPLETED_OK
}

type TaskUpdatedResponse struct {
	Type TaskReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.TaskReplyType" json:"type,omitempty"`
}

func (m *TaskUpdatedResponse) Reset()                    { *m = TaskUpdatedResponse{} }
func (m *TaskUpdatedResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskUpdatedResponse) ProtoMessage()               {}
func (*TaskUpdatedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *TaskUpdatedResponse) GetType() TaskReplyType {
	if m != nil {
		return m.Type
	}
	return TaskReplyType_TASK_COMPLETED_OK
}

type NodeAddedResponse struct {
	Type NodeReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.NodeReplyType" json:"type,omitempty"`
}

func (m *NodeAddedResponse) Reset()                    { *m = NodeAddedResponse{} }
func (m *NodeAddedResponse) String() string            { return proto.CompactTextString(m) }
func (*NodeAddedResponse) ProtoMessage()               {}
func (*NodeAddedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *NodeAddedResponse) GetType() NodeReplyType {
	if m != nil {
		return m.Type
	}
	return NodeReplyType_NODE_ADDED_OK
}

type NodeRemovedResponse struct {
	Type NodeReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.NodeReplyType" json:"type,omitempty"`
}

func (m *NodeRemovedResponse) Reset()                    { *m = NodeRemovedResponse{} }
func (m *NodeRemovedResponse) String() string            { return proto.CompactTextString(m) }
func (*NodeRemovedResponse) ProtoMessage()               {}
func (*NodeRemovedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *NodeRemovedResponse) GetType() NodeReplyType {
	if m != nil {
		return m.Type
	}
	return NodeReplyType_NODE_ADDED_OK
}

type NodeFailedResponse struct {
	Type NodeReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.NodeReplyType" json:"type,omitempty"`
}

func (m *NodeFailedResponse) Reset()                    { *m = NodeFailedResponse{} }
func (m *NodeFailedResponse) String() string            { return proto.CompactTextString(m) }
func (*NodeFailedResponse) ProtoMessage()               {}
func (*NodeFailedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *NodeFailedResponse) GetType() NodeReplyType {
	if m != nil {
		return m.Type
	}
	return NodeReplyType_NODE_ADDED_OK
}

type NodeUpdatedResponse struct {
	Type NodeReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.NodeReplyType" json:"type,omitempty"`
}

func (m *NodeUpdatedResponse) Reset()                    { *m = NodeUpdatedResponse{} }
func (m *NodeUpdatedResponse) String() string            { return proto.CompactTextString(m) }
func (*NodeUpdatedResponse) ProtoMessage()               {}
func (*NodeUpdatedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *NodeUpdatedResponse) GetType() NodeReplyType {
	if m != nil {
		return m.Type
	}
	return NodeReplyType_NODE_ADDED_OK
}

type TaskStatsResponse struct {
	Type TaskReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.TaskReplyType" json:"type,omitempty"`
}

func (m *TaskStatsResponse) Reset()                    { *m = TaskStatsResponse{} }
func (m *TaskStatsResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskStatsResponse) ProtoMessage()               {}
func (*TaskStatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *TaskStatsResponse) GetType() TaskReplyType {
	if m != nil {
		return m.Type
	}
	return TaskReplyType_TASK_COMPLETED_OK
}

type ResourceStatsResponse struct {
	Type NodeReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.NodeReplyType" json:"type,omitempty"`
}

func (m *ResourceStatsResponse) Reset()                    { *m = ResourceStatsResponse{} }
func (m *ResourceStatsResponse) String() string            { return proto.CompactTextString(m) }
func (*ResourceStatsResponse) ProtoMessage()               {}
func (*ResourceStatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *ResourceStatsResponse) GetType() NodeReplyType {
	if m != nil {
		return m.Type
	}
	return NodeReplyType_NODE_ADDED_OK
}

type TaskUID struct {
	TaskUid uint64 `protobuf:"varint,1,opt,name=task_uid,json=taskUid" json:"task_uid,omitempty"`
}

func (m *TaskUID) Reset()                    { *m = TaskUID{} }
func (m *TaskUID) String() string            { return proto.CompactTextString(m) }
func (*TaskUID) ProtoMessage()               {}
func (*TaskUID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *TaskUID) GetTaskUid() uint64 {
	if m != nil {
		return m.TaskUid
	}
	return 0
}

type ResourceUID struct {
	ResourceUid string `protobuf:"bytes,1,opt,name=resource_uid,json=resourceUid" json:"resource_uid,omitempty"`
}

func (m *ResourceUID) Reset()                    { *m = ResourceUID{} }
func (m *ResourceUID) String() string            { return proto.CompactTextString(m) }
func (*ResourceUID) ProtoMessage()               {}
func (*ResourceUID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *ResourceUID) GetResourceUid() string {
	if m != nil {
		return m.ResourceUid
	}
	return ""
}

type HealthCheckRequest struct {
	GrpcService string `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService" json:"grpc_service,omitempty"`
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

func (m *HealthCheckRequest) GetGrpcService() string {
	if m != nil {
		return m.GrpcService
	}
	return ""
}

type HealthCheckResponse struct {
	Status ServingStatus `protobuf:"varint,1,opt,name=status,enum=firmament.ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{17} }

func (m *HealthCheckResponse) GetStatus() ServingStatus {
	if m != nil {
		return m.Status
	}
	return ServingStatus_UNKNOWN
}

func init() {
	proto.RegisterType((*ScheduleRequest)(nil), "firmament.ScheduleRequest")
	proto.RegisterType((*SchedulingDeltas)(nil), "firmament.SchedulingDeltas")
	proto.RegisterType((*TaskCompletedResponse)(nil), "firmament.TaskCompletedResponse")
	proto.RegisterType((*TaskDescription)(nil), "firmament.TaskDescription")
	proto.RegisterType((*TaskSubmittedResponse)(nil), "firmament.TaskSubmittedResponse")
	proto.RegisterType((*TaskRemovedResponse)(nil), "firmament.TaskRemovedResponse")
	proto.RegisterType((*TaskFailedResponse)(nil), "firmament.TaskFailedResponse")
	proto.RegisterType((*TaskUpdatedResponse)(nil), "firmament.TaskUpdatedResponse")
	proto.RegisterType((*NodeAddedResponse)(nil), "firmament.NodeAddedResponse")
	proto.RegisterType((*NodeRemovedResponse)(nil), "firmament.NodeRemovedResponse")
	proto.RegisterType((*NodeFailedResponse)(nil), "firmament.NodeFailedResponse")
	proto.RegisterType((*NodeUpdatedResponse)(nil), "firmament.NodeUpdatedResponse")
	proto.RegisterType((*TaskStatsResponse)(nil), "firmament.TaskStatsResponse")
	proto.RegisterType((*ResourceStatsResponse)(nil), "firmament.ResourceStatsResponse")
	proto.RegisterType((*TaskUID)(nil), "firmament.TaskUID")
	proto.RegisterType((*ResourceUID)(nil), "firmament.ResourceUID")
	proto.RegisterType((*HealthCheckRequest)(nil), "firmament.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "firmament.HealthCheckResponse")
	proto.RegisterEnum("firmament.TaskReplyType", TaskReplyType_name, TaskReplyType_value)
	proto.RegisterEnum("firmament.NodeReplyType", NodeReplyType_name, NodeReplyType_value)
	proto.RegisterEnum("firmament.ServingStatus", ServingStatus_name, ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FirmamentScheduler service

type FirmamentSchedulerClient interface {
	Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*SchedulingDeltas, error)
	TaskCompleted(ctx context.Context, in *TaskUID, opts ...grpc.CallOption) (*TaskCompletedResponse, error)
	TaskFailed(ctx context.Context, in *TaskUID, opts ...grpc.CallOption) (*TaskFailedResponse, error)
	TaskRemoved(ctx context.Context, in *TaskUID, opts ...grpc.CallOption) (*TaskRemovedResponse, error)
	TaskSubmitted(ctx context.Context, in *TaskDescription, opts ...grpc.CallOption) (*TaskSubmittedResponse, error)
	TaskUpdated(ctx context.Context, in *TaskDescription, opts ...grpc.CallOption) (*TaskUpdatedResponse, error)
	NodeAdded(ctx context.Context, in *ResourceTopologyNodeDescriptor, opts ...grpc.CallOption) (*NodeAddedResponse, error)
	NodeFailed(ctx context.Context, in *ResourceUID, opts ...grpc.CallOption) (*NodeFailedResponse, error)
	NodeRemoved(ctx context.Context, in *ResourceUID, opts ...grpc.CallOption) (*NodeRemovedResponse, error)
	NodeUpdated(ctx context.Context, in *ResourceTopologyNodeDescriptor, opts ...grpc.CallOption) (*NodeUpdatedResponse, error)
	AddTaskStats(ctx context.Context, in *TaskStats, opts ...grpc.CallOption) (*TaskStatsResponse, error)
	AddNodeStats(ctx context.Context, in *ResourceStats, opts ...grpc.CallOption) (*ResourceStatsResponse, error)
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type firmamentSchedulerClient struct {
	cc *grpc.ClientConn
}

func NewFirmamentSchedulerClient(cc *grpc.ClientConn) FirmamentSchedulerClient {
	return &firmamentSchedulerClient{cc}
}

func (c *firmamentSchedulerClient) Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*SchedulingDeltas, error) {
	out := new(SchedulingDeltas)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/Schedule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) TaskCompleted(ctx context.Context, in *TaskUID, opts ...grpc.CallOption) (*TaskCompletedResponse, error) {
	out := new(TaskCompletedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/TaskCompleted", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) TaskFailed(ctx context.Context, in *TaskUID, opts ...grpc.CallOption) (*TaskFailedResponse, error) {
	out := new(TaskFailedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/TaskFailed", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) TaskRemoved(ctx context.Context, in *TaskUID, opts ...grpc.CallOption) (*TaskRemovedResponse, error) {
	out := new(TaskRemovedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/TaskRemoved", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) TaskSubmitted(ctx context.Context, in *TaskDescription, opts ...grpc.CallOption) (*TaskSubmittedResponse, error) {
	out := new(TaskSubmittedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/TaskSubmitted", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) TaskUpdated(ctx context.Context, in *TaskDescription, opts ...grpc.CallOption) (*TaskUpdatedResponse, error) {
	out := new(TaskUpdatedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/TaskUpdated", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) NodeAdded(ctx context.Context, in *ResourceTopologyNodeDescriptor, opts ...grpc.CallOption) (*NodeAddedResponse, error) {
	out := new(NodeAddedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/NodeAdded", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) NodeFailed(ctx context.Context, in *ResourceUID, opts ...grpc.CallOption) (*NodeFailedResponse, error) {
	out := new(NodeFailedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/NodeFailed", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) NodeRemoved(ctx context.Context, in *ResourceUID, opts ...grpc.CallOption) (*NodeRemovedResponse, error) {
	out := new(NodeRemovedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/NodeRemoved", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) NodeUpdated(ctx context.Context, in *ResourceTopologyNodeDescriptor, opts ...grpc.CallOption) (*NodeUpdatedResponse, error) {
	out := new(NodeUpdatedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/NodeUpdated", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) AddTaskStats(ctx context.Context, in *TaskStats, opts ...grpc.CallOption) (*TaskStatsResponse, error) {
	out := new(TaskStatsResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/AddTaskStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) AddNodeStats(ctx context.Context, in *ResourceStats, opts ...grpc.CallOption) (*ResourceStatsResponse, error) {
	out := new(ResourceStatsResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/AddNodeStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FirmamentScheduler service

type FirmamentSchedulerServer interface {
	Schedule(context.Context, *ScheduleRequest) (*SchedulingDeltas, error)
	TaskCompleted(context.Context, *TaskUID) (*TaskCompletedResponse, error)
	TaskFailed(context.Context, *TaskUID) (*TaskFailedResponse, error)
	TaskRemoved(context.Context, *TaskUID) (*TaskRemovedResponse, error)
	TaskSubmitted(context.Context, *TaskDescription) (*TaskSubmittedResponse, error)
	TaskUpdated(context.Context, *TaskDescription) (*TaskUpdatedResponse, error)
	NodeAdded(context.Context, *ResourceTopologyNodeDescriptor) (*NodeAddedResponse, error)
	NodeFailed(context.Context, *ResourceUID) (*NodeFailedResponse, error)
	NodeRemoved(context.Context, *ResourceUID) (*NodeRemovedResponse, error)
	NodeUpdated(context.Context, *ResourceTopologyNodeDescriptor) (*NodeUpdatedResponse, error)
	AddTaskStats(context.Context, *TaskStats) (*TaskStatsResponse, error)
	AddNodeStats(context.Context, *ResourceStats) (*ResourceStatsResponse, error)
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterFirmamentSchedulerServer(s *grpc.Server, srv FirmamentSchedulerServer) {
	s.RegisterService(&_FirmamentScheduler_serviceDesc, srv)
}

func _FirmamentScheduler_Schedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).Schedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/Schedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).Schedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_TaskCompleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).TaskCompleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/TaskCompleted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).TaskCompleted(ctx, req.(*TaskUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_TaskFailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).TaskFailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/TaskFailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).TaskFailed(ctx, req.(*TaskUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_TaskRemoved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).TaskRemoved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/TaskRemoved",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).TaskRemoved(ctx, req.(*TaskUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_TaskSubmitted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskDescription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).TaskSubmitted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/TaskSubmitted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).TaskSubmitted(ctx, req.(*TaskDescription))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_TaskUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskDescription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).TaskUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/TaskUpdated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).TaskUpdated(ctx, req.(*TaskDescription))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_NodeAdded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceTopologyNodeDescriptor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).NodeAdded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/NodeAdded",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).NodeAdded(ctx, req.(*ResourceTopologyNodeDescriptor))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_NodeFailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).NodeFailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/NodeFailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).NodeFailed(ctx, req.(*ResourceUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_NodeRemoved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).NodeRemoved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/NodeRemoved",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).NodeRemoved(ctx, req.(*ResourceUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_NodeUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceTopologyNodeDescriptor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).NodeUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/NodeUpdated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).NodeUpdated(ctx, req.(*ResourceTopologyNodeDescriptor))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_AddTaskStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskStats)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).AddTaskStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/AddTaskStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).AddTaskStats(ctx, req.(*TaskStats))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_AddNodeStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceStats)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).AddNodeStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/AddNodeStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).AddNodeStats(ctx, req.(*ResourceStats))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FirmamentScheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "firmament.FirmamentScheduler",
	HandlerType: (*FirmamentSchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Schedule",
			Handler:    _FirmamentScheduler_Schedule_Handler,
		},
		{
			MethodName: "TaskCompleted",
			Handler:    _FirmamentScheduler_TaskCompleted_Handler,
		},
		{
			MethodName: "TaskFailed",
			Handler:    _FirmamentScheduler_TaskFailed_Handler,
		},
		{
			MethodName: "TaskRemoved",
			Handler:    _FirmamentScheduler_TaskRemoved_Handler,
		},
		{
			MethodName: "TaskSubmitted",
			Handler:    _FirmamentScheduler_TaskSubmitted_Handler,
		},
		{
			MethodName: "TaskUpdated",
			Handler:    _FirmamentScheduler_TaskUpdated_Handler,
		},
		{
			MethodName: "NodeAdded",
			Handler:    _FirmamentScheduler_NodeAdded_Handler,
		},
		{
			MethodName: "NodeFailed",
			Handler:    _FirmamentScheduler_NodeFailed_Handler,
		},
		{
			MethodName: "NodeRemoved",
			Handler:    _FirmamentScheduler_NodeRemoved_Handler,
		},
		{
			MethodName: "NodeUpdated",
			Handler:    _FirmamentScheduler_NodeUpdated_Handler,
		},
		{
			MethodName: "AddTaskStats",
			Handler:    _FirmamentScheduler_AddTaskStats_Handler,
		},
		{
			MethodName: "AddNodeStats",
			Handler:    _FirmamentScheduler_AddNodeStats_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _FirmamentScheduler_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "firmament_scheduler.proto",
}

func init() { proto.RegisterFile("firmament_scheduler.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 926 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0xa5, 0xfc, 0x25, 0x7b, 0x14, 0x49, 0xd4, 0x38, 0x76, 0x6d, 0xb5, 0x0d, 0x1c, 0xa2, 0x87,
	0xd4, 0x28, 0x8c, 0x40, 0x3d, 0x14, 0xe8, 0xa5, 0xa0, 0x44, 0xca, 0x55, 0x6c, 0x8b, 0x01, 0x49,
	0xb9, 0x1f, 0x17, 0x82, 0x16, 0xb7, 0x36, 0x63, 0x49, 0x64, 0x49, 0x2a, 0x80, 0x7f, 0x44, 0xaf,
	0xfd, 0x7b, 0xbd, 0xf7, 0x57, 0x14, 0xbb, 0x24, 0x57, 0xfc, 0x72, 0xa0, 0x2a, 0xc7, 0x7d, 0x3b,
	0xf3, 0x76, 0xde, 0x2c, 0xf7, 0x0d, 0xe1, 0xf4, 0x0f, 0x37, 0x98, 0xdb, 0x73, 0xb2, 0x88, 0xac,
	0x70, 0xfa, 0x40, 0x9c, 0xe5, 0x8c, 0x04, 0x17, 0x7e, 0xe0, 0x45, 0x1e, 0x1e, 0xf0, 0xad, 0x6e,
	0xeb, 0x83, 0x77, 0x67, 0x39, 0x24, 0x9c, 0xc6, 0x5b, 0xdd, 0x97, 0x01, 0x09, 0xbd, 0x65, 0x30,
	0x25, 0x56, 0x18, 0xd9, 0x51, 0x98, 0xa0, 0xaf, 0x39, 0x1a, 0x79, 0xbe, 0x37, 0xf3, 0xee, 0x9f,
	0xac, 0x85, 0xe7, 0x90, 0x6c, 0x62, 0x3b, 0xb2, 0xc3, 0xc7, 0x2c, 0x20, 0x32, 0x20, 0xcb, 0x72,
	0x9c, 0xd4, 0xe1, 0x2e, 0xee, 0x2d, 0x87, 0xcc, 0x22, 0x3b, 0xc6, 0xa5, 0x0e, 0xb4, 0x8d, 0xa4,
	0x42, 0x9d, 0xfc, 0xb9, 0x24, 0x61, 0x24, 0x0d, 0x41, 0x34, 0x78, 0xb0, 0x42, 0x63, 0x43, 0xec,
	0xc1, 0x1e, 0xcb, 0x0a, 0x4f, 0x6a, 0x67, 0xdb, 0x6f, 0x1a, 0xbd, 0xee, 0x05, 0x97, 0x71, 0x51,
	0x08, 0xd6, 0x93, 0x48, 0x49, 0x85, 0x23, 0xd3, 0x0e, 0x1f, 0x07, 0xde, 0xdc, 0x9f, 0x91, 0x88,
	0x38, 0x3a, 0x09, 0x7d, 0x6f, 0x11, 0x12, 0xfc, 0x0e, 0x76, 0xa2, 0x27, 0x9f, 0x9c, 0xd4, 0xce,
	0x6a, 0x6f, 0x5a, 0xbd, 0x93, 0x0c, 0x15, 0x8d, 0xd7, 0x89, 0x3f, 0x7b, 0x32, 0x9f, 0x7c, 0xa2,
	0xb3, 0x28, 0xe9, 0xef, 0x1a, 0xb4, 0x29, 0xae, 0x90, 0x70, 0x1a, 0xb8, 0x7e, 0xe4, 0x7a, 0x0b,
	0xec, 0xc3, 0x4a, 0x32, 0xc5, 0xbc, 0x80, 0x91, 0x35, 0x7a, 0xa7, 0x05, 0x32, 0x85, 0x07, 0xe8,
	0xad, 0x28, 0xb7, 0xc6, 0x9f, 0x80, 0xf7, 0x3f, 0xa1, 0xd8, 0x62, 0x14, 0xd9, 0x7a, 0xde, 0x79,
	0x77, 0x19, 0x86, 0xe6, 0x87, 0xec, 0x32, 0xd5, 0x67, 0x2c, 0xef, 0xe6, 0x6e, 0xb4, 0xb9, 0xbe,
	0x01, 0x1c, 0xc6, 0xf0, 0xdc, 0xfb, 0xb8, 0x31, 0x49, 0x1f, 0x90, 0xc2, 0x43, 0xdb, 0x9d, 0x7d,
	0x6e, 0x21, 0x13, 0xdf, 0xb1, 0x37, 0x57, 0x23, 0x43, 0x67, 0xec, 0x39, 0x44, 0x76, 0x9c, 0xb5,
	0x28, 0x68, 0x6c, 0x45, 0x1d, 0x31, 0xbc, 0x6e, 0x43, 0xaa, 0x48, 0xfa, 0x80, 0x14, 0x5e, 0xbb,
	0x21, 0x9f, 0x28, 0x64, 0xfd, 0x86, 0x54, 0x91, 0xc8, 0xd0, 0x61, 0x5f, 0x09, 0x7d, 0x8b, 0x1b,
	0xf6, 0x54, 0x85, 0x23, 0x3d, 0xf1, 0x80, 0x75, 0x69, 0xaa, 0x2a, 0xf9, 0x06, 0xea, 0xec, 0x7e,
	0x47, 0x0a, 0x9e, 0xc2, 0x3e, 0x7b, 0x3f, 0x4b, 0xd7, 0x61, 0xc9, 0x3b, 0x7a, 0x9d, 0xae, 0x27,
	0xae, 0x23, 0xbd, 0x85, 0x46, 0x7a, 0x18, 0x8d, 0x7c, 0x0d, 0x2f, 0xb8, 0xff, 0xa4, 0xd1, 0x07,
	0x7a, 0x23, 0xc5, 0x68, 0xc6, 0x0f, 0x80, 0x3f, 0x13, 0x7b, 0x16, 0x3d, 0x0c, 0x1e, 0xc8, 0xf4,
	0x31, 0x71, 0x11, 0x9a, 0x78, 0x1f, 0xf8, 0x53, 0x2b, 0x24, 0xc1, 0x47, 0x77, 0x4a, 0xd2, 0x44,
	0x8a, 0x19, 0x31, 0x24, 0x5d, 0xc2, 0x61, 0x2e, 0x31, 0x51, 0xf5, 0x16, 0xf6, 0xa8, 0x73, 0x2d,
	0xc3, 0x0a, 0x5d, 0x2c, 0x75, 0x71, 0x6f, 0xb0, 0x7d, 0x3d, 0x89, 0x3b, 0xff, 0xa7, 0x06, 0xcd,
	0x5c, 0xe3, 0xf0, 0x08, 0x3a, 0xa6, 0x6c, 0x5c, 0x59, 0x03, 0xed, 0xe6, 0xfd, 0xb5, 0x6a, 0xaa,
	0x8a, 0xa5, 0x5d, 0x89, 0x02, 0x87, 0x8d, 0x49, 0xff, 0x66, 0x64, 0x26, 0x70, 0x0d, 0x0f, 0xa1,
	0xcd, 0x60, 0x5d, 0xbd, 0xd1, 0x6e, 0x63, 0x70, 0x0b, 0x11, 0x5a, 0x0c, 0x1c, 0xca, 0xa3, 0xeb,
	0x18, 0xdb, 0xe6, 0x81, 0x93, 0xf7, 0x8a, 0x9c, 0x64, 0xef, 0xf0, 0xc0, 0xb1, 0x66, 0x5a, 0x43,
	0x6d, 0x32, 0x56, 0xc4, 0x5d, 0x3c, 0x06, 0x64, 0xd8, 0x3b, 0xad, 0x9f, 0xc1, 0xf7, 0xb0, 0x0b,
	0xc7, 0x0c, 0x97, 0xaf, 0x75, 0x55, 0x56, 0x7e, 0x5b, 0x15, 0x22, 0xd6, 0xf9, 0x9e, 0x61, 0xca,
	0xa6, 0xca, 0xb2, 0x06, 0xba, 0x4a, 0x8f, 0x11, 0xf7, 0xcf, 0xff, 0xaa, 0x41, 0x33, 0x77, 0xa7,
	0xd8, 0x81, 0xe6, 0x58, 0x53, 0x54, 0x4b, 0x56, 0x94, 0x54, 0x1d, 0x42, 0x8b, 0x41, 0xab, 0x8a,
	0x99, 0x34, 0x86, 0xe5, 0xa4, 0xa5, 0x60, 0x46, 0xc6, 0x36, 0xcf, 0x5e, 0x95, 0xbb, 0x83, 0x5f,
	0xc0, 0x61, 0x7c, 0x48, 0x52, 0xae, 0xfa, 0xeb, 0xc8, 0x30, 0x0d, 0x71, 0xf7, 0xfc, 0x47, 0x68,
	0xe6, 0xae, 0x02, 0x1b, 0x50, 0x9f, 0x8c, 0xaf, 0xc6, 0xda, 0x2f, 0x63, 0x51, 0xa0, 0x0b, 0x43,
	0xd5, 0x6f, 0x47, 0xe3, 0x4b, 0xb1, 0x86, 0x6d, 0x68, 0x50, 0xca, 0x14, 0xd8, 0xea, 0xfd, 0x5b,
	0x07, 0x1c, 0xa6, 0x37, 0x9a, 0x0e, 0x9f, 0x00, 0x55, 0xd8, 0x4f, 0x17, 0x58, 0x31, 0x5e, 0xd2,
	0xf1, 0xd4, 0xfd, 0xf2, 0xf9, 0xd1, 0x13, 0x4a, 0x02, 0x5e, 0xc6, 0x9f, 0x02, 0x9f, 0x3a, 0x88,
	0x85, 0xd7, 0x35, 0x19, 0x29, 0xdd, 0xb3, 0x02, 0x56, 0x9a, 0x51, 0x92, 0x80, 0x32, 0xc0, 0xca,
	0x52, 0x2b, 0x59, 0xbe, 0x2e, 0x60, 0x79, 0xb3, 0x91, 0x04, 0x1c, 0x40, 0x23, 0x63, 0xed, 0x95,
	0x1c, 0xaf, 0x4a, 0x6f, 0x3f, 0xe7, 0x7a, 0x92, 0x80, 0x5a, 0x2c, 0x88, 0x8f, 0x99, 0x5c, 0x73,
	0x0a, 0x83, 0xb1, 0x24, 0xac, 0x34, 0x9c, 0x24, 0x01, 0xaf, 0xe2, 0xaa, 0x12, 0x5b, 0xfb, 0x24,
	0x5d, 0xb1, 0xba, 0x82, 0x15, 0x4a, 0x02, 0xde, 0xc2, 0x01, 0xf7, 0x7b, 0xfc, 0x36, 0x13, 0x9e,
	0x9a, 0x88, 0x99, 0xfc, 0xb4, 0xd0, 0xa8, 0xd5, 0xf0, 0xec, 0x7e, 0x55, 0x30, 0xab, 0xdc, 0xc0,
	0x90, 0x04, 0x54, 0x01, 0x56, 0xfe, 0x8d, 0xc7, 0x15, 0xc4, 0xc5, 0x1b, 0x28, 0xdb, 0x3d, 0xfb,
	0x1a, 0x1a, 0x99, 0x59, 0xf2, 0x2c, 0xcf, 0xab, 0x92, 0x75, 0x16, 0x6f, 0xe1, 0xf7, 0x98, 0x28,
	0x6d, 0xda, 0xff, 0x50, 0x5a, 0xe4, 0x2e, 0xf7, 0x50, 0x81, 0x17, 0xb2, 0xe3, 0xf0, 0x29, 0x81,
	0x2f, 0x8b, 0x97, 0x48, 0xd1, 0x5c, 0xc7, 0x4a, 0x13, 0x45, 0x12, 0xf0, 0x9a, 0xb1, 0xd0, 0x13,
	0x62, 0x96, 0x93, 0x8a, 0x12, 0x63, 0xa6, 0xb3, 0xe7, 0x76, 0x32, 0x6c, 0x43, 0xd8, 0x65, 0xae,
	0x8c, 0xd9, 0x16, 0x97, 0x6d, 0x3e, 0xa7, 0xae, 0xc2, 0xcc, 0xef, 0xf6, 0xd8, 0x6f, 0xe6, 0xf7,
	0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xea, 0x0c, 0x38, 0xbd, 0x12, 0x0b, 0x00, 0x00,
}
