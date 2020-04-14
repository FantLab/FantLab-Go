package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
	"time"
)

func (api *API) DeleteForumMessageFile(r *http.Request) (int, proto.Message) {
	var params struct {
		// id сообщения
		MessageId uint64 `http:"id,path"`
		// id файла
		FileId uint64 `http:"file_id,form"`
	}

	api.bindParams(&params, r)

	if params.MessageId == 0 {
		return api.badParam("id")
	}

	if params.FileId == 0 {
		return api.badParam("file_id")
	}

	availableForums := api.getAvailableForums(r)

	dbMessage, err := api.services.DB().FetchForumMessage(r.Context(), params.MessageId, availableForums)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.MessageId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	dbFile, err := api.services.DB().FetchForumMessageFile(r.Context(), dbMessage.MessageID, params.FileId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.FileId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// Все дальнейшие проверки аналогичны таковым при редактировании сообщения

	if dbMessage.UserID == 0 {
		// В базе есть сообщения, у которых user_id = 0. Визуально помечается как "Автор удален"
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Запрещено удалять аттачи сообщений удаленных пользователей",
		}
	}

	user := api.getUser(r)

	userIsForumModerator, err := api.services.DB().FetchUserIsForumModerator(r.Context(), user.UserId, dbMessage.TopicId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// TODO:
	//  1. В коде метода Forum.pm#EditMessageOk есть логика, касающаяся переноса сообщений между темами. Есть смысл
	//     вынести этот функционал отдельным endpoint-ом.
	//  2. Пропущена обработка Profile->workgroup_referee, поскольку оно реализовано хардкодом в Auth.pm
	//  3. Пропущен хардкод про права rusty_cat править FAQ

	forum, err := api.services.DB().FetchShortForum(r.Context(), dbMessage.ForumId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	topicStarterCanEditFirstMessage, err := api.services.DB().FetchTopicStarterCanEditFirstMessage(r.Context(), dbMessage.MessageID)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	isTimeUp := uint64(time.Since(dbMessage.DateOfAdd).Seconds()) > api.config.MaxForumMessageEditTimeout
	userCanEditOwnForumMessages := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_EDIT_OWN_FORUM_MESSAGES)

	// Еще не вышло время редактирования
	//  или пользователь может редактировать свои сообщения без ограничения по времени
	//  или это первое сообщение темы и модератор разрешил его автору правки
	canUserEditMessage := !isTimeUp || userCanEditOwnForumMessages || (dbMessage.Number == 1 && topicStarterCanEditFirstMessage)

	isMessageEditable := dbMessage.IsCensored == 0 && dbMessage.IsRed == 0

	if !(user.UserId == dbMessage.UserID && canUserEditMessage && isMessageEditable) && !userIsForumModerator && forum.OnlyForAdmins == 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Вы не можете удалить аттач данного сообщения",
		}
	}

	dbMessage, err = api.services.DB().DeleteForumMessageFile(r.Context(), dbMessage.MessageID, dbFile.FileId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// Удаляем файл, ошибку игнорим
	_ = api.services.DeleteFile(r.Context(), dbFile.FileName, dbFile.DateOfAdd)

	messageResponse := converters.GetForumTopicMessage(dbMessage, api.config)

	return http.StatusOK, messageResponse
}
