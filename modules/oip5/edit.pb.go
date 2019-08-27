// Code generated by protoc-gen-go. DO NOT EDIT.
// source: edit.proto

package oip5

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import oipProto "."

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Step_Action int32

const (
	// 0 is reserved for errors.
	Step_ACTION_ERROR        Step_Action = 0
	Step_ACTION_REPLACE_ALL  Step_Action = 1
	Step_ACTION_APPEND_ALL   Step_Action = 2
	Step_ACTION_REMOVE_ALL   Step_Action = 3
	Step_ACTION_REMOVE_ONE   Step_Action = 4
	Step_ACTION_REPLACE_ONE  Step_Action = 5
	Step_ACTION_STRING_PATCH Step_Action = 6
	Step_ACTION_STEP_INTO    Step_Action = 7
)

var Step_Action_name = map[int32]string{
	0: "ACTION_ERROR",
	1: "ACTION_REPLACE_ALL",
	2: "ACTION_APPEND_ALL",
	3: "ACTION_REMOVE_ALL",
	4: "ACTION_REMOVE_ONE",
	5: "ACTION_REPLACE_ONE",
	6: "ACTION_STRING_PATCH",
	7: "ACTION_STEP_INTO",
}
var Step_Action_value = map[string]int32{
	"ACTION_ERROR":        0,
	"ACTION_REPLACE_ALL":  1,
	"ACTION_APPEND_ALL":   2,
	"ACTION_REMOVE_ALL":   3,
	"ACTION_REMOVE_ONE":   4,
	"ACTION_REPLACE_ONE":  5,
	"ACTION_STRING_PATCH": 6,
	"ACTION_STEP_INTO":    7,
}

func (x Step_Action) String() string {
	return proto.EnumName(Step_Action_name, int32(x))
}
func (Step_Action) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{2, 0} }

type EditProto struct {
	Reference *oipProto.Txid `protobuf:"bytes,1,opt,name=reference" json:"reference,omitempty"`
	NewValues *RecordProto   `protobuf:"bytes,2,opt,name=newValues" json:"newValues,omitempty"`
	Ops       []*Op          `protobuf:"bytes,3,rep,name=ops" json:"ops,omitempty"`
}

func (m *EditProto) Reset()                    { *m = EditProto{} }
func (m *EditProto) String() string            { return proto.CompactTextString(m) }
func (*EditProto) ProtoMessage()               {}
func (*EditProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *EditProto) GetReference() *oipProto.Txid {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *EditProto) GetNewValues() *RecordProto {
	if m != nil {
		return m.NewValues
	}
	return nil
}

func (m *EditProto) GetOps() []*Op {
	if m != nil {
		return m.Ops
	}
	return nil
}

type Op struct {
	Path []*Step `protobuf:"bytes,1,rep,name=Path" json:"Path,omitempty"`
}

func (m *Op) Reset()                    { *m = Op{} }
func (m *Op) String() string            { return proto.CompactTextString(m) }
func (*Op) ProtoMessage()               {}
func (*Op) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *Op) GetPath() []*Step {
	if m != nil {
		return m.Path
	}
	return nil
}

type Step struct {
	Tag      int32       `protobuf:"varint,1,opt,name=tag" json:"tag,omitempty"`
	Action   Step_Action `protobuf:"varint,2,opt,name=action,enum=oipProto.Step_Action" json:"action,omitempty"`
	SrcIndex int32       `protobuf:"varint,3,opt,name=srcIndex" json:"srcIndex,omitempty"`
	DstIndex int32       `protobuf:"varint,4,opt,name=dstIndex" json:"dstIndex,omitempty"`
}

func (m *Step) Reset()                    { *m = Step{} }
func (m *Step) String() string            { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()               {}
func (*Step) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *Step) GetTag() int32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Step) GetAction() Step_Action {
	if m != nil {
		return m.Action
	}
	return Step_ACTION_ERROR
}

func (m *Step) GetSrcIndex() int32 {
	if m != nil {
		return m.SrcIndex
	}
	return 0
}

func (m *Step) GetDstIndex() int32 {
	if m != nil {
		return m.DstIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*EditProto)(nil), "oipProto.EditProto")
	proto.RegisterType((*Op)(nil), "oipProto.Op")
	proto.RegisterType((*Step)(nil), "oipProto.Step")
	proto.RegisterEnum("oipProto.Step_Action", Step_Action_name, Step_Action_value)
}

