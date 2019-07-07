// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common_models.proto

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
	Common_UNKNOWN Common_Gender = 0
	Common_MALE    Common_Gender = 1
	Common_FEMALE  Common_Gender = 2
)

var Common_Gender_name = map[int32]string{
	0: "UNKNOWN",
	1: "MALE",
	2: "FEMALE",
}

var Common_Gender_value = map[string]int32{
	"UNKNOWN": 0,
	"MALE":    1,
	"FEMALE":  2,
}

func (x Common_Gender) String() string {
	return proto.EnumName(Common_Gender_name, int32(x))
}

func (Common_Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6151d991ff354fae, []int{0, 0}
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
	return fileDescriptor_6151d991ff354fae, []int{0}
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
	IsDeleted            bool          `protobuf:"varint,8,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Common_UserLink) Reset()         { *m = Common_UserLink{} }
func (m *Common_UserLink) String() string { return proto.CompactTextString(m) }
func (*Common_UserLink) ProtoMessage()    {}
func (*Common_UserLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_6151d991ff354fae, []int{0, 0}
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
	return Common_UNKNOWN
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

func (m *Common_UserLink) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
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
	return fileDescriptor_6151d991ff354fae, []int{0, 1}
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

type Common_ErrorResponse struct {
	ErrorCode            int32    `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Common_ErrorResponse) Reset()         { *m = Common_ErrorResponse{} }
func (m *Common_ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*Common_ErrorResponse) ProtoMessage()    {}
func (*Common_ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6151d991ff354fae, []int{0, 2}
}

