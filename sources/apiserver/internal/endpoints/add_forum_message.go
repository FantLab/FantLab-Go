package endpoints

import (
	"fantlab/core/converters"
	"fantlab/core/helpers"
	"fantlab/pb"
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

	isTopicExists, err := api.services.DB().FetchForumTopicExists(r.Context(), params.TopicId, availableForums)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if !isTopicExists {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: strconv.FormatUint(params.TopicId, 10),
		}
	}

	dbTopic, err := api.services.DB().FetchForumTopic(r.Context(), params.TopicId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if dbTopic.IsClosed == 1 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Тема закрыта",
		}
	}

	user := api.getUser(r)

	userIsForumModerator, err := api.services.DB().FetchUserIsForumModerator(r.Context(), user.UserId, dbTopic.ForumId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	userCanPerformAdminActions := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_PERFORM_ADMIN_ACTIONS)
	userCanEditOwnForumMessages := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_EDIT_OWN_FORUM_MESSAGES_WITHOUT_TIME_RESTRICTION)

	formattedMessage := helpers.FormatMessage(params.Message)

	messageContainsModerTags := helpers.ContainsModerTags(formattedMessage)

	if !userIsForumModerator && messageContainsModerTags {
		formattedMessage = helpers.RemoveModerTags(formattedMessage)
	}

	formattedMessageLength := uint64(len(formattedMessage))

	if formattedMessageLength == 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Текст сообщения пустой (после форматирования)",
		}
	}

	if formattedMessageLength > api.services.AppConfig().MaxMessageLength && user.UserId != api.services.AppConfig().BotUserId {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: fmt.Sprintf("Текст сообщения слишком длинный (больше %d символов после форматирования)", api.services.AppConfig().MaxMessageLength),
		}
	}

	var isRed uint8
	if userIsForumModerator && messageContainsModerTags {
		isRed = 1
	}

	dbMessage, err := api.services.DB().InsertForumMessage(r.Context(), dbTopic, user.UserId, user.Login,
		formattedMessage, isRed, api.services.AppConfig().ForumMessagesInPage)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// Инвалидируем кеши подписчиков и текущего юзера (запрос не фейлим в случае ошибки)
	{
		subscribers, _ := api.services.DB().FetchForumTopicSubscribers(r.Context(), dbTopic.TopicId)

		for _, subscriber := range subscribers {
			_ = api.services.DeleteUserCache(r.Context(), subscriber)
		}

		_ = api.services.DeleteUserCache(r.Context(), user.UserId)
	}

	// У свежесозданного сообщения еще нет аттачей
	var attaches []*pb.Common_Attachment

	// NOTE Возможен баг. Поскольку потенциальные ошибки игнорируются (а иначе необходимо откатывать обратно транзакцию
	// БД и все операции с аттачами), вычисление прав текущего юзера может отработать слегка неадекватно. Хоть это и
	// маловероятно на практике
	additionalInfo, _ := api.services.DB().FetchAdditionalMessageInfo(r.Context(), dbMessage.MessageId, dbMessage.TopicId,
		dbMessage.ForumId, user.UserId)

	messageResponse := converters.GetForumTopicMessage(dbMessage, attaches, additionalInfo, api.services.AppConfig(),
		user, userCanPerformAdminActions, userCanEditOwnForumMessages)

	return http.StatusOK, messageResponse
}
