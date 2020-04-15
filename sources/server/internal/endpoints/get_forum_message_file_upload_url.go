package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fantlab/server/internal/helpers"
	"fmt"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
	"time"
)

func (api *API) GetForumMessageFileUploadUrl(r *http.Request) (int, proto.Message) {
	var params struct {
		// id сообщения
		MessageId uint64 `http:"id,path"`
		// полное имя файла (с расширением)
		FileName string `http:"file_name,query"`
	}

	api.bindParams(&params, r)

	if params.MessageId == 0 {
		return api.badParam("id")
	}

	if !helpers.IsValidFileName(params.FileName) {
		return api.badParam("file_name")
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

	// Все дальнейшие проверки аналогичны таковым при редактировании сообщения

	if dbMessage.UserID == 0 {
		// В базе есть сообщения, у которых user_id = 0. Визуально помечается как "Автор удален"
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Запрещено добавлять аттачи к сообщениям удаленных пользователей",
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
			Context: "Вы не можете добавить аттач к данному сообщению",
		}
	}

	files, err := api.services.GetFiles(r.Context(), app.ForumMessageFileGroup, dbMessage.MessageID)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	fileCount := uint64(len(files))

	if fileCount >= api.config.MaxAttachCountPerMessage {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: fmt.Sprintf("К сообщению уже приаттачено %d файлов, это максимум", api.config.MaxAttachCountPerMessage),
		}
	}

	uploadUrl, err := api.services.GetFileUploadUrl(r.Context(), app.ForumMessageFileGroup, dbMessage.MessageID, params.FileName)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_FileUploadResponse{
		Url: uploadUrl,
	}
}
