package endpoints

import (
	"fantlab/core/app"
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) ConfirmForumMessageDraft(r *http.Request) (int, proto.Message) {
	var params struct {
		// id темы
		TopicId uint64 `http:"id,path"`
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

	dbMessageDraft, err := api.services.DB().FetchForumMessageDraft(r.Context(), dbTopic.TopicId, user.UserId)

	if err != nil {
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: "Message draft",
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	formattedMessage := helpers.FormatMessage(dbMessageDraft.Message)

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

	dbMessage, err := api.services.DB().ConfirmForumMessageDraft(r.Context(), dbTopic, user.UserId, user.Login,
		formattedMessage, isRed, api.services.AppConfig().ForumMessagesInPage)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// Инвалидируем кэши подписчиков и текущего юзера (запрос не фейлим в случае ошибки)
	{
		subscribers, _ := api.services.DB().FetchForumTopicSubscribers(r.Context(), dbTopic.TopicId)

		for _, subscriber := range subscribers {
			_ = api.services.DeleteUserCache(r.Context(), subscriber)
		}

		_ = api.services.DeleteUserCache(r.Context(), user.UserId)
	}

	// Аттачи черновика переливаем из их директории в файловой системе в Minio
	draftAttachments, _ := app.GetForumMessageDraftAttachments(user.UserId, dbTopic.TopicId)
	for _, draftAttachment := range draftAttachments {
		file, err := os.Open(fmt.Sprintf("%s/%s", app.GetForumMessageDraftAttachmentsDir(user.UserId, dbTopic.TopicId), draftAttachment.Name))
		if err != nil {
			continue
		}
		_ = api.services.MoveFileFromFSToMinio(r.Context(), app.ForumMessageFileGroup, dbMessage.MessageId, file)
		_ = file.Close()
	}
	// Удаляем директорию с аттачами черновика
	app.DeleteForumMessageDraftAttachments(user.UserId, dbTopic.TopicId)
	// Уже имевшиеся в Minio аттачи черновика превращаем в аттачи сообщения
	_ = api.services.MoveFilesInsideMinio(r.Context(), app.ForumMessageDraftFileGroup, dbMessageDraft.DraftId, app.ForumMessageFileGroup, dbMessage.MessageId)

	var attaches []*pb.Common_Attachment
	files, _ := api.services.GetMinioFiles(r.Context(), app.ForumMessageFileGroup, dbMessage.MessageId)
	for _, file := range files {
		attaches = append(attaches, &pb.Common_Attachment{
			Url:  api.services.GetMinioForumMessageAttachmentUrl(dbMessage.MessageId, file.Name),
			Size: file.Size,
		})
	}

	// NOTE Возможен баг. Поскольку потенциальные ошибки игнорируются (а иначе необходимо откатывать обратно транзакцию
	// БД и все операции с аттачами), вычисление прав текущего юзера может отработать слегка неадекватно. Хоть это и
	// маловероятно на практике
	additionalInfo, _ := api.services.DB().FetchAdditionalMessageInfo(r.Context(), dbMessage.MessageId, dbMessage.TopicId,
		dbMessage.ForumId, user.UserId)

	messageResponse := converters.GetForumTopicMessage(dbMessage, attaches, additionalInfo, api.services.AppConfig(),
		user, userCanPerformAdminActions, userCanEditOwnForumMessages)

	return http.StatusOK, messageResponse
}
