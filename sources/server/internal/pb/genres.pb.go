// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/genres.proto

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

type Genre struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Genre) Reset()         { *m = Genre{} }
func (m *Genre) String() string { return proto.CompactTextString(m) }
func (*Genre) ProtoMessage()    {}
func (*Genre) Descriptor() ([]byte, []int) {
	return fileDescriptor_d42afaacf88a7096, []int{0}
}

func (m *Genre) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genre.Unmarshal(m, b)
}
func (m *Genre) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genre.Marshal(b, m, deterministic)
}
func (m *Genre) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genre.Merge(m, src)
}
func (m *Genre) XXX_Size() int {
	return xxx_messageInfo_Genre.Size(m)
}
func (m *Genre) XXX_DiscardUnknown() {
	xxx_messageInfo_Genre.DiscardUnknown(m)
}

var xxx_messageInfo_Genre proto.InternalMessageInfo

type Genre_Genre struct {
	// id жанра
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// название
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// информация
	Info string `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	// коичество произведений
	WorkCount uint64 `protobuf:"varint,4,opt,name=work_count,json=workCount,proto3" json:"work_count,omitempty"`
	// поджанры
	Subgenres            []*Genre_Genre `protobuf:"bytes,5,rep,name=subgenres,proto3" json:"subgenres,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Genre_Genre) Reset()         { *m = Genre_Genre{} }
func (m *Genre_Genre) String() string { return proto.CompactTextString(m) }
func (*Genre_Genre) ProtoMessage()    {}
func (*Genre_Genre) Descriptor() ([]byte, []int) {
	return fileDescriptor_d42afaacf88a7096, []int{0, 0}
}

func (m *Genre_Genre) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genre_Genre.Unmarshal(m, b)
}
func (m *Genre_Genre) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genre_Genre.Marshal(b, m, deterministic)
}
func (m *Genre_Genre) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genre_Genre.Merge(m, src)
}
func (m *Genre_Genre) XXX_Size() int {
	return xxx_messageInfo_Genre_Genre.Size(m)
}
func (m *Genre_Genre) XXX_DiscardUnknown() {
	xxx_messageInfo_Genre_Genre.DiscardUnknown(m)
}

var xxx_messageInfo_Genre_Genre proto.InternalMessageInfo

func (m *Genre_Genre) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Genre_Genre) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Genre_Genre) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Genre_Genre) GetWorkCount() uint64 {
	if m != nil {
		return m.WorkCount
	}
	return 0
}

func (m *Genre_Genre) GetSubgenres() []*Genre_Genre {
	if m != nil {
		return m.Subgenres
	}
	return nil
}

type Genre_Group struct {
	// id группы жанров
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// название
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// жанры
	Genres               []*Genre_Genre `protobuf:"bytes,5,rep,name=genres,proto3" json:"genres,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Genre_Group) Reset()         { *m = Genre_Group{} }
func (m *Genre_Group) String() string { return proto.CompactTextString(m) }
func (*Genre_Group) ProtoMessage()    {}
func (*Genre_Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_d42afaacf88a7096, []int{0, 1}
}

func (m *Genre_Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genre_Group.Unmarshal(m, b)
}
func (m *Genre_Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genre_Group.Marshal(b, m, deterministic)
}
func (m *Genre_Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genre_Group.Merge(m, src)
}
func (m *Genre_Group) XXX_Size() int {
	return xxx_messageInfo_Genre_Group.Size(m)
}
func (m *Genre_Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Genre_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Genre_Group proto.InternalMessageInfo

func (m *Genre_Group) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Genre_Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Genre_Group) GetGenres() []*Genre_Genre {
	if m != nil {
		return m.Genres
	}
	return nil
}

type Genre_Response struct {
	// группы жанров
	Groups               []*Genre_Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Genre_Response) Reset()         { *m = Genre_Response{} }
func (m *Genre_Response) String() string { return proto.CompactTextString(m) }
func (*Genre_Response) ProtoMessage()    {}
func (*Genre_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_d42afaacf88a7096, []int{0, 2}
}

func (m *Genre_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genre_Response.Unmarshal(m, b)
}
func (m *Genre_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genre_Response.Marshal(b, m, deterministic)
}
func (m *Genre_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genre_Response.Merge(m, src)
}
func (m *Genre_Response) XXX_Size() int {
	return xxx_messageInfo_Genre_Response.Size(m)
}
func (m *Genre_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Genre_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Genre_Response proto.InternalMessageInfo

func (m *Genre_Response) GetGroups() []*Genre_Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*Genre)(nil), "Genre")
	proto.RegisterType((*Genre_Genre)(nil), "Genre.Genre")
	proto.RegisterType((*Genre_Group)(nil), "Genre.Group")
	proto.RegisterType((*Genre_Response)(nil), "Genre.Response")
}

func init() { proto.RegisterFile("proto/genres.proto", fileDescriptor_d42afaacf88a7096) }

var fileDescriptor_d42afaacf88a7096 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4f, 0xcd, 0x2b, 0x4a, 0x2d, 0xd6, 0x03, 0x73, 0x94, 0x66, 0x32, 0x71, 0xb1,
	0xba, 0x83, 0x04, 0xa4, 0xba, 0x18, 0xa1, 0x2c, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x1b, 0x24, 0x96, 0x99, 0x97, 0x96, 0x2f, 0xc1,
	0x0c, 0x11, 0x03, 0xb1, 0x85, 0x64, 0xb9, 0xb8, 0xca, 0xf3, 0x8b, 0xb2, 0xe3, 0x93, 0xf3, 0x4b,
	0xf3, 0x4a, 0x24, 0x58, 0xc0, 0xfa, 0x39, 0x41, 0x22, 0xce, 0x20, 0x01, 0x21, 0x2d, 0x2e, 0xce,
	0xe2, 0xd2, 0x24, 0x88, 0xed, 0x12, 0xac, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x3c, 0x7a, 0x60, 0x1b,
	0x21, 0x64, 0x10, 0x42, 0x5a, 0x2a, 0x90, 0x8b, 0xd5, 0xbd, 0x28, 0xbf, 0xb4, 0x80, 0x28, 0xb7,
	0xa8, 0x70, 0xb1, 0xe1, 0x31, 0x15, 0x2a, 0x27, 0x65, 0xc0, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90,
	0x9f, 0x57, 0x0c, 0xd1, 0x01, 0x32, 0xbe, 0x58, 0x82, 0x11, 0x55, 0x07, 0x48, 0x30, 0x08, 0x2a,
	0xe7, 0x24, 0x13, 0x25, 0x55, 0x9c, 0x5f, 0x5a, 0x94, 0x9c, 0x5a, 0xac, 0x5f, 0x9c, 0x5a, 0x54,
	0x96, 0x5a, 0xa4, 0x9f, 0x99, 0x57, 0x92, 0x5a, 0x94, 0x97, 0x98, 0xa3, 0x5f, 0x90, 0x94, 0xc4,
	0x06, 0x0e, 0x40, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0x5d, 0x80, 0xb9, 0x56, 0x01,
	0x00, 0x00,
}