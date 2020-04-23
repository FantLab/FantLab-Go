package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"fantlab/server/internal/db"
	"fantlab/server/internal/helpers"
	"net/http"
	"strconv"

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
		// сортировать по: порядку - order (по умолчанию), автору - author, названию - title, году - year
		SortBy string `http:"sort,query"`
	}{
		Page:   1,
		Limit:  api.config.BookcaseItemInPage,
		SortBy: "order",
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
	if _, ok := db.EditionSortMap[params.SortBy]; !ok {
		return api.badParam("sort")
	}

	dbBookcase, err := api.services.DB().FetchBookcase(r.Context(), db.BookcaseEditionType, params.BookcaseId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.BookcaseId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	user := api.getUser(r)

	if dbBookcase.BookcaseShared == 0 && (user == nil || user.UserId != dbBookcase.UserId) {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Приватная книжная полка",
		}
	}

	offset := params.Limit * (params.Page - 1)

	dbResponse, err := api.services.DB().FetchEditionBookcase(r.Context(), dbBookcase.BookcaseId, params.Limit, offset, params.SortBy)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	editionBookcase := converters.GetEditionBookcase(dbResponse, dbBookcase, params.Page, params.Limit, api.config)

	return http.StatusOK, editionBookcase
}
