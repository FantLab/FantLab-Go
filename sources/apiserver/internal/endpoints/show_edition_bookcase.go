package endpoints

import (
	"fantlab/apiserver/internal/postprocessing"
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
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

	switch params.SortBy {
	case "author":
		{
			err = postprocessing.SortBookcaseEditionsByAuthor(dbResponse.Editions, func(authorIds []uint64) ([]db.Autor, error) {
				return api.services.DB().FetchAutors(r.Context(), authorIds)
			})
			if err != nil {
				return http.StatusInternalServerError, &pb.Error_Response{
					Status: pb.Error_SOMETHING_WENT_WRONG,
				}
			}
		}
	case "title":
		postprocessing.SortBookcaseEditionsByTitle(dbResponse.Editions)
	case "year":
		postprocessing.SortBookcaseEditionsByYear(dbResponse.Editions)
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
