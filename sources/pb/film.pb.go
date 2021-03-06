// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1-devel
// 	protoc        v3.11.4
// source: proto/film.proto

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

// TODO Список взят из Film.pm; если он изменится в основном репозитории, здесь тоже придется дорабатывать
// https://github.com/parserpro/fantlab/blob/2933533d76806584d52533b6026b8a3411193732/pm/Film.pm#L11-L18
type FilmType int32

const (
	FilmType_FILM_TYPE_UNKNOWN FilmType = 0
	// Фильм
	FilmType_FILM_TYPE_FILM FilmType = 1
	// Сериал
	FilmType_FILM_TYPE_SERIES FilmType = 2
	// Эпизод
	FilmType_FILM_TYPE_EPISODE FilmType = 3
	// Документальный
	FilmType_FILM_TYPE_DOCUMENTARY FilmType = 4
	// Анимационный
	FilmType_FILM_TYPE_ANIMATION FilmType = 5
	// Короткометражный
	FilmType_FILM_TYPE_SHORT FilmType = 6
	// Телеспектакль
	FilmType_FILM_TYPE_SPECTACLE FilmType = 7
)

// Enum value maps for FilmType.
var (
	FilmType_name = map[int32]string{
		0: "FILM_TYPE_UNKNOWN",
		1: "FILM_TYPE_FILM",
		2: "FILM_TYPE_SERIES",
		3: "FILM_TYPE_EPISODE",
		4: "FILM_TYPE_DOCUMENTARY",
		5: "FILM_TYPE_ANIMATION",
		6: "FILM_TYPE_SHORT",
		7: "FILM_TYPE_SPECTACLE",
	}
	FilmType_value = map[string]int32{
		"FILM_TYPE_UNKNOWN":     0,
		"FILM_TYPE_FILM":        1,
		"FILM_TYPE_SERIES":      2,
		"FILM_TYPE_EPISODE":     3,
		"FILM_TYPE_DOCUMENTARY": 4,
		"FILM_TYPE_ANIMATION":   5,
		"FILM_TYPE_SHORT":       6,
		"FILM_TYPE_SPECTACLE":   7,
	}
)

func (x FilmType) Enum() *FilmType {
	p := new(FilmType)
	*p = x
	return p
}

func (x FilmType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilmType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_film_proto_enumTypes[0].Descriptor()
}

func (FilmType) Type() protoreflect.EnumType {
	return &file_proto_film_proto_enumTypes[0]
}

func (x FilmType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilmType.Descriptor instead.
func (FilmType) EnumDescriptor() ([]byte, []int) {
	return file_proto_film_proto_rawDescGZIP(), []int{0}
}

var File_proto_film_proto protoreflect.FileDescriptor

var file_proto_film_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0xc4, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x46, 0x49, 0x4c, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x49, 0x4c, 0x4d, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x4d, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x49,
	0x4c, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x49, 0x45, 0x53, 0x10, 0x02,
	0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x4c, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x50,
	0x49, 0x53, 0x4f, 0x44, 0x45, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x49, 0x4c, 0x4d, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x43, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x41, 0x52, 0x59,
	0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x4c, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x41, 0x4e, 0x49, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x46,
	0x49, 0x4c, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x48, 0x4f, 0x52, 0x54, 0x10, 0x06,
	0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x4c, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x50,
	0x45, 0x43, 0x54, 0x41, 0x43, 0x4c, 0x45, 0x10, 0x07, 0x42, 0x0c, 0x5a, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_film_proto_rawDescOnce sync.Once
	file_proto_film_proto_rawDescData = file_proto_film_proto_rawDesc
)

func file_proto_film_proto_rawDescGZIP() []byte {
	file_proto_film_proto_rawDescOnce.Do(func() {
		file_proto_film_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_film_proto_rawDescData)
	})
	return file_proto_film_proto_rawDescData
}

var file_proto_film_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_film_proto_goTypes = []interface{}{
	(FilmType)(0), // 0: FilmType
}
var file_proto_film_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_film_proto_init() }
func file_proto_film_proto_init() {
	if File_proto_film_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_film_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_film_proto_goTypes,
		DependencyIndexes: file_proto_film_proto_depIdxs,
		EnumInfos:         file_proto_film_proto_enumTypes,
	}.Build()
	File_proto_film_proto = out.File
	file_proto_film_proto_rawDesc = nil
	file_proto_film_proto_goTypes = nil
	file_proto_film_proto_depIdxs = nil
}
