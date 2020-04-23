package db

import (
	"context"
	"fantlab/base/codeflow"
	"fantlab/base/dbtools/sqlbuilder"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"time"
)

const (
	BookcaseWorkType    = "work"
	BookcaseEditionType = "edition"
	BookcaseFilmType    = "film"
)

var (
	EditionSortMap = map[string]string{
		"order":  "item_sort",
		"author": "autors",
		"title":  "name",
		"year":   "year",
	}

	FilmSortMap = map[string]string{
		"order":      "item_sort",
		"title":      "rusname",
		"orig_title": "name",
	}
)

type Bookcase struct {
	BookcaseId      uint64 `db:"bookcase_id"`
	UserId          uint64 `db:"user_id"`
	BookcaseType    string `db:"bookcase_type"`
	BookcaseGroup   string `db:"bookcase_group"`
	BookcaseName    string `db:"bookcase_name"`
	BookcaseComment string `db:"bookcase_comment"`
	BookcaseShared  uint8  `db:"bookcase_shared"`
	Sort            uint64 `db:"sort"`
	ItemCount       uint64 `db:"item_count"`
}

type Edition struct {
	EditionId         uint64 `db:"edition_id"`
	Name              string `db:"name"`
	Autors            string `db:"autors"`
	Type              uint64 `db:"type"`
	Year              uint64 `db:"year"`
	Publisher         string `db:"publisher"`
	Description       string `db:"description"`
	Correct           uint64 `db:"correct"`
	PlanDate          string `db:"plan_date"`
	OzonId            uint64 `db:"ozon_id"`
	OzonCost          uint64 `db:"ozon_cost"`
	OzonAvailable     uint8  `db:"ozon_available"`
	LabirintId        uint64 `db:"labirint_id"`
	LabirintCost      uint64 `db:"labirint_cost"`
	LabirintAvailable uint8  `db:"labirint_available"`
	Comment           string `db:"comment"`
}

type Film struct {
	FilmId       uint64 `db:"film_id"`
	Name         string `db:"name"`
	RusName      string `db:"rusname"`
	Type         uint64 `db:"type"`
	Year         uint64 `db:"year"`
	Year2        uint64 `db:"year2"`
	Country      string `db:"country"`
	Genre        string `db:"genre"`
	Director     string `db:"director"`
	ScreenWriter string `db:"screenwriter"`
	Actors       string `db:"actors"`
	Description  string `db:"description"`
	Comment      string `db:"comment"`
}

type EditionBookcaseDbResponse struct {
	Editions   []Edition
	TotalCount uint64
}

type FilmBookcaseDbResponse struct {
	Films      []Film
	TotalCount uint64
}

func (db *DB) FetchBookcases(ctx context.Context, userId uint64, isOwner bool) ([]Bookcase, error) {
	var bookcases []Bookcase

	var availabilityCondition string
	if isOwner {
		availabilityCondition = "1"
	} else {
		availabilityCondition = "bookcase_shared = 1"
	}

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.BookcaseGetBookcases).WithArgs(userId).Inject(availabilityCondition)).Scan(&bookcases)

	if err != nil {
		return nil, err
	}

	return bookcases, nil
}

func (db *DB) FetchBookcase(ctx context.Context, bookcaseType string, bookcaseId uint64) (Bookcase, error) {
	var bookcase Bookcase

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.BookcaseGetBookcase).WithArgs(bookcaseType, bookcaseId)).Scan(&bookcase)

	if err != nil {
		return Bookcase{}, err
	}

	return bookcase, nil
}

func (db *DB) FetchEditionBookcase(ctx context.Context, bookcaseId, limit, offset uint64, sort string) (EditionBookcaseDbResponse, error) {
	var editions []Edition
	var count uint64

	err := db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Получаем список изданий на полке
				sortOrder := EditionSortMap[sort]
				return rw.Read(ctx, sqlr.NewQuery(queries.BookcaseGetEditionBookcaseItems).WithArgs(bookcaseId, limit, offset).Inject(sortOrder)).Scan(&editions)
			},
			func() error { // Получаем общее количество изданий на полке
				return rw.Read(ctx, sqlr.NewQuery(queries.BookcaseGetBookcaseItemCount).WithArgs(bookcaseId)).Scan(&count)
			},
		)
	})

	if err != nil {
		return EditionBookcaseDbResponse{}, err
	}

	response := EditionBookcaseDbResponse{
		Editions:   editions,
		TotalCount: count,
	}

	return response, nil
}

