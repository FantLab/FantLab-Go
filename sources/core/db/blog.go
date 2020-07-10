package db

import (
	"context"
	"fantlab/core/db/queries"
	"fantlab/core/helpers"
	"fmt"
	"strings"
	"time"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
	"github.com/FantLab/go-kit/database/sqlbuilder"
)

const (
	blogArticleCommentsMaxDepth = 50
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
	BlogId        uint64    `db:"blog_id"`
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

type BlogTopicComment struct {
	MessageId       uint64    `db:"message_id"`
	TopicId         uint64    `db:"topic_id"`
	UserId          uint64    `db:"user_id"`
	ParentMessageId uint64    `db:"parent_message_id"`
	MessageLength   uint64    `db:"message_length"`
	IsCensored      uint8     `db:"is_censored"`
	DateOfAdd       time.Time `db:"date_of_add"`
	TopicType       uint64    `db:"topic_type"`
	Text            string    `db:"content"`
	UserLogin       string    `db:"login"`
	UserSex         uint8     `db:"sex"`
	UserPhotoNumber uint64    `db:"photo_number"`
}

func (db *DB) FetchCommunities(ctx context.Context) ([]Community, error) {
	var communities []Community

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.Communities), &communities)

	if err != nil {
		return nil, err
	}

	return communities, nil
}

func (db *DB) FetchCommunity(ctx context.Context, communityID uint64) (*Community, error) {
	var community Community

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.Community).WithArgs(communityID), &community)

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
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.ShortCommunity).WithArgs(communityID), &community)
		},
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.CommunityModerators).WithArgs(communityID), &moderators)
		},
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.CommunityAuthors).WithArgs(communityID), &authors)
		},
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.CommunityTopics).WithArgs(communityID), &topics)
		},
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.CommunityTopicCount).WithArgs(communityID), &count)
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
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.Blogs).Inject(sortOption).WithArgs(limit, offset), &blogs)
		},
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.BlogCount), &count)
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

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.Blog).WithArgs(blogId), &blog)

	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (db *DB) FetchBlogExists(ctx context.Context, blogId uint64) (exists bool, err error) {
	var blogExists uint8
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.BlogExists).WithArgs(blogId), &blogExists)
	return blogExists == 1, err
}

func (db *DB) FetchIsUserReadOnly(ctx context.Context, userId, topicId, blogId uint64) (isReadOnly bool, err error) {
	var blogs []uint64
	var readOnlyCount uint8
	err = codeflow.Try(
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.BlogGetRelatedBlogs).WithArgs(topicId), &blogs)
		},
		func() error {
			blogs = append(blogs, blogId)
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.BlogIsUserReadOnly).WithArgs(userId, blogs).FlatArgs(), &readOnlyCount)
		},
	)
	isReadOnly = readOnlyCount > 0
	return
}

func (db *DB) FetchBlogTopics(ctx context.Context, blogID, limit, offset uint64) (response *BlogTopicsDBResponse, err error) {
	var topics []BlogTopic
	var count uint64

	err = codeflow.Try(
		func() error {
			_, err := db.FetchBlogExists(ctx, blogID)
			return err
		},
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.BlogTopics).WithArgs(blogID, limit, offset), &topics)
		},
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.BlogTopicCount).WithArgs(blogID), &count)
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

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.BlogTopic).WithArgs(topicId), &topic)

	if err != nil {
		return nil, err
	}

	return &topic, nil
}

func (db *DB) FetchBlogTopicCommentsCount(ctx context.Context, topicId uint64) (result uint64, err error) {
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.BlogGetTopicMessagesCount).WithArgs(topicId), &result)
	return
}

func (db *DB) FetchBlogTopicComments(ctx context.Context, topicId uint64, after time.Time, sortAsc bool, count uint8) (response []BlogTopicComment, err error) {
	var sortDirection string
	if sortAsc {
		sortDirection = "ASC"
	} else {
		sortDirection = "DESC"
	}
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.BlogTopicMessages).Inject(sortDirection).WithArgs(topicId, after, count, blogArticleCommentsMaxDepth), &response)
	return
}

