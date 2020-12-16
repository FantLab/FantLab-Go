// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1-devel
// 	protoc        v3.11.4
// source: proto/work.proto

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

// TODO Список взят из таблицы work_types; если он изменится в базе, здесь тоже придется дорабатывать
type WorkType int32

const (
	WorkType_WORK_TYPE_UNKNOWN WorkType = 0
	// Роман
	WorkType_WORK_TYPE_NOVEL WorkType = 1
	// Сборник
	WorkType_WORK_TYPE_COMPILATION WorkType = 2
	// Цикл
	WorkType_WORK_TYPE_SERIES WorkType = 3
	// Стихотворение
	WorkType_WORK_TYPE_VERSE WorkType = 4
	// Произведение (прочее)
	WorkType_WORK_TYPE_OTHER WorkType = 5
	// Сказка
	WorkType_WORK_TYPE_FAIRY_TALE WorkType = 6
	// Эссе
	WorkType_WORK_TYPE_ESSAY WorkType = 7
	// Статья
	WorkType_WORK_TYPE_ARTICLE WorkType = 8
	// Роман-эпопея
	WorkType_WORK_TYPE_EPIC_NOVEL WorkType = 9
	// Антология
	WorkType_WORK_TYPE_ANTHOLOGY WorkType = 10
	// Пьеса
	WorkType_WORK_TYPE_PLAY WorkType = 11
	// Киносценарий
	WorkType_WORK_TYPE_SCREENPLAY WorkType = 12
	// Документальное произведение
	WorkType_WORK_TYPE_DOCUMENTARY WorkType = 13
	// Микрорассказ
	WorkType_WORK_TYPE_MICROTALE WorkType = 14
	// Диссертация
	WorkType_WORK_TYPE_DISSERTATION WorkType = 15
	// Монография
	WorkType_WORK_TYPE_MONOGRAPH WorkType = 16
	// Учебное издание
	WorkType_WORK_TYPE_EDUCATIONAL_PUBLICATION WorkType = 17
	// Энциклопедия/справочник
	WorkType_WORK_TYPE_ENCYCLOPEDIA WorkType = 18
	// Журнал
	WorkType_WORK_TYPE_MAGAZINE WorkType = 19
	// Поэма
	WorkType_WORK_TYPE_POEM WorkType = 20
	// Стихотворения
	WorkType_WORK_TYPE_POETRY WorkType = 21
	// Стихотворение в прозе
	WorkType_WORK_TYPE_PROSE_VERSE WorkType = 22
	// Комикс
	WorkType_WORK_TYPE_COMIC_BOOK WorkType = 23
	// Манга
	WorkType_WORK_TYPE_MANGA WorkType = 24
	// Графический роман
	WorkType_WORK_TYPE_GRAPHIC_NOVEL WorkType = 25
	// Повесть
	WorkType_WORK_TYPE_NOVELETTE WorkType = 26
	// Рассказ
	WorkType_WORK_TYPE_STORY WorkType = 27
	// Очерк
	WorkType_WORK_TYPE_FEATURE_ARTICLE WorkType = 28
	// Репортаж
	WorkType_WORK_TYPE_REPORTAGE WorkType = 29
	// Условный цикл
	WorkType_WORK_TYPE_CONDITIONAL_SERIES WorkType = 30
	// Отрывок
	WorkType_WORK_TYPE_EXCERPT WorkType = 31
	// Интервью
	WorkType_WORK_TYPE_INTERVIEW WorkType = 32
	// Рецензия
	WorkType_WORK_TYPE_REVIEW WorkType = 33
	// Научно-популярная книга
	WorkType_WORK_TYPE_POPULAR_SCIENCE_BOOK WorkType = 34
	// Артбук
	WorkType_WORK_TYPE_ARTBOOK WorkType = 35
	// Либретто
	WorkType_WORK_TYPE_LIBRETTO WorkType = 36
)

