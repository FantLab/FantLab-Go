package endpoints

import (
	"fantlab/core/app"
	"fantlab/core/helpers"
	"fantlab/pb"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
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

	// Все дальнейшие проверки аналогичны таковым при редактировании сообщения

	if dbMessage.UserId == 0 {
		// В базе есть сообщения, у которых user_id = 0. Визуально помечается как "Автор удален"
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Запрещено добавлять аттачи к сообщениям удаленных пользователей",
		}
	}

	user := api.getUser(r)

	userIsForumModerator, err := api.services.DB().FetchUserIsForumModerator(r.Context(), user.UserId, dbMessage.ForumId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// NOTE
	// 1. Пропущена обработка Profile->workgroup_referee, поскольку оно реализовано хардкодом в Auth.pm
	// 2. Пропущен хардкод прав на редактирование FAQ (это тоже тема в форуме)

	forum, err := api.services.DB().FetchForum(r.Context(), dbMessage.ForumId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	onlyForAdminsForum := forum.OnlyForAdmins == 1

	shortTopic, err := api.services.DB().FetchForumTopicShort(r.Context(), dbMessage.TopicId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	topicStarterCanEditFirstMessage := shortTopic.IsEditTopicStarter == 1

	isTimeUp := uint64(time.Since(dbMessage.DateOfAdd).Seconds()) > api.services.AppConfig().MaxForumMessageEditTimeout
	userCanEditOwnForumMessages := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_EDIT_OWN_FORUM_MESSAGES_WITHOUT_TIME_RESTRICTION)

	// Еще не вышло время редактирования
	//  или пользователь может редактировать свои сообщения без ограничения по времени
	//  или это первое сообщение темы и модератор разрешил его автору правки
	canUserEditMessage := !isTimeUp || userCanEditOwnForumMessages || (dbMessage.Number == 1 && topicStarterCanEditFirstMessage)

	isMessageEditable := dbMessage.IsCensored == 0 && dbMessage.IsRed == 0

	if !(user.UserId == dbMessage.UserId && canUserEditMessage && isMessageEditable) && !userIsForumModerator && !onlyForAdminsForum {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Вы не можете добавить аттач к данному сообщению",
		}
	}

	attachments, err := api.services.DB().FetchForumMessageAttachments(r.Context(), dbMessage.MessageId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	for _, attachment := range attachments {
		if attachment.FileName == params.FileName {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status:  pb.Error_ACTION_FORBIDDEN,
				Context: "К сообщению уже приаттачен файл с таким именем",
			}
		}
	}

	files, err := api.services.GetMinioFiles(r.Context(), app.ForumMessageFileGroup, dbMessage.MessageId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	for _, file := range files {
		if file.Name == params.FileName {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status:  pb.Error_ACTION_FORBIDDEN,
				Context: "К сообщению уже приаттачен файл с таким именем",
			}
		}
	}

	fileCount := uint64(len(attachments) + len(files))

	if fileCount >= api.services.AppConfig().MaxAttachCountPerMessage {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: fmt.Sprintf("К сообщению уже приаттачено %d файлов, это максимум", api.services.AppConfig().MaxAttachCountPerMessage),
		}
	}

	uploadUrl, err := api.services.GetMinioFileUploadUrl(r.Context(), app.ForumMessageFileGroup, dbMessage.MessageId, params.FileName)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_FileUploadResponse{
		Url: uploadUrl,
	}
}
