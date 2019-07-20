// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Common_Gender int32

const (
	Common_UNKNOWN_GENDER Common_Gender = 0
	Common_MALE           Common_Gender = 1
	Common_FEMALE         Common_Gender = 2
)

var Common_Gender_name = map[int32]string{
	0: "UNKNOWN_GENDER",
	1: "MALE",
	2: "FEMALE",
}

var Common_Gender_value = map[string]int32{
	"UNKNOWN_GENDER": 0,
	"MALE":           1,
	"FEMALE":         2,
}

func (x Common_Gender) String() string {
	return proto.EnumName(Common_Gender_name, int32(x))
}

func (Common_Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0, 0}
}

type Common struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Common) Reset()         { *m = Common{} }
func (m *Common) String() string { return proto.CompactTextString(m) }
func (*Common) ProtoMessage()    {}
func (*Common) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

func (m *Common) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Common.Unmarshal(m, b)
}
func (m *Common) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Common.Marshal(b, m, deterministic)
}
func (m *Common) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Common.Merge(m, src)
}
func (m *Common) XXX_Size() int {
	return xxx_messageInfo_Common.Size(m)
}
func (m *Common) XXX_DiscardUnknown() {
	xxx_messageInfo_Common.DiscardUnknown(m)
}

var xxx_messageInfo_Common proto.InternalMessageInfo

type Common_UserLink struct {
	Id                   uint32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Login                string        `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Name                 string        `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Gender               Common_Gender `protobuf:"varint,4,opt,name=gender,proto3,enum=fantlab.Common_Gender" json:"gender,omitempty"`
	Avatar               string        `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Class                uint32        `protobuf:"varint,6,opt,name=class,proto3" json:"class,omitempty"`
	Sign                 string        `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Common_UserLink) Reset()         { *m = Common_UserLink{} }
func (m *Common_UserLink) String() string { return proto.CompactTextString(m) }
func (*Common_UserLink) ProtoMessage()    {}
func (*Common_UserLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0, 0}
}

func (m *Common_UserLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Common_UserLink.Unmarshal(m, b)
}
func (m *Common_UserLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Common_UserLink.Marshal(b, m, deterministic)
}
func (m *Common_UserLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Common_UserLink.Merge(m, src)
}
func (m *Common_UserLink) XXX_Size() int {
	return xxx_messageInfo_Common_UserLink.Size(m)
}
func (m *Common_UserLink) XXX_DiscardUnknown() {
	xxx_messageInfo_Common_UserLink.DiscardUnknown(m)
}

var xxx_messageInfo_Common_UserLink proto.InternalMessageInfo

func (m *Common_UserLink) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Common_UserLink) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *Common_UserLink) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Common_UserLink) GetGender() Common_Gender {
	if m != nil {
		return m.Gender
	}
	return Common_UNKNOWN_GENDER
}

func (m *Common_UserLink) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Common_UserLink) GetClass() uint32 {
	if m != nil {
		return m.Class
	}
	return 0
}

func (m *Common_UserLink) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

type Common_Creation struct {
	User                 *Common_UserLink     `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Common_Creation) Reset()         { *m = Common_Creation{} }
func (m *Common_Creation) String() string { return proto.CompactTextString(m) }
func (*Common_Creation) ProtoMessage()    {}
func (*Common_Creation) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0, 1}
}

func (m *Common_Creation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Common_Creation.Unmarshal(m, b)
}
func (m *Common_Creation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Common_Creation.Marshal(b, m, deterministic)
}
func (m *Common_Creation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Common_Creation.Merge(m, src)
}
func (m *Common_Creation) XXX_Size() int {
	return xxx_messageInfo_Common_Creation.Size(m)
}
func (m *Common_Creation) XXX_DiscardUnknown() {
	xxx_messageInfo_Common_Creation.DiscardUnknown(m)
}

var xxx_messageInfo_Common_Creation proto.InternalMessageInfo

func (m *Common_Creation) GetUser() *Common_UserLink {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Common_Creation) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type Common_Pages struct {
	Current              uint32   `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	Count                uint32   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Common_Pages) Reset()         { *m = Common_Pages{} }
func (m *Common_Pages) String() string { return proto.CompactTextString(m) }
func (*Common_Pages) ProtoMessage()    {}
func (*Common_Pages) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0, 2}
}

func (m *Common_Pages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Common_Pages.Unmarshal(m, b)
}
func (m *Common_Pages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Common_Pages.Marshal(b, m, deterministic)
}
func (m *Common_Pages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Common_Pages.Merge(m, src)
}
func (m *Common_Pages) XXX_Size() int {
	return xxx_messageInfo_Common_Pages.Size(m)
}
func (m *Common_Pages) XXX_DiscardUnknown() {
	xxx_messageInfo_Common_Pages.DiscardUnknown(m)
}

var xxx_messageInfo_Common_Pages proto.InternalMessageInfo

func (m *Common_Pages) GetCurrent() uint32 {
	if m != nil {
		return m.Current
	}
	return 0
}

func (m *Common_Pages) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Common_SuccessResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Common_SuccessResponse) Reset()         { *m = Common_SuccessResponse{} }
func (m *Common_SuccessResponse) String() string { return proto.CompactTextString(m) }
func (*Common_SuccessResponse) ProtoMessage()    {}
func (*Common_SuccessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0, 3}
}

