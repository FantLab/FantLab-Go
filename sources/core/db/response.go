package db

import (
	"context"
	"fantlab/core/db/queries"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
)

var (
	ResponseVoteMap = map[string]int8{
		"plus":  1,
		"minus": -1,
	}
)

type Response struct {
	ResponseId uint64 `db:"response_id"`
	UserId     uint64 `db:"user_id"`
	WorkId     uint64 `db:"work_id"`
	VotePlus   uint64 `db:"vote_plus"`
	VoteMinus  uint64 `db:"vote_minus"`
}

func (db *DB) FetchResponse(ctx context.Context, responseId uint64) (Response, error) {
	var response Response

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ResponseGetResponse).WithArgs(responseId), &response)

	if err != nil {
		return Response{}, err
	}

	return response, nil
}

func (db *DB) FetchResponseUserVoteCount(ctx context.Context, userId, responseId uint64) (uint64, error) {
	var count uint64

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ResponseGetResponseUserVoteCount).WithArgs(userId, responseId), &count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *DB) FetchUserResponseCountForWork(ctx context.Context, userId, workId uint64) (uint64, error) {
	var count uint64

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ResponseGetUserResponseCountForWork).WithArgs(userId, workId), &count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *DB) FetchSuchUserResponseCountForWork(ctx context.Context, userId, workId uint64, response string) (uint64, error) {
	var suchResponseCount uint64

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.ResponseGetSuchUserResponseCountForWork).WithArgs(userId, workId, response), &suchResponseCount)

	if err != nil {
		return 0, err
	}

	return suchResponseCount, nil
}

func (db *DB) InsertResponse(ctx context.Context, userId, workId uint64, workAuthorIds []uint64, response string) error {
	var registeredWorkAutorIds []uint64

	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Удаляем черновик
				return rw.Write(ctx, sqlapi.NewQuery(queries.ResponseDeleteResponsePreview).WithArgs(userId, workId)).Error
			},
			func() error { // Создаем отзыв
				return rw.Write(ctx, sqlapi.NewQuery(queries.ResponseInsertResponse).WithArgs(userId, workId, response)).Error
			},
			func() error { // Получаем список авторов произведения, которые зарегистрированы на сайте
				return rw.Read(ctx, sqlapi.NewQuery(queries.WorkGetRegisteredWorkAutorIds).WithArgs(workId), &registeredWorkAutorIds)
			},
			func() error { // Инкрементим счетчик количества новых отзывов у авторов, зарегистрированных на сайте
				if len(registeredWorkAutorIds) > 0 && registeredWorkAutorIds[0] != 0 {
					return rw.Write(ctx, sqlapi.NewQuery(queries.AutorIncrementAutorsNewResponseCount).WithArgs(registeredWorkAutorIds).FlatArgs()).Error
				} else {
					return nil
				}
			},
			func() error { // Выставляем флаги для Cron-а о необходимости пересчета статистики авторов
				if len(workAuthorIds) > 0 {
					return rw.Write(ctx, sqlapi.NewQuery(queries.AutorMarkAutorsNeedRecalcStats).WithArgs(workAuthorIds).FlatArgs()).Error
				} else {
					return nil
				}
			},
			func() error { // Инкрементим счетчик количества отзывов пользователя
				return rw.Write(ctx, sqlapi.NewQuery(queries.UserIncrementResponseCount).WithArgs(userId)).Error
			},
			func() error { // Выставляем флаг для Cron-а о необходимости пересчета уровня развития пользователя
				return rw.Write(ctx, sqlapi.NewQuery(queries.UserMarkUserNeedLevelRecalc).WithArgs(userId)).Error
			},
			func() error { // Инкрементим счетчик количества отзывов на произведение
				return rw.Write(ctx, sqlapi.NewQuery(queries.WorkStatsInsertResponseCount).WithArgs(workId)).Error
			},
		)
	})
}

func (db *DB) UpdateResponse(ctx context.Context, responseId uint64, response string, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Обновляем отзыв
				return rw.Write(ctx, sqlapi.NewQuery(queries.ResponseUpdateResponse).WithArgs(response, responseId)).Error
			},
			func() error { // Выставляем флаг для Cron-а о необходимости пересчета уровня развития пользователя
				return rw.Write(ctx, sqlapi.NewQuery(queries.UserMarkUserNeedLevelRecalc).WithArgs(userId)).Error
			},
		)
	})
}

func (db *DB) UpdateResponseVotes(ctx context.Context, responseId, userId uint64, userVote string) (Response, error) {
	var response Response

	vote := ResponseVoteMap[userVote]
	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.ResponseInsertResponseVote).WithArgs(userId, responseId, vote)).Error
			},
			func() error {
				if vote == 1 {
					return rw.Write(ctx, sqlapi.NewQuery(queries.ResponseUpdateResponseVotePlus).WithArgs(responseId)).Error
				} else {
					return rw.Write(ctx, sqlapi.NewQuery(queries.ResponseUpdateResponseVoteMinus).WithArgs(responseId)).Error
				}
			},
			func() error {
				return rw.Read(ctx, sqlapi.NewQuery(queries.ResponseGetResponse).WithArgs(responseId), &response)
			},
		)
	})

	if err != nil {
		return Response{}, err
	}

	return response, nil
}

func (db *DB) DeleteResponse(ctx context.Context, responseId, workId, userId uint64) error {
	var registeredWorkAutorIds []uint64

	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Удаляем отзыв
				return rw.Write(ctx, sqlapi.NewQuery(queries.ResponseDeleteResponse).WithArgs(responseId)).Error
			},
			func() error { // Получаем список авторов произведения, которые зарегистрированы на сайте
				return rw.Read(ctx, sqlapi.NewQuery(queries.WorkGetRegisteredWorkAutorIds).WithArgs(workId), &registeredWorkAutorIds)
			},
			func() error { // Уменьшаем счетчик количества новых отзывов у авторов, зарегистрированных на сайте
				if len(registeredWorkAutorIds) > 0 && registeredWorkAutorIds[0] != 0 /* издержки сканирования в слайс */ {
					// NOTE Нет никакой уверенности, что удаленный отзыв не был прочитан автором произведения, поэтому
					// декремент счетчика может увести его в область отрицательных значений. Это кривая логика Perl-бэка,
					// в базе такие записи действительно есть.
					return rw.Write(ctx, sqlapi.NewQuery(queries.AutorDecrementAutorsNewResponseCount).WithArgs(registeredWorkAutorIds).FlatArgs()).Error
				} else {
					return nil
				}
			},
			func() error { // Уменьшаем счетчик количества отзывов пользователя
				return rw.Write(ctx, sqlapi.NewQuery(queries.UserDecrementResponseCount).WithArgs(userId)).Error
			},
			func() error { // Выставляем флаг для Cron-а о необходимости пересчета уровня развития пользователя
				return rw.Write(ctx, sqlapi.NewQuery(queries.UserMarkUserNeedLevelRecalc).WithArgs(userId)).Error
			},
			func() error { // Уменьшаем счетчик количества отзывов на произведение
				return rw.Write(ctx, sqlapi.NewQuery(queries.WorkStatsDecrementResponseCount).WithArgs(workId)).Error
			},
		)
	})
}
