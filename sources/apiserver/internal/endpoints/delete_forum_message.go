package endpoints

import (
	"fantlab/core/app"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
)

func (api *API) DeleteForumMessage(r *http.Request) (int, proto.Message) {
	var params struct {
		// id сообщения
		MessageId uint64 `http:"id,path"`
	}

	api.bindParams(&params, r)

	if params.MessageId == 0 {
		return api.badParam("id")
	}

	availableForums := api.getAvailableForums(r)

	dbMessage, err := api.services.DB().FetchForumMessage(r.Context(), params.MessageId, availableForums)

	if err != nil {
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.MessageId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if dbMessage.UserID == 0 {
		// В базе есть сообщения, у которых user_id = 0. Визуально помечается как "Автор удален"
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Запрещено удалять сообщения удаленных пользователей",
		}
	}

	userId := api.getUserId(r)

	userIsForumModerator, err := api.services.DB().FetchUserIsForumModerator(r.Context(), userId, dbMessage.TopicId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// TODO:
	//  1. Пропущена обработка Profile->workgroup_referee, поскольку оно реализовано хардкодом в Auth.pm
	//  2. Пропущен хардкод про то, что creator и vad - модераторы

	isTimeUp := uint64(time.Since(dbMessage.DateOfAdd).Seconds()) > api.services.AppConfig().MaxForumMessageEditTimeout
	userCanEditOwnForumMessages := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_EDIT_OWN_FORUM_MESSAGES_WITHOUT_TIME_RESTRICTION)

	// Еще не вышло время редактирования
	//  или пользователь может редактировать свои сообщения без ограничения по времени
	canUserEditMessage := !isTimeUp || userCanEditOwnForumMessages

	isMessageEditable := dbMessage.IsCensored == 0 && dbMessage.IsRed == 0

	// Из логики кода получается, что, в отличие от редактирования, удалять сообщения в админских форумах могут
	// только модераторы этих форумов.
	if !(userId == dbMessage.UserID && canUserEditMessage && isMessageEditable) && !userIsForumModerator {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Вы не можете удалить данное сообщение",
		}
	}

	err = api.services.DB().DeleteForumMessage(r.Context(), dbMessage.MessageID, dbMessage.TopicId, dbMessage.ForumId,
		dbMessage.DateOfAdd, api.services.AppConfig().ForumMessagesInPage)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	helpers.DeleteForumMessageTextCache(dbMessage.MessageID)
	helpers.DeleteForumMessageAttachments(dbMessage.MessageID)
	api.services.DeleteFiles(r.Context(), app.ForumMessageFileGroup, dbMessage.MessageID)

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
