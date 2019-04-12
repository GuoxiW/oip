// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RecordTemplateProto.proto

package oip5

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RecordTemplateProto struct {
	// Human readable name to quickly identify type (non-unique)
	FriendlyName string `protobuf:"bytes,1,opt,name=friendlyName" json:"friendlyName,omitempty"`
	// Description of the purpose behind this new type
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Compiled descriptor including dependencies
	DescriptorSetProto []byte `protobuf:"bytes,4,opt,name=DescriptorSetProto,proto3" json:"DescriptorSetProto,omitempty"`
	// Populated by oipd with the unique identifier for this type
	Identifier int64 `protobuf:"fixed64,10,opt,name=identifier" json:"identifier,omitempty"`
	// List of unique template identifiers recommended for use with this template
	Recommended []uint64 `protobuf:"fixed64,11,rep,packed,name=recommended" json:"recommended,omitempty"`
	// List of unique template identifiers required for use with this template
	Required []uint64 `protobuf:"fixed64,12,rep,packed,name=required" json:"required,omitempty"`
}

func (m *RecordTemplateProto) Reset()                    { *m = RecordTemplateProto{} }
func (m *RecordTemplateProto) String() string            { return proto.CompactTextString(m) }
func (*RecordTemplateProto) ProtoMessage()               {}
func (*RecordTemplateProto) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *RecordTemplateProto) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func (m *RecordTemplateProto) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RecordTemplateProto) GetDescriptorSetProto() []byte {
	if m != nil {
		return m.DescriptorSetProto
	}
	return nil
}

func (m *RecordTemplateProto) GetIdentifier() int64 {
	if m != nil {
		return m.Identifier
	}
	return 0
}

func (m *RecordTemplateProto) GetRecommended() []uint64 {
	if m != nil {
		return m.Recommended
	}
	return nil
}

func (m *RecordTemplateProto) GetRequired() []uint64 {
	if m != nil {
		return m.Required
	}
	return nil
}

func init() {
	proto.RegisterType((*RecordTemplateProto)(nil), "oip5.record.RecordTemplateProto")
}

func init() { proto.RegisterFile("RecordTemplateProto.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3d, 0x4b, 0xc7, 0x30,
	0x10, 0xc6, 0x89, 0x96, 0xa0, 0xd7, 0x4e, 0x71, 0x89, 0x0e, 0x12, 0x3a, 0x65, 0xea, 0x22, 0x7e,
	0x01, 0x71, 0x16, 0x89, 0x4e, 0x6e, 0xb5, 0xb9, 0xc2, 0x41, 0xf3, 0xe2, 0x19, 0x07, 0x3f, 0xb2,
	0xdf, 0x42, 0x12, 0x44, 0xfa, 0x87, 0x2e, 0x07, 0xcf, 0xef, 0x79, 0xb8, 0x37, 0xb8, 0x76, 0xb8,
	0x24, 0xf6, 0xaf, 0x18, 0xf2, 0x36, 0x17, 0x7c, 0xe6, 0x54, 0xd2, 0x94, 0x6b, 0x55, 0x7d, 0xa2,
	0x7c, 0x3f, 0x71, 0xf3, 0xc7, 0x1f, 0x01, 0x57, 0x07, 0x51, 0x35, 0xc2, 0xb0, 0x32, 0x61, 0xf4,
	0xdb, 0xf7, 0xd3, 0x1c, 0x50, 0x0b, 0x23, 0xec, 0xa5, 0x3b, 0x61, 0xca, 0x40, 0xef, 0xf1, 0x73,
	0x61, 0xca, 0x85, 0x52, 0xd4, 0x67, 0x2d, 0xb2, 0x47, 0x6a, 0x02, 0xf5, 0xf8, 0x27, 0x13, 0xbf,
	0x60, 0x69, 0xbd, 0x75, 0x67, 0x84, 0x1d, 0xdc, 0x81, 0xa3, 0x6e, 0x01, 0xc8, 0x63, 0x2c, 0xb4,
	0x12, 0xb2, 0x06, 0x23, 0xac, 0x74, 0x3b, 0x52, 0x27, 0xd6, 0xbd, 0x43, 0xc0, 0xe8, 0xd1, 0xeb,
	0xde, 0x9c, 0x5b, 0xe9, 0xf6, 0x48, 0xdd, 0xc0, 0x05, 0xe3, 0xc7, 0x17, 0x31, 0x7a, 0x3d, 0x34,
	0xfb, 0x5f, 0x3f, 0xc8, 0xb7, 0xae, 0x9e, 0xfe, 0x2e, 0xdb, 0x1f, 0xee, 0x7e, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x59, 0x7f, 0x5c, 0x18, 0x24, 0x01, 0x00, 0x00,
}
