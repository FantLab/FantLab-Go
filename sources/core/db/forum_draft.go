package db

import (
	"context"
	"fantlab/core/db/queries"
	"time"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
)

type ForumMessageDraft struct {
	DraftId     uint64    `db:"preview_id"`
	TopicId     uint64    `db:"topic_id"`
	Message     string    `db:"message"`
	DateOfAdd   time.Time `db:"date_of_add"`
	DateOfEdit  time.Time `db:"date_of_edit"`
	UserID      uint64    `db:"user_id"`
	Login       string    `db:"login"`
	Sex         uint8     `db:"sex"`
	PhotoNumber uint64    `db:"photo_number"`
	UserClass   uint8     `db:"user_class"`
	Sign        string    `db:"sign"`
}

// TODO Переделать всю работу с черновиками, чтобы опираться на draftId вместо topicId+userId
func (db *DB) InsertForumMessageDraft(ctx context.Context, message string, topicId, userId uint64) (*ForumMessageDraft, error) {
	var messageDraft ForumMessageDraft

	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Создаем черновик сообщения
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumInsertMessagePreview).WithArgs(message, userId, topicId, message)).Error
			},
			func() error { // Получаем черновик
				return rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicMessagePreview).WithArgs(topicId, userId), &messageDraft)
			},
		)
	})

	if err != nil {
		return nil, err
	}

	return &messageDraft, nil
}

func (db *DB) FetchForumMessageDraft(ctx context.Context, topicId, userId uint64) (*ForumMessageDraft, error) {
	var messageDraft ForumMessageDraft

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicMessagePreview).WithArgs(topicId, userId), &messageDraft)

	if err != nil {
		return nil, err
	}

	return &messageDraft, nil
}

func (db *DB) ConfirmForumMessageDraft(ctx context.Context, topic *ForumTopic, userId uint64, login, text string, isRed uint8, forumMessagesInPage uint64) (*ForumMessage, error) {
	var messageId uint64
	var message ForumMessage

	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Создаем сообщение
				message, err := db.InsertForumMessage(ctx, topic, userId, login, text, isRed, forumMessagesInPage)
				if err == nil {
					messageId = message.MessageID
				}
				return err
			},
			func() error { // Получаем сообщение
				return rw.Read(ctx, sqlapi.NewQuery(queries.ForumGetTopicMessage).WithArgs(messageId), &message)
			},
			func() error { // Удаляем черновик
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumDeleteForumMessagePreview).WithArgs(topic.TopicId, userId)).Error
			},
		)
	})

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (db *DB) DeleteForumMessageDraft(ctx context.Context, topicId, draftId, userId uint64) error {
	return db.engine.Write(ctx, sqlapi.NewQuery(queries.ForumDeleteForumMessagePreview).WithArgs(topicId, userId)).Error
}
