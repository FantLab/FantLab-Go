// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/blog.proto

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
	return fileDescriptor_fc5203cdc85000bc, []int{0}
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
	Id                   uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
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
	return fileDescriptor_fc5203cdc85000bc, []int{0, 0}
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

func (m *Blog_LastArticle) GetId() uint64 {
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
	Id                   uint64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
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
	return fileDescriptor_fc5203cdc85000bc, []int{0, 1}
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

func (m *Blog_Community) GetId() uint64 {
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
	ArticleCount         uint64   `protobuf:"varint,1,opt,name=article_count,json=articleCount,proto3" json:"article_count,omitempty"`
	SubscriberCount      uint64   `protobuf:"varint,2,opt,name=subscriber_count,json=subscriberCount,proto3" json:"subscriber_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog_Community_Stats) Reset()         { *m = Blog_Community_Stats{} }
func (m *Blog_Community_Stats) String() string { return proto.CompactTextString(m) }
func (*Blog_Community_Stats) ProtoMessage()    {}
func (*Blog_Community_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{0, 1, 0}
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

func (m *Blog_Community_Stats) GetArticleCount() uint64 {
	if m != nil {
		return m.ArticleCount
	}
	return 0
}

func (m *Blog_Community_Stats) GetSubscriberCount() uint64 {
	if m != nil {
		return m.SubscriberCount
	}
	return 0
}

type Blog_Article struct {
	Id                   uint64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
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
	return fileDescriptor_fc5203cdc85000bc, []int{0, 2}
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

func (m *Blog_Article) GetId() uint64 {
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
	LikeCount            uint64   `protobuf:"varint,1,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	ViewCount            uint64   `protobuf:"varint,2,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	CommentCount         uint64   `protobuf:"varint,3,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog_Article_Stats) Reset()         { *m = Blog_Article_Stats{} }
func (m *Blog_Article_Stats) String() string { return proto.CompactTextString(m) }
func (*Blog_Article_Stats) ProtoMessage()    {}
func (*Blog_Article_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{0, 2, 0}
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

func (m *Blog_Article_Stats) GetLikeCount() uint64 {
	if m != nil {
		return m.LikeCount
	}
	return 0
}

func (m *Blog_Article_Stats) GetViewCount() uint64 {
	if m != nil {
		return m.ViewCount
	}
	return 0
}

func (m *Blog_Article_Stats) GetCommentCount() uint64 {
	if m != nil {
		return m.CommentCount
	}
	return 0
}

type Blog_Blog struct {
	Id                   uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
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
	return fileDescriptor_fc5203cdc85000bc, []int{0, 3}
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

func (m *Blog_Blog) GetId() uint64 {
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
	ArticleCount         uint64   `protobuf:"varint,1,opt,name=article_count,json=articleCount,proto3" json:"article_count,omitempty"`
	SubscriberCount      uint64   `protobuf:"varint,2,opt,name=subscriber_count,json=subscriberCount,proto3" json:"subscriber_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog_Blog_Stats) Reset()         { *m = Blog_Blog_Stats{} }
func (m *Blog_Blog_Stats) String() string { return proto.CompactTextString(m) }
func (*Blog_Blog_Stats) ProtoMessage()    {}
func (*Blog_Blog_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{0, 3, 0}
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

func (m *Blog_Blog_Stats) GetArticleCount() uint64 {
	if m != nil {
		return m.ArticleCount
	}
	return 0
}

func (m *Blog_Blog_Stats) GetSubscriberCount() uint64 {
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
	return fileDescriptor_fc5203cdc85000bc, []int{0, 4}
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
	Pages                *Common_Pages      `protobuf:"bytes,5,opt,name=pages,proto3" json:"pages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Blog_CommunityResponse) Reset()         { *m = Blog_CommunityResponse{} }
func (m *Blog_CommunityResponse) String() string { return proto.CompactTextString(m) }
func (*Blog_CommunityResponse) ProtoMessage()    {}
func (*Blog_CommunityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{0, 5}
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

func (m *Blog_CommunityResponse) GetPages() *Common_Pages {
	if m != nil {
		return m.Pages
	}
	return nil
}

type Blog_BlogsResponse struct {
	Blogs                []*Blog_Blog  `protobuf:"bytes,1,rep,name=blogs,proto3" json:"blogs,omitempty"`
	Pages                *Common_Pages `protobuf:"bytes,2,opt,name=pages,proto3" json:"pages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Blog_BlogsResponse) Reset()         { *m = Blog_BlogsResponse{} }
func (m *Blog_BlogsResponse) String() string { return proto.CompactTextString(m) }
func (*Blog_BlogsResponse) ProtoMessage()    {}
func (*Blog_BlogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{0, 6}
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

func (m *Blog_BlogsResponse) GetPages() *Common_Pages {
	if m != nil {
		return m.Pages
	}
	return nil
}

type Blog_BlogResponse struct {
	Articles             []*Blog_Article `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
	Pages                *Common_Pages   `protobuf:"bytes,2,opt,name=pages,proto3" json:"pages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Blog_BlogResponse) Reset()         { *m = Blog_BlogResponse{} }
func (m *Blog_BlogResponse) String() string { return proto.CompactTextString(m) }
func (*Blog_BlogResponse) ProtoMessage()    {}
func (*Blog_BlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{0, 7}
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

func (m *Blog_BlogResponse) GetPages() *Common_Pages {
	if m != nil {
		return m.Pages
	}
	return nil
}

type Blog_BlogArticleResponse struct {
	Article              *Blog_Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Blog_BlogArticleResponse) Reset()         { *m = Blog_BlogArticleResponse{} }
func (m *Blog_BlogArticleResponse) String() string { return proto.CompactTextString(m) }
func (*Blog_BlogArticleResponse) ProtoMessage()    {}
func (*Blog_BlogArticleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{0, 8}
}

func (m *Blog_BlogArticleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_BlogArticleResponse.Unmarshal(m, b)
}
func (m *Blog_BlogArticleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_BlogArticleResponse.Marshal(b, m, deterministic)
}
func (m *Blog_BlogArticleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_BlogArticleResponse.Merge(m, src)
}
func (m *Blog_BlogArticleResponse) XXX_Size() int {
	return xxx_messageInfo_Blog_BlogArticleResponse.Size(m)
}
func (m *Blog_BlogArticleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_BlogArticleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_BlogArticleResponse proto.InternalMessageInfo

func (m *Blog_BlogArticleResponse) GetArticle() *Blog_Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type Blog_BlogSubscriptionResponse struct {
	IsSubscribed         bool     `protobuf:"varint,1,opt,name=is_subscribed,json=isSubscribed,proto3" json:"is_subscribed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog_BlogSubscriptionResponse) Reset()         { *m = Blog_BlogSubscriptionResponse{} }
func (m *Blog_BlogSubscriptionResponse) String() string { return proto.CompactTextString(m) }
func (*Blog_BlogSubscriptionResponse) ProtoMessage()    {}
func (*Blog_BlogSubscriptionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{0, 9}
}

func (m *Blog_BlogSubscriptionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_BlogSubscriptionResponse.Unmarshal(m, b)
}
func (m *Blog_BlogSubscriptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_BlogSubscriptionResponse.Marshal(b, m, deterministic)
}
func (m *Blog_BlogSubscriptionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_BlogSubscriptionResponse.Merge(m, src)
}
func (m *Blog_BlogSubscriptionResponse) XXX_Size() int {
	return xxx_messageInfo_Blog_BlogSubscriptionResponse.Size(m)
}
func (m *Blog_BlogSubscriptionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_BlogSubscriptionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_BlogSubscriptionResponse proto.InternalMessageInfo

func (m *Blog_BlogSubscriptionResponse) GetIsSubscribed() bool {
	if m != nil {
		return m.IsSubscribed
	}
	return false
}

type Blog_BlogArticleLikeResponse struct {
	LikeCount            uint64   `protobuf:"varint,1,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog_BlogArticleLikeResponse) Reset()         { *m = Blog_BlogArticleLikeResponse{} }
func (m *Blog_BlogArticleLikeResponse) String() string { return proto.CompactTextString(m) }
func (*Blog_BlogArticleLikeResponse) ProtoMessage()    {}
func (*Blog_BlogArticleLikeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{0, 10}
}

func (m *Blog_BlogArticleLikeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog_BlogArticleLikeResponse.Unmarshal(m, b)
}
func (m *Blog_BlogArticleLikeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog_BlogArticleLikeResponse.Marshal(b, m, deterministic)
}
func (m *Blog_BlogArticleLikeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog_BlogArticleLikeResponse.Merge(m, src)
}
func (m *Blog_BlogArticleLikeResponse) XXX_Size() int {
	return xxx_messageInfo_Blog_BlogArticleLikeResponse.Size(m)
}
func (m *Blog_BlogArticleLikeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog_BlogArticleLikeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Blog_BlogArticleLikeResponse proto.InternalMessageInfo

func (m *Blog_BlogArticleLikeResponse) GetLikeCount() uint64 {
	if m != nil {
		return m.LikeCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Blog)(nil), "Blog")
	proto.RegisterType((*Blog_LastArticle)(nil), "Blog.LastArticle")
	proto.RegisterType((*Blog_Community)(nil), "Blog.Community")
	proto.RegisterType((*Blog_Community_Stats)(nil), "Blog.Community.Stats")
	proto.RegisterType((*Blog_Article)(nil), "Blog.Article")
	proto.RegisterType((*Blog_Article_Stats)(nil), "Blog.Article.Stats")
	proto.RegisterType((*Blog_Blog)(nil), "Blog.Blog")
	proto.RegisterType((*Blog_Blog_Stats)(nil), "Blog.Blog.Stats")
	proto.RegisterType((*Blog_CommunitiesResponse)(nil), "Blog.CommunitiesResponse")
	proto.RegisterType((*Blog_CommunityResponse)(nil), "Blog.CommunityResponse")
	proto.RegisterType((*Blog_BlogsResponse)(nil), "Blog.BlogsResponse")
	proto.RegisterType((*Blog_BlogResponse)(nil), "Blog.BlogResponse")
	proto.RegisterType((*Blog_BlogArticleResponse)(nil), "Blog.BlogArticleResponse")
	proto.RegisterType((*Blog_BlogSubscriptionResponse)(nil), "Blog.BlogSubscriptionResponse")
	proto.RegisterType((*Blog_BlogArticleLikeResponse)(nil), "Blog.BlogArticleLikeResponse")
}

func init() { proto.RegisterFile("proto/blog.proto", fileDescriptor_fc5203cdc85000bc) }

var fileDescriptor_fc5203cdc85000bc = []byte{
	// 730 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x96, 0x1d, 0xe7, 0xe2, 0x93, 0xe4, 0x6f, 0x3b, 0x6d, 0xff, 0xdf, 0xf2, 0xaf, 0x8a, 0x88,
	0x20, 0x48, 0xb9, 0x38, 0xa8, 0x65, 0xc1, 0x0a, 0x44, 0xc3, 0xb2, 0x0b, 0x34, 0xe5, 0x22, 0xb1,
	0x20, 0x9a, 0xc4, 0x43, 0x18, 0xe2, 0xd8, 0x96, 0x67, 0x52, 0xe0, 0x11, 0xd8, 0xf1, 0x00, 0x3c,
	0x14, 0xcf, 0xc1, 0x8a, 0x47, 0x40, 0x73, 0xb3, 0x13, 0x5a, 0xa2, 0xb2, 0x62, 0x13, 0x79, 0xbe,
	0xf3, 0x9d, 0x33, 0xe7, 0x7c, 0xf3, 0xcd, 0x04, 0xb6, 0xf3, 0x22, 0x13, 0xd9, 0x70, 0x92, 0x64,
	0xb3, 0x48, 0x7d, 0x86, 0x48, 0x23, 0xd3, 0x6c, 0xb1, 0xc8, 0x52, 0x83, 0x5d, 0x9b, 0x65, 0xd9,
	0x2c, 0xa1, 0x43, 0xb5, 0x9a, 0x2c, 0xdf, 0x0e, 0x05, 0x5b, 0x50, 0x2e, 0xc8, 0x22, 0xd7, 0x84,
	0xeb, 0xdf, 0xbb, 0xe0, 0x9d, 0x24, 0xd9, 0x2c, 0xfc, 0xec, 0x40, 0xfb, 0x94, 0x70, 0xf1, 0xa4,
	0x10, 0x6c, 0x9a, 0x50, 0xf4, 0x0f, 0xb8, 0x2c, 0x0e, 0x9c, 0x9e, 0x33, 0xf0, 0xb0, 0xcb, 0x62,
	0xb4, 0x07, 0x75, 0xc1, 0x44, 0x42, 0x03, 0xb7, 0xe7, 0x0c, 0x7c, 0xac, 0x17, 0xe8, 0x06, 0x78,
	0x4b, 0x4e, 0x8b, 0xa0, 0xd6, 0x73, 0x06, 0xed, 0xa3, 0xed, 0x68, 0xa4, 0x37, 0x7f, 0xc1, 0x69,
	0x71, 0xca, 0xd2, 0x39, 0x56, 0x51, 0x14, 0x81, 0x17, 0x13, 0x41, 0x03, 0x4f, 0xb1, 0xc2, 0x48,
	0x37, 0x15, 0xd9, 0xa6, 0xa2, 0xe7, 0xb6, 0x29, 0xac, 0x78, 0xe1, 0x37, 0x17, 0x7c, 0x59, 0x69,
	0x99, 0x32, 0xf1, 0xe9, 0x8a, 0x9d, 0x1c, 0xc3, 0xfe, 0xd4, 0xa6, 0x8c, 0x63, 0xca, 0xa7, 0x05,
	0xcb, 0x05, 0xcb, 0x52, 0xd5, 0x9a, 0x8f, 0xf7, 0xca, 0xe0, 0xd3, 0x2a, 0x26, 0x4b, 0x15, 0xcb,
	0x84, 0x72, 0xd5, 0x99, 0x8f, 0xf5, 0x02, 0xfd, 0x0b, 0x0d, 0x72, 0x4e, 0x04, 0x29, 0x82, 0xba,
	0x82, 0xcd, 0x0a, 0xdd, 0x81, 0x3a, 0x17, 0x44, 0xf0, 0xa0, 0xa1, 0xe6, 0xd8, 0x8f, 0xa4, 0x70,
	0x51, 0xd9, 0x68, 0x74, 0x26, 0x83, 0x58, 0x73, 0xd0, 0x03, 0xe8, 0x24, 0x84, 0x8b, 0x31, 0xd1,
	0x7a, 0x06, 0x4d, 0x95, 0xb3, 0xa3, 0x73, 0x56, 0x84, 0xc6, 0xed, 0xa4, 0x5a, 0x84, 0xaf, 0xa0,
	0xae, 0xaa, 0xa0, 0x3e, 0x74, 0x4d, 0xe6, 0x78, 0x9a, 0x2d, 0x53, 0x61, 0xe6, 0xef, 0x18, 0x70,
	0x24, 0x31, 0x74, 0x08, 0xdb, 0x7c, 0x39, 0x91, 0xe3, 0x4c, 0x68, 0x61, 0x78, 0xae, 0xe2, 0x6d,
	0x55, 0xb8, 0xa2, 0x86, 0x5f, 0x5d, 0x68, 0xfe, 0xd9, 0xd1, 0xde, 0x85, 0xd6, 0xb4, 0xa0, 0xa4,
	0xd4, 0x70, 0xe5, 0x78, 0x47, 0x06, 0xc7, 0x25, 0x03, 0x21, 0xf0, 0x04, 0xfd, 0x28, 0x8c, 0x90,
	0xea, 0x5b, 0x61, 0x64, 0xc6, 0x8d, 0x8a, 0xea, 0x1b, 0x1d, 0xae, 0x6b, 0xb8, 0xab, 0xf5, 0x30,
	0x9d, 0xad, 0x29, 0x18, 0xbe, 0xb7, 0x5a, 0x1c, 0x00, 0x24, 0x6c, 0xbe, 0x2e, 0x84, 0x2f, 0x11,
	0xad, 0xc2, 0x01, 0xc0, 0x39, 0xa3, 0x1f, 0xd6, 0xe6, 0xf7, 0x25, 0xa2, 0xc3, 0x7d, 0xe8, 0xca,
	0xb3, 0xa7, 0xa9, 0x30, 0x8c, 0x9a, 0x56, 0xd2, 0x80, 0x5a, 0x9e, 0x2f, 0xae, 0xbe, 0x06, 0x17,
	0xb4, 0xb1, 0x06, 0x77, 0x37, 0x1a, 0xfc, 0x7f, 0xf0, 0x19, 0x1f, 0x4f, 0x93, 0x8c, 0xd3, 0x58,
	0xd5, 0x6f, 0xe1, 0x16, 0xe3, 0x23, 0xb5, 0x46, 0x37, 0xed, 0xc8, 0x9e, 0xa9, 0xa1, 0x46, 0x56,
	0x3f, 0x1b, 0x1d, 0x53, 0xff, 0xbb, 0x8e, 0x99, 0xc3, 0xae, 0xb5, 0x36, 0xa3, 0x1c, 0x53, 0x9e,
	0x67, 0x29, 0xa7, 0xa8, 0x0f, 0xde, 0x82, 0xb0, 0x34, 0x70, 0x7a, 0xb5, 0x41, 0xfb, 0x68, 0xeb,
	0x97, 0x3b, 0x80, 0x55, 0x10, 0x0d, 0x01, 0x48, 0x1c, 0x33, 0xe9, 0x0c, 0x92, 0x04, 0xee, 0xe5,
	0xd4, 0x15, 0x4a, 0xf8, 0xc3, 0x81, 0x9d, 0x2a, 0x62, 0xf7, 0xba, 0x07, 0x7e, 0x79, 0x6d, 0xd5,
	0x38, 0x97, 0x54, 0xa9, 0x18, 0xe8, 0x3e, 0xc0, 0x22, 0x8b, 0x69, 0x41, 0x44, 0x56, 0x70, 0xb3,
	0xeb, 0xc5, 0x13, 0x5b, 0xe1, 0xa0, 0xdb, 0xd0, 0x24, 0x4b, 0xf1, 0x4e, 0xd2, 0x6b, 0xbf, 0xa1,
	0x5b, 0x02, 0x3a, 0x84, 0x96, 0x91, 0x52, 0x9e, 0xa4, 0x24, 0x77, 0xd7, 0xcc, 0x8b, 0xcb, 0x30,
	0xea, 0x43, 0x3d, 0x27, 0x33, 0xca, 0xcd, 0x11, 0x76, 0x6d, 0xd1, 0x67, 0x12, 0xc4, 0x3a, 0x16,
	0xbe, 0x84, 0xae, 0x4c, 0xaf, 0x94, 0xed, 0x41, 0x5d, 0xbe, 0xe6, 0xdc, 0x48, 0x0b, 0x95, 0x4f,
	0xb0, 0x0e, 0x54, 0x75, 0xdd, 0x0d, 0x75, 0xdf, 0x40, 0x47, 0xe5, 0xd8, 0xb2, 0xab, 0x7d, 0x3b,
	0x57, 0xec, 0x7b, 0x53, 0xfd, 0x47, 0xb0, 0x2b, 0xd3, 0x6d, 0xb6, 0xdd, 0xe6, 0x16, 0x34, 0xad,
	0x71, 0x1d, 0x93, 0xbd, 0xb6, 0x8b, 0x8d, 0x86, 0x8f, 0x21, 0x90, 0x81, 0x33, 0x6d, 0x37, 0xf5,
	0x0e, 0xaf, 0x98, 0xab, 0xcb, 0xf8, 0xb8, 0x74, 0xa2, 0xbe, 0x88, 0x2d, 0xdc, 0x61, 0xfc, 0xac,
	0xc4, 0xc2, 0x87, 0xf0, 0xdf, 0x4a, 0x03, 0xa7, 0x6c, 0x5e, 0x35, 0xb1, 0xf9, 0xa5, 0x38, 0x69,
	0xbd, 0x6e, 0x90, 0x3c, 0x1f, 0xe6, 0x93, 0x49, 0x43, 0xfd, 0xf7, 0x1c, 0xff, 0x0c, 0x00, 0x00,
	0xff, 0xff, 0x2b, 0x37, 0xb9, 0xbb, 0x46, 0x07, 0x00, 0x00,
}
