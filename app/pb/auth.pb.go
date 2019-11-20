// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth.proto

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

type Auth struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d38a2cdbb4f144, []int{0}
}

func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (m *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(m, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

type Auth_LoginResponse struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionToken         string   `protobuf:"bytes,2,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth_LoginResponse) Reset()         { *m = Auth_LoginResponse{} }
func (m *Auth_LoginResponse) String() string { return proto.CompactTextString(m) }
func (*Auth_LoginResponse) ProtoMessage()    {}
func (*Auth_LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d38a2cdbb4f144, []int{0, 0}
}

func (m *Auth_LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth_LoginResponse.Unmarshal(m, b)
}
func (m *Auth_LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth_LoginResponse.Marshal(b, m, deterministic)
}
func (m *Auth_LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth_LoginResponse.Merge(m, src)
}
func (m *Auth_LoginResponse) XXX_Size() int {
	return xxx_messageInfo_Auth_LoginResponse.Size(m)
}
func (m *Auth_LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Auth_LoginResponse proto.InternalMessageInfo

func (m *Auth_LoginResponse) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Auth_LoginResponse) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func init() {
	proto.RegisterType((*Auth)(nil), "Auth")
	proto.RegisterType((*Auth_LoginResponse)(nil), "Auth.LoginResponse")
}

func init() { proto.RegisterFile("proto/auth.proto", fileDescriptor_a9d38a2cdbb4f144) }

var fileDescriptor_a9d38a2cdbb4f144 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0x03, 0x33, 0x95, 0x42, 0xb9, 0x58, 0x1c, 0x4b, 0x4b,
	0x32, 0xa4, 0x7c, 0xb9, 0x78, 0x7d, 0xf2, 0xd3, 0x33, 0xf3, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x85, 0xc4, 0xb9, 0xd8, 0x4b, 0x8b, 0x53, 0x8b, 0xe2, 0x33, 0x53, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x58, 0x82, 0xd8, 0x40, 0x5c, 0xcf, 0x14, 0x21, 0x65, 0x2e, 0xde, 0xe2, 0xd4, 0xe2,
	0xe2, 0xcc, 0xfc, 0xbc, 0xf8, 0x92, 0xfc, 0xec, 0xd4, 0x3c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x1e, 0xa8, 0x60, 0x08, 0x48, 0xcc, 0x89, 0x23, 0x8a, 0x2d, 0xb1, 0xa0, 0x40, 0xbf, 0x20,
	0x29, 0x89, 0x0d, 0x6c, 0x8f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xc9, 0x2e, 0x98, 0x7b,
	0x00, 0x00, 0x00,
}
