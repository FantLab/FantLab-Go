package endpoints

import (
	"fantlab/core/app"
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
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
		Limit:   api.services.AppConfig().ForumMessagesInPage,
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
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.TopicId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	attachmentsMap := map[uint64][]helpers.File{}
	if dbResponse.PinnedFirstMessage != (db.ForumMessage{}) {
		pinnedFirstMessageId := dbResponse.PinnedFirstMessage.MessageID
		for _, attachment := range dbResponse.PinnedFirstMessageAttachments {
			attachmentsMap[pinnedFirstMessageId] = append(attachmentsMap[pinnedFirstMessageId], helpers.File{
				Name: attachment.FileName,
				Size: attachment.FileSize,
			})
		}
		messageFiles, _ := api.services.GetFiles(r.Context(), app.ForumMessageFileGroup, pinnedFirstMessageId)
		attachmentsMap[pinnedFirstMessageId] = append(attachmentsMap[pinnedFirstMessageId], messageFiles...)
	}
	for _, attachment := range dbResponse.Attachments {
		attachmentsMap[attachment.MessageId] = append(attachmentsMap[attachment.MessageId], helpers.File{
			Name: attachment.FileName,
			Size: attachment.FileSize,
		})
	}
	for _, message := range dbResponse.Messages {
		messageFiles, _ := api.services.GetFiles(r.Context(), app.ForumMessageFileGroup, message.MessageID)
		attachmentsMap[message.MessageID] = append(attachmentsMap[message.MessageID], messageFiles...)
	}

	topicMessages := converters.GetTopic(dbResponse, attachmentsMap, params.Page, params.Limit, api.services.AppConfig())
	return http.StatusOK, topicMessages
}
