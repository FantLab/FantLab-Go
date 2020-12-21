// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1-devel
// 	protoc        v3.11.4
// source: proto/common.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Common_Gender int32

const (
	Common_GENDER_UNKNOWN Common_Gender = 0
	// мужчина
	Common_GENDER_MALE Common_Gender = 1
	// женщина
	Common_GENDER_FEMALE Common_Gender = 2
)

// Enum value maps for Common_Gender.
var (
	Common_Gender_name = map[int32]string{
		0: "GENDER_UNKNOWN",
		1: "GENDER_MALE",
		2: "GENDER_FEMALE",
	}
	Common_Gender_value = map[string]int32{
		"GENDER_UNKNOWN": 0,
		"GENDER_MALE":    1,
		"GENDER_FEMALE":  2,
	}
)

func (x Common_Gender) Enum() *Common_Gender {
	p := new(Common_Gender)
	*p = x
	return p
}

func (x Common_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Common_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_common_proto_enumTypes[0].Descriptor()
}

func (Common_Gender) Type() protoreflect.EnumType {
	return &file_proto_common_proto_enumTypes[0]
}

func (x Common_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Common_Gender.Descriptor instead.
func (Common_Gender) EnumDescriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{0, 0}
}

type Common_UserClass int32

const (
	Common_USERCLASS_UNKNOWN Common_UserClass = 0
	// новичок
	Common_USERCLASS_BEGINNER Common_UserClass = 1
	// активист
	Common_USERCLASS_ACTIVIST Common_UserClass = 2
	// авторитет
	Common_USERCLASS_AUTHORITY Common_UserClass = 3
	// философ
	Common_USERCLASS_PHILOSOPHER Common_UserClass = 4
	// магистр
	Common_USERCLASS_MASTER Common_UserClass = 5
	// гранд-мастер
	Common_USERCLASS_GRANDMASTER Common_UserClass = 6
	// миродержец
	Common_USERCLASS_PEACEKEEPER Common_UserClass = 7
	// миротворец
	Common_USERCLASS_PEACEMAKER Common_UserClass = 8
)

// Enum value maps for Common_UserClass.
var (
	Common_UserClass_name = map[int32]string{
		0: "USERCLASS_UNKNOWN",
		1: "USERCLASS_BEGINNER",
		2: "USERCLASS_ACTIVIST",
		3: "USERCLASS_AUTHORITY",
		4: "USERCLASS_PHILOSOPHER",
		5: "USERCLASS_MASTER",
		6: "USERCLASS_GRANDMASTER",
		7: "USERCLASS_PEACEKEEPER",
		8: "USERCLASS_PEACEMAKER",
	}
	Common_UserClass_value = map[string]int32{
		"USERCLASS_UNKNOWN":     0,
		"USERCLASS_BEGINNER":    1,
		"USERCLASS_ACTIVIST":    2,
		"USERCLASS_AUTHORITY":   3,
		"USERCLASS_PHILOSOPHER": 4,
		"USERCLASS_MASTER":      5,
		"USERCLASS_GRANDMASTER": 6,
		"USERCLASS_PEACEKEEPER": 7,
		"USERCLASS_PEACEMAKER":  8,
	}
)

func (x Common_UserClass) Enum() *Common_UserClass {
	p := new(Common_UserClass)
	*p = x
	return p
}

func (x Common_UserClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Common_UserClass) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_common_proto_enumTypes[1].Descriptor()
}

func (Common_UserClass) Type() protoreflect.EnumType {
	return &file_proto_common_proto_enumTypes[1]
}

func (x Common_UserClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Common_UserClass.Descriptor instead.
func (Common_UserClass) EnumDescriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{0, 1}
}

type Common struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Common) Reset() {
	*x = Common{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Common) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Common) ProtoMessage() {}

func (x *Common) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Common.ProtoReflect.Descriptor instead.
func (*Common) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{0}
}

