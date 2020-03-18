package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/helpers"
	"fmt"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) AddForumMessage(r *http.Request) (int, proto.Message) {
	var params struct {
		// id темы
		TopicId uint64 `http:"id,path"`
		// текст сообщения
		Message string `http:"message,form"`
	}

	api.bindParams(&params, r)

	if params.TopicId == 0 {
		return api.badParam("id")
	}

	availableForums := api.getAvailableForums(r)

	dbTopic, err := api.services.DB().FetchForumTopic(r.Context(), availableForums, params.TopicId)

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

	if dbTopic.IsClosed == 1 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Тема закрыта",
		}
	}

	user := api.getUser(r)

	userIsForumModerator, err := api.services.DB().FetchUserIsForumModerator(r.Context(), user.UserId, dbTopic.TopicId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	formattedMessage := helpers.FormatMessage(params.Message)

	messageContainsModerTags := helpers.ContainsModerTags(formattedMessage)

	if !userIsForumModerator && messageContainsModerTags {
		formattedMessage = helpers.RemoveModerTags(formattedMessage)
	}

	formattedMessageLength := uint64(len(formattedMessage))

	if formattedMessageLength == 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Текст сообщения пустой (после форматирования)",
		}
	}

	if formattedMessageLength > api.config.MaxForumMessageLength && user.UserId != api.config.BotUserId {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: fmt.Sprintf("Текст сообщения слишком длинный (больше %d символов после форматирования)", api.config.MaxForumMessageLength),
		}
	}

	var isRed uint8
	if userIsForumModerator && messageContainsModerTags {
		isRed = 1
	}

	err = api.services.DB().InsertForumMessage(r.Context(), dbTopic, user.UserId, user.Login, formattedMessage, isRed, api.config.ForumMessagesInPage)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// инвалидируем кэши подписчиков и текущего юзера (запрос не фейлим в случае ошибки)
	{
		subscribers, _ := api.services.DB().FetchForumTopicSubscribers(r.Context(), dbTopic.TopicId)

		for _, subscriber := range subscribers {
			_ = api.services.DeleteUserCache(r.Context(), subscriber)
		}

		_ = api.services.DeleteUserCache(r.Context(), user.UserId)
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}