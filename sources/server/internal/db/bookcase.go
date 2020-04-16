package db

import (
	"context"
	"fantlab/base/codeflow"
	"fantlab/base/dbtools/sqlbuilder"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"time"
)

type Bookcase struct {
	BookcaseId      uint64 `db:"bookcase_id"`
	BookcaseType    string `db:"bookcase_type"`
	BookcaseGroup   string `db:"bookcase_group"`
	BookcaseName    string `db:"bookcase_name"`
	BookcaseComment string `db:"bookcase_comment"`
	BookcaseShared  uint8  `db:"bookcase_shared"`
	Sort            uint64 `db:"sort"`
	ItemCount       uint64 `db:"item_count"`
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
			BookcaseType:    "edition",
			BookcaseGroup:   "free",
			BookcaseName:    "Мои книги",
			BookcaseComment: "Книги, имеющиеся в моей библиотеке",
			BookcaseShared:  1,
			Sort:            1,
			DateOfAdd:       now,
		},
		newBookcaseEntry{
			UserId:          userId,
			BookcaseType:    "edition",
			BookcaseGroup:   "sale",
			BookcaseName:    "Продаю",
			BookcaseComment: "Книги, которые я готов продать или обменять",
			BookcaseShared:  1,
			Sort:            2,
			DateOfAdd:       now,
		},
		newBookcaseEntry{
			UserId:          userId,
			BookcaseType:    "edition",
			BookcaseGroup:   "buy",
			BookcaseName:    "Куплю",
			BookcaseComment: "Имею желание приобрести эти книги",
			BookcaseShared:  1,
			Sort:            3,
			DateOfAdd:       now,
		},
		newBookcaseEntry{
			UserId:          userId,
			BookcaseType:    "work",
			BookcaseGroup:   "read",
			BookcaseName:    "Прочитать",
			BookcaseComment: "В очереди на прочтение",
			BookcaseShared:  1,
			Sort:            1,
			DateOfAdd:       now,
		},
		newBookcaseEntry{
			UserId:          userId,
			BookcaseType:    "work",
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