func (db *DB) FetchFilmBookcase(ctx context.Context, bookcaseId, limit, offset uint64, sort string) (FilmBookcaseDbResponse, error) {
	var films []Film
	var count uint64

	err := db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Получаем список фильмов на полке
				sortOrder := FilmSortMap[sort]
				return rw.Read(ctx, sqlr.NewQuery(queries.BookcaseGetFilmBookcaseItems).WithArgs(bookcaseId, limit, offset).Inject(sortOrder)).Scan(&films)
			},
			func() error { // Получаем общее количество фильмов на полке
				return rw.Read(ctx, sqlr.NewQuery(queries.BookcaseGetBookcaseItemCount).WithArgs(bookcaseId)).Scan(&count)
			},
		)
	})

	if err != nil {
		return FilmBookcaseDbResponse{}, err
	}

	response := FilmBookcaseDbResponse{
		Films:      films,
		TotalCount: count,
	}

	return response, nil
}

func (db *DB) InsertDefaultBookcases(ctx context.Context, userId uint64) ([]Bookcase, error) {
	var bookcases []Bookcase

	type newBookcaseEntry struct {
		UserId          uint64    `db:"user_id"`
		BookcaseType    string    `db:"bookcase_type"`
		BookcaseGroup   string    `db:"bookcase_group"`
		BookcaseName    string    `db:"bookcase_name"`
		BookcaseComment string    `db:"bookcase_comment"`
		BookcaseShared  uint8     `db:"bookcase_shared"`
		Sort            uint64    `db:"sort"`
		DateOfAdd       time.Time `db:"date_of_add"`
	}

	now := time.Now()

	entries := []interface{}{
		newBookcaseEntry{
			UserId:          userId,
			BookcaseType:    BookcaseEditionType,
			BookcaseGroup:   "free",
			BookcaseName:    "Мои книги",
			BookcaseComment: "Книги, имеющиеся в моей библиотеке",
			BookcaseShared:  1,
			Sort:            1,
			DateOfAdd:       now,
		},
		newBookcaseEntry{
			UserId:          userId,
			BookcaseType:    BookcaseEditionType,
			BookcaseGroup:   "sale",
			BookcaseName:    "Продаю",
			BookcaseComment: "Книги, которые я готов продать или обменять",
			BookcaseShared:  1,
			Sort:            2,
			DateOfAdd:       now,
		},
		newBookcaseEntry{
			UserId:          userId,
			BookcaseType:    BookcaseEditionType,
			BookcaseGroup:   "buy",
			BookcaseName:    "Куплю",
			BookcaseComment: "Имею желание приобрести эти книги",
			BookcaseShared:  1,
			Sort:            3,
			DateOfAdd:       now,
		},
		newBookcaseEntry{
			UserId:          userId,
			BookcaseType:    BookcaseWorkType,
			BookcaseGroup:   "read",
			BookcaseName:    "Прочитать",
			BookcaseComment: "В очереди на прочтение",
			BookcaseShared:  1,
			Sort:            1,
			DateOfAdd:       now,
		},
		newBookcaseEntry{
			UserId:          userId,
			BookcaseType:    BookcaseWorkType,
			BookcaseGroup:   "wait",
			BookcaseName:    "Ожидаю",
			BookcaseComment: "Ожидаю выхода издания этого произведения, либо его перевод",
			BookcaseShared:  1,
			Sort:            2,
			DateOfAdd:       now,
		},
	}

	err := db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Создаем полки
				return rw.Write(ctx, sqlbuilder.InsertInto(queries.BookcasesTable, entries...)).Error
			},
			func() error { // Получаем полки
				return rw.Read(ctx, sqlr.NewQuery(queries.BookcaseGetBookcases).WithArgs(userId).Inject("1")).Scan(&bookcases)
			},
		)
	})

	if err != nil {
		return nil, err
	}

	return bookcases, nil
}
