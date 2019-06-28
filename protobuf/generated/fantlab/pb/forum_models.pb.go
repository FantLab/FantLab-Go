// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema/forum_models.proto

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

type Forum_Topic_Type int32

const (
	Forum_Topic_unknown Forum_Topic_Type = 0
	Forum_Topic_topic   Forum_Topic_Type = 1
	Forum_Topic_poll    Forum_Topic_Type = 2
)

var Forum_Topic_Type_name = map[int32]string{
	0: "unknown",
	1: "topic",
	2: "poll",
}

var Forum_Topic_Type_value = map[string]int32{
	"unknown": 0,
	"topic":   1,
	"poll":    2,
}

func (x Forum_Topic_Type) String() string {
	return proto.EnumName(Forum_Topic_Type_name, int32(x))
}

func (Forum_Topic_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 5, 0}
}

type Forum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Forum) Reset()         { *m = Forum{} }
func (m *Forum) String() string { return proto.CompactTextString(m) }
func (*Forum) ProtoMessage()    {}
func (*Forum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0}
}

func (m *Forum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum.Unmarshal(m, b)
}
func (m *Forum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum.Marshal(b, m, deterministic)
}
func (m *Forum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum.Merge(m, src)
}
func (m *Forum) XXX_Size() int {
	return xxx_messageInfo_Forum.Size(m)
}
func (m *Forum) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum.DiscardUnknown(m)
}

var xxx_messageInfo_Forum proto.InternalMessageInfo

type Forum_UserLink struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Login                string   `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Gender               Gender   `protobuf:"varint,3,opt,name=gender,proto3,enum=fantlab.Gender" json:"gender,omitempty"`
	Avatar               string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Class                uint32   `protobuf:"varint,5,opt,name=class,proto3" json:"class,omitempty"`
	Sign                 string   `protobuf:"bytes,6,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Forum_UserLink) Reset()         { *m = Forum_UserLink{} }
func (m *Forum_UserLink) String() string { return proto.CompactTextString(m) }
func (*Forum_UserLink) ProtoMessage()    {}
func (*Forum_UserLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 0}
}

func (m *Forum_UserLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_UserLink.Unmarshal(m, b)
}
func (m *Forum_UserLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_UserLink.Marshal(b, m, deterministic)
}
func (m *Forum_UserLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_UserLink.Merge(m, src)
}
func (m *Forum_UserLink) XXX_Size() int {
	return xxx_messageInfo_Forum_UserLink.Size(m)
}
func (m *Forum_UserLink) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_UserLink.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_UserLink proto.InternalMessageInfo

func (m *Forum_UserLink) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Forum_UserLink) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *Forum_UserLink) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_unknown
}

func (m *Forum_UserLink) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Forum_UserLink) GetClass() uint32 {
	if m != nil {
		return m.Class
	}
	return 0
}

func (m *Forum_UserLink) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

