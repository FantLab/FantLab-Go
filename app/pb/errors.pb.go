// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/errors.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Error_Status int32

const (
	Error_UNKNOWN              Error_Status = 0
	Error_SOMETHING_WENT_WRONG Error_Status = 1
	Error_INVALID_PARAMETER    Error_Status = 2
	Error_NOT_FOUND            Error_Status = 3
	Error_INVALID_PASSWORD     Error_Status = 4
	Error_LOG_OUT_FIRST        Error_Status = 5
)

var Error_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SOMETHING_WENT_WRONG",
	2: "INVALID_PARAMETER",
	3: "NOT_FOUND",
	4: "INVALID_PASSWORD",
	5: "LOG_OUT_FIRST",
}

var Error_Status_value = map[string]int32{
	"UNKNOWN":              0,
	"SOMETHING_WENT_WRONG": 1,
	"INVALID_PARAMETER":    2,
	"NOT_FOUND":            3,
	"INVALID_PASSWORD":     4,
	"LOG_OUT_FIRST":        5,
}

func (x Error_Status) String() string {
	return proto.EnumName(Error_Status_name, int32(x))
}

func (Error_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_176b67d994589ed1, []int{0, 0}
}

type Error struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_176b67d994589ed1, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

type Error_Response struct {
	Status               Error_Status `protobuf:"varint,1,opt,name=status,proto3,enum=fantlab.Error_Status" json:"status,omitempty"`
	Context              string       `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Error_Response) Reset()         { *m = Error_Response{} }
func (m *Error_Response) String() string { return proto.CompactTextString(m) }
func (*Error_Response) ProtoMessage()    {}
func (*Error_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_176b67d994589ed1, []int{0, 0}
}

func (m *Error_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error_Response.Unmarshal(m, b)
}
func (m *Error_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error_Response.Marshal(b, m, deterministic)
}
func (m *Error_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error_Response.Merge(m, src)
}
func (m *Error_Response) XXX_Size() int {
	return xxx_messageInfo_Error_Response.Size(m)
}
func (m *Error_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Error_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Error_Response proto.InternalMessageInfo

func (m *Error_Response) GetStatus() Error_Status {
	if m != nil {
		return m.Status
	}
	return Error_UNKNOWN
}

func (m *Error_Response) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func init() {
	proto.RegisterEnum("fantlab.Error_Status", Error_Status_name, Error_Status_value)
	proto.RegisterType((*Error)(nil), "fantlab.Error")
	proto.RegisterType((*Error_Response)(nil), "fantlab.Error.Response")
}

func init() { proto.RegisterFile("proto/errors.proto", fileDescriptor_176b67d994589ed1) }

var fileDescriptor_176b67d994589ed1 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xcf, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x71, 0x3b, 0x5d, 0xbb, 0xbd, 0x32, 0xc9, 0x5e, 0x36, 0x28, 0x9e, 0xc6, 0x4e, 0xbb,
	0xd8, 0x81, 0x7e, 0x82, 0x4a, 0xb3, 0x5a, 0xdc, 0x12, 0x49, 0x52, 0x0b, 0x5e, 0x42, 0x2b, 0xf5,
	0x24, 0x4d, 0x68, 0x22, 0x78, 0xf2, 0x93, 0xfa, 0x61, 0xc4, 0x3a, 0xd9, 0xf1, 0xe1, 0xff, 0xbb,
	0x3c, 0x80, 0xb6, 0x37, 0xde, 0x6c, 0xdb, 0xbe, 0x37, 0xbd, 0x4b, 0x86, 0x81, 0xd1, 0x5b, 0xdd,
	0xf9, 0xf7, 0xba, 0x59, 0x7f, 0x07, 0x30, 0xa6, 0xbf, 0xe5, 0x5a, 0xc2, 0x44, 0xb4, 0xce, 0x9a,
	0xce, 0xb5, 0x78, 0x03, 0xa1, 0xf3, 0xb5, 0xff, 0x70, 0x71, 0xb0, 0x0a, 0x36, 0x57, 0xb7, 0xcb,
	0xe4, 0xe8, 0x93, 0xc1, 0x26, 0x72, 0x88, 0xe2, 0x88, 0x30, 0x86, 0xe8, 0xd5, 0x74, 0xbe, 0xfd,
	0xf4, 0xf1, 0x68, 0x15, 0x6c, 0xa6, 0xe2, 0x7f, 0xae, 0xbf, 0x20, 0xfc, 0xb3, 0x78, 0x09, 0x51,
	0xc9, 0x1e, 0x19, 0xaf, 0x18, 0x39, 0xc3, 0x18, 0x16, 0x92, 0x1f, 0xa8, 0x7a, 0x28, 0x58, 0xae,
	0x2b, 0xca, 0x94, 0xae, 0x04, 0x67, 0x39, 0x09, 0x70, 0x09, 0xf3, 0x82, 0x3d, 0xa7, 0xfb, 0x22,
	0xd3, 0x4f, 0xa9, 0x48, 0x0f, 0x54, 0x51, 0x41, 0x46, 0x38, 0x83, 0x29, 0xe3, 0x4a, 0xef, 0x78,
	0xc9, 0x32, 0x72, 0x8e, 0x0b, 0x20, 0x27, 0x25, 0x65, 0xc5, 0x45, 0x46, 0x2e, 0x70, 0x0e, 0xb3,
	0x3d, 0xcf, 0x35, 0x2f, 0x95, 0xde, 0x15, 0x42, 0x2a, 0x32, 0xbe, 0x9f, 0xbc, 0x84, 0xb5, 0xb5,
	0x5b, 0xdb, 0x34, 0xe1, 0x70, 0xfc, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xc4, 0xc5, 0x9e,
	0x0e, 0x01, 0x00, 0x00,
}
