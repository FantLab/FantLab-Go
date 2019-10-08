package db

import (
	"context"
	"time"

	"fantlab/dbtools/sqlr"
)

type Community struct {
	BlogId          uint32    `db:"blog_id"`
	Name            string    `db:"name"`
	Description     string    `db:"description"`
	Rules           string    `db:"rules"`
	TopicsCount     uint32    `db:"topics_count"`
	IsPublic        uint8     `db:"is_public"`
	LastTopicDate   time.Time `db:"last_topic_date"`
	LastTopicHead   string    `db:"last_topic_head"`
	LastTopicId     uint32    `db:"last_topic_id"`
	SubscriberCount uint32    `db:"subscriber_count"`
	LastUserId      uint32    `db:"last_user_id"`
	LastLogin       string    `db:"last_login"`
	LastSex         uint8     `db:"last_sex"`
	LastPhotoNumber uint32    `db:"last_photo_number"`
}

type CommunityModerator struct {
	UserID      uint32 `db:"user_id"`
	Login       string `db:"login"`
	Sex         uint8  `db:"sex"`
	PhotoNumber uint32 `db:"photo_number"`
}

type CommunityAuthor struct {
	UserID      uint32    `db:"user_id"`
	DateOfAdd   time.Time `db:"date_of_add"`
	Login       string    `db:"login"`
	Sex         uint8     `db:"sex"`
	PhotoNumber uint32    `db:"photo_number"`
}

type Blog struct {
	BlogId          uint32    `db:"blog_id"`
	UserId          uint32    `db:"user_id"`
	Login           string    `db:"login"`
	Fio             string    `db:"fio"`
	Sex             uint8     `db:"sex"`
	PhotoNumber     uint32    `db:"photo_number"`
	TopicsCount     uint32    `db:"topics_count"`
	SubscriberCount uint32    `db:"subscriber_count"`
	IsClose         uint8     `db:"is_close"`
	LastTopicDate   time.Time `db:"last_topic_date"`
	LastTopicHead   string    `db:"last_topic_head"`
	LastTopicId     uint32    `db:"last_topic_id"`
}

type BlogTopic struct {
	TopicId       uint32    `db:"topic_id"`
	HeadTopic     string    `db:"head_topic"`
	DateOfAdd     time.Time `db:"date_of_add"`
	UserId        uint32    `db:"user_id"`
	Login         string    `db:"login"`
	Sex           uint8     `db:"sex"`
	PhotoNumber   uint16    `db:"photo_number"`
	MessageText   string    `db:"message_text"`
	Tags          string    `db:"tags"`
	LikesCount    uint32    `db:"likes_count"`
	Views         uint32    `db:"views"`
	CommentsCount uint32    `db:"comments_count"`
}

type BlogTopicLike struct {
	TopicLikeId uint32    `db:"topic_like_id"`
	TopicId     uint32    `db:"topic_id"`
	UserId      uint32    `db:"user_id"`
	DateOfAdd   time.Time `db:"date_of_add"`
}

type CommunityTopicsDBResponse struct {
	Community        Community
	Moderators       []CommunityModerator
	Authors          []CommunityAuthor
	Topics           []BlogTopic
	TotalTopicsCount uint32
}

type BlogsDBResponse struct {
	Blogs      []Blog
	TotalCount uint32
}

type BlogTopicsDBResponse struct {
	Topics           []BlogTopic
	TotalTopicsCount uint32
}

