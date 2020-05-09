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
		Limit: api.services.AppConfig().BlogTopicsInPage,
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
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.BlogId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	topicIds := make([]uint64, 0, len(dbResponse.Topics))

	for _, topic := range dbResponse.Topics {
		topicIds = append(topicIds, topic.TopicId)
	}

	viewCounts := api.services.BlogTopicsViewCount(r.Context(), topicIds)

	blog := converters.GetBlog(dbResponse, viewCounts, params.Page, params.Limit, api.services.AppConfig())
	return http.StatusOK, blog
}
