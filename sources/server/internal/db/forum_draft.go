package db

import (
	"context"
	"fantlab/base/codeflow"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"time"
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

type ForumMessageDraftFile struct {
	FileId    uint64    `db:"file_id"`
	FileGroup string    `db:"file_group"`
	DraftId   uint64    `db:"draft_id"`
	FileName  string    `db:"file_name"`
	FileSize  uint64    `db:"file_size"`
	DateOfAdd time.Time `db:"date_of_add"`
	UserId    uint64    `db:"user_id"`
}

// TODO Переделать всю работу с черновиками, чтобы опираться на draftId вместо topicId+userId
func (db *DB) InsertForumMessageDraft(ctx context.Context, message string, topicId, userId uint64) (*ForumMessageDraft, error) {
	var messageDraft ForumMessageDraft

	err := db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Создаем черновик сообщения
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumInsertMessagePreview).WithArgs(message, userId, topicId, message)).Error
			},
			func() error { // Получаем черновик
				return rw.Read(ctx, sqlr.NewQuery(queries.ForumGetTopicMessagePreview).WithArgs(topicId, userId)).Scan(&messageDraft)
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

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.ForumGetTopicMessagePreview).WithArgs(topicId, userId)).Scan(&messageDraft)

	if err != nil {
		return nil, err
	}

	return &messageDraft, nil
}

func (db *DB) ConfirmForumMessageDraft(ctx context.Context, topic *ForumTopic, userId uint64, login, text string, isRed uint8, forumMessagesInPage uint64) (*ForumMessage, error) {
	var messageId uint64
	var message ForumMessage

	err := db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Создаем сообщение
				message, err := db.InsertForumMessage(ctx, topic, userId, login, text, isRed, forumMessagesInPage)
				if err == nil {
					messageId = message.MessageID
				}
				return err
			},
			func() error { // Получаем сообщение
				return rw.Read(ctx, sqlr.NewQuery(queries.ForumGetTopicMessage).WithArgs(messageId)).Scan(&message)
			},
			func() error { // Удаляем черновик
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumDeleteForumMessagePreview).WithArgs(topic.TopicId, userId)).Error
			},
		)
	})

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (db *DB) DeleteForumMessageDraft(ctx context.Context, topicId, draftId, userId uint64) ([]ForumMessageDraftFile, error) {
	var files []ForumMessageDraftFile

	err := db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Получаем записи об аттачах черновика
				return rw.Read(ctx, sqlr.NewQuery(queries.ForumGetMessageDraftMinioFiles).WithArgs(draftId)).Scan(&files)
			},
			func() error { // Удаляем черновик
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumDeleteForumMessagePreview).WithArgs(topicId, userId)).Error
			},
		)
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func (db *DB) FetchForumMessageDraftFileCount(ctx context.Context, draftId uint64) (uint64, error) {
	var count uint64

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.ForumGetMessageDraftMinioFileCount).WithArgs(draftId)).Error()

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *DB) FetchForumMessageDraftFile(ctx context.Context, draftId, fileId uint64) (*ForumMessageDraftFile, error) {
	var file ForumMessageDraftFile

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.ForumGetMessageDraftMinioFile).WithArgs(draftId, fileId)).Scan(&file)

	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (db *DB) InsertForumMessageDraftFile(ctx context.Context, topicId, draftId uint64, fileName string, fileSize uint64, fileDate time.Time, userId uint64) (*ForumMessageDraft, error) {
	var draft ForumMessageDraft

	err := db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Создаем запись о новом аттаче
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumInsertMessageDraftMinioFile).WithArgs(draftId, fileName, fileSize, fileDate, userId)).Error
			},
			func() error { // Получаем черновик
				// TODO Возвращать в т.ч. список аттачей черновика
				return rw.Read(ctx, sqlr.NewQuery(queries.ForumGetTopicMessagePreview).WithArgs(topicId, userId)).Scan(&draft)
			},
		)
	})

	if err != nil {
		return nil, err
	}

	return &draft, nil
}

func (db *DB) DeleteForumMessageDraftFile(ctx context.Context, draftId, fileId, topicId, userId uint64) (*ForumMessageDraft, error) {
	var draft ForumMessageDraft

	err := db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Удаляем запись об аттаче
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumDeleteMessageDraftMinioFile).WithArgs(draftId, fileId)).Error
			},
			func() error { // Получаем черновик
				// TODO Возвращать в т.ч. список аттачей черновика
				return rw.Read(ctx, sqlr.NewQuery(queries.ForumGetTopicMessagePreview).WithArgs(topicId, userId)).Scan(&draft)
			},
		)
	})

	if err != nil {
		return nil, err
	}

	return &draft, nil
}
