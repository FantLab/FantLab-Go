package endpoints

import (
	"fantlab/core/app"
	"fantlab/core/converters"
	"fantlab/core/helpers"
	"fantlab/pb"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
)

func (api *API) EditForumMessage(r *http.Request) (int, proto.Message) {
	var params struct {
		// id сообщения
		MessageId uint64 `http:"id,path"`
		// новый текст сообщения
		Message string `http:"message,form"`
	}

	api.bindParams(&params, r)

	if params.MessageId == 0 {
		return api.badParam("id")
	}

	availableForums := api.getAvailableForums(r)

	isMessageExists, err := api.services.DB().FetchForumMessageExists(r.Context(), params.MessageId, availableForums)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if !isMessageExists {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: strconv.FormatUint(params.MessageId, 10),
		}
	}

	dbMessage, err := api.services.DB().FetchForumMessage(r.Context(), params.MessageId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if dbMessage.UserId == 0 {
		// В базе есть сообщения, у которых user_id = 0. Визуально помечается как "Автор удален"
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Запрещено редактировать сообщения удаленных пользователей",
		}
	}

	user := api.getUser(r)

	additionalInfo, err := api.services.DB().FetchAdditionalMessageInfo(r.Context(), dbMessage.MessageId, dbMessage.TopicId,
		dbMessage.ForumId, user.UserId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	userIsForumModerator := additionalInfo.ForumModerators[user.UserId]
	topicStarterCanEditFirstMessage := additionalInfo.TopicStarterCanEditFirstMessage
	onlyForAdminsForum := additionalInfo.OnlyForAdminsForum

	userCanPerformAdminActions := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_PERFORM_ADMIN_ACTIONS)
	userCanEditOwnForumMessages := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_EDIT_OWN_FORUM_MESSAGES_WITHOUT_TIME_RESTRICTION)

	// TODO:
	//  1. В коде метода Forum.pm#EditMessageOk есть логика, касающаяся переноса сообщений между темами. Есть смысл
	//     вынести этот функционал отдельным endpoint-ом.
	//  2. Пропущена обработка Profile->workgroup_referee, поскольку оно реализовано хардкодом в Auth.pm
	//  3. Пропущен хардкод про права rusty_cat править FAQ

	isTimeUp := uint64(time.Since(dbMessage.DateOfAdd).Seconds()) > api.services.AppConfig().MaxForumMessageEditTimeout

	// Еще не вышло время редактирования
	//  или пользователь может редактировать свои сообщения без ограничения по времени
	//  или это первое сообщение темы и модератор разрешил его автору правки
	canUserEditMessage := !isTimeUp || userCanEditOwnForumMessages || (dbMessage.Number == 1 && topicStarterCanEditFirstMessage)

	isMessageEditable := dbMessage.IsCensored == 0 && dbMessage.IsRed == 0

	if !(user.UserId == dbMessage.UserId && canUserEditMessage && isMessageEditable) && !userIsForumModerator && !onlyForAdminsForum {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Вы не можете отредактировать данное сообщение",
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

	if formattedMessageLength > api.services.AppConfig().MaxForumMessageLength && user.UserId != api.services.AppConfig().BotUserId {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: fmt.Sprintf("Текст сообщения слишком длинный (больше %d символов после форматирования)", api.services.AppConfig().MaxForumMessageLength),
		}
	}

	var isRed uint8
	if userIsForumModerator && messageContainsModerTags {
		isRed = 1
	}

	dbMessage, err = api.services.DB().UpdateForumMessage(r.Context(), dbMessage.MessageId, dbMessage.TopicId, formattedMessage, isRed)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	helpers.DeleteForumMessageTextCache(dbMessage.MessageId)

	var attaches []*pb.Common_Attachment

	attachments, _ := api.services.DB().FetchForumMessageAttachments(r.Context(), dbMessage.MessageId)
	for _, attachment := range attachments {
		attaches = append(attaches, &pb.Common_Attachment{
			Url:  api.services.GetFSForumMessageAttachmentUrl(dbMessage.MessageId, attachment.FileName),
			Size: attachment.FileSize,
		})
	}

	files, _ := api.services.GetMinioFiles(r.Context(), app.ForumMessageFileGroup, dbMessage.MessageId)
	for _, file := range files {
		attaches = append(attaches, &pb.Common_Attachment{
			Url:  api.services.GetMinioForumMessageAttachmentUrl(dbMessage.MessageId, file.Name),
			Size: file.Size,
		})
	}

	messageResponse := converters.GetForumTopicMessage(dbMessage, attaches, additionalInfo, api.services.AppConfig(),
		user, userCanPerformAdminActions, userCanEditOwnForumMessages)

	return http.StatusOK, messageResponse
}