type Forum_Creation struct {
	User                 *Forum_UserLink      `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Forum_Creation) Reset()         { *m = Forum_Creation{} }
func (m *Forum_Creation) String() string { return proto.CompactTextString(m) }
func (*Forum_Creation) ProtoMessage()    {}
func (*Forum_Creation) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 1}
}

func (m *Forum_Creation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_Creation.Unmarshal(m, b)
}
func (m *Forum_Creation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_Creation.Marshal(b, m, deterministic)
}
func (m *Forum_Creation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_Creation.Merge(m, src)
}
func (m *Forum_Creation) XXX_Size() int {
	return xxx_messageInfo_Forum_Creation.Size(m)
}
func (m *Forum_Creation) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_Creation.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_Creation proto.InternalMessageInfo

func (m *Forum_Creation) GetUser() *Forum_UserLink {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Forum_Creation) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type Forum_TopicLink struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Forum_TopicLink) Reset()         { *m = Forum_TopicLink{} }
func (m *Forum_TopicLink) String() string { return proto.CompactTextString(m) }
func (*Forum_TopicLink) ProtoMessage()    {}
func (*Forum_TopicLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 2}
}

func (m *Forum_TopicLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_TopicLink.Unmarshal(m, b)
}
func (m *Forum_TopicLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_TopicLink.Marshal(b, m, deterministic)
}
func (m *Forum_TopicLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_TopicLink.Merge(m, src)
}
func (m *Forum_TopicLink) XXX_Size() int {
	return xxx_messageInfo_Forum_TopicLink.Size(m)
}
func (m *Forum_TopicLink) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_TopicLink.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_TopicLink proto.InternalMessageInfo

func (m *Forum_TopicLink) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Forum_TopicLink) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type Forum_LastMessage struct {
	Id                   uint32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic                *Forum_TopicLink     `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	User                 *Forum_UserLink      `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Forum_LastMessage) Reset()         { *m = Forum_LastMessage{} }
func (m *Forum_LastMessage) String() string { return proto.CompactTextString(m) }
func (*Forum_LastMessage) ProtoMessage()    {}
func (*Forum_LastMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 3}
}

func (m *Forum_LastMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_LastMessage.Unmarshal(m, b)
}
func (m *Forum_LastMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_LastMessage.Marshal(b, m, deterministic)
}
func (m *Forum_LastMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_LastMessage.Merge(m, src)
}
func (m *Forum_LastMessage) XXX_Size() int {
	return xxx_messageInfo_Forum_LastMessage.Size(m)
}
func (m *Forum_LastMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_LastMessage.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_LastMessage proto.InternalMessageInfo

func (m *Forum_LastMessage) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Forum_LastMessage) GetTopic() *Forum_TopicLink {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *Forum_LastMessage) GetUser() *Forum_UserLink {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Forum_LastMessage) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type Forum_TopicMessage struct {
	Id                   uint32                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creation             *Forum_Creation           `protobuf:"bytes,2,opt,name=creation,proto3" json:"creation,omitempty"`
	Text                 string                    `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	IsCensored           bool                      `protobuf:"varint,4,opt,name=is_censored,json=isCensored,proto3" json:"is_censored,omitempty"`
	IsModerTagWorks      bool                      `protobuf:"varint,5,opt,name=is_moder_tag_works,json=isModerTagWorks,proto3" json:"is_moder_tag_works,omitempty"`
	Stats                *Forum_TopicMessage_Stats `protobuf:"bytes,6,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Forum_TopicMessage) Reset()         { *m = Forum_TopicMessage{} }
func (m *Forum_TopicMessage) String() string { return proto.CompactTextString(m) }
func (*Forum_TopicMessage) ProtoMessage()    {}
func (*Forum_TopicMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 4}
}

func (m *Forum_TopicMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_TopicMessage.Unmarshal(m, b)
}
func (m *Forum_TopicMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_TopicMessage.Marshal(b, m, deterministic)
}
func (m *Forum_TopicMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_TopicMessage.Merge(m, src)
}
func (m *Forum_TopicMessage) XXX_Size() int {
	return xxx_messageInfo_Forum_TopicMessage.Size(m)
}
func (m *Forum_TopicMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_TopicMessage.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_TopicMessage proto.InternalMessageInfo

func (m *Forum_TopicMessage) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Forum_TopicMessage) GetCreation() *Forum_Creation {
	if m != nil {
		return m.Creation
	}
	return nil
}

func (m *Forum_TopicMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Forum_TopicMessage) GetIsCensored() bool {
	if m != nil {
		return m.IsCensored
	}
	return false
}

func (m *Forum_TopicMessage) GetIsModerTagWorks() bool {
	if m != nil {
		return m.IsModerTagWorks
	}
	return false
}

func (m *Forum_TopicMessage) GetStats() *Forum_TopicMessage_Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type Forum_TopicMessage_Stats struct {
	PlusCount            uint32   `protobuf:"varint,1,opt,name=plus_count,json=plusCount,proto3" json:"plus_count,omitempty"`
	MinusCount           uint32   `protobuf:"varint,2,opt,name=minus_count,json=minusCount,proto3" json:"minus_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Forum_TopicMessage_Stats) Reset()         { *m = Forum_TopicMessage_Stats{} }
func (m *Forum_TopicMessage_Stats) String() string { return proto.CompactTextString(m) }
func (*Forum_TopicMessage_Stats) ProtoMessage()    {}
func (*Forum_TopicMessage_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 4, 0}
}

