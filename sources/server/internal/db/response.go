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
	WorkId     uint64 `db:"work_id"`
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
			func() error { // Выставляем флаг для Cron-а о необходимости пересчета уровня развития пользователя
				return rw.Write(ctx, sqlr.NewQuery(queries.UserMarkUserNeedLevelRecalc).WithArgs(userId)).Error
			},
		)
	})
}

func (db *DB) DeleteResponse(ctx context.Context, responseId, workId, userId uint64) error {
	var registeredWorkAutorIds []uint64

	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Удаляем отзыв
				return rw.Write(ctx, sqlr.NewQuery(queries.ResponseDeleteResponse).WithArgs(responseId)).Error
			},
			func() error { // Получаем список авторов произведения, которые зарегистрированы на сайте
				return rw.Read(ctx, sqlr.NewQuery(queries.WorkGetRegisteredWorkAutorIds).WithArgs(workId)).Scan(&registeredWorkAutorIds)
			},
			func() error { // Уменьшаем счетчик количества новых отзывов у авторов, зарегистрированных на сайте
				if len(registeredWorkAutorIds) > 0 && registeredWorkAutorIds[0] != 0 /* издержки сканирования в слайс */ {
					// NOTE Нет никакой уверенности, что удаленный отзыв не был прочитан автором произведения, поэтому
					// декремент счетчика может увести его в область отрицательных значений. Это кривая логика Perl-бэка,
					// в базе такие записи действительно есть.
					return rw.Write(ctx, sqlr.NewQuery(queries.AutorDecrementAutorsNewResponseCount).WithArgs(registeredWorkAutorIds).FlatArgs()).Error
				} else {
					return nil
				}
			},
			func() error { // Уменьшаем счетчик количества отзывов пользователя
				return rw.Write(ctx, sqlr.NewQuery(queries.UserDecrementResponseCount).WithArgs(userId)).Error
			},
			func() error { // Уменьшаем счетчик количества отзывов на произведение
				return rw.Write(ctx, sqlr.NewQuery(queries.WorkDecrementResponseCount).WithArgs(workId)).Error
			},
			func() error { // Выставляем флаг для Cron-а о необходимости пересчета уровня развития пользователя
				return rw.Write(ctx, sqlr.NewQuery(queries.UserMarkUserNeedLevelRecalc).WithArgs(userId)).Error
			},
		)
	})
}