func (m *Common_ErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Common_ErrorResponse.Unmarshal(m, b)
}
func (m *Common_ErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Common_ErrorResponse.Marshal(b, m, deterministic)
}
func (m *Common_ErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Common_ErrorResponse.Merge(m, src)
}
func (m *Common_ErrorResponse) XXX_Size() int {
	return xxx_messageInfo_Common_ErrorResponse.Size(m)
}
func (m *Common_ErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Common_ErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Common_ErrorResponse proto.InternalMessageInfo

func (m *Common_ErrorResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *Common_ErrorResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("fantlab.Common_Gender", Common_Gender_name, Common_Gender_value)
	proto.RegisterType((*Common)(nil), "fantlab.Common")
	proto.RegisterType((*Common_UserLink)(nil), "fantlab.Common.UserLink")
	proto.RegisterType((*Common_Creation)(nil), "fantlab.Common.Creation")
	proto.RegisterType((*Common_ErrorResponse)(nil), "fantlab.Common.ErrorResponse")
}

func init() { proto.RegisterFile("proto/common_models.proto", fileDescriptor_6151d991ff354fae) }

var fileDescriptor_6151d991ff354fae = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0xdd, 0x8a, 0x13, 0x31,
	0x14, 0x76, 0xba, 0xd3, 0x74, 0x7a, 0xca, 0x2e, 0x4b, 0x90, 0x25, 0x0e, 0x88, 0x65, 0xaf, 0x0a,
	0x4a, 0x0a, 0xf5, 0x09, 0xb4, 0x56, 0x05, 0xd7, 0x0a, 0xc1, 0x45, 0xf0, 0xa6, 0xa4, 0xcd, 0xd9,
	0x31, 0x38, 0x93, 0x0c, 0x49, 0xd6, 0xf7, 0xf3, 0x21, 0x7c, 0x1f, 0x99, 0x93, 0xce, 0xcd, 0xde,
	0x9d, 0xef, 0xcb, 0x47, 0xbe, 0x1f, 0x78, 0xd1, 0x07, 0x9f, 0xfc, 0xfa, 0xe4, 0xbb, 0xce, 0xbb,
	0x43, 0xe7, 0x0d, 0xb6, 0x51, 0x12, 0xc7, 0x67, 0x0f, 0xda, 0xa5, 0x56, 0x1f, 0xeb, 0x57, 0x8d,
	0xf7, 0x4d, 0x8b, 0x6b, 0xa2, 0x8f, 0x8f, 0x0f, 0xeb, 0x64, 0x3b, 0x8c, 0x49, 0x77, 0x7d, 0x56,
	0xde, 0xfe, 0xbd, 0x00, 0xb6, 0xa5, 0x1f, 0xea, 0x7f, 0x05, 0x54, 0xf7, 0x11, 0xc3, 0x9d, 0x75,
	0xbf, 0xf9, 0x15, 0x4c, 0xac, 0x11, 0xc5, 0xb2, 0x58, 0x5d, 0xaa, 0x89, 0x35, 0xfc, 0x39, 0x4c,
	0x5b, 0xdf, 0x58, 0x27, 0x26, 0xcb, 0x62, 0x35, 0x57, 0x19, 0x70, 0x0e, 0xa5, 0xd3, 0x1d, 0x8a,
	0x0b, 0x22, 0xe9, 0xe6, 0x12, 0x58, 0x83, 0xce, 0x60, 0x10, 0xe5, 0xb2, 0x58, 0x5d, 0x6d, 0x6e,
	0xe4, 0x39, 0x8c, 0xcc, 0x3e, 0xf2, 0x13, 0xbd, 0xaa, 0xb3, 0x8a, 0xdf, 0x00, 0xd3, 0x7f, 0x74,
	0xd2, 0x41, 0x4c, 0xe9, 0x97, 0x33, 0x1a, 0x1c, 0x4f, 0xad, 0x8e, 0x51, 0x30, 0x0a, 0x91, 0xc1,
	0xe0, 0x18, 0x6d, 0xe3, 0xc4, 0x2c, 0x3b, 0x0e, 0x37, 0x7f, 0x09, 0x60, 0xe3, 0xc1, 0x60, 0x8b,
	0x09, 0x8d, 0xa8, 0x96, 0xc5, 0xaa, 0x52, 0x73, 0x1b, 0x3f, 0x64, 0xa2, 0xfe, 0x05, 0xd5, 0x36,
	0xa0, 0x4e, 0xd6, 0x3b, 0xfe, 0x06, 0xca, 0xc7, 0x88, 0x81, 0x8a, 0x2d, 0x36, 0xe2, 0x69, 0xb4,
	0xb1, 0xbe, 0x22, 0x15, 0x97, 0x50, 0x1a, 0x9d, 0x90, 0x3a, 0x2f, 0x36, 0xb5, 0xcc, 0x63, 0xca,
	0x71, 0x4c, 0xf9, 0x7d, 0x1c, 0x53, 0x91, 0xae, 0xfe, 0x0c, 0x97, 0xbb, 0x10, 0x7c, 0x50, 0x18,
	0x7b, 0xef, 0x22, 0x0e, 0xc9, 0x70, 0x20, 0x0e, 0x27, 0x6f, 0x90, 0x4c, 0xa7, 0x6a, 0x4e, 0xcc,
	0xd6, 0x1b, 0xe4, 0x02, 0x66, 0x1d, 0xc6, 0xa8, 0x1b, 0x3c, 0xcf, 0x3a, 0xc2, 0xdb, 0xd7, 0xc0,
	0xf2, 0x4c, 0x7c, 0x01, 0xb3, 0xfb, 0xfd, 0x97, 0xfd, 0xb7, 0x1f, 0xfb, 0xeb, 0x67, 0xbc, 0x82,
	0xf2, 0xeb, 0xbb, 0xbb, 0xdd, 0x75, 0xc1, 0x01, 0xd8, 0xc7, 0x1d, 0xdd, 0x93, 0xf7, 0xd5, 0x4f,
	0xa6, 0xfb, 0x7e, 0xdd, 0x1f, 0x8f, 0x8c, 0xa2, 0xbd, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x8f,
	0x80, 0xc6, 0x2e, 0x1b, 0x02, 0x00, 0x00,
}
