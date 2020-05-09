package endpoints

import (
	"fantlab/core/converters"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func (api *API) ShowBlogs(r *http.Request) (int, proto.Message) {
	params := struct {
		// номер страницы (по умолчанию - 1)
		Page uint64 `http:"page,query"`
		// кол-во записей на странице (по умолчанию - 5)
		Limit uint64 `http:"limit,query"`
		// сортировать по (кол-ву тем в блоге - article, кол-ву подписчиков - subscriber, дате обновления от новых к старым - update (по умолчанию))
		SortBy string `http:"sort,query"`
	}{
		Page:   1,
		Limit:  api.services.AppConfig().BlogsInPage,
		SortBy: "update",
	}

	api.bindParams(&params, r)

	if params.Page == 0 {
		return api.badParam("page")
	}
	if !helpers.IsValidLimit(params.Limit) {
		return api.badParam("limit")
	}
	if !(params.SortBy == "article" || params.SortBy == "subscriber" || params.SortBy == "update") {
		return api.badParam("sort")
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

	blogs := converters.GetBlogs(dbResponse, params.Page, params.Limit, api.services.AppConfig())
	return http.StatusOK, blogs
}
