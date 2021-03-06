// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1-devel
// 	protoc        v3.11.4
// source: proto/genres.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type Genre struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Genre) Reset() {
	*x = Genre{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_genres_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre) ProtoMessage() {}

func (x *Genre) ProtoReflect() protoreflect.Message {
	mi := &file_proto_genres_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre.ProtoReflect.Descriptor instead.
func (*Genre) Descriptor() ([]byte, []int) {
	return file_proto_genres_proto_rawDescGZIP(), []int{0}
}

type Genre_Genre struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id жанра
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// название
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// информация
	Info string `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	// поджанры
	Subgenres []*Genre_Genre `protobuf:"bytes,4,rep,name=subgenres,proto3" json:"subgenres,omitempty"`
	// количество произведений (опционально)
	WorkCount uint64 `protobuf:"varint,5,opt,name=work_count,json=workCount,proto3" json:"work_count,omitempty"`
	// количество голосов (опционально)
	VoteCount uint64 `protobuf:"varint,6,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
}

func (x *Genre_Genre) Reset() {
	*x = Genre_Genre{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_genres_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genre_Genre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre_Genre) ProtoMessage() {}

func (x *Genre_Genre) ProtoReflect() protoreflect.Message {
	mi := &file_proto_genres_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre_Genre.ProtoReflect.Descriptor instead.
func (*Genre_Genre) Descriptor() ([]byte, []int) {
	return file_proto_genres_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Genre_Genre) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Genre_Genre) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Genre_Genre) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *Genre_Genre) GetSubgenres() []*Genre_Genre {
	if x != nil {
		return x.Subgenres
	}
	return nil
}

func (x *Genre_Genre) GetWorkCount() uint64 {
	if x != nil {
		return x.WorkCount
	}
	return 0
}

func (x *Genre_Genre) GetVoteCount() uint64 {
	if x != nil {
		return x.VoteCount
	}
	return 0
}

type Genre_Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id группы жанров
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// название
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// жанры
	Genres []*Genre_Genre `protobuf:"bytes,5,rep,name=genres,proto3" json:"genres,omitempty"`
}

func (x *Genre_Group) Reset() {
	*x = Genre_Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_genres_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genre_Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre_Group) ProtoMessage() {}

func (x *Genre_Group) ProtoReflect() protoreflect.Message {
	mi := &file_proto_genres_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre_Group.ProtoReflect.Descriptor instead.
func (*Genre_Group) Descriptor() ([]byte, []int) {
	return file_proto_genres_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Genre_Group) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Genre_Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Genre_Group) GetGenres() []*Genre_Genre {
	if x != nil {
		return x.Genres
	}
	return nil
}

type Genre_GenresResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// группы жанров
	Groups []*Genre_Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *Genre_GenresResponse) Reset() {
	*x = Genre_GenresResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_genres_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genre_GenresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre_GenresResponse) ProtoMessage() {}

func (x *Genre_GenresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_genres_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre_GenresResponse.ProtoReflect.Descriptor instead.
func (*Genre_GenresResponse) Descriptor() ([]byte, []int) {
	return file_proto_genres_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Genre_GenresResponse) GetGroups() []*Genre_Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type Genre_ClassificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// группы жанров
	Groups []*Genre_Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	// сколько раз пользователи классифицировали произведение
	ClassificationCount uint64 `protobuf:"varint,2,opt,name=classification_count,json=classificationCount,proto3" json:"classification_count,omitempty"`
}

func (x *Genre_ClassificationResponse) Reset() {
	*x = Genre_ClassificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_genres_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genre_ClassificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre_ClassificationResponse) ProtoMessage() {}

func (x *Genre_ClassificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_genres_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre_ClassificationResponse.ProtoReflect.Descriptor instead.
func (*Genre_ClassificationResponse) Descriptor() ([]byte, []int) {
	return file_proto_genres_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Genre_ClassificationResponse) GetGroups() []*Genre_Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *Genre_ClassificationResponse) GetClassificationCount() uint64 {
	if x != nil {
		return x.ClassificationCount
	}
	return 0
}

var File_proto_genres_proto protoreflect.FileDescriptor

var file_proto_genres_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x03, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x1a, 0xa9,
	0x01, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x2a, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x52, 0x09, 0x73, 0x75, 0x62, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x76, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x51, 0x0a, 0x05, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x2e,
	0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x1a, 0x36, 0x0a,
	0x0e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x71, 0x0a, 0x16, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x13, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_genres_proto_rawDescOnce sync.Once
	file_proto_genres_proto_rawDescData = file_proto_genres_proto_rawDesc
)

func file_proto_genres_proto_rawDescGZIP() []byte {
	file_proto_genres_proto_rawDescOnce.Do(func() {
		file_proto_genres_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_genres_proto_rawDescData)
	})
	return file_proto_genres_proto_rawDescData
}

var file_proto_genres_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_genres_proto_goTypes = []interface{}{
	(*Genre)(nil),                        // 0: Genre
	(*Genre_Genre)(nil),                  // 1: Genre.Genre
	(*Genre_Group)(nil),                  // 2: Genre.Group
	(*Genre_GenresResponse)(nil),         // 3: Genre.GenresResponse
	(*Genre_ClassificationResponse)(nil), // 4: Genre.ClassificationResponse
}
var file_proto_genres_proto_depIdxs = []int32{
	1, // 0: Genre.Genre.subgenres:type_name -> Genre.Genre
	1, // 1: Genre.Group.genres:type_name -> Genre.Genre
	2, // 2: Genre.GenresResponse.groups:type_name -> Genre.Group
	2, // 3: Genre.ClassificationResponse.groups:type_name -> Genre.Group
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_genres_proto_init() }
func file_proto_genres_proto_init() {
	if File_proto_genres_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_genres_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genre); i {
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
		file_proto_genres_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genre_Genre); i {
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
		file_proto_genres_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genre_Group); i {
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
		file_proto_genres_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genre_GenresResponse); i {
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
		file_proto_genres_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genre_ClassificationResponse); i {
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
			RawDescriptor: file_proto_genres_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_genres_proto_goTypes,
		DependencyIndexes: file_proto_genres_proto_depIdxs,
		MessageInfos:      file_proto_genres_proto_msgTypes,
	}.Build()
	File_proto_genres_proto = out.File
	file_proto_genres_proto_rawDesc = nil
	file_proto_genres_proto_goTypes = nil
	file_proto_genres_proto_depIdxs = nil
}
