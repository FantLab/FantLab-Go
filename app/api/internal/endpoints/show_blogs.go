package endpoints

import (
	"fantlab/api/internal/endpoints/internal/datahelpers"
	"fantlab/helpers"
	"fantlab/pb"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowBlogs(r *http.Request) (int, proto.Message) {
	params := struct {
		// номер страницы (по умолчанию - 1)
		Page uint64 `http:"page,query"`
		// кол-во записей на странице (по умолчанию - 5)
		Limit uint64 `http:"limit,query"`
		// сортировать по (кол-ву тем в блоге - article, кол-ву подписчиков - subscriber, дате обновления - update (по умолчанию))
		SortBy string `http:"sort,query"`
	}{
		Page:   1,
		Limit:  api.config.BlogTopicsInPage,
		SortBy: "update",
	}

	api.bindParams(&params, r)

	if params.Page == 0 {
		return api.badParam("page")
	}
	if !helpers.IsValidLimit(params.Limit) {
		return api.badParam("limit")
	}

	dbResponse, err := api.services.DB().FetchBlogs(
		r.Context(),
		params.Limit,
		params.Limit*(params.Page-1),
		params.SortBy,
	)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	blogs := datahelpers.GetBlogs(dbResponse, params.Page, params.Limit, api.config)
	return http.StatusOK, blogs
}
