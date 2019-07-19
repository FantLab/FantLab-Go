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
	Error_UNKNOWN_STATUS       Error_Status = 0
	Error_SOMETHING_WENT_WRONG Error_Status = 1
	Error_INVALID_PARAMETER    Error_Status = 2
	Error_NOT_FOUND            Error_Status = 3
	Error_INVALID_PASSWORD     Error_Status = 4
	Error_LOG_OUT_FIRST        Error_Status = 5
	Error_INVALID_SESSION      Error_Status = 6
)

var Error_Status_name = map[int32]string{
	0: "UNKNOWN_STATUS",
	1: "SOMETHING_WENT_WRONG",
	2: "INVALID_PARAMETER",
	3: "NOT_FOUND",
	4: "INVALID_PASSWORD",
	5: "LOG_OUT_FIRST",
	6: "INVALID_SESSION",
}

var Error_Status_value = map[string]int32{
	"UNKNOWN_STATUS":       0,
	"SOMETHING_WENT_WRONG": 1,
	"INVALID_PARAMETER":    2,
	"NOT_FOUND":            3,
	"INVALID_PASSWORD":     4,
	"LOG_OUT_FIRST":        5,
	"INVALID_SESSION":      6,
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
	return Error_UNKNOWN_STATUS
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
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd0, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x61, 0xa7, 0xda, 0xb4, 0xbd, 0xd0, 0x3a, 0xbd, 0xb6, 0x10, 0x5c, 0x95, 0xae, 0xba,
	0x31, 0x05, 0x7d, 0x82, 0x48, 0xa6, 0x31, 0xd8, 0xce, 0xc8, 0xdc, 0x89, 0x01, 0x37, 0x43, 0x22,
	0x71, 0x25, 0x49, 0x48, 0x46, 0xf0, 0x59, 0x7c, 0x53, 0x77, 0x62, 0x6c, 0x71, 0x79, 0xf8, 0xbf,
	0xd5, 0x01, 0x6c, 0xda, 0xda, 0xd5, 0xdb, 0xb2, 0x6d, 0xeb, 0xb6, 0x0b, 0xfa, 0x81, 0xa3, 0xb7,
	0xbc, 0x72, 0xef, 0x79, 0xb1, 0xfe, 0x66, 0x30, 0x14, 0xbf, 0xe5, 0x9a, 0x60, 0xac, 0xcb, 0xae,
	0xa9, 0xab, 0xae, 0xc4, 0x1b, 0xf0, 0x3a, 0x97, 0xbb, 0x8f, 0xce, 0x67, 0x2b, 0xb6, 0x99, 0xdd,
	0x2e, 0x83, 0xa3, 0x0f, 0x7a, 0x1b, 0x50, 0x1f, 0xf5, 0x11, 0xa1, 0x0f, 0xa3, 0xd7, 0xba, 0x72,
	0xe5, 0xa7, 0xf3, 0x07, 0x2b, 0xb6, 0x99, 0xe8, 0xd3, 0x5c, 0x7f, 0x31, 0xf0, 0xfe, 0x30, 0x22,
	0xcc, 0x52, 0xf9, 0x28, 0x55, 0x26, 0x2d, 0x99, 0xd0, 0xa4, 0xc4, 0xcf, 0xd0, 0x87, 0x05, 0xa9,
	0x83, 0x30, 0x0f, 0x89, 0x8c, 0x6d, 0x26, 0xa4, 0xb1, 0x99, 0x56, 0x32, 0xe6, 0x0c, 0x97, 0x30,
	0x4f, 0xe4, 0x73, 0xb8, 0x4f, 0x22, 0xfb, 0x14, 0xea, 0xf0, 0x20, 0x8c, 0xd0, 0x7c, 0x80, 0x53,
	0x98, 0x48, 0x65, 0xec, 0x4e, 0xa5, 0x32, 0xe2, 0xe7, 0xb8, 0x00, 0xfe, 0xaf, 0x88, 0x32, 0xa5,
	0x23, 0x7e, 0x81, 0x73, 0x98, 0xee, 0x55, 0x6c, 0x55, 0x6a, 0xec, 0x2e, 0xd1, 0x64, 0xf8, 0x10,
	0xaf, 0xe0, 0xf2, 0x04, 0x49, 0x10, 0x25, 0x4a, 0x72, 0xef, 0x7e, 0xfc, 0xe2, 0xe5, 0x4d, 0xb3,
	0x6d, 0x8a, 0xc2, 0xeb, 0x5f, 0xb9, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xe0, 0xf6, 0x2d,
	0x2b, 0x01, 0x00, 0x00,
}
