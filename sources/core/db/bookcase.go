package db

import (
	"context"
	"fantlab/core/db/queries"
	"time"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
	"github.com/FantLab/go-kit/database/sqlbuilder"
)

const (
	BookcaseWorkType    = "work"
	BookcaseEditionType = "edition"
	BookcaseFilmType    = "film"
)

var (
	EditionDefaultSortMap = map[string]string{
		"order":  "default",
		"author": "autor",
		"title":  "name",
		"year":   "year",
	}

	EditionSortMap = map[string]string{
		"order":  "bi.item_sort",
		"author": "e.autors",
		"title":  "e.name",
		"year":   "e.year",
	}

	WorkDefaultSortMap = map[string]string{
		"order":      "default",
		"author":     "autor",
		"title":      "name",
		"orig_title": "name_orig",
		"year":       "year",
		"mark_count": "mark_count",
		"avg_mark":   "rating",
	}

	WorkSortMap = map[string]string{
		"order":      "bi.item_sort",
		"author":     "a.shortrusname",
		"title":      "w.rusname",
		"orig_title": "w.name",
		"year":       "w.year",
		"mark_count": "ws.markcount DESC",
		"avg_mark":   "ws.midmark DESC",
	}

	FilmDefaultSortMap = map[string]string{
		"order":      "default",
		"title":      "name",
		"orig_title": "name_orig",
	}

	FilmSortMap = map[string]string{
		"order":      "bi.item_sort",
		"title":      "f.rusname",
		"orig_title": "f.name",
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

type BookcaseItem struct {
	BookcaseItemId uint64 `db:"bookcase_item_id"`
	BookcaseId     uint64 `db:"bookcase_id"`
	ItemId         uint64 `db:"item_id"`
}

type BookcaseEdition struct {
	ItemId            uint64 `db:"item_id"`
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

type BookcaseWork struct {
	ItemId        uint64  `db:"item_id"`
	WorkId        uint64  `db:"work_id"`
	Name          string  `db:"name"`
	AltName       string  `db:"altname"`
	RusName       string  `db:"rusname"`
	Year          int64   `db:"year"` // не uint64, поскольку может быть отрицательным (до н.э.)
	BonusText     string  `db:"bonus_text"`
	Description   string  `db:"description"`
	Published     uint8   `db:"published"`
	WorkTypeId    uint64  `db:"work_type_id"`
	AutorId       uint64  `db:"autor_id"`
	Autor2Id      uint64  `db:"autor2_id"`
	Autor3Id      uint64  `db:"autor3_id"`
	Autor4Id      uint64  `db:"autor4_id"`
	Autor5Id      uint64  `db:"autor5_id"`
	MidMark       float64 `db:"midmark"`
	MarkCount     uint64  `db:"markcount"`
	ResponseCount uint64  `db:"response_count"`
	Comment       string  `db:"comment"`
}

type BookcaseAutor struct {
	AutorId  uint64 `db:"autor_id"`
	RusName  string `db:"rusname"`
	IsOpened uint8  `db:"is_opened"`
}

type BookcaseFilm struct {
	ItemId       uint64 `db:"item_id"`
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
	Editions   []BookcaseEdition
	TotalCount uint64
}

type WorkBookcaseDbResponse struct {
	Works            []BookcaseWork
	Autors           map[uint64]BookcaseAutor
	OwnWorkMarks     map[uint64]uint64
	OwnWorkResponses map[uint64]uint64
	TotalCount       uint64
}

type FilmBookcaseDbResponse struct {
	Films      []BookcaseFilm
	TotalCount uint64
}

func (db *DB) FetchAllUserBookcases(ctx context.Context, userId uint64, isOwner bool) ([]Bookcase, error) {
	var bookcases []Bookcase

	var availabilityCondition string
	if isOwner {
		availabilityCondition = "1"
	} else {
		availabilityCondition = "bookcase_shared = 1"
	}

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetAllUserBookcases).WithArgs(userId).Inject(availabilityCondition), &bookcases)

	if err != nil {
		return nil, err
	}

	return bookcases, nil
}

func (db *DB) FetchUserBookcasesOrder(ctx context.Context, userId uint64, bookcaseIds []uint64) (map[uint64]uint64, error) {
	bookcasesSort := map[uint64]uint64{}

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetUserBookcasesSort).WithArgs(userId, bookcaseIds).FlatArgs(), &bookcasesSort)

	if err != nil {
		return nil, err
	}

	return bookcasesSort, nil
}

func (db *DB) FetchBookcase(ctx context.Context, bookcaseId uint64) (Bookcase, error) {
	var bookcase Bookcase

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetBookcase).WithArgs(bookcaseId), &bookcase)

	if err != nil {
		return Bookcase{}, err
	}

	return bookcase, nil
}

