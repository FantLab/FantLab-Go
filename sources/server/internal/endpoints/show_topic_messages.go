package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"fantlab/server/internal/helpers"
	"net/http"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (api *API) ShowTopicMessages(r *http.Request) (int, proto.Message) {
	params := struct {
		// айди темы
		TopicId uint64 `http:"id,path"`
		// номер страницы (по умолчанию - 1)
		Page uint64 `http:"page,query"`
		// кол-во записей на странице (по умолчанию - 20)
		Limit uint64 `http:"limit,query"`
		// порядок выдачи (asc - по умолчанию, desc)
		Order string `http:"order,query"`
	}{
		Page:  1,
		Limit: api.config.ForumMessagesInPage,
		Order: "asc",
	}

	api.bindParams(&params, r)

	if params.TopicId == 0 {
		return api.badParam("id")
	}
	if params.Page == 0 {
		return api.badParam("page")
	}
	if !helpers.IsValidLimit(params.Limit) {
		return api.badParam("limit")
	}

	availableForums := api.getAvailableForums(r)

	dbResponse, err := api.services.DB().FetchTopicMessages(
		r.Context(),
		availableForums,
		params.TopicId,
		params.Limit,
		params.Limit*(params.Page-1),
		strings.ToUpper(params.Order) == "ASC",
	)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.TopicId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	topicMessages := converters.GetTopic(dbResponse, params.Page, params.Limit, api.config)
	return http.StatusOK, topicMessages
}
