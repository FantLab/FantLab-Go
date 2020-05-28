package endpoints

import (
	"encoding/json"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (api *API) EditBookcase(r *http.Request) (int, proto.Message) {
	params := struct {
		// id полки
		BookcaseId uint64 `http:"id,path"`
		// название
		Title string `http:"title,form"`
		// тип полки (sale - на продажу, buy - купить, read - читать, wait - ожидаю, free - прочее)
		Type string `http:"type,form"`
		// описание, до 50 символов (иначе будет обрезано), может быть пустым
		Description string `http:"description,form"`
		// приватная?
		IsPrivate bool `http:"is_private,form"`
		// сортировка, издания по: порядку - order (по умолчанию), автору - author, названию - title, году - year
		// произведения по: порядку - order (по умолчанию), автору - author, названию - title, оригинальному названию - orig_title, году - year, количеству оценок - mark_count, средней оценке - avg_mark
		// фильмы по: порядку - order (по умолчанию), названию - title, оригинальному названию - orig_title
		SortBy string `http:"sort,form"`
		// item-ы в формате [{"id1":"comment1"},...,{"idN":"commentN"}], id - это editionId для изданий etc, commentN может быть пустым
		Items string `http:"items,form"`
	}{
		SortBy: "order",
	}

	api.bindParams(&params, r)

	title := strings.Join(strings.Fields(params.Title), " ")

	if len(title) == 0 {
		return api.badParam("title")
	}
	if _, ok := helpers.BookcaseTypeMap[params.Type]; !ok {
		return api.badParam("type")
	}

	dbBookcase, err := api.services.DB().FetchBookcase(r.Context(), params.BookcaseId)

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

	if userId != dbBookcase.UserId {
		if dbBookcase.BookcaseShared == 1 {
			return http.StatusForbidden, &pb.Error_Response{
				Status:  pb.Error_ACTION_PERMITTED,
				Context: "Невозможно отредактировать чужую книжную полку",
			}
		} else {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.BookcaseId, 10),
			}
		}
	}

	var itemsInfo []map[uint64]string

	err = json.Unmarshal([]byte(params.Items), &itemsInfo)

	if err != nil {
		return api.badParam("items")
	}

	itemIds := make([]uint64, 0, len(itemsInfo))

	for _, itemInfo := range itemsInfo {
		for itemId := range itemInfo {
			itemIds = append(itemIds, itemId)
		}
	}

	var dbItemCount int
	switch dbBookcase.BookcaseType {
	case db.BookcaseEditionType:
		if _, ok := db.EditionDefaultSortMap[params.SortBy]; !ok {
			return api.badParam("sort")
		}
		var dbEditions []db.Edition
		dbEditions, err = api.services.DB().FetchEditions(r.Context(), itemIds)
		dbItemCount = len(dbEditions)
	case db.BookcaseWorkType:
		if _, ok := db.WorkDefaultSortMap[params.SortBy]; !ok {
			return api.badParam("sort")
		}
		var dbWorks []db.Work
		dbWorks, err = api.services.DB().FetchWorks(r.Context(), itemIds)
		dbItemCount = len(dbWorks)
	case db.BookcaseFilmType:
		if _, ok := db.FilmDefaultSortMap[params.SortBy]; !ok {
			return api.badParam("sort")
		}
		var dbFilms []db.Film
		dbFilms, err = api.services.DB().FetchFilms(r.Context(), itemIds)
		dbItemCount = len(dbFilms)
	}

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if dbItemCount != len(itemIds) {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: "Не все id item-ов указаны верно",
		}
	}

	description := strings.TrimSpace(params.Description)
	maxDescriptionLength := 50
	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength]
	}

	err = api.services.DB().UpdateBookcase(r.Context(), dbBookcase.BookcaseId, dbBookcase.BookcaseType, params.Type,
		title, description, params.SortBy, params.IsPrivate, itemsInfo)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
