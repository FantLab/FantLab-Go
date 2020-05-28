package db

import (
	"context"
	"fantlab/core/db/queries"
	"time"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
	"github.com/FantLab/go-kit/database/sqlbuilder"
)

type WorkGenre struct {
	Id       uint64 `db:"work_genre_id"`
	ParentId uint64 `db:"parent_work_genre_id"`
	GroupId  uint64 `db:"work_genre_group_id"`
	Name     string `db:"name"`
	Info     string `db:"description"`
}

type WorkGenreGroup struct {
	Id   uint64 `db:"work_genre_group_id"`
	Name string `db:"name"`
}

type WorkGenresDBResponse struct {
	Genres      []WorkGenre
	GenreGroups []WorkGenreGroup
}

type userWorkGenreEntry struct {
	UserId      uint64    `db:"user_id"`
	WorkId      uint64    `db:"work_id"`
	WorkGenreId uint64    `db:"work_genre_id"`
	DateOfAdd   time.Time `db:"date_of_add"`
}

type workGenreCacheEntry struct {
	WorkId                uint64 `db:"work_id"`
	GenreId               uint64 `db:"work_genre_id"`
	VoteCount             uint64 `db:"vote_count"`
	AllVotesAfterGenreAdd uint64 `db:"all_votes_after_genre_add"`
}

func (db *DB) FetchGenres(ctx context.Context) (response *WorkGenresDBResponse, err error) {
	var genres []WorkGenre
	var genreGroups []WorkGenreGroup

	err = codeflow.Try(
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.Genres), &genres)
		},
		func() error {
			return db.engine.Read(ctx, sqlapi.NewQuery(queries.GenreGroups), &genreGroups)
		},
	)

	if err == nil {
		response = &WorkGenresDBResponse{
			Genres:      genres,
			GenreGroups: genreGroups,
		}
	}

	return
}

func (db *DB) GetUserWorkGenreIds(ctx context.Context, workId, userId uint64) (ids []uint64, err error) {
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.UserWorkGenreIds).WithArgs(workId, userId), &ids)
	return
}

func (db *DB) FetchGenreWorkCounts(ctx context.Context) (stat map[uint64]uint64, err error) {
	stat = make(map[uint64]uint64)
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.GenreWorkCounts), &stat)
	return
}

func (db *DB) FetchWorkGenreVotes(ctx context.Context, workId uint64) (stat map[uint64]uint64, err error) {
	stat = make(map[uint64]uint64)
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.WorkGenreVotes).WithArgs(workId), &stat)
	return
}

func (db *DB) GetWorkClassificationCount(ctx context.Context, workId uint64) (count uint64, err error) {
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.WorkClassificationCount).WithArgs(workId), &count)
	return
}

func (db *DB) GenreVote(ctx context.Context, workId, userId uint64, genreIds []uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		var userClassifCount uint64

		return codeflow.Try(
			func() error { // Удаляем предыдущие выбранные жанры
				return rw.Write(ctx, sqlapi.NewQuery(queries.DeleteUserGenres).WithArgs(userId, workId)).Error
			},
			func() error { // Записываем новые
				now := time.Now()

				newEntries := make([]interface{}, 0, len(genreIds))

				for _, genreId := range genreIds {
					newEntries = append(newEntries, userWorkGenreEntry{
						UserId:      userId,
						WorkId:      workId,
						WorkGenreId: genreId,
						DateOfAdd:   now,
					})
				}

				return rw.Write(ctx, sqlbuilder.InsertInto(queries.UserWorkGenresTable, newEntries...)).Error
			},
			func() error { // Получаем кол-во классификаций пользователя
				return rw.Read(ctx, sqlapi.NewQuery(queries.UserClassifCount).WithArgs(userId), &userClassifCount)
			},
			func() error { // Выставляем флажок пересчета уровня пользователя если изменилось кол-во классификаций
				return rw.Write(ctx, sqlapi.NewQuery(queries.UpdateUserClassifCount).WithArgs(userClassifCount, userId, userClassifCount)).Error
			},
			func() error { // Обновляем кол-во проголосовавших у произведения
				return rw.Write(ctx, sqlapi.NewQuery(queries.UpdateWorkVoterCount).WithArgs(workId)).Error
			},
			func() error { // Обновляем кэш жанров у произведения
				return updateWorkGenreCache(ctx, rw, workId)
			},
		)
	})
}

// https://github.com/parserpro/fantlab/blob/6313324869ceee0ff8cb274d251bc097dfa7e45d/pm/Work.pm#L5144
func updateWorkGenreCache(ctx context.Context, rw sqlapi.ReaderWriter, workId uint64) error {
	voteCounts := make(map[uint64]uint64)
	workClassifsAfterGenreAdd := make(map[uint64]uint64)

	return codeflow.Try(
		func() error { // Получаем распределение голосов по жанрам для выбранного произведения
			return rw.Read(ctx, sqlapi.NewQuery(queries.WorkGenreVoteCounts).WithArgs(workId), &voteCounts)
		},
		func() error { // Посчитаем сколько было всего классификаций (не только по выбранному жанру) с момента добавления каждого жанра (https://fantlab.ru/forum/forum19page1/topic8224page22#msg3179144)
			return rw.Read(ctx, sqlapi.NewQuery(queries.WorkClassifCountAfterGenreAdd).WithArgs(workId), &workClassifsAfterGenreAdd)
		},
		func() error { // Удаляем все текущие записи
			return rw.Write(ctx, sqlapi.NewQuery(queries.DeleteWorkGenreCache).WithArgs(workId)).Error
		},
		func() error { // Записываем новые
			entries := make([]interface{}, 0, len(voteCounts))
			for genreId, voteCount := range voteCounts {
				entries = append(entries, workGenreCacheEntry{
					WorkId:                workId,
					GenreId:               genreId,
					VoteCount:             voteCount,
					AllVotesAfterGenreAdd: workClassifsAfterGenreAdd[genreId],
				})
			}
			return rw.Write(ctx, sqlbuilder.InsertInto(queries.WorkGenreCacheTable, entries...)).Error
		},
	)
}
