// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node_affinity.proto

package firmament

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NodeSelectorRequirement struct {
	Key      string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Operator string   `protobuf:"bytes,2,opt,name=operator" json:"operator,omitempty"`
	Values   []string `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
}

func (m *NodeSelectorRequirement) Reset()                    { *m = NodeSelectorRequirement{} }
func (m *NodeSelectorRequirement) String() string            { return proto.CompactTextString(m) }
func (*NodeSelectorRequirement) ProtoMessage()               {}
func (*NodeSelectorRequirement) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *NodeSelectorRequirement) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *NodeSelectorRequirement) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *NodeSelectorRequirement) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type NodeSelectorTerm struct {
	// A list of node selector requirements by node's labels.
	// +optional
	MatchExpressions []*NodeSelectorRequirement `protobuf:"bytes,1,rep,name=matchExpressions" json:"matchExpressions,omitempty"`
	// A list of node selector requirements by node's fields.
	// +optional
	MatchFields []*NodeSelectorRequirement `protobuf:"bytes,2,rep,name=matchFields" json:"matchFields,omitempty"`
}

func (m *NodeSelectorTerm) Reset()                    { *m = NodeSelectorTerm{} }
func (m *NodeSelectorTerm) String() string            { return proto.CompactTextString(m) }
func (*NodeSelectorTerm) ProtoMessage()               {}
func (*NodeSelectorTerm) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *NodeSelectorTerm) GetMatchExpressions() []*NodeSelectorRequirement {
	if m != nil {
		return m.MatchExpressions
	}
	return nil
}

func (m *NodeSelectorTerm) GetMatchFields() []*NodeSelectorRequirement {
	if m != nil {
		return m.MatchFields
	}
	return nil
}

type NodeSelector struct {
	// Required. A list of node selector terms. The terms are ORed.
	NodeSelectorTerms []*NodeSelectorTerm `protobuf:"bytes,1,rep,name=nodeSelectorTerms" json:"nodeSelectorTerms,omitempty"`
}

func (m *NodeSelector) Reset()                    { *m = NodeSelector{} }
func (m *NodeSelector) String() string            { return proto.CompactTextString(m) }
func (*NodeSelector) ProtoMessage()               {}
func (*NodeSelector) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *NodeSelector) GetNodeSelectorTerms() []*NodeSelectorTerm {
	if m != nil {
		return m.NodeSelectorTerms
	}
	return nil
}

type PreferredSchedulingTerm struct {
	// Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
	Weight int32 `protobuf:"varint,1,opt,name=weight" json:"weight,omitempty"`
	// A node selector term, associated with the corresponding weight.
	Preference *NodeSelectorTerm `protobuf:"bytes,2,opt,name=preference" json:"preference,omitempty"`
}

func (m *PreferredSchedulingTerm) Reset()                    { *m = PreferredSchedulingTerm{} }
func (m *PreferredSchedulingTerm) String() string            { return proto.CompactTextString(m) }
func (*PreferredSchedulingTerm) ProtoMessage()               {}
func (*PreferredSchedulingTerm) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *PreferredSchedulingTerm) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *PreferredSchedulingTerm) GetPreference() *NodeSelectorTerm {
	if m != nil {
		return m.Preference
	}
	return nil
}

type NodeAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  *NodeSelector              `protobuf:"bytes,1,opt,name=requiredDuringSchedulingIgnoredDuringExecution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredSchedulingTerm `protobuf:"bytes,2,rep,name=preferredDuringSchedulingIgnoredDuringExecution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

func (m *NodeAffinity) Reset()                    { *m = NodeAffinity{} }
func (m *NodeAffinity) String() string            { return proto.CompactTextString(m) }
func (*NodeAffinity) ProtoMessage()               {}
func (*NodeAffinity) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *NodeAffinity) GetRequiredDuringSchedulingIgnoredDuringExecution() *NodeSelector {
	if m != nil {
		return m.RequiredDuringSchedulingIgnoredDuringExecution
	}
	return nil
}

func (m *NodeAffinity) GetPreferredDuringSchedulingIgnoredDuringExecution() []*PreferredSchedulingTerm {
	if m != nil {
		return m.PreferredDuringSchedulingIgnoredDuringExecution
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeSelectorRequirement)(nil), "firmament.NodeSelectorRequirement")
	proto.RegisterType((*NodeSelectorTerm)(nil), "firmament.NodeSelectorTerm")
	proto.RegisterType((*NodeSelector)(nil), "firmament.NodeSelector")
	proto.RegisterType((*PreferredSchedulingTerm)(nil), "firmament.PreferredSchedulingTerm")
	proto.RegisterType((*NodeAffinity)(nil), "firmament.NodeAffinity")
}