// Enum value maps for WorkType.
var (
	WorkType_name = map[int32]string{
		0:  "WORK_TYPE_UNKNOWN",
		1:  "WORK_TYPE_NOVEL",
		2:  "WORK_TYPE_COMPILATION",
		3:  "WORK_TYPE_SERIES",
		4:  "WORK_TYPE_VERSE",
		5:  "WORK_TYPE_OTHER",
		6:  "WORK_TYPE_FAIRY_TALE",
		7:  "WORK_TYPE_ESSAY",
		8:  "WORK_TYPE_ARTICLE",
		9:  "WORK_TYPE_EPIC_NOVEL",
		10: "WORK_TYPE_ANTHOLOGY",
		11: "WORK_TYPE_PLAY",
		12: "WORK_TYPE_SCREENPLAY",
		13: "WORK_TYPE_DOCUMENTARY",
		14: "WORK_TYPE_MICROTALE",
		15: "WORK_TYPE_DISSERTATION",
		16: "WORK_TYPE_MONOGRAPH",
		17: "WORK_TYPE_EDUCATIONAL_PUBLICATION",
		18: "WORK_TYPE_ENCYCLOPEDIA",
		19: "WORK_TYPE_MAGAZINE",
		20: "WORK_TYPE_POEM",
		21: "WORK_TYPE_POETRY",
		22: "WORK_TYPE_PROSE_VERSE",
		23: "WORK_TYPE_COMIC_BOOK",
		24: "WORK_TYPE_MANGA",
		25: "WORK_TYPE_GRAPHIC_NOVEL",
		26: "WORK_TYPE_NOVELETTE",
		27: "WORK_TYPE_STORY",
		28: "WORK_TYPE_FEATURE_ARTICLE",
		29: "WORK_TYPE_REPORTAGE",
		30: "WORK_TYPE_CONDITIONAL_SERIES",
		31: "WORK_TYPE_EXCERPT",
		32: "WORK_TYPE_INTERVIEW",
		33: "WORK_TYPE_REVIEW",
		34: "WORK_TYPE_POPULAR_SCIENCE_BOOK",
		35: "WORK_TYPE_ARTBOOK",
		36: "WORK_TYPE_LIBRETTO",
	}
	WorkType_value = map[string]int32{
		"WORK_TYPE_UNKNOWN":                 0,
		"WORK_TYPE_NOVEL":                   1,
		"WORK_TYPE_COMPILATION":             2,
		"WORK_TYPE_SERIES":                  3,
		"WORK_TYPE_VERSE":                   4,
		"WORK_TYPE_OTHER":                   5,
		"WORK_TYPE_FAIRY_TALE":              6,
		"WORK_TYPE_ESSAY":                   7,
		"WORK_TYPE_ARTICLE":                 8,
		"WORK_TYPE_EPIC_NOVEL":              9,
		"WORK_TYPE_ANTHOLOGY":               10,
		"WORK_TYPE_PLAY":                    11,
		"WORK_TYPE_SCREENPLAY":              12,
		"WORK_TYPE_DOCUMENTARY":             13,
		"WORK_TYPE_MICROTALE":               14,
		"WORK_TYPE_DISSERTATION":            15,
		"WORK_TYPE_MONOGRAPH":               16,
		"WORK_TYPE_EDUCATIONAL_PUBLICATION": 17,
		"WORK_TYPE_ENCYCLOPEDIA":            18,
		"WORK_TYPE_MAGAZINE":                19,
		"WORK_TYPE_POEM":                    20,
		"WORK_TYPE_POETRY":                  21,
		"WORK_TYPE_PROSE_VERSE":             22,
		"WORK_TYPE_COMIC_BOOK":              23,
		"WORK_TYPE_MANGA":                   24,
		"WORK_TYPE_GRAPHIC_NOVEL":           25,
		"WORK_TYPE_NOVELETTE":               26,
		"WORK_TYPE_STORY":                   27,
		"WORK_TYPE_FEATURE_ARTICLE":         28,
		"WORK_TYPE_REPORTAGE":               29,
		"WORK_TYPE_CONDITIONAL_SERIES":      30,
		"WORK_TYPE_EXCERPT":                 31,
		"WORK_TYPE_INTERVIEW":               32,
		"WORK_TYPE_REVIEW":                  33,
		"WORK_TYPE_POPULAR_SCIENCE_BOOK":    34,
		"WORK_TYPE_ARTBOOK":                 35,
		"WORK_TYPE_LIBRETTO":                36,
	}
)

