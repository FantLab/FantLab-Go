package db

import (
	"context"
	"fantlab/base/codeflow"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"strings"
	"time"
)

type Community struct {
	BlogId          uint64    `db:"blog_id"`
	Name            string    `db:"name"`
	Description     string    `db:"description"`
	Rules           string    `db:"rules"`
	TopicsCount     uint64    `db:"topics_count"`
	IsPublic        uint8     `db:"is_public"`
	LastTopicDate   time.Time `db:"last_topic_date"`
	LastTopicHead   string    `db:"last_topic_head"`
	LastTopicId     uint64    `db:"last_topic_id"`
	SubscriberCount uint64    `db:"subscriber_count"`
	LastUserId      uint64    `db:"last_user_id"`
	LastLogin       string    `db:"last_login"`
	LastSex         uint8     `db:"last_sex"`
	LastPhotoNumber uint64    `db:"last_photo_number"`
}

type CommunityModerator struct {
	UserID      uint64 `db:"user_id"`
	Login       string `db:"login"`
	Sex         uint8  `db:"sex"`
	PhotoNumber uint64 `db:"photo_number"`
}

type CommunityAuthor struct {
	UserID      uint64    `db:"user_id"`
	DateOfAdd   time.Time `db:"date_of_add"`
	Login       string    `db:"login"`
	Sex         uint8     `db:"sex"`
	PhotoNumber uint64    `db:"photo_number"`
}

type Blog struct {
	BlogId          uint64    `db:"blog_id"`
	UserId          uint64    `db:"user_id"`
	Login           string    `db:"login"`
	Fio             string    `db:"fio"`
	Sex             uint8     `db:"sex"`
	PhotoNumber     uint64    `db:"photo_number"`
	TopicsCount     uint64    `db:"topics_count"`
	SubscriberCount uint64    `db:"subscriber_count"`
	IsClose         uint8     `db:"is_close"`
	LastTopicDate   time.Time `db:"last_topic_date"`
	LastTopicHead   string    `db:"last_topic_head"`
	LastTopicId     uint64    `db:"last_topic_id"`
}

type BlogTopic struct {
	TopicId       uint64    `db:"topic_id"`
	HeadTopic     string    `db:"head_topic"`
	DateOfAdd     time.Time `db:"date_of_add"`
	UserId        uint64    `db:"user_id"`
	Login         string    `db:"login"`
	Sex           uint8     `db:"sex"`
	PhotoNumber   uint64    `db:"photo_number"`
	MessageText   string    `db:"message_text"`
	Tags          string    `db:"tags"`
	LikesCount    uint64    `db:"likes_count"`
	Views         uint64    `db:"views"`
	CommentsCount uint64    `db:"comments_count"`
}

type BlogTopicLike struct {
	TopicLikeId uint64    `db:"topic_like_id"`
	TopicId     uint64    `db:"topic_id"`
	UserId      uint64    `db:"user_id"`
	DateOfAdd   time.Time `db:"date_of_add"`
}

type CommunityTopicsDBResponse struct {
	Community        Community
	Moderators       []CommunityModerator
	Authors          []CommunityAuthor
	Topics           []BlogTopic
	TotalTopicsCount uint64
}

type BlogsDBResponse struct {
	Blogs      []Blog
	TotalCount uint64
}

type BlogTopicsDBResponse struct {
	Topics           []BlogTopic
	TotalTopicsCount uint64
}

func (db *DB) FetchCommunities(ctx context.Context) ([]Community, error) {
	var communities []Community

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.Communities)).Scan(&communities)

	if err != nil {
		return nil, err
	}

	return communities, nil
}

func (db *DB) FetchCommunity(ctx context.Context, communityID uint64) (*Community, error) {
	var community Community

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.Community).WithArgs(communityID)).Scan(&community)

	if err != nil {
		return nil, err
	}

	return &community, nil
}

func (db *DB) FetchCommunityTopics(ctx context.Context, communityID, limit, offset uint64) (response *CommunityTopicsDBResponse, err error) {
	var community Community
	var moderators []CommunityModerator
	var authors []CommunityAuthor
	var topics []BlogTopic
	var count uint64

	err = codeflow.Try(
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.ShortCommunity).WithArgs(communityID)).Scan(&community)
		},
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.CommunityModerators).WithArgs(communityID)).Scan(&moderators)
		},
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.CommunityAuthors).WithArgs(communityID)).Scan(&authors)
		},
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.CommunityTopics).WithArgs(communityID)).Scan(&topics)
		},
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.CommunityTopicCount).WithArgs(communityID)).Scan(&count)
		},
	)

	if err == nil {
		response = &CommunityTopicsDBResponse{
			Community:        community,
			Moderators:       moderators,
			Authors:          authors,
			Topics:           topics,
			TotalTopicsCount: count,
		}
	}

	return
}

func (db *DB) FetchBlogs(ctx context.Context, limit, offset uint64, sort string) (response *BlogsDBResponse, err error) {
	var sortOption string
	switch strings.ToLower(sort) {
	case "article":
		sortOption = "topics_count"
	case "subscriber":
		sortOption = "subscriber_count"
	default: // "update"
		sortOption = "last_topic_date"
	}

	var blogs []Blog
	var count uint64

	err = codeflow.Try(
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.Blogs).Inject(sortOption).WithArgs(limit, offset)).Scan(&blogs)
		},
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.BlogCount)).Scan(&count)
		},
	)

	if err == nil {
		response = &BlogsDBResponse{
			Blogs:      blogs,
			TotalCount: count,
		}
	}

	return
}

func (db *DB) FetchBlog(ctx context.Context, blogId uint64) (*Blog, error) {
	var blog Blog

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.Blog).WithArgs(blogId)).Scan(&blog)

	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (db *DB) FetchBlogTopics(ctx context.Context, blogID, limit, offset uint64) (response *BlogTopicsDBResponse, err error) {
	var blogExists uint8
	var topics []BlogTopic
	var count uint64

	err = codeflow.Try(
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.BlogExists).WithArgs(blogID)).Scan(&blogExists)
		},
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.BlogTopics).WithArgs(blogID)).Scan(&topics)
		},
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.BlogTopicCount).WithArgs(blogID)).Scan(&count)
		},
	)

	if err == nil {
		response = &BlogTopicsDBResponse{
			Topics:           topics,
			TotalTopicsCount: count,
		}
	}

	return
}

func (db *DB) FetchBlogTopic(ctx context.Context, topicId uint64) (*BlogTopic, error) {
	var topic BlogTopic

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.BlogTopic).WithArgs(topicId)).Scan(&topic)

	if err != nil {
		return nil, err
	}

	return &topic, nil
}