func (db *DB) FetchBlogTopicComment(ctx context.Context, commentId uint64) (comment BlogTopicComment, err error) {
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.BlogGetTopicMessage).WithArgs(commentId), &comment)
	return
}

func (db *DB) InsertBlogTopicComment(ctx context.Context, userId, topicId, topicUserId, parentCommentId, parentCommentUserId uint64,
	text string, blogArticleCommentsInPage uint64) (*BlogTopicComment, error) {
	var commentId uint64
	var comment BlogTopicComment

	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Создаем новый комментарий
				// Примечания:
				//  - message_length всегда равен 0. Вообще неясно, зачем существует это поле. Возможно, предполагалось,
				//    что в будущем за написание новых комментариев тоже будут начисляться баллы развития.
				result := rw.Write(ctx, sqlapi.NewQuery(queries.BlogTopicInsertNewMessage).WithArgs(topicId, userId, parentCommentId, 0, 0, time.Now()))
				commentId = uint64(result.LastInsertId)
				return result.Error
			},
			func() error { // Сохраняем текст комментария
				return rw.Write(ctx, sqlapi.NewQuery(queries.BlogSetMessageText).WithArgs(commentId, text)).Error
			},
			func() error { // Получаем комментарий
				return rw.Read(ctx, sqlapi.NewQuery(queries.BlogGetTopicMessage).WithArgs(commentId), &comment)
			},
			func() error { // Обновляем статистику статьи
				return updateBlogTopicStatAfterNewMessage(ctx, rw, topicId)
			},
		)
	})

	if err != nil {
		return nil, err
	}

	var parentUsersToNotify []uint64

	if userId == topicUserId {
		if parentCommentUserId > 0 {
			// Автор комментария - автор статьи, 1+ уровень вложенности
			parentUsersToNotify = append(parentUsersToNotify, parentCommentUserId)
		}
	} else {
		if parentCommentUserId > 0 {
			if parentCommentUserId == topicUserId {
				// Автор комментария - не автор статьи, 1+ уровень вложенности, автор родительского комментария - автор статьи
				parentUsersToNotify = append(parentUsersToNotify, topicUserId)
			} else {
				// Автор комментария - не автор статьи, 1+ уровень вложенности, автор родительского комментария - тоже не автор статьи
				parentUsersToNotify = append(parentUsersToNotify, parentCommentUserId, topicUserId)
			}
		} else {
			// Автор комментария - не автор статьи, 1 уровень вложенности
			parentUsersToNotify = append(parentUsersToNotify, topicUserId)
		}
	}

	for _, userId := range parentUsersToNotify {
		err = notifyParentUserAboutNewBlogTopicMessage(ctx, db.engine, userId, topicId, commentId, blogArticleCommentsInPage)

		if err != nil {
			return nil, err
		}
	}

	err = notifyBlogTopicSubscribersAboutNewMessage(ctx, db.engine, userId, topicId, topicUserId, commentId, parentCommentUserId, blogArticleCommentsInPage)

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func updateBlogTopicStatAfterNewMessage(ctx context.Context, rw sqlapi.ReaderWriter, topicId uint64) error {
	var commentCount uint64

	return codeflow.Try(
		func() error {
			// is_actual = 0 - флаг для Cron'а, что надо пересчитать количество непрочитанных сообщений к статье для
			// данного пользователя. В настоящий момент скрипт располагается в script/cron/hulk/update_b_topic_read_cache.pl,
			// запускается через каждые 2 минуты. Используется для актуализации информации в журнале на Главной.
			return rw.Write(ctx, sqlapi.NewQuery(queries.BlogUpdateLastCommentReadActuality).WithArgs(topicId)).Error
		},
		func() error { // Получаем количество комментариев к статье
			return rw.Read(ctx, sqlapi.NewQuery(queries.BlogGetTopicMessagesCount).WithArgs(topicId), &commentCount)
		},
		func() error { // Обновляем количество комментариев в записи о самой статье
			return rw.Write(ctx, sqlapi.NewQuery(queries.BlogUpdateTopicCommentCount).WithArgs(commentCount, topicId)).Error
		},
	)
}