func (x WorkType) Enum() *WorkType {
	p := new(WorkType)
	*p = x
	return p
}

func (x WorkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_work_proto_enumTypes[0].Descriptor()
}

func (WorkType) Type() protoreflect.EnumType {
	return &file_proto_work_proto_enumTypes[0]
}

func (x WorkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkType.Descriptor instead.
func (WorkType) EnumDescriptor() ([]byte, []int) {
	return file_proto_work_proto_rawDescGZIP(), []int{0}
}

type Work_PublishStatus int32

const (
	Work_PUBLISH_STATUS_UNKNOWN Work_PublishStatus = 0
	// Не закончено
	Work_PUBLISH_STATUS_NOT_FINISHED Work_PublishStatus = 1
	// Не опубликовано
	Work_PUBLISH_STATUS_NOT_PUBLISHED Work_PublishStatus = 2
	// Сетевая публикация
	Work_PUBLISH_STATUS_NETWORK_PUBLICATION Work_PublishStatus = 3
	// Доступно в сети
	Work_PUBLISH_STATUS_AVAILABLE_ONLINE Work_PublishStatus = 4
	// В планах автора
	Work_PUBLISH_STATUS_PLANNED_BY_THE_AUTHOR Work_PublishStatus = 5
)

// Enum value maps for Work_PublishStatus.
var (
	Work_PublishStatus_name = map[int32]string{
		0: "PUBLISH_STATUS_UNKNOWN",
		1: "PUBLISH_STATUS_NOT_FINISHED",
		2: "PUBLISH_STATUS_NOT_PUBLISHED",
		3: "PUBLISH_STATUS_NETWORK_PUBLICATION",
		4: "PUBLISH_STATUS_AVAILABLE_ONLINE",
		5: "PUBLISH_STATUS_PLANNED_BY_THE_AUTHOR",
	}
	Work_PublishStatus_value = map[string]int32{
		"PUBLISH_STATUS_UNKNOWN":               0,
		"PUBLISH_STATUS_NOT_FINISHED":          1,
		"PUBLISH_STATUS_NOT_PUBLISHED":         2,
		"PUBLISH_STATUS_NETWORK_PUBLICATION":   3,
		"PUBLISH_STATUS_AVAILABLE_ONLINE":      4,
		"PUBLISH_STATUS_PLANNED_BY_THE_AUTHOR": 5,
	}
)

func (x Work_PublishStatus) Enum() *Work_PublishStatus {
	p := new(Work_PublishStatus)
	*p = x
	return p
}

func (x Work_PublishStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Work_PublishStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_work_proto_enumTypes[1].Descriptor()
}

func (Work_PublishStatus) Type() protoreflect.EnumType {
	return &file_proto_work_proto_enumTypes[1]
}

func (x Work_PublishStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Work_PublishStatus.Descriptor instead.
func (Work_PublishStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_work_proto_rawDescGZIP(), []int{0, 0}
}

type Work struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Work) Reset() {
	*x = Work{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_work_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Work) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Work) ProtoMessage() {}

func (x *Work) ProtoReflect() protoreflect.Message {
	mi := &file_proto_work_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Work.ProtoReflect.Descriptor instead.
func (*Work) Descriptor() ([]byte, []int) {
	return file_proto_work_proto_rawDescGZIP(), []int{0}
}

type Work_SubWork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// идентификатор произведения
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// оригинальное название
	OrigName string `protobuf:"bytes,2,opt,name=orig_name,json=origName,proto3" json:"orig_name,omitempty"`
	// название на русском
	RusName string `protobuf:"bytes,3,opt,name=rus_name,json=rusName,proto3" json:"rus_name,omitempty"`
	// год публикации
	Year uint64 `protobuf:"varint,4,opt,name=year,proto3" json:"year,omitempty"`
	// тип произведения
	WorkType WorkType `protobuf:"varint,5,opt,name=work_type,json=workType,proto3,enum=WorkType" json:"work_type,omitempty"`
	// рейтинг
	Rating float64 `protobuf:"fixed64,6,opt,name=rating,proto3" json:"rating,omitempty"`
	// кол-во оценок
	Marks uint64 `protobuf:"varint,7,opt,name=marks,proto3" json:"marks,omitempty"`
	// кол-во отзывов
	Reviews uint64 `protobuf:"varint,8,opt,name=reviews,proto3" json:"reviews,omitempty"`
	// является ли произведение дополнительным
	Plus bool `protobuf:"varint,9,opt,name=plus,proto3" json:"plus,omitempty"`
	// статус публикации (не закончено, в планах, etc.)
	PublishStatus []Work_PublishStatus `protobuf:"varint,10,rep,packed,name=publish_status,json=publishStatus,proto3,enum=Work_PublishStatus" json:"publish_status,omitempty"`
	// дочерние произведения
	Subworks []*Work_SubWork `protobuf:"bytes,11,rep,name=subworks,proto3" json:"subworks,omitempty"`
}

func (x *Work_SubWork) Reset() {
	*x = Work_SubWork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_work_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Work_SubWork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Work_SubWork) ProtoMessage() {}

func (x *Work_SubWork) ProtoReflect() protoreflect.Message {
	mi := &file_proto_work_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Work_SubWork.ProtoReflect.Descriptor instead.
func (*Work_SubWork) Descriptor() ([]byte, []int) {
	return file_proto_work_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Work_SubWork) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Work_SubWork) GetOrigName() string {
	if x != nil {
		return x.OrigName
	}
	return ""
}

func (x *Work_SubWork) GetRusName() string {
	if x != nil {
		return x.RusName
	}
	return ""
}

func (x *Work_SubWork) GetYear() uint64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Work_SubWork) GetWorkType() WorkType {
	if x != nil {
		return x.WorkType
	}
	return WorkType_WORK_TYPE_UNKNOWN
}