func (m *Forum_TopicMessage_Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_TopicMessage_Stats.Unmarshal(m, b)
}
func (m *Forum_TopicMessage_Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_TopicMessage_Stats.Marshal(b, m, deterministic)
}
func (m *Forum_TopicMessage_Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_TopicMessage_Stats.Merge(m, src)
}
func (m *Forum_TopicMessage_Stats) XXX_Size() int {
	return xxx_messageInfo_Forum_TopicMessage_Stats.Size(m)
}
func (m *Forum_TopicMessage_Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_TopicMessage_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_TopicMessage_Stats proto.InternalMessageInfo

func (m *Forum_TopicMessage_Stats) GetPlusCount() uint32 {
	if m != nil {
		return m.PlusCount
	}
	return 0
}

func (m *Forum_TopicMessage_Stats) GetMinusCount() uint32 {
	if m != nil {
		return m.MinusCount
	}
	return 0
}

type Forum_Topic struct {
	Id                   uint32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string             `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	TopicType            Forum_Topic_Type   `protobuf:"varint,3,opt,name=topic_type,json=topicType,proto3,enum=fantlab.forum.Forum_Topic_Type" json:"topic_type,omitempty"`
	Creation             *Forum_Creation    `protobuf:"bytes,4,opt,name=creation,proto3" json:"creation,omitempty"`
	IsClosed             bool               `protobuf:"varint,5,opt,name=is_closed,json=isClosed,proto3" json:"is_closed,omitempty"`
	IsPinned             bool               `protobuf:"varint,6,opt,name=is_pinned,json=isPinned,proto3" json:"is_pinned,omitempty"`
	Stats                *Forum_Topic_Stats `protobuf:"bytes,7,opt,name=stats,proto3" json:"stats,omitempty"`
	LastMessage          *Forum_LastMessage `protobuf:"bytes,8,opt,name=last_message,json=lastMessage,proto3" json:"last_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Forum_Topic) Reset()         { *m = Forum_Topic{} }
func (m *Forum_Topic) String() string { return proto.CompactTextString(m) }
func (*Forum_Topic) ProtoMessage()    {}
func (*Forum_Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 5}
}

func (m *Forum_Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_Topic.Unmarshal(m, b)
}
func (m *Forum_Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_Topic.Marshal(b, m, deterministic)
}
func (m *Forum_Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_Topic.Merge(m, src)
}
func (m *Forum_Topic) XXX_Size() int {
	return xxx_messageInfo_Forum_Topic.Size(m)
}
func (m *Forum_Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_Topic proto.InternalMessageInfo

func (m *Forum_Topic) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Forum_Topic) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Forum_Topic) GetTopicType() Forum_Topic_Type {
	if m != nil {
		return m.TopicType
	}
	return Forum_Topic_unknown
}

func (m *Forum_Topic) GetCreation() *Forum_Creation {
	if m != nil {
		return m.Creation
	}
	return nil
}

func (m *Forum_Topic) GetIsClosed() bool {
	if m != nil {
		return m.IsClosed
	}
	return false
}

func (m *Forum_Topic) GetIsPinned() bool {
	if m != nil {
		return m.IsPinned
	}
	return false
}

func (m *Forum_Topic) GetStats() *Forum_Topic_Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Forum_Topic) GetLastMessage() *Forum_LastMessage {
	if m != nil {
		return m.LastMessage
	}
	return nil
}

type Forum_Topic_Stats struct {
	MessageCount         uint32   `protobuf:"varint,1,opt,name=message_count,json=messageCount,proto3" json:"message_count,omitempty"`
	ViewsCount           uint32   `protobuf:"varint,2,opt,name=views_count,json=viewsCount,proto3" json:"views_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Forum_Topic_Stats) Reset()         { *m = Forum_Topic_Stats{} }
func (m *Forum_Topic_Stats) String() string { return proto.CompactTextString(m) }
func (*Forum_Topic_Stats) ProtoMessage()    {}
func (*Forum_Topic_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 5, 0}
}

func (m *Forum_Topic_Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_Topic_Stats.Unmarshal(m, b)
}
func (m *Forum_Topic_Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_Topic_Stats.Marshal(b, m, deterministic)
}
func (m *Forum_Topic_Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_Topic_Stats.Merge(m, src)
}
func (m *Forum_Topic_Stats) XXX_Size() int {
	return xxx_messageInfo_Forum_Topic_Stats.Size(m)
}
func (m *Forum_Topic_Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_Topic_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_Topic_Stats proto.InternalMessageInfo

func (m *Forum_Topic_Stats) GetMessageCount() uint32 {
	if m != nil {
		return m.MessageCount
	}
	return 0
}

func (m *Forum_Topic_Stats) GetViewsCount() uint32 {
	if m != nil {
		return m.ViewsCount
	}
	return 0
}

type Forum_Forum struct {
	Id                   uint32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string             `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string             `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Moderators           []*Forum_UserLink  `protobuf:"bytes,4,rep,name=moderators,proto3" json:"moderators,omitempty"`
	Stats                *Forum_Forum_Stats `protobuf:"bytes,5,opt,name=stats,proto3" json:"stats,omitempty"`
	LastMessage          *Forum_LastMessage `protobuf:"bytes,6,opt,name=last_message,json=lastMessage,proto3" json:"last_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Forum_Forum) Reset()         { *m = Forum_Forum{} }
func (m *Forum_Forum) String() string { return proto.CompactTextString(m) }
func (*Forum_Forum) ProtoMessage()    {}
func (*Forum_Forum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 6}
}