func (m *Common_SuccessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Common_SuccessResponse.Unmarshal(m, b)
}
func (m *Common_SuccessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Common_SuccessResponse.Marshal(b, m, deterministic)
}
func (m *Common_SuccessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Common_SuccessResponse.Merge(m, src)
}
func (m *Common_SuccessResponse) XXX_Size() int {
	return xxx_messageInfo_Common_SuccessResponse.Size(m)
}
func (m *Common_SuccessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Common_SuccessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Common_SuccessResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("fantlab.Common_Gender", Common_Gender_name, Common_Gender_value)
	proto.RegisterType((*Common)(nil), "fantlab.Common")
	proto.RegisterType((*Common_UserLink)(nil), "fantlab.Common.UserLink")
	proto.RegisterType((*Common_Creation)(nil), "fantlab.Common.Creation")
	proto.RegisterType((*Common_Pages)(nil), "fantlab.Common.Pages")
	proto.RegisterType((*Common_SuccessResponse)(nil), "fantlab.Common.SuccessResponse")
}

func init() { proto.RegisterFile("proto/common.proto", fileDescriptor_1747d3070a2311a0) }

var fileDescriptor_1747d3070a2311a0 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0x41, 0xcf, 0x93, 0x40,
	0x14, 0x14, 0x3e, 0xba, 0xc5, 0xd7, 0x58, 0xeb, 0xc6, 0x34, 0x1b, 0x2e, 0x36, 0x3d, 0xf5, 0x60,
	0x96, 0x04, 0x0f, 0x9e, 0xb5, 0x62, 0x0f, 0x56, 0x34, 0x68, 0x63, 0xe2, 0xc5, 0x2c, 0xb0, 0x45,
	0x22, 0xec, 0x92, 0xdd, 0xc5, 0x3f, 0xe5, 0xdd, 0xdf, 0x67, 0x78, 0xc0, 0xc5, 0xdb, 0xcc, 0xdb,
	0xd9, 0x37, 0xf3, 0x06, 0x68, 0x6f, 0xb4, 0xd3, 0x71, 0xa9, 0xbb, 0x4e, 0x2b, 0x8e, 0x84, 0xae,
	0xef, 0x42, 0xb9, 0x56, 0x14, 0xd1, 0x8b, 0x5a, 0xeb, 0xba, 0x95, 0x31, 0x8e, 0x8b, 0xe1, 0x1e,
	0xbb, 0xa6, 0x93, 0xd6, 0x89, 0xae, 0x9f, 0x94, 0xc7, 0x3f, 0x0f, 0x40, 0xce, 0xf8, 0x35, 0xfa,
	0xeb, 0x41, 0x78, 0xb3, 0xd2, 0x5c, 0x1b, 0xf5, 0x8b, 0x6e, 0xc1, 0x6f, 0x2a, 0xe6, 0x1d, 0xbc,
	0xd3, 0x93, 0xdc, 0x6f, 0x2a, 0xfa, 0x1c, 0x56, 0xad, 0xae, 0x1b, 0xc5, 0xfc, 0x83, 0x77, 0x7a,
	0x9c, 0x4f, 0x84, 0x52, 0x08, 0x94, 0xe8, 0x24, 0x7b, 0xc0, 0x21, 0x62, 0xca, 0x81, 0xd4, 0x52,
	0x55, 0xd2, 0xb0, 0xe0, 0xe0, 0x9d, 0xb6, 0xc9, 0x9e, 0xcf, 0x61, 0xf8, 0xe4, 0xc3, 0x2f, 0xf8,
	0x9a, 0xcf, 0x2a, 0xba, 0x07, 0x22, 0x7e, 0x0b, 0x27, 0x0c, 0x5b, 0xe1, 0x96, 0x99, 0x8d, 0x8e,
	0x65, 0x2b, 0xac, 0x65, 0x04, 0x43, 0x4c, 0x64, 0x74, 0xb4, 0x4d, 0xad, 0xd8, 0x7a, 0x72, 0x1c,
	0x71, 0xf4, 0x13, 0xc2, 0xb3, 0x91, 0xc2, 0x35, 0x5a, 0xd1, 0x97, 0x10, 0x0c, 0x56, 0x1a, 0x4c,
	0xbe, 0x49, 0xd8, 0xff, 0xde, 0xcb, 0x7d, 0x39, 0xaa, 0x28, 0x87, 0xa0, 0x12, 0x4e, 0xe2, 0x51,
	0x9b, 0x24, 0xe2, 0x53, 0x5b, 0x7c, 0x69, 0x8b, 0x7f, 0x5d, 0xda, 0xca, 0x51, 0x17, 0xbd, 0x86,
	0xd5, 0x67, 0x51, 0x4b, 0x4b, 0x19, 0xac, 0xcb, 0xc1, 0x18, 0xa9, 0xdc, 0xdc, 0xd1, 0x42, 0x31,
	0xb6, 0x1e, 0x94, 0xc3, 0x9d, 0x63, 0xec, 0x91, 0x44, 0xcf, 0xe0, 0xe9, 0x97, 0xa1, 0x2c, 0xa5,
	0xb5, 0xb9, 0xb4, 0xbd, 0x56, 0x56, 0x1e, 0x13, 0x20, 0x53, 0x13, 0x94, 0xc2, 0xf6, 0x96, 0x7d,
	0xc8, 0x3e, 0x7d, 0xcb, 0x7e, 0x5c, 0xd2, 0xec, 0x5d, 0x9a, 0xef, 0x1e, 0xd1, 0x10, 0x82, 0x8f,
	0x6f, 0xae, 0xe9, 0xce, 0xa3, 0x00, 0xe4, 0x7d, 0x8a, 0xd8, 0x7f, 0x1b, 0x7e, 0x27, 0xa2, 0xef,
	0xe3, 0xbe, 0x28, 0x08, 0x66, 0x7c, 0xf5, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x4d, 0x4e, 0x2a,
	0xfe, 0x01, 0x00, 0x00,
}
