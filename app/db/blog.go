package db

import (
	"database/sql"
	"fantlab/sqlr"
	"time"
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

func (db *DB) FetchCommunities() ([]Community, error) {
	const communitiesQuery = `
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
		b.is_community = 1 AND b.is_hidden = 0
	ORDER BY
		b.is_public DESC, b.last_topic_date DESC`

	var communities []Community

	err := db.R.Query(communitiesQuery).Scan(&communities)

	if err != nil {
		return nil, err
	}

	return communities, nil
}

func (db *DB) FetchCommunity(communityID uint32) (*Community, error) {
	const communityQuery = `
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
		b.blog_id = ? AND b.is_community = 1 AND b.is_hidden = 0`

	var community Community

	err := db.R.Query(communityQuery, communityID).Scan(&community)

	if err != nil {
		return nil, err
	}

	return &community, nil
}

func (db *DB) FetchCommunityTopics(communityID, limit, offset uint32) (*CommunityTopicsDBResponse, error) {
	const communityQuery = `
	SELECT
		blog_id,
		name,
		rules
	FROM
		b_blogs
	WHERE
		blog_id = ? AND is_community = 1`

	var community Community

	err := db.R.Query(communityQuery, communityID).Scan(&community)

	if err != nil {
		return nil, err
	}

	const moderatorsQuery = `
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
		cm.comm_moder_id`

	var moderators []CommunityModerator

	err = db.R.Query(moderatorsQuery, communityID).Scan(&moderators)

	if err != nil {
		return nil, err
	}

	const authorsQuery = `
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
		cu.community_user_id`

	var authors []CommunityAuthor

	err = db.R.Query(authorsQuery, communityID).Scan(&authors)

	if err != nil {
		return nil, err
	}

	const topicsQuery = `
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
		b.date_of_add DESC`

	var topics []BlogTopic

	err = db.R.Query(topicsQuery, communityID).Scan(&topics)

	if err != nil {
		return nil, err
	}

	const countQuery = `
	SELECT
		COUNT(*)
	FROM
		b_topics
	WHERE
		blog_id = ? AND is_opened = 1`

	var count uint32

	err = db.R.Query(countQuery, communityID).Scan(&count)

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

func (db *DB) FetchBlogs(limit, offset uint32, sort string) (*BlogsDBResponse, error) {
	var sortOption string
	switch sort {
	case "article":
		sortOption = "b.topics_count"
	case "subscriber":
		sortOption = "b.subscriber_count"
	default: // "update"
		sortOption = "b.last_topic_date"
	}

	const blogsQuery = `
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
		b.is_close, ? DESC
	LIMIT ?
	OFFSET ?`

	var blogs []Blog

	err := db.R.Query(blogsQuery, sortOption, limit, offset).Scan(&blogs)

	if err != nil {
		return nil, err
	}

	const countQuery = `
	SELECT
		COUNT(*)
	FROM
		b_blogs
	WHERE
		is_community = 0 AND topics_count > 0`

	var count uint32

	err = db.R.Query(countQuery).Scan(&count)

	if err != nil {
		return nil, err
	}

	result := &BlogsDBResponse{
		Blogs:      blogs,
		TotalCount: count,
	}

	return result, nil
}

func (db *DB) FetchBlog(blogId uint32) (*Blog, error) {
	const blogQuery = `
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
		b.blog_id = ? AND b.is_community = 0`

	var blog Blog

	err := db.R.Query(blogQuery, blogId).Scan(&blog)

	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (db *DB) FetchBlogTopics(blogID, limit, offset uint32) (*BlogTopicsDBResponse, error) {
	const blogExistsQuery = `SELECT 1 FROM b_blogs WHERE blog_id = ? AND is_community = 0`

	var blogExists uint8

	err := db.R.Query(blogExistsQuery, blogID).Scan(&blogExists)

	if err != nil {
		return nil, err
	}

	const topicsQuery = `
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
		b.date_of_add DESC`

	var topics []BlogTopic

	err = db.R.Query(topicsQuery, blogID).Scan(&topics)

	if err != nil {
		return nil, err
	}

	const countQuery = `
	SELECT
		COUNT(*)
	FROM
		b_topics
	WHERE
		blog_id = ? AND is_opened = 1`

	var count uint32

	err = db.R.Query(countQuery, blogID).Scan(&count)

	if err != nil {
		return nil, err
	}

	response := &BlogTopicsDBResponse{
		Topics:           topics,
		TotalTopicsCount: count,
	}

	return response, nil
}

func (db *DB) FetchBlogSubscribed(blogId, userId uint32) (bool, error) {
	const blogSubscriptionExistsQuery = `SELECT 1 FROM b_subscribers WHERE blog_id = ? AND user_id = ?`

	var blogSubscriptionExists uint8

	err := db.R.Query(blogSubscriptionExistsQuery, blogId, userId).Scan(&blogSubscriptionExists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (db *DB) UpdateBlogSubscribed(blogId, userId uint32) error {
	return db.R.InTransaction(func(dbrw sqlr.DBReaderWriter) error {
		const blogSubscriptionInsert = `
		INSERT INTO
			b_subscribers
			(user_id, blog_id, date_of_add)
		VALUES
			(?, ?, ?)`

		err := dbrw.Exec(blogSubscriptionInsert, userId, blogId, time.Now()).Error

		if err != nil {
			return err
		}

		err = updateBlogSubscriberCount(dbrw, blogId)

		return err
	})
}

func (db *DB) UpdateBlogUnsubscribed(blogId, userId uint32) error {
	return db.R.InTransaction(func(dbrw sqlr.DBReaderWriter) error {
		const blogSubscriptionDelete = `
		DELETE FROM
			b_subscribers
		WHERE
			blog_id = ? AND user_id = ?`

		err := dbrw.Exec(blogSubscriptionDelete, blogId, userId).Error

		if err != nil {
			return err
		}

		err = updateBlogSubscriberCount(dbrw, blogId)

		return err
	})
}

func updateBlogSubscriberCount(dbrw sqlr.DBReaderWriter, blogId uint32) error {
	const blogSubscriberUpdate = `
	UPDATE
		b_blogs b
	SET
		b.subscriber_count = (SELECT COUNT(DISTINCT bs.user_id) FROM b_subscribers bs WHERE bs.blog_id = b.blog_id)
	WHERE
		b.blog_id = ?`

	err := dbrw.Exec(blogSubscriberUpdate, blogId).Error

	return err
}

func (db *DB) FetchBlogTopic(topicId uint32) (*BlogTopic, error) {
	const topicQuery = `
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
		b.topic_id = ? AND b.is_opened > 0`

	var topic BlogTopic

	err := db.R.Query(topicQuery, topicId).Scan(&topic)

	if err != nil {
		return nil, err
	}

	return &topic, nil
}

func (db *DB) FetchBlogTopicSubscribed(topicId, userId uint32) (bool, error) {
	const topicSubscriptionExistsQuery = `SELECT 1 FROM b_topics_subscribers WHERE topic_id = ? AND user_id = ?`

	var topicSubscriptionExists uint8

	err := db.R.Query(topicSubscriptionExistsQuery, topicId, userId).Scan(&topicSubscriptionExists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (db *DB) UpdateBlogTopicSubscribed(topicId, userId uint32) error {
	const topicSubscriptionInsert = `
	INSERT INTO
		b_topics_subscribers
		(user_id, topic_id, date_of_add)
	VALUES
		(?, ?, ?)`

	err := db.R.Exec(topicSubscriptionInsert, userId, topicId, time.Now()).Error

	return err
}

func (db *DB) UpdateBlogTopicUnsubscribed(topicId, userId uint32) error {
	const topicSubscriptionDelete = `
	DELETE FROM
		b_topics_subscribers
	WHERE
		topic_id = ? AND user_id = ?`

	err := db.R.Exec(topicSubscriptionDelete, topicId, userId).Error

	return err
}

func (db *DB) FetchBlogTopicLikeCount(topicId uint32) (uint32, error) {
	const topicLikeCountQuery = `
	SELECT
		likes_count
	FROM
		b_topics
	WHERE
		topic_id = ?`

	var likeCount uint32

	err := db.R.Query(topicLikeCountQuery, topicId).Scan(&likeCount)

	if err != nil {
		return 0, err
	}

	return likeCount, nil
}

func (db *DB) FetchBlogTopicLiked(topicId, userId uint32) (bool, error) {
	const topicLikeExistsQuery = `SELECT 1 FROM b_topic_likes WHERE topic_id = ? AND user_id = ?`

	var topicLikeExists uint8

	err := db.R.Query(topicLikeExistsQuery, topicId, userId).Scan(&topicLikeExists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (db *DB) UpdateBlogTopicLiked(topicId, userId uint32) error {
	return db.R.InTransaction(func(dbrw sqlr.DBReaderWriter) error {
		const topicLikeInsert = `
		INSERT INTO
			b_topic_likes
			(topic_id, user_id, date_of_add)
		VALUES
			(?, ?, ?)`

		err := dbrw.Exec(topicLikeInsert, topicId, userId, time.Now()).Error

		if err != nil {
			return err
		}

		err = updateTopicLikesCount(dbrw, topicId)

		return err
	})
}

func (db *DB) UpdateBlogTopicDisliked(topicId, userId uint32) error {
	return db.R.InTransaction(func(dbrw sqlr.DBReaderWriter) error {
		const topicLikeDelete = `
		DELETE FROM
			b_topic_likes
		WHERE
			topic_id = ? AND user_id = ?`

		err := dbrw.Exec(topicLikeDelete, topicId, userId).Error

		if err != nil {
			return err
		}

		err = updateTopicLikesCount(dbrw, topicId)

		return err
	})
}

func updateTopicLikesCount(dbrw sqlr.DBReaderWriter, topicId uint32) error {
	const topicLikeUpdate = `
	UPDATE
		b_topics b
	SET
		b.likes_count = (SELECT COUNT(DISTINCT btl.user_id) FROM b_topic_likes btl WHERE btl.topic_id = b.topic_id)
	WHERE
		b.topic_id = ?`

	err := dbrw.Exec(topicLikeUpdate, topicId).Error

	return err
}