func (x *Work_SubWork) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Work_SubWork) GetMarks() uint64 {
	if x != nil {
		return x.Marks
	}
	return 0
}

func (x *Work_SubWork) GetReviews() uint64 {
	if x != nil {
		return x.Reviews
	}
	return 0
}

func (x *Work_SubWork) GetPlus() bool {
	if x != nil {
		return x.Plus
	}
	return false
}

func (x *Work_SubWork) GetPublishStatus() []Work_PublishStatus {
	if x != nil {
		return x.PublishStatus
	}
	return nil
}

func (x *Work_SubWork) GetSubworks() []*Work_SubWork {
	if x != nil {
		return x.Subworks
	}
	return nil
}

type Work_SubWorksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// айди произведения, для которого был запрос
	WorkId uint64 `protobuf:"varint,1,opt,name=work_id,json=workId,proto3" json:"work_id,omitempty"`
	// произведения, входящие в запрашиваемое
	Subworks []*Work_SubWork `protobuf:"bytes,2,rep,name=subworks,proto3" json:"subworks,omitempty"`
}

func (x *Work_SubWorksResponse) Reset() {
	*x = Work_SubWorksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_work_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Work_SubWorksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Work_SubWorksResponse) ProtoMessage() {}

func (x *Work_SubWorksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_work_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Work_SubWorksResponse.ProtoReflect.Descriptor instead.
func (*Work_SubWorksResponse) Descriptor() ([]byte, []int) {
	return file_proto_work_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Work_SubWorksResponse) GetWorkId() uint64 {
	if x != nil {
		return x.WorkId
	}
	return 0
}

func (x *Work_SubWorksResponse) GetSubworks() []*Work_SubWork {
	if x != nil {
		return x.Subworks
	}
	return nil
}

var File_proto_work_proto protoreflect.FileDescriptor

var file_proto_work_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x99, 0x05, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x6b, 0x1a, 0xd0, 0x02, 0x0a, 0x07,
	0x53, 0x75, 0x62, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x75, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x12, 0x26, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x70, 0x6c, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x2e, 0x53, 0x75, 0x62,
	0x57, 0x6f, 0x72, 0x6b, 0x52, 0x08, 0x73, 0x75, 0x62, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x1a, 0x56,
	0x0a, 0x10, 0x53, 0x75, 0x62, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x73,
	0x75, 0x62, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x2e, 0x53, 0x75, 0x62, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x08, 0x73, 0x75,
	0x62, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x22, 0xe5, 0x01, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x53, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53,
	0x48, 0x45, 0x44, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x12, 0x26, 0x0a, 0x22, 0x50, 0x55, 0x42, 0x4c, 0x49,
	0x53, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52,
	0x4b, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12,
	0x23, 0x0a, 0x1f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x4f, 0x4e, 0x4c, 0x49,
	0x4e, 0x45, 0x10, 0x04, 0x12, 0x28, 0x0a, 0x24, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x4e, 0x45, 0x44, 0x5f, 0x42,
	0x59, 0x5f, 0x54, 0x48, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x10, 0x05, 0x2a, 0xae,
	0x07, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x57,
	0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4e, 0x4f, 0x56, 0x45, 0x4c, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x57, 0x4f, 0x52, 0x4b, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x49, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x45, 0x52, 0x49, 0x45, 0x53, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x4f, 0x52, 0x4b,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x45, 0x10, 0x04, 0x12, 0x13, 0x0a,
	0x0f, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52,
	0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x46, 0x41, 0x49, 0x52, 0x59, 0x5f, 0x54, 0x41, 0x4c, 0x45, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f,
	0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x53, 0x53, 0x41, 0x59, 0x10,
	0x07, 0x12, 0x15, 0x0a, 0x11, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41,
	0x52, 0x54, 0x49, 0x43, 0x4c, 0x45, 0x10, 0x08, 0x12, 0x18, 0x0a, 0x14, 0x57, 0x4f, 0x52, 0x4b,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x50, 0x49, 0x43, 0x5f, 0x4e, 0x4f, 0x56, 0x45, 0x4c,
	0x10, 0x09, 0x12, 0x17, 0x0a, 0x13, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x41, 0x4e, 0x54, 0x48, 0x4f, 0x4c, 0x4f, 0x47, 0x59, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x57,
	0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x10, 0x0b, 0x12,
	0x18, 0x0a, 0x14, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x43, 0x52,
	0x45, 0x45, 0x4e, 0x50, 0x4c, 0x41, 0x59, 0x10, 0x0c, 0x12, 0x19, 0x0a, 0x15, 0x57, 0x4f, 0x52,
	0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x43, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x41,
	0x52, 0x59, 0x10, 0x0d, 0x12, 0x17, 0x0a, 0x13, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x49, 0x43, 0x52, 0x4f, 0x54, 0x41, 0x4c, 0x45, 0x10, 0x0e, 0x12, 0x1a, 0x0a,
	0x16, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x53, 0x45,
	0x52, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0f, 0x12, 0x17, 0x0a, 0x13, 0x57, 0x4f, 0x52,
	0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x4f, 0x47, 0x52, 0x41, 0x50, 0x48,
	0x10, 0x10, 0x12, 0x25, 0x0a, 0x21, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x11, 0x12, 0x1a, 0x0a, 0x16, 0x57, 0x4f, 0x52,
	0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x43, 0x59, 0x43, 0x4c, 0x4f, 0x50, 0x45,
	0x44, 0x49, 0x41, 0x10, 0x12, 0x12, 0x16, 0x0a, 0x12, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4d, 0x41, 0x47, 0x41, 0x5a, 0x49, 0x4e, 0x45, 0x10, 0x13, 0x12, 0x12, 0x0a,
	0x0e, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x45, 0x4d, 0x10,
	0x14, 0x12, 0x14, 0x0a, 0x10, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x4f, 0x45, 0x54, 0x52, 0x59, 0x10, 0x15, 0x12, 0x19, 0x0a, 0x15, 0x57, 0x4f, 0x52, 0x4b, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x53, 0x45, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x45,
	0x10, 0x16, 0x12, 0x18, 0x0a, 0x14, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x4f, 0x4d, 0x49, 0x43, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x17, 0x12, 0x13, 0x0a, 0x0f,
	0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x47, 0x41, 0x10,
	0x18, 0x12, 0x1b, 0x0a, 0x17, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47,
	0x52, 0x41, 0x50, 0x48, 0x49, 0x43, 0x5f, 0x4e, 0x4f, 0x56, 0x45, 0x4c, 0x10, 0x19, 0x12, 0x17,
	0x0a, 0x13, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x56, 0x45,
	0x4c, 0x45, 0x54, 0x54, 0x45, 0x10, 0x1a, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x4f, 0x52, 0x4b, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x1b, 0x12, 0x1d, 0x0a, 0x19,
	0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52,
	0x45, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x43, 0x4c, 0x45, 0x10, 0x1c, 0x12, 0x17, 0x0a, 0x13, 0x57,
	0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x41,
	0x47, 0x45, 0x10, 0x1d, 0x12, 0x20, 0x0a, 0x1c, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45,
	0x52, 0x49, 0x45, 0x53, 0x10, 0x1e, 0x12, 0x15, 0x0a, 0x11, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x52, 0x50, 0x54, 0x10, 0x1f, 0x12, 0x17, 0x0a,
	0x13, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x56, 0x49, 0x45, 0x57, 0x10, 0x20, 0x12, 0x14, 0x0a, 0x10, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x21, 0x12, 0x22, 0x0a, 0x1e,
	0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x50, 0x55, 0x4c, 0x41,
	0x52, 0x5f, 0x53, 0x43, 0x49, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x22,
	0x12, 0x15, 0x0a, 0x11, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52,
	0x54, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x23, 0x12, 0x16, 0x0a, 0x12, 0x57, 0x4f, 0x52, 0x4b, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49, 0x42, 0x52, 0x45, 0x54, 0x54, 0x4f, 0x10, 0x24, 0x42,
	0x0c, 0x5a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_work_proto_rawDescOnce sync.Once
	file_proto_work_proto_rawDescData = file_proto_work_proto_rawDesc
)

