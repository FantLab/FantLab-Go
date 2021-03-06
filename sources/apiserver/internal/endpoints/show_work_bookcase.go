package endpoints

import (
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) ShowWorkBookcase(r *http.Request) (int, proto.Message) {
	params := struct {
		// id книжной полки
		BookcaseId uint64 `http:"id,path"`
		// номер страницы (>0, по умолчанию - 1)
		Page uint64 `http:"page,query"`
		// кол-во элементов на странице ([5..50], по умолчанию - 50)
		Limit uint64 `http:"limit,query"`
		// сортировать по: порядку - order (по умолчанию, если иное не задано в настройках полки), автору - author, названию - title, оригинальному названию - orig_title, году - year, количеству оценок - mark_count, средней оценке - avg_mark
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
		if _, ok := db.WorkSortMap[sortBy]; !ok {
			return api.badParam("sort")
		}
	}

	dbBookcase, err := api.services.DB().FetchTypedBookcase(r.Context(), db.BookcaseWorkType, params.BookcaseId)

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
		for order, defaultSort := range db.WorkDefaultSortMap {
			if defaultSort == dbBookcase.DefaultSort {
				sortBy = order
				break
			}
		}
	}

	offset := params.Limit * (params.Page - 1)

	dbResponse, err := api.services.DB().FetchWorkBookcase(r.Context(), dbBookcase.BookcaseId, params.Limit, offset, sortBy, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	editionBookcase := converters.GetWorkBookcase(dbResponse, dbBookcase, params.Page, params.Limit)

	return http.StatusOK, editionBookcase
}
