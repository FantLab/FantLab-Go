package db

import (
	"context"
	"fantlab/base/codeflow"
	"fantlab/base/dbtools/sqlbuilder"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"fantlab/server/internal/helpers"
	"time"
)

type Forum struct {
	ForumID         uint64    `db:"forum_id"`
	Name            string    `db:"name"`
	Description     string    `db:"description"`
	TopicCount      uint64    `db:"topic_count"`
	MessageCount    uint64    `db:"message_count"`
	LastTopicID     uint64    `db:"last_topic_id"`
	LastTopicName   string    `db:"last_topic_name"`
	UserID          uint64    `db:"user_id"`
	Login           string    `db:"login"`
	Sex             uint8     `db:"sex"`
	PhotoNumber     uint64    `db:"photo_number"`
	LastMessageID   uint64    `db:"last_message_id"`
	LastMessageText string    `db:"last_message_text"`
	LastMessageDate time.Time `db:"last_message_date"`
	ForumBlockID    uint64    `db:"forum_block_id"`
	ForumBlockName  string    `db:"forum_block_name"`
}

type ForumTopic struct {
	TopicId         uint64    `db:"topic_id"`
	ForumId         uint64    `db:"forum_id"`
	Name            string    `db:"name"`
	DateOfAdd       time.Time `db:"date_of_add"`
	Views           uint64    `db:"views"`
	UserID          uint64    `db:"user_id"`
	Login           string    `db:"login"`
	Sex             uint8     `db:"sex"`
	PhotoNumber     uint64    `db:"photo_number"`
	TopicTypeID     uint64    `db:"topic_type_id"`
	IsClosed        uint8     `db:"is_closed"`
	IsPinned        uint8     `db:"is_pinned"`
	MessageCount    uint64    `db:"message_count"`
	LastMessageID   uint64    `db:"last_message_id"`
	LastUserID      uint64    `db:"last_user_id"`
	LastLogin       string    `db:"last_login"`
	LastSex         uint8     `db:"last_sex"`
	LastPhotoNumber uint64    `db:"last_photo_number"`
	LastMessageText string    `db:"last_message_text"`
	LastMessageDate time.Time `db:"last_message_date"`
	IsModerated     uint8     `db:"moderated"`
}

type ShortForumTopic struct {
	TopicID              uint64 `db:"topic_id"`
	TopicName            string `db:"topic_name"`
	IsFirstMessagePinned uint8  `db:"is_first_message_pinned"`
	ForumID              uint64 `db:"forum_id"`
	ForumName            string `db:"forum_name"`
}

type ForumMessage struct {
	MessageID   uint64    `db:"message_id"`
	TopicId     uint64    `db:"topic_id"`
	DateOfAdd   time.Time `db:"date_of_add"`
	UserID      uint64    `db:"user_id"`
	IsRed       uint8     `db:"is_red"` // модераторское?
	Login       string    `db:"login"`
	Sex         uint8     `db:"sex"`
	PhotoNumber uint64    `db:"photo_number"`
	UserClass   uint8     `db:"user_class"`
	Sign        string    `db:"sign"`
	MessageText string    `db:"message_text"`
	IsCensored  uint8     `db:"is_censored"`
	VotePlus    uint64    `db:"vote_plus"`
	VoteMinus   uint64    `db:"vote_minus"`
}

type ForumModerator struct {
	UserID      uint64 `db:"user_id"`
	Login       string `db:"login"`
	Sex         uint8  `db:"sex"`
	PhotoNumber uint64 `db:"photo_number"`
	ForumID     uint64 `db:"forum_id"`
}

type ForumTopicsDBResponse struct {
	Topics           []ForumTopic
	TotalTopicsCount uint64
}

type ForumTopicMessagesDBResponse struct {
	Topic              ShortForumTopic
	PinnedFirstMessage ForumMessage
	Messages           []ForumMessage
	TotalMessagesCount uint64
}

func (db *DB) FetchForums(ctx context.Context, availableForums []uint64) ([]Forum, error) {
	var forums []Forum

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.Forums).WithArgs(availableForums).FlatArgs()).Scan(&forums)

	if err != nil {
		return nil, err
	}

	return forums, nil
}

func (db *DB) FetchModerators(ctx context.Context) (map[uint64][]ForumModerator, error) {
	var moderators []ForumModerator

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.ForumModerators)).Scan(&moderators)

	if err != nil {
		return nil, err
	}

	moderatorsMap := map[uint64][]ForumModerator{}

	for _, moderator := range moderators {
		moderatorsMap[moderator.ForumID] = append(moderatorsMap[moderator.ForumID], moderator)
	}

	return moderatorsMap, nil
}