type Common_UserLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id пользователя
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// логин
	Login string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	// имя
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// пол
	Gender Common_Gender `protobuf:"varint,4,opt,name=gender,proto3,enum=Common_Gender" json:"gender,omitempty"`
	// аватар
	Avatar string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	// класс
	Class Common_UserClass `protobuf:"varint,6,opt,name=class,proto3,enum=Common_UserClass" json:"class,omitempty"`
	// подпись на форуме
	Sign string `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *Common_UserLink) Reset() {
	*x = Common_UserLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Common_UserLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Common_UserLink) ProtoMessage() {}

func (x *Common_UserLink) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Common_UserLink.ProtoReflect.Descriptor instead.
func (*Common_UserLink) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Common_UserLink) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Common_UserLink) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Common_UserLink) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Common_UserLink) GetGender() Common_Gender {
	if x != nil {
		return x.Gender
	}
	return Common_GENDER_UNKNOWN
}

func (x *Common_UserLink) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Common_UserLink) GetClass() Common_UserClass {
	if x != nil {
		return x.Class
	}
	return Common_USERCLASS_UNKNOWN
}

func (x *Common_UserLink) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

type Common_Creation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// пользователь
	User *Common_UserLink `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// дата создания
	Date *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Common_Creation) Reset() {
	*x = Common_Creation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Common_Creation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Common_Creation) ProtoMessage() {}