var (
	communitiesQuery = sqlr.NewQuery(`
		SELECT
			b.blog_id,
			b.name,
			b.description,
			b.topics_count,
			b.is_public,
			b.last_topic_date,
			b.last_topic_head,
			b.last_topic_id,
			b.subscriber_count,
			u.user_id AS last_user_id,
			u.login AS last_login,
			u.sex AS last_sex,
			u.photo_number AS last_photo_number
		FROM
			b_blogs b
		LEFT JOIN
			users u ON u.user_id = b.last_user_id
		WHERE
			b.is_community = 1 AND b.is_hidden = 0
		ORDER BY
			b.is_public DESC, b.last_topic_date DESC
	`)

	communityQuery = sqlr.NewQuery(`
		SELECT
			b.name,
			b.description,
			b.topics_count,
			b.is_public,
			b.last_topic_date,
			b.last_topic_head,
			b.last_topic_id,
			b.subscriber_count,
			u.user_id AS last_user_id,
			u.login AS last_login,
			u.sex AS last_sex,
			u.photo_number AS last_photo_number
		FROM
			b_blogs b
		LEFT JOIN
			users u ON u.user_id = b.last_user_id
		WHERE
			b.blog_id = ? AND b.is_community = 1 AND b.is_hidden = 0
	`)

	shortCommunityQuery = sqlr.NewQuery(`
		SELECT
			blog_id,
			name,
			rules
		FROM
			b_blogs
		WHERE
			blog_id = ? AND is_community = 1
	`)

	communityModeratorsQuery = sqlr.NewQuery(`
		SELECT
			cm.user_id,
			u.login,
			u.sex,
			u.photo_number
		FROM
			b_community_moderators cm
		LEFT JOIN
			users u ON u.user_id = cm.user_id
		WHERE
			cm.blog_id = ?
		ORDER BY
			cm.comm_moder_id
	`)

	communityAuthorsQuery = sqlr.NewQuery(`
		SELECT
			cu.user_id,
			u.login,
			u.sex,
			u.photo_number
		FROM
			b_community_users cu
		LEFT JOIN
			users u ON u.user_id = cu.user_id
		WHERE
			cu.blog_id = ? AND cu.accepted = 1
		ORDER BY
			cu.community_user_id
	`)

	communityTopicsQuery = sqlr.NewQuery(`
		SELECT
			b.topic_id,
			b.head_topic,
			b.date_of_add,
			u.user_id,
			u.login,
			u.sex,
			u.photo_number,
			t.message_text,
			b.tags,
			b.likes_count,
			b.comments_count
		FROM
			b_topics b
		JOIN
			b_topics_text t ON t.message_id = b.topic_id
		LEFT JOIN
			users u ON u.user_id = b.user_id
		WHERE
			b.blog_id = ? AND b.is_opened = 1
		ORDER BY
			b.date_of_add DESC
	`)

	communityTopicCountQuery = sqlr.NewQuery(`
		SELECT
			COUNT(*)
		FROM
			b_topics
		WHERE
			blog_id = ? AND is_opened = 1
	`)

	blogsQuery = sqlr.NewQuery(`
		SELECT
			b.blog_id,
			u.user_id,
			u.login,
			u.fio,
			u.sex,
			u.photo_number,
			b.topics_count,
			b.subscriber_count,
			b.is_close,
			b.last_topic_date,
			b.last_topic_head,
			b.last_topic_id
		FROM
			b_blogs b
		LEFT JOIN
			users u ON u.user_id = b.user_id
		WHERE
			b.is_community = 0 AND b.topics_count > 0
		ORDER BY
			b.is_close, b.%s DESC
		LIMIT ?
		OFFSET ?
	`)

	blogCountQuery = sqlr.NewQuery(`
		SELECT
			COUNT(*)
		FROM
			b_blogs
		WHERE
			is_community = 0 AND topics_count > 0
	`)

	blogQuery = sqlr.NewQuery(`
		SELECT
			b.blog_id,
			u.user_id,
			u.login,
			u.fio,
			u.sex,
			u.photo_number,
			b.topics_count,
			b.subscriber_count,
			b.is_close,
			b.last_topic_date,
			b.last_topic_head,
			b.last_topic_id
		FROM
			b_blogs b
		LEFT JOIN
			users u ON u.user_id = b.user_id
		WHERE
			b.blog_id = ? AND b.is_community = 0
	`)

	blogExistsQuery = sqlr.NewQuery("SELECT 1 FROM b_blogs WHERE blog_id = ? AND is_community = 0")

	blogTopicsQuery = sqlr.NewQuery(`
		SELECT
			b.topic_id,
			b.head_topic,
			b.date_of_add,
			u.user_id,
			u.login,
			u.sex,
			u.photo_number,
			t.message_text,
			b.tags,
			b.likes_count,
			b.comments_count
		FROM
			b_topics b
		JOIN
			b_topics_text t ON t.message_id = b.topic_id
		LEFT JOIN
			users u ON u.user_id = b.user_id
		WHERE
			b.blog_id = ? AND b.is_opened = 1
		ORDER BY
			b.date_of_add DESC
	`)

	blogTopicCountQuery = sqlr.NewQuery(`
		SELECT
			COUNT(*)
		FROM
			b_topics
		WHERE
			blog_id = ? AND is_opened = 1
	`)

	topicQuery = sqlr.NewQuery(`
		SELECT
			b.topic_id,
			b.head_topic,
			b.date_of_add,
			u.user_id,
			u.login,
			u.sex,
			u.photo_number,
			t.message_text,
			b.tags,
			b.likes_count,
			b.comments_count
		FROM
			b_topics b
		JOIN
			b_topics_text t ON t.message_id = b.topic_id
		LEFT JOIN
			users u ON u.user_id = b.user_id
		WHERE
			b.topic_id = ? AND b.is_opened > 0
	`)
)