func (db *DB) FetchTypedBookcase(ctx context.Context, bookcaseType string, bookcaseId uint64) (Bookcase, error) {
	var bookcase Bookcase

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetTypedBookcase).WithArgs(bookcaseType, bookcaseId), &bookcase)

	if err != nil {
		return Bookcase{}, err
	}

	return bookcase, nil
}

func (db *DB) FetchBookcaseItem(ctx context.Context, bookcaseItemId uint64) (BookcaseItem, error) {
	var bookcaseItem BookcaseItem

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetBookcaseItem).WithArgs(bookcaseItemId), &bookcaseItem)

	if err != nil {
		return BookcaseItem{}, err
	}

	return bookcaseItem, nil
}

func (db *DB) FetchEditionBookcase(ctx context.Context, bookcaseId, limit, offset uint64, sort string) (EditionBookcaseDbResponse, error) {
	var editions []BookcaseEdition
	var count uint64

	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Получаем список изданий на полке
				sortOrder := EditionSortMap[sort]
				return rw.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetEditionBookcaseItems).WithArgs(bookcaseId, limit, offset).Inject(sortOrder), &editions)
			},
			func() error { // Получаем общее количество изданий на полке
				return rw.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetBookcaseItemCount).WithArgs(bookcaseId), &count)
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

func (db *DB) FetchWorkBookcase(ctx context.Context, bookcaseId, limit, offset uint64, sort string, userId uint64) (WorkBookcaseDbResponse, error) {
	var works []BookcaseWork
	var workIds []uint64
	var autorIds []uint64
	var autors []BookcaseAutor
	var autorsMap = map[uint64]BookcaseAutor{}
	var ownWorkMarks = map[uint64]uint64{}
	var ownWorkResponses = map[uint64]uint64{}
	var count uint64

	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Получаем список произведений на полке
				sortOrder := WorkSortMap[sort]
				err := rw.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetWorkBookcaseItems).WithArgs(bookcaseId, limit, offset).Inject(sortOrder), &works)
				if err == nil {
					for _, work := range works {
						workIds = append(workIds, work.WorkId)
						autorIds = append(autorIds, work.AutorId)
						if work.Autor2Id != 0 {
							autorIds = append(autorIds, work.Autor2Id)
						}
						if work.Autor3Id != 0 {
							autorIds = append(autorIds, work.Autor3Id)
						}
						if work.Autor4Id != 0 {
							autorIds = append(autorIds, work.Autor4Id)
						}
						if work.Autor5Id != 0 {
							autorIds = append(autorIds, work.Autor5Id)
						}
					}
				}
				return err
			},
			func() error { // Получаем данные по авторам
				err := rw.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetWorksAutors).WithArgs(autorIds).FlatArgs(), &autors)
				if err == nil {
					for _, autor := range autors {
						autorsMap[autor.AutorId] = autor
					}
				}
				return err
			},
			func() error {
				if userId != 0 { // Получаем список оценок самого пользователя произведениям с полки
					return rw.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetOwnWorkMarks).WithArgs(workIds, userId).FlatArgs(), &ownWorkMarks)
				} else {
					return nil
				}
			},
			func() error {
				if userId != 0 { // Получаем список произведений с полки, на которые пользователь написал отзыв
					return rw.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetOwnWorkResponses).WithArgs(workIds, userId).FlatArgs(), &ownWorkResponses)
				} else {
					return nil
				}
			},
			func() error { // Получаем общее количество произведений на полке
				return rw.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetBookcaseItemCount).WithArgs(bookcaseId), &count)
			},
		)
	})

	if err != nil {
		return WorkBookcaseDbResponse{}, err
	}

	response := WorkBookcaseDbResponse{
		Works:            works,
		Autors:           autorsMap,
		OwnWorkMarks:     ownWorkMarks,
		OwnWorkResponses: ownWorkResponses,
		TotalCount:       count,
	}

	return response, nil
}

func (db *DB) FetchFilmBookcase(ctx context.Context, bookcaseId, limit, offset uint64, sort string) (FilmBookcaseDbResponse, error) {
	var films []BookcaseFilm
	var count uint64

	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Получаем список фильмов на полке
				sortOrder := FilmSortMap[sort]
				return rw.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetFilmBookcaseItems).WithArgs(bookcaseId, limit, offset).Inject(sortOrder), &films)
			},
			func() error { // Получаем общее количество фильмов на полке
				return rw.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetBookcaseItemCount).WithArgs(bookcaseId), &count)
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

	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Создаем полки
				return rw.Write(ctx, sqlbuilder.InsertInto(queries.BookcasesTable, entries...)).Error
			},
			func() error { // Получаем полки
				return rw.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetAllUserBookcases).WithArgs(userId).Inject("1"), &bookcases)
			},
		)
	})

	if err != nil {
		return nil, err
	}

	return bookcases, nil
}

