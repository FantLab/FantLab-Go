package endpoints

import (
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (api *API) ShowEditionBookcase(r *http.Request) (int, proto.Message) {
	params := struct {
		// id книжной полки
		BookcaseId uint64 `http:"id,path"`
		// номер страницы (>0, по умолчанию - 1)
		Page uint64 `http:"page,query"`
		// кол-во элементов на странице ([5..50], по умолчанию - 50)
		Limit uint64 `http:"limit,query"`
		// сортировать по: порядку - order (по умолчанию, если иное не задано в настройках полки), автору - author, названию - title, году - year
		SortBy string `http:"sort,query"`
	}{
		Page:  1,
		Limit: api.services.AppConfig().BookcaseItemInPage,
	}

	api.bindParams(&params, r)

	if params.BookcaseId == 0 {
		return api.badParam("id")
	}
	if params.Page == 0 {
		return api.badParam("page")
	}
	if !helpers.IsValidLimit(params.Limit) {
		return api.badParam("limit")
	}

	sortBy := params.SortBy

	if len(sortBy) != 0 {
		if _, ok := db.EditionSortMap[sortBy]; !ok {
			return api.badParam("sort")
		}
	}

	dbBookcase, err := api.services.DB().FetchTypedBookcase(r.Context(), db.BookcaseEditionType, params.BookcaseId)

	if err != nil {
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.BookcaseId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	userId := api.getUserId(r)

	if dbBookcase.BookcaseShared == 0 && userId != dbBookcase.UserId {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: strconv.FormatUint(params.BookcaseId, 10),
		}
	}

	if len(sortBy) == 0 {
		for order, defaultSort := range db.EditionDefaultSortMap {
			if defaultSort == dbBookcase.DefaultSort {
				sortBy = order
				break
			}
		}
	}

	dbResponse, err := api.services.DB().FetchEditionBookcase(r.Context(), dbBookcase.BookcaseId, sortBy)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// TODO В Perl издания повторно сортируются и в случае других вариантов сортировки, но это бессмысленно, поскольку
	//  фактический порядок не меняется
	if params.SortBy == "author" {
		err = sortEditionsByAuthor(dbResponse.Editions, func(authorIds []uint64) ([]db.Autor, error) {
			return api.services.DB().FetchAutors(r.Context(), authorIds)
		})

		if err != nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}
	}

	offset := params.Limit * (params.Page - 1)

	leftBound := offset
	if leftBound > uint64(len(dbResponse.Editions)) {
		leftBound = uint64(len(dbResponse.Editions))
	}
	rightBound := offset + params.Limit
	if rightBound > uint64(len(dbResponse.Editions)) {
		rightBound = uint64(len(dbResponse.Editions))
	}

	dbResponse.Editions = dbResponse.Editions[leftBound:rightBound]

	editionBookcase := converters.GetEditionBookcase(dbResponse, dbBookcase, params.Page, params.Limit, api.services.AppConfig())

	return http.StatusOK, editionBookcase
}

// Сортируем: 1. по автору, 2. по порядку сортировки
// TODO В Perl зачем-то сортируется еще и по году (хотя порядок сортировки у всех изданий заведомо разный)
func sortEditionsByAuthor(editions []db.BookcaseEdition, fetchAuthors func([]uint64) ([]db.Autor, error)) error {
	linkedAuthorRegex := regexp.MustCompile(`\[autor=([0-9]+)](.+?)\[/autor]`)

	editionsSort := make(map[uint64]int, len(editions))

	// Для каждого издания составляем список авторов (если у автора нет страницы, пишем 0)
	editionsAuthorIds := make([][]uint64, len(editions))
	for index, edition := range editions {
		// Запоминаем изначальный порядок издания, пригодится при сортировке изданий одного и того же автора
		editionsSort[edition.EditionId] = index
		if len(edition.Autors) > 0 {
			authors := strings.Split(edition.Autors, ", ")
			for _, author := range authors {
				if linkedAuthorRegex.MatchString(author) {
					// В теге autor, по идее, всегда валидный id, поэтому ошибки конвертации быть не может
					authorId, _ := strconv.ParseUint(linkedAuthorRegex.ReplaceAllString(author, "$1"), 10, 0)
					editionsAuthorIds[index] = append(editionsAuthorIds[index], authorId)
				} else {
					editionsAuthorIds[index] = append(editionsAuthorIds[index], 0)
				}
			}
		}
	}

	// Составляем список уникальных авторов
	uniqueAuthorIdsMap := make(map[uint64]bool)
	for _, editionAuthorIds := range editionsAuthorIds {
		for _, editionAuthorId := range editionAuthorIds {
			uniqueAuthorIdsMap[editionAuthorId] = editionAuthorId > 0
		}
	}

	// Перегоняем в слайс
	var uniqueAuthorIds []uint64
	for uniqueAuthorId := range uniqueAuthorIdsMap {
		uniqueAuthorIds = append(uniqueAuthorIds, uniqueAuthorId)
	}

	// Запрашиваем сортировочные имена по каждому автору, у которого есть страница
	dbAuthors, err := fetchAuthors(uniqueAuthorIds)

	if err != nil {
		return err
	}

	// Загоняем сортировочные имена в мапу для быстрого доступа
	authorsSortNames := make(map[uint64]string, len(dbAuthors))
	for _, dbAuthor := range dbAuthors {
		authorsSortNames[dbAuthor.AutorId] = dbAuthor.ShortRusName
	}

	authorNames := make(map[uint64]string, len(editions))
	for editionIndex, edition := range editions {
		if len(editionsAuthorIds[editionIndex]) == 0 {
			// Костыль для вывода в самом верху списка
			authorNames[edition.EditionId] = "   "
		} else {
			// TODO В Perl логика значительно отличается:
			//  1. случаи с авторами в тегах и без обрабатываются по-разному
			//  2. если есть автор в тегах, то учитывается только он, все остальные игнорятся
			//  3. если все авторы без тегов, то бьются по запятой и реверсируются для порядка "фамилия имя"
			//     (но логика кривая, см. https://fantlab.ru/edition3066)
			//  4. есть (вернее, должна была быть) отдельная обработка антологий
			authors := strings.Split(edition.Autors, ", ")
			names := make([]string, len(authors))
			for authorIndex, author := range authors {
				if editionsAuthorIds[editionIndex][authorIndex] > 0 {
					authorId := editionsAuthorIds[editionIndex][authorIndex]
					authorSortName := authorsSortNames[authorId]
					if len(authorSortName) > 0 {
						names[authorIndex] = authorSortName
						continue
					}
				}
				// Ориентируемся на порядок "имя фамилия". Попадаем сюда или когда у автора нет страницы,
				// или когда нет сортировочного имени
				dividedName := strings.Split(author, " ")
				revertedName := []string{dividedName[len(dividedName)-1]}
				revertedName = append(revertedName, dividedName[:len(dividedName)-1]...)
				names[authorIndex] = strings.Join(revertedName, " ")
			}
			authorNames[edition.EditionId] = strings.Join(names, ", ")
		}
	}

	sort.Slice(editions, func(i, j int) bool {
		e1 := editions[i]
		e2 := editions[j]
		if authorNames[e1.EditionId] < authorNames[e2.EditionId] {
			return true
		} else if authorNames[e1.EditionId] > authorNames[e2.EditionId] {
			return false
		}
		return editionsSort[e1.EditionId] < editionsSort[e2.EditionId]
	})

	return nil
}