func file_proto_work_proto_rawDescGZIP() []byte {
	file_proto_work_proto_rawDescOnce.Do(func() {
		file_proto_work_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_work_proto_rawDescData)
	})
	return file_proto_work_proto_rawDescData
}

var file_proto_work_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_work_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_work_proto_goTypes = []interface{}{
	(WorkType)(0),                 // 0: WorkType
	(Work_PublishStatus)(0),       // 1: Work.PublishStatus
	(*Work)(nil),                  // 2: Work
	(*Work_SubWork)(nil),          // 3: Work.SubWork
	(*Work_SubWorksResponse)(nil), // 4: Work.SubWorksResponse
}
var file_proto_work_proto_depIdxs = []int32{
	0, // 0: Work.SubWork.work_type:type_name -> WorkType
	1, // 1: Work.SubWork.publish_status:type_name -> Work.PublishStatus
	3, // 2: Work.SubWork.subworks:type_name -> Work.SubWork
	3, // 3: Work.SubWorksResponse.subworks:type_name -> Work.SubWork
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_work_proto_init() }
func file_proto_work_proto_init() {
	if File_proto_work_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_work_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Work); i {
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
		file_proto_work_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Work_SubWork); i {
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
		file_proto_work_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Work_SubWorksResponse); i {
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
			RawDescriptor: file_proto_work_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_work_proto_goTypes,
		DependencyIndexes: file_proto_work_proto_depIdxs,
		EnumInfos:         file_proto_work_proto_enumTypes,
		MessageInfos:      file_proto_work_proto_msgTypes,
	}.Build()
	File_proto_work_proto = out.File
	file_proto_work_proto_rawDesc = nil
	file_proto_work_proto_goTypes = nil
	file_proto_work_proto_depIdxs = nil
}