func notifyParentUserAboutNewBlogTopicMessage(ctx context.Context, rw sqlapi.ReaderWriter, parentUserId, topicId, commentId, blogArticleCommentsInPage uint64) error {
	var firstLevelCommentCount uint64

	return codeflow.Try(
		func() error { // Инкрементим счетчик количества новых комментариев в блогах для parentCommentUserId
			return rw.Write(ctx, sqlapi.NewQuery(queries.BlogIncrementNewBlogCommentsCount).WithArgs(parentUserId)).Error
		},
		func() error { // Получаем количество комментариев первого уровня в данной статье
			return rw.Read(ctx, sqlapi.NewQuery(queries.BlogGetFirstLevelMessageCount).WithArgs(topicId, commentId), &firstLevelCommentCount)
		},
		func() error {
			pageCount := helpers.CalculatePageCount(firstLevelCommentCount, blogArticleCommentsInPage)
			endpoint := fmt.Sprintf("blogarticle%dpage%d", topicId, pageCount)
			// Добавляем оповещение для parentUser (автор статьи или автор родительского комментария) о новом комментарии в статье
			return rw.Write(ctx, sqlapi.NewQuery(queries.BlogInsertNewMessageNotification).WithArgs(parentUserId, commentId, endpoint, parentUserId)).Error
		},
	)
}

func notifyBlogTopicSubscribersAboutNewMessage(ctx context.Context, rw sqlapi.ReaderWriter, userId, topicId, topicUserId,
	commentId, parentCommentUserId, blogArticleCommentsInPage uint64) error {
	var topicSubscribers []uint64
	var firstLevelCommentCount uint64

	type newBlogMessageEntry struct {
		UserId    uint64 `db:"user_id"`
		MessageId uint64 `db:"message_id"`
		// Поле используется для хранения endpoint-а для быстрого перехода из области уведомлений о новых сообщениях в
		// блогах. Что подобное делает в базе, останется загадкой на долгие годы.
		Endpoint string `db:"action"`
		// Сводное поле. Если текущий комментарий - первого уровня, в поле хранится id автора поста. Если комментарий
		// вложенный, то хранится id автора родительского комментария.
		ParentUserId uint64    `db:"parent_user_id"`
		MessageDate  time.Time `db:"date_of_add"`
	}

	err := codeflow.Try(
		func() error { // Получаем список подписчиков на обновления статьи
			excludedUserIds := []uint64{userId, parentCommentUserId, topicUserId}
			return rw.Read(ctx, sqlapi.NewQuery(queries.BlogGetTopicSubscribers).WithArgs(topicId, excludedUserIds).FlatArgs(), &topicSubscribers)
		},
		func() error { // Инкрементим счетчик количества новых комментариев в блогах для подписчиков
			if len(topicSubscribers) != 0 {
				return rw.Write(ctx, sqlapi.NewQuery(queries.BlogIncrementNewBlogCommentsCount).WithArgs(topicSubscribers).FlatArgs()).Error
			}
			return nil
		},
		func() error { // Получаем количество комментариев первого уровня в данной статье
			return rw.Read(ctx, sqlapi.NewQuery(queries.BlogGetFirstLevelMessageCount).WithArgs(topicId, commentId), &firstLevelCommentCount)
		},
	)

	if err != nil {
		return err
	}

	if len(topicSubscribers) == 0 {
		return nil
	}

	pageCount := helpers.CalculatePageCount(firstLevelCommentCount, blogArticleCommentsInPage)
	endpoint := fmt.Sprintf("blogarticle%dpage%d", topicId, pageCount)

	var parentUserId uint64
	if parentCommentUserId > 0 {
		parentUserId = parentCommentUserId
	} else {
		parentUserId = topicUserId
	}

	// Добавляем оповещение для подписчиков о новом комментарии в статье. Не в транзакции, поскольку запрос тяжелый.
	entries := make([]interface{}, 0, len(topicSubscribers))
	for _, userId := range topicSubscribers {
		entries = append(entries, newBlogMessageEntry{
			UserId:       userId,
			MessageId:    commentId,
			Endpoint:     endpoint,
			ParentUserId: parentUserId,
			MessageDate:  time.Now(),
		})
	}

	return rw.Write(ctx, sqlbuilder.InsertInto(queries.NewBlogTopicMessagesTable, entries...)).Error
}