func init() { proto.RegisterFile("edit.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x4e, 0xab, 0x40,
	0x14, 0xc6, 0x2f, 0x7f, 0xca, 0x6d, 0x4f, 0x9b, 0x66, 0xee, 0x5c, 0xab, 0xa4, 0x0b, 0xd3, 0xb0,
	0xea, 0x42, 0x59, 0xb4, 0xf1, 0x01, 0xb0, 0x4e, 0x94, 0xa4, 0x02, 0x99, 0x92, 0x2e, 0xdc, 0x10,
	0x84, 0x51, 0x49, 0x0c, 0x10, 0x18, 0x63, 0x9f, 0xc2, 0x47, 0x32, 0x3e, 0x9a, 0x99, 0x81, 0x96,
	0x26, 0xba, 0x3b, 0xe7, 0xf7, 0xfd, 0x38, 0x7c, 0x8b, 0x01, 0x60, 0x69, 0xc6, 0xed, 0xb2, 0x2a,
	0x78, 0x81, 0xfb, 0x45, 0x56, 0x06, 0x62, 0x9a, 0x8e, 0x28, 0x4b, 0x8a, 0x2a, 0x6d, 0xf8, 0x14,
	0xf8, 0x2e, 0x6b, 0x67, 0xeb, 0x43, 0x81, 0x01, 0x49, 0x33, 0x2e, 0x3d, 0x7c, 0x01, 0x83, 0x8a,
	0x3d, 0xb1, 0x8a, 0xe5, 0x09, 0x33, 0x95, 0x99, 0x32, 0x1f, 0x2e, 0xc6, 0xf6, 0xfe, 0x8a, 0x1d,
	0xee, 0xb2, 0x94, 0x76, 0x02, 0x5e, 0xc2, 0x20, 0x67, 0xef, 0xdb, 0xf8, 0xf5, 0x8d, 0xd5, 0xa6,
	0x2a, 0xed, 0x49, 0x67, 0x37, 0xbf, 0x94, 0x33, 0xed, 0x3c, 0x7c, 0x0e, 0x5a, 0x51, 0xd6, 0xa6,
	0x36, 0xd3, 0xe6, 0xc3, 0xc5, 0xa8, 0xd3, 0xfd, 0x92, 0x8a, 0xc0, 0x9a, 0x83, 0xea, 0x97, 0xd8,
	0x02, 0x3d, 0x88, 0xf9, 0x8b, 0xa9, 0x48, 0xed, 0xa8, 0xc3, 0x86, 0xb3, 0x92, 0xca, 0xcc, 0xfa,
	0x52, 0x41, 0x17, 0x2b, 0x46, 0xa0, 0xf1, 0xf8, 0x59, 0xf6, 0xed, 0x51, 0x31, 0xe2, 0x4b, 0x30,
	0xe2, 0x84, 0x67, 0x45, 0x2e, 0x6b, 0x8d, 0x8f, 0x6b, 0x89, 0x2f, 0x6c, 0x47, 0x86, 0xb4, 0x95,
	0xf0, 0x14, 0xfa, 0x75, 0x95, 0xb8, 0x79, 0xca, 0x76, 0xa6, 0x26, 0xaf, 0x1c, 0x76, 0x91, 0xa5,
	0x35, 0x6f, 0x32, 0xbd, 0xc9, 0xf6, 0xbb, 0xf5, 0xa9, 0x80, 0xd1, 0x9c, 0xc2, 0x08, 0x46, 0xce,
	0x2a, 0x74, 0x7d, 0x2f, 0x22, 0x94, 0xfa, 0x14, 0xfd, 0xc1, 0xa7, 0x80, 0x5b, 0x42, 0x49, 0xb0,
	0x76, 0x56, 0x24, 0x72, 0xd6, 0x6b, 0xa4, 0xe0, 0x09, 0xfc, 0x6b, 0xb9, 0x13, 0x04, 0xc4, 0xbb,
	0x91, 0x58, 0x3d, 0xc2, 0x94, 0xdc, 0xfb, 0xdb, 0xc6, 0xd6, 0x7e, 0x62, 0xdf, 0x23, 0x48, 0xff,
	0xe5, 0xb8, 0xe0, 0x3d, 0x7c, 0x06, 0xff, 0x5b, 0xbe, 0x09, 0xa9, 0xeb, 0xdd, 0x46, 0x81, 0x13,
	0xae, 0xee, 0x90, 0x81, 0x4f, 0x00, 0x1d, 0x02, 0x12, 0x44, 0xae, 0x17, 0xfa, 0xe8, 0xef, 0xb5,
	0xf1, 0xa0, 0x17, 0x59, 0x79, 0xf5, 0x68, 0xc8, 0xc7, 0xb0, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xf8, 0x25, 0x7e, 0x63, 0x3e, 0x02, 0x00, 0x00,
}