func init() { proto.RegisterFile("node_affinity.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x53, 0x1a, 0x89, 0x0c, 0x1e, 0x70, 0x4d, 0xa0, 0xd1, 0x0b, 0xe9, 0x89, 0x53, 0x4d,
	0xf0, 0xe8, 0xc9, 0x04, 0x4c, 0xb8, 0x10, 0xb3, 0x78, 0xf1, 0x44, 0x6a, 0x3b, 0x2d, 0x1b, 0xdb,
	0xdd, 0x3a, 0xdd, 0x2a, 0x9c, 0x7c, 0x01, 0x9f, 0xc1, 0xb3, 0x8f, 0x69, 0x58, 0xfe, 0xd8, 0xd8,
	0x34, 0xc2, 0xad, 0x33, 0x9d, 0xef, 0xdb, 0xdf, 0x7e, 0x93, 0x85, 0x0b, 0xa9, 0x42, 0x9c, 0xfb,
	0x51, 0x24, 0xa4, 0xd0, 0x2b, 0x2f, 0x23, 0xa5, 0x15, 0x6b, 0x45, 0x82, 0x52, 0x3f, 0x45, 0xa9,
	0xdd, 0x39, 0xf4, 0xa6, 0x2a, 0xc4, 0x19, 0x26, 0x18, 0x68, 0x45, 0x1c, 0x5f, 0x0b, 0x41, 0xb8,
	0xfe, 0xc5, 0x3a, 0x60, 0xbf, 0xe0, 0xca, 0xb1, 0xfa, 0xd6, 0xa0, 0xc5, 0xd7, 0x9f, 0xec, 0x12,
	0x4e, 0x55, 0x86, 0xe4, 0x6b, 0x45, 0x4e, 0xc3, 0xb4, 0xf7, 0x35, 0xeb, 0x42, 0xf3, 0xcd, 0x4f,
	0x0a, 0xcc, 0x1d, 0xbb, 0x6f, 0x0f, 0x5a, 0x7c, 0x5b, 0xb9, 0xdf, 0x16, 0x74, 0xca, 0x27, 0x3c,
	0x22, 0xa5, 0x6c, 0x0a, 0x9d, 0xd4, 0xd7, 0xc1, 0x62, 0xbc, 0xcc, 0x08, 0xf3, 0x5c, 0x28, 0x99,
	0x3b, 0x56, 0xdf, 0x1e, 0xb4, 0x87, 0xae, 0xb7, 0x67, 0xf3, 0x6a, 0xc0, 0x78, 0x45, 0xcb, 0x46,
	0xd0, 0x36, 0xbd, 0x7b, 0x81, 0x49, 0x98, 0x3b, 0x8d, 0x83, 0xad, 0xca, 0x32, 0xf7, 0x09, 0xce,
	0xca, 0x73, 0x6c, 0x02, 0xe7, 0xf2, 0x0f, 0xf9, 0x0e, 0xf3, 0xaa, 0xc6, 0x7b, 0x3d, 0xc3, 0xab,
	0x2a, 0x57, 0x42, 0xef, 0x81, 0x30, 0x42, 0x22, 0x0c, 0x67, 0xc1, 0x02, 0xc3, 0x22, 0x11, 0x32,
	0x36, 0x59, 0x74, 0xa1, 0xf9, 0x8e, 0x22, 0x5e, 0x68, 0x93, 0xf4, 0x09, 0xdf, 0x56, 0xec, 0x16,
	0x20, 0x33, 0x12, 0x94, 0x01, 0x9a, 0xb8, 0xff, 0x39, 0xb6, 0x34, 0xee, 0x7e, 0x35, 0x36, 0x77,
	0xb9, 0xdb, 0x2e, 0x9e, 0x7d, 0x80, 0x47, 0x9b, 0x7b, 0x87, 0xa3, 0x82, 0x84, 0x8c, 0x7f, 0x29,
	0x26, 0xb1, 0x54, 0xfb, 0xf6, 0x78, 0x89, 0x41, 0xa1, 0x85, 0x92, 0x86, 0xa6, 0x3d, 0xec, 0xd5,
	0x85, 0x78, 0xa4, 0x1d, 0xfb, 0xb4, 0xe0, 0x3a, 0xdb, 0x45, 0x70, 0x20, 0x42, 0x75, 0x8f, 0x35,
	0x21, 0xf2, 0x63, 0xad, 0x9f, 0x9b, 0xe6, 0x25, 0xdc, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb7,
	0xa4, 0x7e, 0x1b, 0x20, 0x03, 0x00, 0x00,
}