func (db *DB) FetchBlogTopicSubscribers(ctx context.Context, topicId uint64, excludedUserIds []uint64) (subscribers []uint64, err error) {
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.BlogGetTopicSubscribers).WithArgs(topicId, excludedUserIds).FlatArgs(), &subscribers)
	return
}

func (db *DB) FetchUserIsCommunityModerator(ctx context.Context, userId, blogId, topicId uint64) (bool, error) {
	var userIsCommunityModerator uint8
	var userIsCommunityTopicModerator uint8

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.BlogGetUserIsCommunityModerator).WithArgs(blogId, userId), &userIsCommunityModerator)

	if err != nil {
		return false, err
	}

	if userIsCommunityModerator == 1 {
		return true, nil
	}

	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.BlogGetUserIsCommunityTopicModerator).WithArgs(topicId, userId), &userIsCommunityTopicModerator)

	if err != nil {
		return false, err
	}

	return userIsCommunityTopicModerator == 1, nil
}

func (db *DB) UpdateBlogTopicComment(ctx context.Context, commentId uint64, text string) (*BlogTopicComment, error) {
	var comment BlogTopicComment

	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.BlogSetMessageText).WithArgs(commentId, text)).Error
			},
			func() error {
				return rw.Read(ctx, sqlapi.NewQuery(queries.BlogGetTopicMessage).WithArgs(commentId), &comment)
			},
		)
	})

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (db *DB) DeleteBlogTopicComment(ctx context.Context, commentId, parentCommentId, topicId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Удаляем комментарий
				return rw.Write(ctx, sqlapi.NewQuery(queries.BlogDeleteMessage).WithArgs(commentId)).Error
			},
			func() error { // Удаляем текст комментария
				return rw.Write(ctx, sqlapi.NewQuery(queries.BlogDeleteMessageText).WithArgs(commentId)).Error
			},
			func() error { // Поднимаем дочерние комментарии на уровень выше
				return rw.Write(ctx, sqlapi.NewQuery(queries.BlogUpdateMessagesParent).WithArgs(parentCommentId, commentId)).Error
			},
			func() error {
				return updateBlogTopicStatAfterCommentDeleting(ctx, db.engine, topicId)
			},
			func() error {
				return notifyBlogTopicSubscribersAboutCommentDeleting(ctx, db.engine, commentId, topicId)
			},
		)
	})
}

func updateBlogTopicStatAfterCommentDeleting(ctx context.Context, rw sqlapi.ReaderWriter, topicId uint64) error {
	var commentCount uint64

	return codeflow.Try(
		func() error { // Получаем количество комментариев к статье
			return rw.Read(ctx, sqlapi.NewQuery(queries.BlogGetTopicMessagesCount).WithArgs(topicId), &commentCount)
		},
		func() error { // Обновляем количество комментариев в записи о самой статье
			return rw.Write(ctx, sqlapi.NewQuery(queries.BlogUpdateTopicCommentCount).WithArgs(commentCount, topicId)).Error
		},
	)
}

func notifyBlogTopicSubscribersAboutCommentDeleting(ctx context.Context, rw sqlapi.ReaderWriter, commentId, topicId uint64) error {
	return codeflow.Try(
		func() error { // Удаляем оповещение о комментарии
			return rw.Write(ctx, sqlapi.NewQuery(queries.BlogDeleteNewMessage).WithArgs(commentId)).Error
		},
		func() error {
			// Заносим id статьи в отдельную таблицу, чтобы Cron пересчитал ссылки на новые комментарии в статьях
			// (script/cron/hulk/update_b_topic_comments.pl)
			return rw.Write(ctx, sqlapi.NewQuery(queries.BlogInsertMessageDeleted).WithArgs(topicId)).Error
		},
	)
}
