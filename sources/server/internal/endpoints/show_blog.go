package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/server/internal/convers"
	"fantlab/server/internal/helpers"
	"fantlab/server/internal/pb"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowBlog(r *http.Request) (int, proto.Message) {
	params := struct {
		// айди блога
		BlogId uint64 `http:"id,path"`
		// номер страницы (по умолчанию - 1)
		Page uint64 `http:"page,query"`
		// кол-во записей на странице (по умолчанию - 20)
		Limit uint64 `http:"limit,query"`
	}{
		Page:  1,
		Limit: api.config.BlogTopicsInPage,
	}

	api.bindParams(&params, r)

	if params.BlogId == 0 {
		return api.badParam("id")
	}
	if params.Page == 0 {
		return api.badParam("page")
	}
	if !helpers.IsValidLimit(params.Limit) {
		return api.badParam("limit")
	}

	offset := params.Limit * (params.Page - 1)

	dbResponse, err := api.services.DB().FetchBlogTopics(r.Context(), params.BlogId, params.Limit, offset)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.BlogId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	blog := convers.GetBlog(dbResponse, params.Page, params.Limit, api.config)
	return http.StatusOK, blog
}