func (m *Forum_Forum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_Forum.Unmarshal(m, b)
}
func (m *Forum_Forum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_Forum.Marshal(b, m, deterministic)
}
func (m *Forum_Forum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_Forum.Merge(m, src)
}
func (m *Forum_Forum) XXX_Size() int {
	return xxx_messageInfo_Forum_Forum.Size(m)
}
func (m *Forum_Forum) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_Forum.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_Forum proto.InternalMessageInfo

func (m *Forum_Forum) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Forum_Forum) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Forum_Forum) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Forum_Forum) GetModerators() []*Forum_UserLink {
	if m != nil {
		return m.Moderators
	}
	return nil
}

func (m *Forum_Forum) GetStats() *Forum_Forum_Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Forum_Forum) GetLastMessage() *Forum_LastMessage {
	if m != nil {
		return m.LastMessage
	}
	return nil
}

type Forum_Forum_Stats struct {
	TopicCount           uint32   `protobuf:"varint,1,opt,name=topic_count,json=topicCount,proto3" json:"topic_count,omitempty"`
	MessageCount         uint32   `protobuf:"varint,2,opt,name=message_count,json=messageCount,proto3" json:"message_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Forum_Forum_Stats) Reset()         { *m = Forum_Forum_Stats{} }
func (m *Forum_Forum_Stats) String() string { return proto.CompactTextString(m) }
func (*Forum_Forum_Stats) ProtoMessage()    {}
func (*Forum_Forum_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 6, 0}
}

func (m *Forum_Forum_Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_Forum_Stats.Unmarshal(m, b)
}
func (m *Forum_Forum_Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_Forum_Stats.Marshal(b, m, deterministic)
}
func (m *Forum_Forum_Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_Forum_Stats.Merge(m, src)
}
func (m *Forum_Forum_Stats) XXX_Size() int {
	return xxx_messageInfo_Forum_Forum_Stats.Size(m)
}
func (m *Forum_Forum_Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_Forum_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_Forum_Stats proto.InternalMessageInfo

func (m *Forum_Forum_Stats) GetTopicCount() uint32 {
	if m != nil {
		return m.TopicCount
	}
	return 0
}

func (m *Forum_Forum_Stats) GetMessageCount() uint32 {
	if m != nil {
		return m.MessageCount
	}
	return 0
}

type Forum_ForumBlock struct {
	// hack to omit field (https://github.com/gogo/protobuf/blob/dadb625850898f31a8e40e83492f4a7132e520a2/jsonpb/jsonpb.go#L286)
	XXX_Id               uint32         `protobuf:"varint,1,opt,name=XXX__id,json=XXXId,proto3" json:"XXX__id,omitempty"`
	Title                string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Forums               []*Forum_Forum `protobuf:"bytes,3,rep,name=forums,proto3" json:"forums,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Forum_ForumBlock) Reset()         { *m = Forum_ForumBlock{} }
func (m *Forum_ForumBlock) String() string { return proto.CompactTextString(m) }
func (*Forum_ForumBlock) ProtoMessage()    {}
func (*Forum_ForumBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 7}
}

func (m *Forum_ForumBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_ForumBlock.Unmarshal(m, b)
}
func (m *Forum_ForumBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_ForumBlock.Marshal(b, m, deterministic)
}
func (m *Forum_ForumBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_ForumBlock.Merge(m, src)
}
func (m *Forum_ForumBlock) XXX_Size() int {
	return xxx_messageInfo_Forum_ForumBlock.Size(m)
}
func (m *Forum_ForumBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_ForumBlock.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_ForumBlock proto.InternalMessageInfo

func (m *Forum_ForumBlock) GetXXX_Id() uint32 {
	if m != nil {
		return m.XXX_Id
	}
	return 0
}

func (m *Forum_ForumBlock) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Forum_ForumBlock) GetForums() []*Forum_Forum {
	if m != nil {
		return m.Forums
	}
	return nil
}