func (db *DB) FetchForumTopics(ctx context.Context, availableForums []uint64, forumID, limit, offset uint64) (response *ForumTopicsDBResponse, err error) {
	var forumExists uint8
	var topics []ForumTopic
	var count uint64

	err = codeflow.Try(
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.ForumExists).WithArgs(forumID, availableForums).FlatArgs()).Scan(&forumExists)
		},
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.ForumTopics).WithArgs(forumID, limit, offset)).Scan(&topics)
		},
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.ForumTopicsCount).WithArgs(forumID)).Scan(&count)
		},
	)

	if err == nil {
		response = &ForumTopicsDBResponse{
			Topics:           topics,
			TotalTopicsCount: count,
		}
	}

	return
}

func (db *DB) FetchForumTopic(ctx context.Context, availableForums []uint64, topicID uint64) (*ForumTopic, error) {
	var topic ForumTopic

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.ForumTopic).WithArgs(topicID, availableForums).FlatArgs()).Scan(&topic)

	if err != nil {
		return nil, err
	}

	return &topic, nil
}

func (db *DB) FetchTopicMessages(ctx context.Context, availableForums []uint64, topicID, limit, offset uint64, asc bool) (response *ForumTopicMessagesDBResponse, err error) {
	var shortTopic ShortForumTopic
	var pinnedFirstMessage ForumMessage
	var messages []ForumMessage
	var count uint64

	err = codeflow.Try(
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.ShortForumTopic).WithArgs(topicID, availableForums).FlatArgs()).Scan(&shortTopic)
		},
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.ForumTopicMessagesCount).WithArgs(topicID)).Scan(&count)
		},
		func() error {
			finalOffset := int64(offset)
			if !asc {
				finalOffset = int64(count) - int64(offset) - int64(limit)
			}

			var sortDirection string
			if asc {
				sortDirection = "ASC"
			} else {
				sortDirection = "DESC"
			}

			return db.engine.Read(ctx, sqlr.NewQuery(queries.ForumTopicMessages).Inject(sortDirection).WithArgs(topicID, finalOffset+1, finalOffset+int64(limit))).Scan(&messages)
		},
	)

	if shortTopic.IsFirstMessagePinned == 1 {
		err = db.engine.Read(ctx, sqlr.NewQuery(queries.ForumTopicFirstMessage).WithArgs(topicID)).Scan(&pinnedFirstMessage)
	}

	if err == nil {
		response = &ForumTopicMessagesDBResponse{
			Topic:              shortTopic,
			PinnedFirstMessage: pinnedFirstMessage,
			Messages:           messages,
			TotalMessagesCount: count,
		}
	}

	return
}

func (db *DB) FetchForumMessage(ctx context.Context, messageId uint64, availableForums []uint64) (*ForumMessage, error) {
	var message ForumMessage

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.ForumGetShortMessage).WithArgs(messageId, availableForums).FlatArgs()).Scan(&message)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (db *DB) FetchForumMessageUserVoteCount(ctx context.Context, userId, messageId uint64) (uint64, error) {
	var count uint64

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.ForumMessageUserVoteCount).WithArgs(userId, messageId)).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *DB) FetchUserIsForumModerator(ctx context.Context, userId, topicId uint64) (bool, error) {
	var userIsForumModerator uint8

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.UserIsForumModerator).WithArgs(userId, topicId)).Scan(&userIsForumModerator)

	if err != nil {
		return false, err
	}

	return userIsForumModerator == 1, nil
}

func (db *DB) InsertNewForumMessage(ctx context.Context, topic *ForumTopic, userId uint64, login string, text string, isRed uint8, forumMessagesInPage uint64) error {
	err := db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		var message ForumMessage

		return codeflow.Try(
			func() error { // Создаем новое сообщение
				// TODO Поле проставляется для подсчета баллов, однако в методе расчета абсолютно не учитывается нерусский текст
				//  (например, если писать в "English forum", рейтинг никак не поменяется). Выглядит как очередной баг.
				messageLength := len(helpers.RemoveImmeasurable(text))
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumInsertNewMessage).WithArgs(messageLength, topic.TopicId, userId, topic.ForumId, isRed, topic.TopicId)).Error
			},
			func() error { // Получаем данные свежесозданного сообщения
				return rw.Read(ctx, sqlr.NewQuery(queries.ForumGetTopicLastMessage).WithArgs(topic.TopicId)).Scan(&message)
			},
			func() error { // Сохраняем текст сообщения
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumSetMessageText).WithArgs(message.MessageID, text)).Error
			},
			func() error { // Удаляем, если есть, черновик сообщения для данной темы
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumCancelTopicMessagePreview).WithArgs(userId, topic.TopicId)).Error
			},
			func() error { // Обновляем статистику пользователя, выставляем флаг для Cron-а
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumUpdateUserStat).WithArgs(userId)).Error
			},
			func() error { // Обновляем статистику форума
				return updateForumStat(ctx, rw, topic, forumMessagesInPage, message.MessageID, userId, login)
			},
		)
	})

	if err != nil {
		return err
	}

	return notifyForumTopicSubscribers(ctx, db.engine, topic.TopicId)
}

