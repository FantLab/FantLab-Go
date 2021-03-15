package endpoints

import (
	"fantlab/core/app"
	"fantlab/core/converters"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
)

func (api *API) DeleteForumMessageFile(r *http.Request) (int, proto.Message) {
	var params struct {
		// id сообщения
		MessageId uint64 `http:"id,path"`
		// полное имя файла (с расширением)
		FileName string `http:"file_name,form"`
	}

	api.bindParams(&params, r)

	if params.MessageId == 0 {
		return api.badParam("id")
	}

	if !helpers.IsValidFileName(params.FileName) {
		return api.badParam("file_name")
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

	attachments, err := api.services.DB().FetchForumMessageAttachments(r.Context(), dbMessage.MessageId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	attachmentExist := false

	for _, attachment := range attachments {
		if attachment.FileName == params.FileName {
			attachmentExist = true
			break
		}
	}

	if !attachmentExist {
		files, err := api.services.GetMinioFiles(r.Context(), app.ForumMessageFileGroup, dbMessage.MessageId)

		if err != nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}

		fileExist := false

		for _, file := range files {
			if file.Name == params.FileName {
				fileExist = true
				break
			}
		}

		if !fileExist {
			return http.StatusForbidden, &pb.Error_Response{
				Status:  pb.Error_ACTION_FORBIDDEN,
				Context: "Не удалось найти аттач с таким именем",
			}
		}
	}

	// Все дальнейшие проверки аналогичны таковым при редактировании сообщения

	if dbMessage.UserId == 0 {
		// В базе есть сообщения, у которых user_id = 0. Визуально помечается как "Автор удален"
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Запрещено удалять аттачи сообщений удаленных пользователей",
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

	// NOTE
	// 1. Пропущена обработка Profile->workgroup_referee, поскольку оно реализовано хардкодом в Auth.pm
	// 2. Пропущен хардкод прав на редактирование FAQ (это тоже тема в форуме)

	isTimeUp := uint64(time.Since(dbMessage.DateOfAdd).Seconds()) > api.services.AppConfig().MaxForumMessageEditTimeout

	// Еще не вышло время редактирования
	//  или пользователь может редактировать свои сообщения без ограничения по времени
	//  или это первое сообщение темы и модератор разрешил его автору правки
	canUserEditMessage := !isTimeUp || userCanEditOwnForumMessages || (dbMessage.Number == 1 && topicStarterCanEditFirstMessage)

	isMessageEditable := dbMessage.IsCensored == 0 && dbMessage.IsRed == 0

	if !(user.UserId == dbMessage.UserId && canUserEditMessage && isMessageEditable) && !userIsForumModerator && !onlyForAdminsForum {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Вы не можете удалить аттач данного сообщения",
		}
	}

	if attachmentExist {
		app.DeleteForumMessageAttachment(dbMessage.MessageId, params.FileName)
	} else { // Minio file exist
		api.services.DeleteMinioFile(r.Context(), app.ForumMessageFileGroup, dbMessage.MessageId, params.FileName)
	}

	var attaches []*pb.Common_Attachment

	attachments, _ = api.services.DB().FetchForumMessageAttachments(r.Context(), dbMessage.MessageId)
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
