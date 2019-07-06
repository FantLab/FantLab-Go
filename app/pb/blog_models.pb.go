// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/blog_models.proto

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

type Blog struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db3de3fa4b90ae2, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

type Blog_LastArticle struct {
	Id                   uint32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	User                 *Common_UserLink     `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Blog_LastArticle) Reset()         { *m = Blog_LastArticle{} }
func (m *Blog_LastArticle) String() string { return proto.CompactTextString(m) }
func (*Blog_LastArticle) ProtoMessage()    {}
func (*Blog_LastArticle) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db3de3fa4b90ae2, []int{0, 0}
}

func (m *Blog_LastArticle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_LastArticle.Unmarshal(m, b)
}
func (m *Blog_LastArticle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_LastArticle.Marshal(b, m, deterministic)
}
func (m *Blog_LastArticle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_LastArticle.Merge(m, src)
}
func (m *Blog_LastArticle) XXX_Size() int {
	return xxx_messageInfo_Blog_LastArticle.Size(m)
}
func (m *Blog_LastArticle) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_LastArticle.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_LastArticle proto.InternalMessageInfo

func (m *Blog_LastArticle) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Blog_LastArticle) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog_LastArticle) GetUser() *Common_UserLink {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Blog_LastArticle) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type Blog_Community struct {
	Id                   uint32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string                `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	CommunityDescription string                `protobuf:"bytes,3,opt,name=community_description,json=communityDescription,proto3" json:"community_description,omitempty"`
	Rules                string                `protobuf:"bytes,4,opt,name=rules,proto3" json:"rules,omitempty"`
	Avatar               string                `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Stats                *Blog_Community_Stats `protobuf:"bytes,6,opt,name=stats,proto3" json:"stats,omitempty"`
	LastArticle          *Blog_LastArticle     `protobuf:"bytes,7,opt,name=last_article,json=lastArticle,proto3" json:"last_article,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Blog_Community) Reset()         { *m = Blog_Community{} }
func (m *Blog_Community) String() string { return proto.CompactTextString(m) }
func (*Blog_Community) ProtoMessage()    {}
func (*Blog_Community) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db3de3fa4b90ae2, []int{0, 1}
}

func (m *Blog_Community) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_Community.Unmarshal(m, b)
}
func (m *Blog_Community) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_Community.Marshal(b, m, deterministic)
}
func (m *Blog_Community) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_Community.Merge(m, src)
}
func (m *Blog_Community) XXX_Size() int {
	return xxx_messageInfo_Blog_Community.Size(m)
}
func (m *Blog_Community) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_Community.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_Community proto.InternalMessageInfo

func (m *Blog_Community) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Blog_Community) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog_Community) GetCommunityDescription() string {
	if m != nil {
		return m.CommunityDescription
	}
	return ""
}

func (m *Blog_Community) GetRules() string {
	if m != nil {
		return m.Rules
	}
	return ""
}

func (m *Blog_Community) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Blog_Community) GetStats() *Blog_Community_Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Blog_Community) GetLastArticle() *Blog_LastArticle {
	if m != nil {
		return m.LastArticle
	}
	return nil
}

