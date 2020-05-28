package db

import (
	"context"
	"fantlab/core/db/queries"
	"time"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
)

type PrivateMessage struct {
	PrivateMessageId uint64    `db:"private_message_id"`
	FromUserId       uint64    `db:"from_user_id"`
	ToUserId         uint64    `db:"to_user_id"`
	IsRead           uint8     `db:"is_read"`
	DateOfAdd        time.Time `db:"date_of_add"`
	Login            string    `db:"login"`
	Sex              uint8     `db:"sex"`
	PhotoNumber      uint64    `db:"photo_number"`
	UserClass        uint8     `db:"user_class"`
	Sign             string    `db:"sign"`
	MessageText      string    `db:"message_text"`
	Number           uint64    `db:"number"`
}

func (db *DB) InsertPrivateMessage(ctx context.Context, fromUserId, toUserId uint64, text string, isRed, sendCopyViaEmail uint8) (PrivateMessage, error) {
	var messageId uint64
	var message PrivateMessage

	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Создаем новое сообщение
				result := rw.Write(ctx, sqlapi.NewQuery(queries.PrivateInsertNewMessage).
					WithArgs(fromUserId, toUserId, sendCopyViaEmail, isRed, fromUserId, toUserId, fromUserId, toUserId))
				messageId = uint64(result.LastInsertId)
				return result.Error
			},
			func() error { // Сохраняем текст сообщения
				return rw.Write(ctx, sqlapi.NewQuery(queries.PrivateInsertMessageText).WithArgs(messageId, text)).Error
			},
			func() error { // Получаем сообщение
				return rw.Read(ctx, sqlapi.NewQuery(queries.PrivateGetMessage).WithArgs(messageId), &message)
			},
			// TODO Здесь пропущен кусок логики про то, что надо поменять ссылки на аттачи, если сообщение создается
			//  путем подтверждения черновика (https://github.com/parserpro/fantlab/blob/8e7f35553b030b798b069bd355be2de0de8fc1c6/pm/Forum.pm#L4055-L4058)
			func() error { // Удаляем, если есть, черновик сообщения
				return rw.Write(ctx, sqlapi.NewQuery(queries.PrivateCancelMessagePreview).WithArgs(fromUserId, toUserId)).Error
			},
		)
	})

	if err != nil {
		return PrivateMessage{}, err
	}

	return message, nil
}
