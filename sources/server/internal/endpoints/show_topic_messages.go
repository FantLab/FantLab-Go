package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"fantlab/server/internal/helpers"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
)

func (api *API) ShowTopicMessages(r *http.Request) (int, proto.Message) {
	params := struct {
		// id темы
		TopicId uint64 `http:"id,path"`
		// номер страницы (по умолчанию - 1)
		Page uint64 `http:"page,query"`
		// кол-во записей на странице (по умолчанию - 20)
		Limit uint64 `http:"limit,query"`
		// порядок выдачи (0 - от новых к старым, 1 - наоборот; по умолчанию - 0)
		SortAsc uint8 `http:"sortAsc,query"`
	}{
		Page:    1,
		Limit:   api.config.ForumMessagesInPage,
		SortAsc: 0,
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
	if !(params.SortAsc == 0 || params.SortAsc == 1) {
		return api.badParam("sortAsc")
	}

	availableForums := api.getAvailableForums(r)

	dbResponse, err := api.services.DB().FetchTopicMessages(
		r.Context(),
		availableForums,
		params.TopicId,
		int64(params.Limit),
		int64(params.Limit*(params.Page-1)),
		params.SortAsc == 1,
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