func (x *Common_Creation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Common_Creation.ProtoReflect.Descriptor instead.
func (*Common_Creation) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Common_Creation) GetUser() *Common_UserLink {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Common_Creation) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type Common_Pages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// текущая
	Current uint64 `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	// количество
	Count uint64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Common_Pages) Reset() {
	*x = Common_Pages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Common_Pages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Common_Pages) ProtoMessage() {}

func (x *Common_Pages) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Common_Pages.ProtoReflect.Descriptor instead.
func (*Common_Pages) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Common_Pages) GetCurrent() uint64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *Common_Pages) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Common_Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ссылка на файл
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// размер (байт)
	Size uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Common_Attachment) Reset() {
	*x = Common_Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Common_Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Common_Attachment) ProtoMessage() {}

func (x *Common_Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Common_Attachment.ProtoReflect.Descriptor instead.
func (*Common_Attachment) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Common_Attachment) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Common_Attachment) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Common_SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Common_SuccessResponse) Reset() {
	*x = Common_SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Common_SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Common_SuccessResponse) ProtoMessage() {}

func (x *Common_SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Common_SuccessResponse.ProtoReflect.Descriptor instead.
func (*Common_SuccessResponse) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{0, 4}
}

type Common_SuccessIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// идентификатор
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Common_SuccessIdResponse) Reset() {
	*x = Common_SuccessIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Common_SuccessIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Common_SuccessIdResponse) ProtoMessage() {}

func (x *Common_SuccessIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Common_SuccessIdResponse.ProtoReflect.Descriptor instead.
func (*Common_SuccessIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{0, 5}
}

func (x *Common_SuccessIdResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Common_FileUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL на загрузку файла
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Common_FileUploadResponse) Reset() {
	*x = Common_FileUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Common_FileUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Common_FileUploadResponse) ProtoMessage() {}

func (x *Common_FileUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Common_FileUploadResponse.ProtoReflect.Descriptor instead.
func (*Common_FileUploadResponse) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{0, 6}
}

func (x *Common_FileUploadResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_proto_common_proto protoreflect.FileDescriptor

var file_proto_common_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x06, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x1a, 0xc1, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x69, 0x67, 0x6e, 0x1a, 0x60, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x37, 0x0a, 0x05, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a,
	0x32, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x1a, 0x11, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x23, 0x0a, 0x11, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x26, 0x0a, 0x12, 0x46,
	0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x40, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x0e, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x45, 0x4d,
	0x41, 0x4c, 0x45, 0x10, 0x02, 0x22, 0xec, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x53, 0x45, 0x52, 0x43, 0x4c, 0x41, 0x53, 0x53,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x53,
	0x45, 0x52, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x42, 0x45, 0x47, 0x49, 0x4e, 0x4e, 0x45, 0x52,
	0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x53, 0x45, 0x52, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x53, 0x54, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x53,
	0x45, 0x52, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x54,
	0x59, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x53, 0x45, 0x52, 0x43, 0x4c, 0x41, 0x53, 0x53,
	0x5f, 0x50, 0x48, 0x49, 0x4c, 0x4f, 0x53, 0x4f, 0x50, 0x48, 0x45, 0x52, 0x10, 0x04, 0x12, 0x14,
	0x0a, 0x10, 0x55, 0x53, 0x45, 0x52, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x4d, 0x41, 0x53, 0x54,
	0x45, 0x52, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x53, 0x45, 0x52, 0x43, 0x4c, 0x41, 0x53,
	0x53, 0x5f, 0x47, 0x52, 0x41, 0x4e, 0x44, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x10, 0x06, 0x12,
	0x19, 0x0a, 0x15, 0x55, 0x53, 0x45, 0x52, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x50, 0x45, 0x41,
	0x43, 0x45, 0x4b, 0x45, 0x45, 0x50, 0x45, 0x52, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x53,
	0x45, 0x52, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x50, 0x45, 0x41, 0x43, 0x45, 0x4d, 0x41, 0x4b,
	0x45, 0x52, 0x10, 0x08, 0x42, 0x0c, 0x5a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_common_proto_rawDescOnce sync.Once
	file_proto_common_proto_rawDescData = file_proto_common_proto_rawDesc
)

func file_proto_common_proto_rawDescGZIP() []byte {
	file_proto_common_proto_rawDescOnce.Do(func() {
		file_proto_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_common_proto_rawDescData)
	})
	return file_proto_common_proto_rawDescData
}

var file_proto_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_common_proto_goTypes = []interface{}{
	(Common_Gender)(0),                // 0: Common.Gender
	(Common_UserClass)(0),             // 1: Common.UserClass
	(*Common)(nil),                    // 2: Common
	(*Common_UserLink)(nil),           // 3: Common.UserLink
	(*Common_Creation)(nil),           // 4: Common.Creation
	(*Common_Pages)(nil),              // 5: Common.Pages
	(*Common_Attachment)(nil),         // 6: Common.Attachment
	(*Common_SuccessResponse)(nil),    // 7: Common.SuccessResponse
	(*Common_SuccessIdResponse)(nil),  // 8: Common.SuccessIdResponse
	(*Common_FileUploadResponse)(nil), // 9: Common.FileUploadResponse
	(*timestamp.Timestamp)(nil),       // 10: google.protobuf.Timestamp
}
var file_proto_common_proto_depIdxs = []int32{
	0,  // 0: Common.UserLink.gender:type_name -> Common.Gender
	1,  // 1: Common.UserLink.class:type_name -> Common.UserClass
	3,  // 2: Common.Creation.user:type_name -> Common.UserLink
	10, // 3: Common.Creation.date:type_name -> google.protobuf.Timestamp
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_common_proto_init() }
func file_proto_common_proto_init() {
	if File_proto_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Common); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Common_UserLink); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Common_Creation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Common_Pages); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Common_Attachment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Common_SuccessResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Common_SuccessIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Common_FileUploadResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_common_proto_goTypes,
		DependencyIndexes: file_proto_common_proto_depIdxs,
		EnumInfos:         file_proto_common_proto_enumTypes,
		MessageInfos:      file_proto_common_proto_msgTypes,
	}.Build()
	File_proto_common_proto = out.File
	file_proto_common_proto_rawDesc = nil
	file_proto_common_proto_goTypes = nil
	file_proto_common_proto_depIdxs = nil
}