type Blog_Community_Stats struct {
	ArticleCount         uint32   `protobuf:"varint,1,opt,name=article_count,json=articleCount,proto3" json:"article_count,omitempty"`
	SubscriberCount      uint32   `protobuf:"varint,2,opt,name=subscriber_count,json=subscriberCount,proto3" json:"subscriber_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog_Community_Stats) Reset()         { *m = Blog_Community_Stats{} }
func (m *Blog_Community_Stats) String() string { return proto.CompactTextString(m) }
func (*Blog_Community_Stats) ProtoMessage()    {}
func (*Blog_Community_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db3de3fa4b90ae2, []int{0, 1, 0}
}

func (m *Blog_Community_Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_Community_Stats.Unmarshal(m, b)
}
func (m *Blog_Community_Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_Community_Stats.Marshal(b, m, deterministic)
}
func (m *Blog_Community_Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_Community_Stats.Merge(m, src)
}
func (m *Blog_Community_Stats) XXX_Size() int {
	return xxx_messageInfo_Blog_Community_Stats.Size(m)
}
func (m *Blog_Community_Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_Community_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_Community_Stats proto.InternalMessageInfo

func (m *Blog_Community_Stats) GetArticleCount() uint32 {
	if m != nil {
		return m.ArticleCount
	}
	return 0
}

func (m *Blog_Community_Stats) GetSubscriberCount() uint32 {
	if m != nil {
		return m.SubscriberCount
	}
	return 0
}

type Blog_Article struct {
	Id                   uint32              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string              `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Creation             *Common_Creation    `protobuf:"bytes,3,opt,name=creation,proto3" json:"creation,omitempty"`
	Text                 string              `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Tags                 string              `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	Stats                *Blog_Article_Stats `protobuf:"bytes,6,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Blog_Article) Reset()         { *m = Blog_Article{} }
func (m *Blog_Article) String() string { return proto.CompactTextString(m) }
func (*Blog_Article) ProtoMessage()    {}
func (*Blog_Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db3de3fa4b90ae2, []int{0, 2}
}

func (m *Blog_Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_Article.Unmarshal(m, b)
}
func (m *Blog_Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_Article.Marshal(b, m, deterministic)
}
func (m *Blog_Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_Article.Merge(m, src)
}
func (m *Blog_Article) XXX_Size() int {
	return xxx_messageInfo_Blog_Article.Size(m)
}
func (m *Blog_Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_Article proto.InternalMessageInfo

func (m *Blog_Article) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Blog_Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog_Article) GetCreation() *Common_Creation {
	if m != nil {
		return m.Creation
	}
	return nil
}

func (m *Blog_Article) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Blog_Article) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *Blog_Article) GetStats() *Blog_Article_Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type Blog_Article_Stats struct {
	LikeCount            uint32   `protobuf:"varint,1,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	ViewCount            uint32   `protobuf:"varint,2,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	CommentCount         uint32   `protobuf:"varint,3,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog_Article_Stats) Reset()         { *m = Blog_Article_Stats{} }
func (m *Blog_Article_Stats) String() string { return proto.CompactTextString(m) }
func (*Blog_Article_Stats) ProtoMessage()    {}
func (*Blog_Article_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db3de3fa4b90ae2, []int{0, 2, 0}
}

func (m *Blog_Article_Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_Article_Stats.Unmarshal(m, b)
}
func (m *Blog_Article_Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_Article_Stats.Marshal(b, m, deterministic)
}
func (m *Blog_Article_Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_Article_Stats.Merge(m, src)
}
func (m *Blog_Article_Stats) XXX_Size() int {
	return xxx_messageInfo_Blog_Article_Stats.Size(m)
}
func (m *Blog_Article_Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_Article_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_Article_Stats proto.InternalMessageInfo

func (m *Blog_Article_Stats) GetLikeCount() uint32 {
	if m != nil {
		return m.LikeCount
	}
	return 0
}

func (m *Blog_Article_Stats) GetViewCount() uint32 {
	if m != nil {
		return m.ViewCount
	}
	return 0
}

func (m *Blog_Article_Stats) GetCommentCount() uint32 {
	if m != nil {
		return m.CommentCount
	}
	return 0
}

type Blog_Blog struct {
	Id                   uint32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 *Common_UserLink  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	IsClosed             bool              `protobuf:"varint,3,opt,name=is_closed,json=isClosed,proto3" json:"is_closed,omitempty"`
	Stats                *Blog_Blog_Stats  `protobuf:"bytes,4,opt,name=stats,proto3" json:"stats,omitempty"`
	LastArticle          *Blog_LastArticle `protobuf:"bytes,5,opt,name=last_article,json=lastArticle,proto3" json:"last_article,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Blog_Blog) Reset()         { *m = Blog_Blog{} }
func (m *Blog_Blog) String() string { return proto.CompactTextString(m) }
func (*Blog_Blog) ProtoMessage()    {}
func (*Blog_Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db3de3fa4b90ae2, []int{0, 3}
}

func (m *Blog_Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_Blog.Unmarshal(m, b)
}
func (m *Blog_Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_Blog.Marshal(b, m, deterministic)
}
func (m *Blog_Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_Blog.Merge(m, src)
}
func (m *Blog_Blog) XXX_Size() int {
	return xxx_messageInfo_Blog_Blog.Size(m)
}
func (m *Blog_Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_Blog proto.InternalMessageInfo

func (m *Blog_Blog) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Blog_Blog) GetUser() *Common_UserLink {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Blog_Blog) GetIsClosed() bool {
	if m != nil {
		return m.IsClosed
	}
	return false
}

func (m *Blog_Blog) GetStats() *Blog_Blog_Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Blog_Blog) GetLastArticle() *Blog_LastArticle {
	if m != nil {
		return m.LastArticle
	}
	return nil
}