func (db *DB) InsertBookcase(ctx context.Context, userId uint64, bookcaseType, group, title, description string, isPrivate bool, items []map[uint64]string) (uint64, error) {
	var maxSort uint64
	var bookcaseId uint64

	itemIds := make([]uint64, 0, len(items))

	for _, item := range items {
		for id := range item {
			itemIds = append(itemIds, id)
		}
	}

	type newBookcaseItemEntry struct {
		BookcaseId  uint64    `db:"bookcase_id"`
		ItemId      uint64    `db:"item_id"`
		ItemSort    uint64    `db:"item_sort"`
		ItemComment string    `db:"item_comment"`
		DateOfAdd   time.Time `db:"date_of_add"`
	}

	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Получаем предыдущий максимальный номер полки в группе
				return rw.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetMaxSortForType).WithArgs(userId, bookcaseType), &maxSort)
			},
			func() error { // Создаем полку
				result := rw.Write(ctx, sqlapi.NewQuery(queries.BookcaseInsertBookcase).
					WithArgs(userId, bookcaseType, group, title, description, !isPrivate, maxSort+1, "default"))
				if result.Error == nil {
					bookcaseId = uint64(result.LastInsertId)
				}
				return result.Error
			},
			func() error { // Добавляем item-ы на полку
				entries := make([]interface{}, 0, len(items))
				for index, item := range items {
					for id, comment := range item {
						entries = append(entries, newBookcaseItemEntry{
							BookcaseId:  bookcaseId,
							ItemId:      id,
							ItemSort:    uint64(index) + 1,
							ItemComment: comment,
							DateOfAdd:   time.Now(),
						})
					}
				}
				return rw.Write(ctx, sqlbuilder.InsertInto(queries.BookcaseItemsTable, entries...)).Error
			},
			func() error {
				switch bookcaseType {
				case BookcaseEditionType: // Выставляем флаг для Cron-а для пересчета популярности издания
					return rw.Write(ctx, sqlapi.NewQuery(queries.EditionMarkEditionsNeedPopularityRecalc).WithArgs(itemIds).FlatArgs()).Error
				case BookcaseWorkType: // Выставляем флаг для Cron-а для пересчета популярности произведения
					return rw.Write(ctx, sqlapi.NewQuery(queries.WorkMarkWorksNeedPopularityRecalc).WithArgs(itemIds).FlatArgs()).Error
				default: // BookcaseFilmType
					return nil
				}
			},
		)
	})

	if err != nil {
		return 0, err
	}
	return bookcaseId, nil
}

func (db *DB) InsertBookcaseItem(ctx context.Context, bookcaseId uint64, bookcaseType string, itemId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Добавляем item на полку
				return rw.Write(ctx, sqlapi.NewQuery(queries.BookcaseInsertItem).WithArgs(bookcaseId, itemId, bookcaseId)).Error
			},
			func() error {
				switch bookcaseType {
				case BookcaseEditionType: // Выставляем флаг для Cron-а для пересчета популярности издания
					return rw.Write(ctx, sqlapi.NewQuery(queries.EditionMarkEditionsNeedPopularityRecalc).WithArgs(itemId)).Error
				case BookcaseWorkType: // Выставляем флаг для Cron-а для пересчета популярности произведения
					return rw.Write(ctx, sqlapi.NewQuery(queries.WorkMarkWorksNeedPopularityRecalc).WithArgs(itemId)).Error
				default: // BookcaseFilmType
					return nil
				}
			},
		)
	})
}