func (db *DB) FetchCommunities(ctx context.Context) ([]Community, error) {
	var communities []Community

	err := db.engine.Read(ctx, communitiesQuery).Scan(&communities)

	if err != nil {
		return nil, err
	}

	return communities, nil
}

func (db *DB) FetchCommunity(ctx context.Context, communityID uint32) (*Community, error) {
	var community Community

	err := db.engine.Read(ctx, communityQuery.WithArgs(communityID)).Scan(&community)

	if err != nil {
		return nil, err
	}

	return &community, nil
}

func (db *DB) FetchCommunityTopics(ctx context.Context, communityID, limit, offset uint32) (*CommunityTopicsDBResponse, error) {
	var community Community

	err := db.engine.Read(ctx, shortCommunityQuery.WithArgs(communityID)).Scan(&community)

	if err != nil {
		return nil, err
	}

	var moderators []CommunityModerator

	err = db.engine.Read(ctx, communityModeratorsQuery.WithArgs(communityID)).Scan(&moderators)

	if err != nil {
		return nil, err
	}

	var authors []CommunityAuthor

	err = db.engine.Read(ctx, communityAuthorsQuery.WithArgs(communityID)).Scan(&authors)

	if err != nil {
		return nil, err
	}

	var topics []BlogTopic

	err = db.engine.Read(ctx, communityTopicsQuery.WithArgs(communityID)).Scan(&topics)

	if err != nil {
		return nil, err
	}

	var count uint32

	err = db.engine.Read(ctx, communityTopicCountQuery.WithArgs(communityID)).Scan(&count)

	if err != nil {
		return nil, err
	}

	result := &CommunityTopicsDBResponse{
		Community:        community,
		Moderators:       moderators,
		Authors:          authors,
		Topics:           topics,
		TotalTopicsCount: count,
	}

	return result, nil
}

func (db *DB) FetchBlogs(ctx context.Context, limit, offset uint32, sort string) (*BlogsDBResponse, error) {
	var sortOption string
	switch sort {
	case "article":
		sortOption = "topics_count"
	case "subscriber":
		sortOption = "subscriber_count"
	default: // "update"
		sortOption = "last_topic_date"
	}

	var blogs []Blog

	err := db.engine.Read(ctx, blogsQuery.Format(sortOption).WithArgs(limit, offset)).Scan(&blogs)

	if err != nil {
		return nil, err
	}

	var count uint32

	err = db.engine.Read(ctx, blogCountQuery).Scan(&count)

	if err != nil {
		return nil, err
	}

	result := &BlogsDBResponse{
		Blogs:      blogs,
		TotalCount: count,
	}

	return result, nil
}

func (db *DB) FetchBlog(ctx context.Context, blogId uint32) (*Blog, error) {
	var blog Blog

	err := db.engine.Read(ctx, blogQuery.WithArgs(blogId)).Scan(&blog)

	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (db *DB) FetchBlogTopics(ctx context.Context, blogID, limit, offset uint32) (*BlogTopicsDBResponse, error) {
	var blogExists uint8

	err := db.engine.Read(ctx, blogExistsQuery.WithArgs(blogID)).Scan(&blogExists)

	if err != nil {
		return nil, err
	}

	var topics []BlogTopic

	err = db.engine.Read(ctx, blogTopicsQuery.WithArgs(blogID)).Scan(&topics)

	if err != nil {
		return nil, err
	}

	var count uint32

	err = db.engine.Read(ctx, blogTopicCountQuery.WithArgs(blogID)).Scan(&count)

	if err != nil {
		return nil, err
	}

	response := &BlogTopicsDBResponse{
		Topics:           topics,
		TotalTopicsCount: count,
	}

	return response, nil
}

func (db *DB) FetchBlogTopic(ctx context.Context, topicId uint32) (*BlogTopic, error) {
	var topic BlogTopic

	err := db.engine.Read(ctx, topicQuery.WithArgs(topicId)).Scan(&topic)

	if err != nil {
		return nil, err
	}

	return &topic, nil
}