func updateForumStat(ctx context.Context, rw sqlr.ReaderWriter, topic *ForumTopic, forumMessagesInPage, messageId, userId uint64, login string) error {
	if topic.IsModerated == 0 {
		return nil
	}

	type forumStat struct {
		TopicCount   uint64 `db:"topic_count"`
		MessageCount uint64 `db:"forum_message_count"`
	}
	var stat forumStat
	var topicMessageCount uint64

	return codeflow.Try(
		func() error { // Обновляем данные о последнем сообщении в теме, выставляем флаги для Cron-а
			return rw.Write(ctx, sqlr.NewQuery(queries.ForumSetTopicLastMessage).WithArgs(messageId, userId, login, topic.TopicId)).Error
		},
		func() error { // Получаем обновленную статистику форума
			return rw.Read(ctx, sqlr.NewQuery(queries.ForumGetStat).WithArgs(topic.ForumId)).Scan(&stat)
		},
		func() error { // Получаем количество сообщений в теме
			return rw.Read(ctx, sqlr.NewQuery(queries.ForumGetTopicMessageCount).WithArgs(topic.TopicId)).Scan(&topicMessageCount)
		},
		func() error { // Обновляем данные о последней теме в форуме
			pageCount := helpers.CalculatePageCount(topicMessageCount, forumMessagesInPage)
			return rw.Write(ctx, sqlr.NewQuery(queries.ForumSetForumLastTopic).
				WithArgs(stat.MessageCount, stat.TopicCount, messageId, userId, login, topic.TopicId, topic.Name, pageCount, topic.ForumId)).Error
		},
	)
}

func notifyForumTopicSubscribers(ctx context.Context, rw sqlr.ReaderWriter, topicId uint64) error {
	var message ForumMessage
	var topicSubscribers []uint64

	type newForumAnswerEntry struct {
		TopicId     uint64    `db:"topic_id"`
		UserId      uint64    `db:"user_id"`
		MessageId   uint64    `db:"message_id"`
		MessageDate time.Time `db:"date_of_add"`
	}

	err := codeflow.Try(
		func() error { // Получаем данные сообщения
			return rw.Read(ctx, sqlr.NewQuery(queries.ForumGetTopicLastMessage).WithArgs(topicId)).Scan(&message)
		},
		func() error { // Получаем список подписчиков на обновления темы (за вычетом автора сообщения)
			return rw.Read(ctx, sqlr.NewQuery(queries.ForumGetTopicSubscribers).WithArgs(topicId, message.UserID)).Scan(&topicSubscribers)
		},
		func() error { // Обновляем статистику новых ответов в форуме для подписчиков
			return rw.Write(ctx, sqlr.NewQuery(queries.ForumUpdateNewForumAnswersCount).WithArgs(topicSubscribers).FlatArgs()).Error
		},
	)

	if err != nil {
		return err
	}

	// Добавляем оповещение для подписчиков о новом сообщении в теме. Не в транзакции, поскольку запрос тяжелый.
	entries := make([]interface{}, len(topicSubscribers))
	for index, userId := range topicSubscribers {
		entries[index] = newForumAnswerEntry{
			TopicId:     topicId,
			UserId:      userId,
			MessageId:   message.MessageID,
			MessageDate: message.DateOfAdd,
		}
	}
	return rw.Write(ctx, sqlbuilder.InsertInto(queries.NewForumAnswersTable, entries...)).Error
}

func (db *DB) FetchForumTopicSubscribers(ctx context.Context, topicId, excludedUserId uint64) (subscribers []uint64, err error) {
	err = db.engine.Read(ctx, sqlr.NewQuery(queries.ForumGetTopicSubscribers).WithArgs(topicId, excludedUserId)).Scan(&subscribers)
	return
}