type Forum_ForumBlocksResponse struct {
	ForumBlocks          []*Forum_ForumBlock `protobuf:"bytes,1,rep,name=forum_blocks,json=forumBlocks,proto3" json:"forum_blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Forum_ForumBlocksResponse) Reset()         { *m = Forum_ForumBlocksResponse{} }
func (m *Forum_ForumBlocksResponse) String() string { return proto.CompactTextString(m) }
func (*Forum_ForumBlocksResponse) ProtoMessage()    {}
func (*Forum_ForumBlocksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 8}
}

func (m *Forum_ForumBlocksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_ForumBlocksResponse.Unmarshal(m, b)
}
func (m *Forum_ForumBlocksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_ForumBlocksResponse.Marshal(b, m, deterministic)
}
func (m *Forum_ForumBlocksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_ForumBlocksResponse.Merge(m, src)
}
func (m *Forum_ForumBlocksResponse) XXX_Size() int {
	return xxx_messageInfo_Forum_ForumBlocksResponse.Size(m)
}
func (m *Forum_ForumBlocksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_ForumBlocksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_ForumBlocksResponse proto.InternalMessageInfo

func (m *Forum_ForumBlocksResponse) GetForumBlocks() []*Forum_ForumBlock {
	if m != nil {
		return m.ForumBlocks
	}
	return nil
}

type Forum_ForumTopicsResponse struct {
	Topics               []*Forum_Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Forum_ForumTopicsResponse) Reset()         { *m = Forum_ForumTopicsResponse{} }
func (m *Forum_ForumTopicsResponse) String() string { return proto.CompactTextString(m) }
func (*Forum_ForumTopicsResponse) ProtoMessage()    {}
func (*Forum_ForumTopicsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 9}
}

func (m *Forum_ForumTopicsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_ForumTopicsResponse.Unmarshal(m, b)
}
func (m *Forum_ForumTopicsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_ForumTopicsResponse.Marshal(b, m, deterministic)
}
func (m *Forum_ForumTopicsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_ForumTopicsResponse.Merge(m, src)
}
func (m *Forum_ForumTopicsResponse) XXX_Size() int {
	return xxx_messageInfo_Forum_ForumTopicsResponse.Size(m)
}
func (m *Forum_ForumTopicsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_ForumTopicsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_ForumTopicsResponse proto.InternalMessageInfo

func (m *Forum_ForumTopicsResponse) GetTopics() []*Forum_Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

type Forum_TopicMessagesResponse struct {
	Messages             []*Forum_TopicMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Forum_TopicMessagesResponse) Reset()         { *m = Forum_TopicMessagesResponse{} }
func (m *Forum_TopicMessagesResponse) String() string { return proto.CompactTextString(m) }
func (*Forum_TopicMessagesResponse) ProtoMessage()    {}
func (*Forum_TopicMessagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb25d04eaceb49e, []int{0, 10}
}

func (m *Forum_TopicMessagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forum_TopicMessagesResponse.Unmarshal(m, b)
}
func (m *Forum_TopicMessagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forum_TopicMessagesResponse.Marshal(b, m, deterministic)
}
func (m *Forum_TopicMessagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forum_TopicMessagesResponse.Merge(m, src)
}
func (m *Forum_TopicMessagesResponse) XXX_Size() int {
	return xxx_messageInfo_Forum_TopicMessagesResponse.Size(m)
}
func (m *Forum_TopicMessagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Forum_TopicMessagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Forum_TopicMessagesResponse proto.InternalMessageInfo

func (m *Forum_TopicMessagesResponse) GetMessages() []*Forum_TopicMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterEnum("fantlab.forum.Forum_Topic_Type", Forum_Topic_Type_name, Forum_Topic_Type_value)
	proto.RegisterType((*Forum)(nil), "fantlab.forum.Forum")
	proto.RegisterType((*Forum_UserLink)(nil), "fantlab.forum.Forum.UserLink")
	proto.RegisterType((*Forum_Creation)(nil), "fantlab.forum.Forum.Creation")
	proto.RegisterType((*Forum_TopicLink)(nil), "fantlab.forum.Forum.TopicLink")
	proto.RegisterType((*Forum_LastMessage)(nil), "fantlab.forum.Forum.LastMessage")
	proto.RegisterType((*Forum_TopicMessage)(nil), "fantlab.forum.Forum.TopicMessage")
	proto.RegisterType((*Forum_TopicMessage_Stats)(nil), "fantlab.forum.Forum.TopicMessage.Stats")
	proto.RegisterType((*Forum_Topic)(nil), "fantlab.forum.Forum.Topic")
	proto.RegisterType((*Forum_Topic_Stats)(nil), "fantlab.forum.Forum.Topic.Stats")
	proto.RegisterType((*Forum_Forum)(nil), "fantlab.forum.Forum.Forum")
	proto.RegisterType((*Forum_Forum_Stats)(nil), "fantlab.forum.Forum.Forum.Stats")
	proto.RegisterType((*Forum_ForumBlock)(nil), "fantlab.forum.Forum.ForumBlock")
	proto.RegisterType((*Forum_ForumBlocksResponse)(nil), "fantlab.forum.Forum.ForumBlocksResponse")
	proto.RegisterType((*Forum_ForumTopicsResponse)(nil), "fantlab.forum.Forum.ForumTopicsResponse")
	proto.RegisterType((*Forum_TopicMessagesResponse)(nil), "fantlab.forum.Forum.TopicMessagesResponse")
}

func init() { proto.RegisterFile("schema/forum_models.proto", fileDescriptor_0bb25d04eaceb49e) }

var fileDescriptor_0bb25d04eaceb49e = []byte{
	// 828 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xd1, 0x6e, 0xe3, 0x44,
	0x14, 0x25, 0x89, 0x9d, 0x3a, 0xd7, 0xe9, 0xb6, 0x1a, 0x60, 0x65, 0x06, 0x2d, 0x09, 0xcb, 0xc3,
	0x56, 0x42, 0x72, 0xb5, 0x01, 0x21, 0xf1, 0x50, 0x1e, 0x1a, 0x89, 0xd5, 0x4a, 0x5b, 0x09, 0x0d,
	0x05, 0x02, 0x2f, 0xd6, 0xc4, 0x9e, 0x86, 0x51, 0xed, 0x19, 0xcb, 0x33, 0xd9, 0xb2, 0xdf, 0xc2,
	0x7f, 0xf0, 0x03, 0xfc, 0x04, 0xe2, 0x1b, 0xf8, 0x08, 0xe4, 0xeb, 0x71, 0x92, 0xd2, 0x75, 0xb6,
	0x15, 0x2f, 0x91, 0xef, 0x9d, 0x73, 0xef, 0x9c, 0x39, 0xe7, 0xce, 0x04, 0x3e, 0x32, 0xe9, 0xaf,
	0xa2, 0xe0, 0xa7, 0x57, 0xba, 0x5a, 0x17, 0x49, 0xa1, 0x33, 0x91, 0x9b, 0xb8, 0xac, 0xb4, 0xd5,
	0xe4, 0xf0, 0x8a, 0x2b, 0x9b, 0xf3, 0x65, 0x8c, 0x6b, 0x94, 0x3a, 0x64, 0xaa, 0x8b, 0x42, 0xab,
	0x5b, 0x50, 0x3a, 0x59, 0x69, 0xbd, 0xca, 0xc5, 0x29, 0x46, 0xcb, 0xf5, 0xd5, 0xa9, 0x95, 0x85,
	0x30, 0x96, 0x17, 0x65, 0x03, 0x78, 0xfa, 0xcf, 0x11, 0xf8, 0xdf, 0x62, 0x9b, 0xdf, 0x7b, 0x10,
	0xfc, 0x60, 0x44, 0xf5, 0x4a, 0xaa, 0x6b, 0xf2, 0x08, 0xfa, 0x32, 0x8b, 0x7a, 0xd3, 0xde, 0xc9,
	0x21, 0xeb, 0xcb, 0x8c, 0x7c, 0x00, 0x7e, 0xae, 0x57, 0x52, 0x45, 0xfd, 0x69, 0xef, 0x64, 0xc4,
	0x9a, 0x80, 0x3c, 0x83, 0xe1, 0x4a, 0xa8, 0x4c, 0x54, 0xd1, 0x60, 0xda, 0x3b, 0x79, 0x34, 0x3b,
	0x8a, 0x5b, 0x66, 0x2f, 0x30, 0xcd, 0xdc, 0x32, 0x79, 0x0c, 0x43, 0xfe, 0x9a, 0x5b, 0x5e, 0x45,
	0x1e, 0xd6, 0xbb, 0xa8, 0x6e, 0x9b, 0xe6, 0xdc, 0x98, 0xc8, 0xc7, 0x9d, 0x9a, 0x80, 0x10, 0xf0,
	0x8c, 0x5c, 0xa9, 0x68, 0x88, 0x58, 0xfc, 0xa6, 0x05, 0x04, 0xf3, 0x4a, 0x70, 0x2b, 0xb5, 0x22,
	0xcf, 0xc1, 0x5b, 0x1b, 0x51, 0x21, 0xbd, 0x70, 0xf6, 0x24, 0xbe, 0x25, 0x47, 0x8c, 0xa7, 0x89,
	0xdb, 0x93, 0x30, 0x84, 0x92, 0x18, 0xbc, 0x8c, 0x5b, 0x81, 0xf4, 0xc3, 0x19, 0x8d, 0x1b, 0x59,
	0xe2, 0x56, 0x96, 0xf8, 0xb2, 0x95, 0x85, 0x21, 0x8e, 0x3e, 0x87, 0xd1, 0xa5, 0x2e, 0x65, 0xda,
	0x25, 0x86, 0x95, 0x36, 0x17, 0xad, 0x18, 0x18, 0xd0, 0x3f, 0x7a, 0x10, 0xbe, 0xe2, 0xc6, 0x5e,
	0x08, 0x63, 0xf8, 0x4a, 0xdc, 0xa9, 0xfa, 0x12, 0x7c, 0x5b, 0xb7, 0x74, 0x1c, 0x3e, 0x79, 0x2b,
	0xed, 0xcd, 0xa6, 0xac, 0x01, 0x6f, 0xce, 0x3a, 0x78, 0xf8, 0x59, 0xbd, 0x7b, 0x9e, 0xf5, 0xcf,
	0x3e, 0x8c, 0x71, 0xdf, 0x2e, 0xe6, 0x5f, 0x43, 0x90, 0x3a, 0xed, 0x1d, 0xf9, 0xb7, 0xf3, 0x68,
	0x0d, 0x62, 0x1b, 0x78, 0x6d, 0xa5, 0x15, 0xbf, 0x59, 0xa4, 0x3f, 0x62, 0xf8, 0x4d, 0x26, 0x10,
	0x4a, 0x93, 0xa4, 0x42, 0x19, 0x5d, 0x89, 0x0c, 0x69, 0x06, 0x0c, 0xa4, 0x99, 0xbb, 0x0c, 0xf9,
	0x1c, 0x88, 0x34, 0x38, 0xc7, 0x55, 0x62, 0xf9, 0x2a, 0xb9, 0xd1, 0xd5, 0x75, 0x33, 0x22, 0x01,
	0x3b, 0x92, 0xe6, 0xa2, 0x5e, 0xb8, 0xe4, 0xab, 0x9f, 0xea, 0x34, 0x39, 0x03, 0xdf, 0x58, 0x6e,
	0x0d, 0x4e, 0x4b, 0x38, 0x7b, 0xd6, 0x2d, 0xab, 0x3b, 0x5e, 0xfc, 0x7d, 0x0d, 0x67, 0x4d, 0x15,
	0x7d, 0x01, 0x3e, 0xc6, 0xe4, 0x09, 0x40, 0x99, 0xaf, 0x4d, 0x92, 0xea, 0xb5, 0xb2, 0xee, 0xf0,
	0xa3, 0x3a, 0x33, 0xaf, 0x13, 0x35, 0xe9, 0x42, 0xaa, 0xcd, 0x7a, 0x1f, 0xd7, 0x01, 0x53, 0x08,
	0xa0, 0x7f, 0x0d, 0xc0, 0xc7, 0x6d, 0xee, 0x37, 0x2e, 0xe4, 0x1b, 0x00, 0x74, 0x38, 0xb1, 0x6f,
	0x4a, 0xe1, 0xee, 0xcf, 0xa4, 0x9b, 0x7c, 0x7c, 0xf9, 0xa6, 0x14, 0x6c, 0x84, 0x25, 0xf5, 0xe7,
	0x2d, 0x53, 0xbc, 0x87, 0x99, 0xf2, 0x31, 0x8c, 0x6a, 0x03, 0x72, 0x6d, 0x44, 0xe6, 0x64, 0x0d,
	0xa4, 0x99, 0x63, 0xec, 0x16, 0x4b, 0xa9, 0x94, 0xc8, 0x50, 0x53, 0x5c, 0xfc, 0x0e, 0x63, 0xf2,
	0x55, 0x2b, 0xf6, 0x01, 0xee, 0x38, 0xdd, 0xc3, 0x77, 0x57, 0x65, 0x32, 0x87, 0x71, 0xce, 0x8d,
	0x4d, 0x8a, 0xc6, 0x82, 0x28, 0xd8, 0x53, 0xbe, 0x73, 0x87, 0x58, 0x98, 0x6f, 0x03, 0x7a, 0xd1,
	0x5a, 0xf5, 0x19, 0x1c, 0xba, 0x46, 0xb7, 0xdc, 0x1a, 0xbb, 0xe4, 0xc6, 0xb0, 0xd7, 0x52, 0xdc,
	0xfc, 0xc7, 0x30, 0x4c, 0x21, 0xe0, 0xe9, 0x09, 0x78, 0x28, 0x64, 0x08, 0x07, 0x6b, 0x75, 0xad,
	0xf4, 0x8d, 0x3a, 0x7e, 0x8f, 0x8c, 0xdc, 0x25, 0x3d, 0xee, 0x91, 0x00, 0xbc, 0x52, 0xe7, 0xf9,
	0x71, 0x9f, 0xfe, 0xdd, 0x77, 0x6f, 0xe4, 0x3d, 0xad, 0x9d, 0x42, 0x98, 0x09, 0x93, 0x56, 0xb2,
	0x44, 0x77, 0x9a, 0xd9, 0xdf, 0x4d, 0x91, 0x33, 0x00, 0x1c, 0x6f, 0x6e, 0x75, 0x65, 0x22, 0x6f,
	0x3a, 0x78, 0xf7, 0xdd, 0xde, 0x29, 0xd8, 0xda, 0xe0, 0xef, 0xd1, 0xb1, 0xf9, 0xdd, 0x6b, 0xc3,
	0xf0, 0x7f, 0xd9, 0x30, 0x81, 0xb0, 0x99, 0xe0, 0x5d, 0x13, 0x9a, 0xa1, 0x6e, 0x2c, 0xb8, 0xe3,
	0x53, 0xff, 0xae, 0x4f, 0x54, 0x01, 0xe0, 0x86, 0xe7, 0xb9, 0x4e, 0xaf, 0xc9, 0x63, 0x38, 0x58,
	0x2c, 0x16, 0x49, 0xb2, 0x51, 0xd9, 0x5f, 0x2c, 0x16, 0x2f, 0xbb, 0x84, 0x9e, 0xc1, 0x10, 0x29,
	0x9b, 0x68, 0x80, 0x12, 0xd2, 0x6e, 0x21, 0x98, 0x43, 0xd2, 0x9f, 0xe1, 0xfd, 0xed, 0x7e, 0x86,
	0x09, 0x53, 0x6a, 0x65, 0x04, 0x39, 0x87, 0x71, 0xf3, 0x4f, 0xbb, 0xc4, 0x7c, 0xd4, 0xc3, 0x86,
	0x93, 0xee, 0x86, 0x58, 0xcf, 0xc2, 0xab, 0x6d, 0x2f, 0xfa, 0xd2, 0xb5, 0xc6, 0x0b, 0xb0, 0x6d,
	0x3d, 0x83, 0x21, 0x8a, 0xd2, 0x36, 0xa5, 0xdd, 0xb7, 0x86, 0x39, 0x24, 0xfd, 0x11, 0x3e, 0xdc,
	0x7d, 0xb3, 0xb6, 0xcd, 0xce, 0x20, 0x70, 0xf2, 0xb5, 0xed, 0x3e, 0x7d, 0xe7, 0x8b, 0xc7, 0x36,
	0x25, 0xe7, 0xe3, 0x5f, 0xc0, 0xa1, 0x4f, 0xcb, 0xe5, 0x72, 0x88, 0xff, 0x09, 0x5f, 0xfc, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x93, 0x71, 0x8b, 0xd5, 0x6c, 0x08, 0x00, 0x00,
}
