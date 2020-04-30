package db

import (
	"context"
	"fantlab/base/codeflow"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
)

type Response struct {
	ResponseId uint64 `db:"response_id"`
	UserId     uint64 `db:"user_id"`
}

func (db *DB) FetchResponse(ctx context.Context, responseId uint64) (Response, error) {
	var response Response

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.ResponseGetResponse).WithArgs(responseId)).Scan(&response)

	if err != nil {
		return Response{}, err
	}

	return response, nil
}

func (db *DB) UpdateResponse(ctx context.Context, responseId uint64, response string, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Обновляем отзыв
				return rw.Write(ctx, sqlr.NewQuery(queries.ResponseUpdateResponse).WithArgs(response, responseId)).Error
			},
			func() error { // Выставляем флаг для Cron-а о необходимости пересчета уровня развития пользователя.
				// NOTE Хотя совершенно непонятно, зачем его пересчитывать при редактировании чужого отзыва; ох уж эта
				// логика Perl-бэка.
				return rw.Write(ctx, sqlr.NewQuery(queries.UserMarkUserNeedLevelRecalc).WithArgs(userId)).Error
			},
		)
	})
}