func (db *DB) UpdateBookcase(ctx context.Context, bookcaseId uint64, bookcaseType, group, title, description, sort string,
	isPrivate bool, items []map[uint64]string) error {
	oldItemsDateOfAdd := make(map[uint64]time.Time, len(items))

	var defaultSort string
	switch bookcaseType {
	case BookcaseEditionType:
		defaultSort = EditionDefaultSortMap[sort]
	case BookcaseWorkType:
		defaultSort = WorkDefaultSortMap[sort]
	case BookcaseFilmType:
		defaultSort = FilmDefaultSortMap[sort]
	}

	itemIds := make([]uint64, 0, len(items))
	for _, item := range items {
		for id := range item {
			itemIds = append(itemIds, id)
		}
	}

	type newBookcaseItemEntry struct {
		BookcaseId  uint64    `db:"bookcase_id"`
		ItemId      uint64    `db:"item_id"`
		ItemSort    uint64    `db:"item_sort"`
		ItemComment string    `db:"item_comment"`
		DateOfAdd   time.Time `db:"date_of_add"`
	}

	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Обновляем полку
				return rw.Write(ctx, sqlapi.NewQuery(queries.BookcaseUpdateBookcase).
					WithArgs(title, description, !isPrivate, group, defaultSort, bookcaseId)).Error
			},
			func() error { // Получаем даты добавления старых item-ов полки
				return rw.Read(ctx, sqlapi.NewQuery(queries.BookcaseGetBookcaseItemsDateOfAdd).WithArgs(bookcaseId), &oldItemsDateOfAdd)
			},
			func() error { // Удаляем item-ы с полки
				return rw.Write(ctx, sqlapi.NewQuery(queries.BookcaseDeleteBookcaseItems).WithArgs(bookcaseId)).Error
			},
			func() error { // Добавляем item-ы на полку
				entries := make([]interface{}, 0, len(items))
				for index, item := range items {
					for id, comment := range item {
						dateOfAdd := oldItemsDateOfAdd[id]
						if dateOfAdd.IsZero() {
							dateOfAdd = time.Now()
						}
						entries = append(entries, newBookcaseItemEntry{
							BookcaseId:  bookcaseId,
							ItemId:      id,
							ItemSort:    uint64(index) + 1,
							ItemComment: comment,
							DateOfAdd:   dateOfAdd,
						})
					}
				}
				return rw.Write(ctx, sqlbuilder.InsertInto(queries.BookcaseItemsTable, entries...)).Error
			},
			func() error {
				switch bookcaseType {
				case BookcaseEditionType: // Выставляем флаг для Cron-а для пересчета популярности издания
					return rw.Write(ctx, sqlapi.NewQuery(queries.EditionMarkEditionsNeedPopularityRecalc).WithArgs(itemIds).FlatArgs()).Error
				case BookcaseWorkType: // Выставляем флаг для Cron-а для пересчета популярности произведения
					return rw.Write(ctx, sqlapi.NewQuery(queries.WorkMarkWorksNeedPopularityRecalc).WithArgs(itemIds).FlatArgs()).Error
				default: // BookcaseFilmType
					return nil
				}
			},
		)
	})
}

func (db *DB) UpdateBookcasesOrder(ctx context.Context, order map[uint64]uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		for bookcaseId, index := range order {
			// Обновляем порядок сортировки у полки
			err := rw.Write(ctx, sqlapi.NewQuery(queries.BookcaseUpdateSort).WithArgs(index, bookcaseId)).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (db *DB) UpdateBookcaseItemComment(ctx context.Context, bookcaseItemId uint64, text string) error {
	return db.engine.Write(ctx, sqlapi.NewQuery(queries.BookcaseUpdateItemComment).WithArgs(text, bookcaseItemId)).Error
}

func (db *DB) DeleteEditionBookcaseItem(ctx context.Context, bookcaseItemId, editionId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Удаляем издание с полки
				return rw.Write(ctx, sqlapi.NewQuery(queries.BookcaseDeleteItem).WithArgs(bookcaseItemId)).Error
			},
			func() error { // Выставляем флаг для Cron-а для пересчета популярности издания
				return rw.Write(ctx, sqlapi.NewQuery(queries.EditionMarkEditionsNeedPopularityRecalc).WithArgs(editionId)).Error
			},
		)
	})
}

func (db *DB) DeleteWorkBookcaseItem(ctx context.Context, bookcaseItemId, workId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Удаляем произведение с полки
				return rw.Write(ctx, sqlapi.NewQuery(queries.BookcaseDeleteItem).WithArgs(bookcaseItemId)).Error
			},
			func() error { // Выставляем флаг для Cron-а для пересчета популярности произведения
				return rw.Write(ctx, sqlapi.NewQuery(queries.WorkMarkWorksNeedPopularityRecalc).WithArgs(workId)).Error
			},
		)
	})
}

func (db *DB) DeleteFilmBookcaseItem(ctx context.Context, bookcaseItemId uint64) error {
	return db.engine.Write(ctx, sqlapi.NewQuery(queries.BookcaseDeleteItem).WithArgs(bookcaseItemId)).Error
}

func (db *DB) DeleteBookcase(ctx context.Context, bookcaseId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Удаляем содержимое полки
				return rw.Write(ctx, sqlapi.NewQuery(queries.BookcaseDeleteBookcaseItems).WithArgs(bookcaseId)).Error
			},
			func() error { // Удаляем саму полку
				return rw.Write(ctx, sqlapi.NewQuery(queries.BookcaseDeleteBookcase).WithArgs(bookcaseId)).Error
			},
		)
	})
}