type Blog_Blog_Stats struct {
	ArticleCount         uint32   `protobuf:"varint,1,opt,name=article_count,json=articleCount,proto3" json:"article_count,omitempty"`
	SubscriberCount      uint32   `protobuf:"varint,2,opt,name=subscriber_count,json=subscriberCount,proto3" json:"subscriber_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog_Blog_Stats) Reset()         { *m = Blog_Blog_Stats{} }
func (m *Blog_Blog_Stats) String() string { return proto.CompactTextString(m) }
func (*Blog_Blog_Stats) ProtoMessage()    {}
func (*Blog_Blog_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db3de3fa4b90ae2, []int{0, 3, 0}
}

func (m *Blog_Blog_Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_Blog_Stats.Unmarshal(m, b)
}
func (m *Blog_Blog_Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_Blog_Stats.Marshal(b, m, deterministic)
}
func (m *Blog_Blog_Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_Blog_Stats.Merge(m, src)
}
func (m *Blog_Blog_Stats) XXX_Size() int {
	return xxx_messageInfo_Blog_Blog_Stats.Size(m)
}
func (m *Blog_Blog_Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_Blog_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_Blog_Stats proto.InternalMessageInfo

func (m *Blog_Blog_Stats) GetArticleCount() uint32 {
	if m != nil {
		return m.ArticleCount
	}
	return 0
}

func (m *Blog_Blog_Stats) GetSubscriberCount() uint32 {
	if m != nil {
		return m.SubscriberCount
	}
	return 0
}

type Blog_CommunitiesResponse struct {
	Main                 []*Blog_Community `protobuf:"bytes,1,rep,name=main,proto3" json:"main,omitempty"`
	Additional           []*Blog_Community `protobuf:"bytes,2,rep,name=additional,proto3" json:"additional,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Blog_CommunitiesResponse) Reset()         { *m = Blog_CommunitiesResponse{} }
func (m *Blog_CommunitiesResponse) String() string { return proto.CompactTextString(m) }
func (*Blog_CommunitiesResponse) ProtoMessage()    {}
func (*Blog_CommunitiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db3de3fa4b90ae2, []int{0, 4}
}

func (m *Blog_CommunitiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_CommunitiesResponse.Unmarshal(m, b)
}
func (m *Blog_CommunitiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_CommunitiesResponse.Marshal(b, m, deterministic)
}
func (m *Blog_CommunitiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_CommunitiesResponse.Merge(m, src)
}
func (m *Blog_CommunitiesResponse) XXX_Size() int {
	return xxx_messageInfo_Blog_CommunitiesResponse.Size(m)
}
func (m *Blog_CommunitiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_CommunitiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_CommunitiesResponse proto.InternalMessageInfo

func (m *Blog_CommunitiesResponse) GetMain() []*Blog_Community {
	if m != nil {
		return m.Main
	}
	return nil
}

func (m *Blog_CommunitiesResponse) GetAdditional() []*Blog_Community {
	if m != nil {
		return m.Additional
	}
	return nil
}

type Blog_CommunityResponse struct {
	Community            *Blog_Community    `protobuf:"bytes,1,opt,name=community,proto3" json:"community,omitempty"`
	Moderators           []*Common_UserLink `protobuf:"bytes,2,rep,name=moderators,proto3" json:"moderators,omitempty"`
	Authors              []*Common_UserLink `protobuf:"bytes,3,rep,name=authors,proto3" json:"authors,omitempty"`
	Articles             []*Blog_Article    `protobuf:"bytes,4,rep,name=articles,proto3" json:"articles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Blog_CommunityResponse) Reset()         { *m = Blog_CommunityResponse{} }
func (m *Blog_CommunityResponse) String() string { return proto.CompactTextString(m) }
func (*Blog_CommunityResponse) ProtoMessage()    {}
func (*Blog_CommunityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db3de3fa4b90ae2, []int{0, 5}
}

func (m *Blog_CommunityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_CommunityResponse.Unmarshal(m, b)
}
func (m *Blog_CommunityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_CommunityResponse.Marshal(b, m, deterministic)
}
func (m *Blog_CommunityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_CommunityResponse.Merge(m, src)
}
func (m *Blog_CommunityResponse) XXX_Size() int {
	return xxx_messageInfo_Blog_CommunityResponse.Size(m)
}
func (m *Blog_CommunityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_CommunityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_CommunityResponse proto.InternalMessageInfo

func (m *Blog_CommunityResponse) GetCommunity() *Blog_Community {
	if m != nil {
		return m.Community
	}
	return nil
}

func (m *Blog_CommunityResponse) GetModerators() []*Common_UserLink {
	if m != nil {
		return m.Moderators
	}
	return nil
}

func (m *Blog_CommunityResponse) GetAuthors() []*Common_UserLink {
	if m != nil {
		return m.Authors
	}
	return nil
}

func (m *Blog_CommunityResponse) GetArticles() []*Blog_Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

type Blog_BlogsResponse struct {
	Blogs                []*Blog_Blog `protobuf:"bytes,1,rep,name=blogs,proto3" json:"blogs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Blog_BlogsResponse) Reset()         { *m = Blog_BlogsResponse{} }
func (m *Blog_BlogsResponse) String() string { return proto.CompactTextString(m) }
func (*Blog_BlogsResponse) ProtoMessage()    {}
func (*Blog_BlogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db3de3fa4b90ae2, []int{0, 6}
}

func (m *Blog_BlogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_BlogsResponse.Unmarshal(m, b)
}
func (m *Blog_BlogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_BlogsResponse.Marshal(b, m, deterministic)
}
func (m *Blog_BlogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_BlogsResponse.Merge(m, src)
}
func (m *Blog_BlogsResponse) XXX_Size() int {
	return xxx_messageInfo_Blog_BlogsResponse.Size(m)
}
func (m *Blog_BlogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_BlogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_BlogsResponse proto.InternalMessageInfo

func (m *Blog_BlogsResponse) GetBlogs() []*Blog_Blog {
	if m != nil {
		return m.Blogs
	}
	return nil
}

type Blog_BlogResponse struct {
	Articles             []*Blog_Article `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Blog_BlogResponse) Reset()         { *m = Blog_BlogResponse{} }
func (m *Blog_BlogResponse) String() string { return proto.CompactTextString(m) }
func (*Blog_BlogResponse) ProtoMessage()    {}
func (*Blog_BlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db3de3fa4b90ae2, []int{0, 7}
}

func (m *Blog_BlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_BlogResponse.Unmarshal(m, b)
}
func (m *Blog_BlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_BlogResponse.Marshal(b, m, deterministic)
}
func (m *Blog_BlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_BlogResponse.Merge(m, src)
}
func (m *Blog_BlogResponse) XXX_Size() int {
	return xxx_messageInfo_Blog_BlogResponse.Size(m)
}
func (m *Blog_BlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_BlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_BlogResponse proto.InternalMessageInfo

func (m *Blog_BlogResponse) GetArticles() []*Blog_Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "fantlab.blog.Blog")
	proto.RegisterType((*Blog_LastArticle)(nil), "fantlab.blog.Blog.LastArticle")
	proto.RegisterType((*Blog_Community)(nil), "fantlab.blog.Blog.Community")
	proto.RegisterType((*Blog_Community_Stats)(nil), "fantlab.blog.Blog.Community.Stats")
	proto.RegisterType((*Blog_Article)(nil), "fantlab.blog.Blog.Article")
	proto.RegisterType((*Blog_Article_Stats)(nil), "fantlab.blog.Blog.Article.Stats")
	proto.RegisterType((*Blog_Blog)(nil), "fantlab.blog.Blog.Blog")
	proto.RegisterType((*Blog_Blog_Stats)(nil), "fantlab.blog.Blog.Blog.Stats")
	proto.RegisterType((*Blog_CommunitiesResponse)(nil), "fantlab.blog.Blog.CommunitiesResponse")
	proto.RegisterType((*Blog_CommunityResponse)(nil), "fantlab.blog.Blog.CommunityResponse")
	proto.RegisterType((*Blog_BlogsResponse)(nil), "fantlab.blog.Blog.BlogsResponse")
	proto.RegisterType((*Blog_BlogResponse)(nil), "fantlab.blog.Blog.BlogResponse")
}

func init() { proto.RegisterFile("proto/blog_models.proto", fileDescriptor_3db3de3fa4b90ae2) }

var fileDescriptor_3db3de3fa4b90ae2 = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x96, 0x9d, 0x3f, 0x8d, 0x27, 0xc9, 0xef, 0x07, 0x4b, 0xa1, 0xc6, 0x50, 0x88, 0xda, 0x4b,
	0x90, 0xc0, 0x41, 0x2d, 0xaa, 0x2a, 0x84, 0x90, 0xda, 0x20, 0x4e, 0x3d, 0x2d, 0x20, 0x24, 0x2e,
	0xd1, 0xda, 0xde, 0x86, 0xa5, 0x8e, 0xd7, 0xf2, 0xae, 0x0b, 0x3c, 0x00, 0x0f, 0xc0, 0xa3, 0xc0,
	0x63, 0x71, 0x47, 0xe2, 0x88, 0xf6, 0x8f, 0xe3, 0x44, 0xa4, 0x6d, 0xb8, 0x71, 0x89, 0x76, 0x66,
	0xbe, 0x99, 0x7c, 0xfb, 0xcd, 0xcc, 0x1a, 0xb6, 0xf2, 0x82, 0x4b, 0x3e, 0x8a, 0x52, 0x3e, 0x9d,
	0xcc, 0x78, 0x42, 0x53, 0x11, 0x6a, 0x0f, 0xea, 0x9d, 0x92, 0x4c, 0xa6, 0x24, 0x0a, 0x55, 0x28,
	0xb8, 0x6d, 0x60, 0x31, 0x9f, 0xcd, 0x78, 0xb6, 0x04, 0x0c, 0xee, 0x4f, 0x39, 0x9f, 0xa6, 0x74,
	0xa4, 0xad, 0xa8, 0x3c, 0x1d, 0x49, 0x36, 0xa3, 0x42, 0x92, 0x59, 0x6e, 0x00, 0x3b, 0xbf, 0x7a,
	0xd0, 0x3c, 0x56, 0x45, 0xbe, 0x3a, 0xd0, 0x3d, 0x21, 0x42, 0x1e, 0x15, 0x92, 0xc5, 0x29, 0x45,
	0xff, 0x81, 0xcb, 0x12, 0xdf, 0x19, 0x38, 0xc3, 0x3e, 0x76, 0x59, 0x82, 0x36, 0xa1, 0x25, 0x99,
	0x4c, 0xa9, 0xef, 0x0e, 0x9c, 0xa1, 0x87, 0x8d, 0x81, 0x1e, 0x42, 0xb3, 0x14, 0xb4, 0xf0, 0x1b,
	0x03, 0x67, 0xd8, 0xdd, 0xf3, 0xc3, 0x8a, 0xd7, 0x58, 0x73, 0x09, 0xdf, 0x08, 0x5a, 0x9c, 0xb0,
	0xec, 0x0c, 0x6b, 0x14, 0x0a, 0xa1, 0x99, 0x10, 0x49, 0xfd, 0xa6, 0x46, 0x07, 0xa1, 0x21, 0x17,
	0x56, 0xe4, 0xc2, 0xd7, 0x15, 0x39, 0xac, 0x71, 0xc1, 0x0f, 0x17, 0x3c, 0x55, 0xa9, 0xcc, 0x98,
	0xfc, 0xbc, 0x26, 0xa3, 0x7d, 0xb8, 0x19, 0x57, 0x29, 0x93, 0x84, 0x8a, 0xb8, 0x60, 0xb9, 0x64,
	0x3c, 0xd3, 0x14, 0x3d, 0xbc, 0x39, 0x0f, 0xbe, 0xa8, 0x63, 0xaa, 0x54, 0x51, 0xa6, 0x54, 0x68,
	0x66, 0x1e, 0x36, 0x06, 0xba, 0x05, 0x6d, 0x72, 0x4e, 0x24, 0x29, 0xfc, 0x96, 0x76, 0x5b, 0x0b,
	0x1d, 0x42, 0x4b, 0x48, 0x22, 0x85, 0xdf, 0xd6, 0xf7, 0xd8, 0x09, 0x17, 0xbb, 0x11, 0x2a, 0x35,
	0xc3, 0x39, 0xeb, 0xf0, 0x95, 0x42, 0x62, 0x93, 0x80, 0x8e, 0xa0, 0x97, 0x12, 0x21, 0x27, 0xc4,
	0x88, 0xec, 0x6f, 0xe8, 0x02, 0xf7, 0x56, 0x14, 0x58, 0x68, 0x05, 0xee, 0xa6, 0xb5, 0x11, 0xbc,
	0x85, 0x96, 0x2e, 0x89, 0x76, 0xa1, 0x6f, 0xcb, 0x4c, 0x62, 0x5e, 0x66, 0xd2, 0x2a, 0xd3, 0xb3,
	0xce, 0xb1, 0xf2, 0xa1, 0x07, 0x70, 0x4d, 0x94, 0x91, 0xba, 0x68, 0x44, 0x0b, 0x8b, 0x73, 0x35,
	0xee, 0xff, 0xda, 0xaf, 0xa1, 0xc1, 0x37, 0x17, 0x36, 0xfe, 0xae, 0xf9, 0x4f, 0xa0, 0x13, 0x17,
	0x94, 0xcc, 0xd5, 0x5d, 0x31, 0x00, 0x63, 0x1b, 0xc7, 0x73, 0x24, 0x42, 0xd0, 0x94, 0xf4, 0x93,
	0xb4, 0x52, 0xeb, 0xb3, 0xf6, 0x91, 0xa9, 0xb0, 0x3a, 0xeb, 0x33, 0x3a, 0x58, 0x56, 0x79, 0xb0,
	0x42, 0x24, 0x4b, 0x77, 0x49, 0xe3, 0xe0, 0x43, 0x25, 0xd0, 0x36, 0x40, 0xca, 0xce, 0x96, 0xd5,
	0xf1, 0x94, 0xc7, 0x48, 0xb3, 0x0d, 0x70, 0xce, 0xe8, 0xc7, 0x25, 0x51, 0x3c, 0xe5, 0x31, 0xe1,
	0x5d, 0xe8, 0xab, 0x51, 0xa1, 0x99, 0xb4, 0x88, 0x86, 0x91, 0xd7, 0x3a, 0x8d, 0x66, 0xdf, 0x5d,
	0xb3, 0x3d, 0x7f, 0x08, 0x56, 0xed, 0x85, 0xbb, 0xd6, 0x5e, 0xdc, 0x01, 0x8f, 0x89, 0x49, 0x9c,
	0x72, 0x41, 0x13, 0xfd, 0x3f, 0x1d, 0xdc, 0x61, 0x62, 0xac, 0x6d, 0xb4, 0x5f, 0xe9, 0x60, 0xb6,
	0x66, 0x7b, 0x85, 0x0e, 0xfa, 0xe7, 0xd2, 0x41, 0x6b, 0xfd, 0x43, 0x83, 0xf6, 0xc5, 0x81, 0x1b,
	0xd5, 0x7e, 0x30, 0x2a, 0x30, 0x15, 0x39, 0xcf, 0x04, 0x45, 0x8f, 0xa1, 0x39, 0x23, 0x2c, 0xf3,
	0x9d, 0x41, 0x63, 0xd8, 0xdd, 0xbb, 0x7b, 0xd9, 0x56, 0x61, 0x8d, 0x44, 0xcf, 0x00, 0x48, 0x92,
	0x30, 0x35, 0x56, 0x24, 0xf5, 0xdd, 0x35, 0xf2, 0x16, 0xf0, 0xc1, 0x4f, 0x07, 0xae, 0xd7, 0x91,
	0x8a, 0xc5, 0x53, 0xf0, 0xe6, 0x4f, 0x84, 0xbe, 0xe9, 0x55, 0x25, 0x6b, 0x38, 0x3a, 0x04, 0x50,
	0xaf, 0x6f, 0x41, 0x24, 0x2f, 0x84, 0xe5, 0x73, 0x71, 0xef, 0x17, 0xb0, 0x68, 0x0f, 0x36, 0x48,
	0x29, 0xdf, 0xab, 0xb4, 0xc6, 0x15, 0x69, 0x15, 0x10, 0x1d, 0x40, 0xc7, 0xb6, 0x40, 0xcd, 0x46,
	0x43, 0xbf, 0xa8, 0x17, 0xee, 0x08, 0x9e, 0x63, 0x83, 0xe7, 0xd0, 0x57, 0x91, 0x5a, 0xf8, 0x47,
	0xd0, 0x52, 0x78, 0x61, 0x95, 0xdf, 0xba, 0x60, 0xc2, 0xb0, 0x41, 0x05, 0x2f, 0xa1, 0xa7, 0xcd,
	0x2a, 0x7d, 0x91, 0x87, 0xb3, 0x3e, 0x8f, 0xe3, 0xce, 0xbb, 0x36, 0xc9, 0xf3, 0x51, 0x1e, 0x45,
	0x6d, 0xfd, 0x05, 0xd8, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x46, 0x57, 0xb1, 0xf0, 0x06,
	0x00, 0x00,
}
