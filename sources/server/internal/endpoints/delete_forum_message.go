package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
	"time"
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

	if dbMessage.UserID == 0 {
		// В базе есть сообщения, у которых user_id = 0. Визуально помечается как "Автор удален"
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Запрещено удалять сообщения удаленных пользователей",
		}
	}

	// Проверка на nil далее опущена, поскольку неавторизованный пользователь сюда не попадет
	user := api.getUser(r)

	userIsForumModerator, err := api.services.DB().FetchUserIsForumModerator(r.Context(), user.UserId, dbMessage.TopicId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// TODO:
	//  1. Пропущена обработка Profile->workgroup_referee, поскольку оно реализовано хардкодом в Auth.pm
	//  2. Пропущен хардкод про то, что creator и vad - модераторы

	timeAgo := time.Since(dbMessage.DateOfAdd).Seconds()
	userCanEditOwnForumMessages := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_EDIT_OWN_FORUM_MESSAGES)

	// Еще не вышло время редактирования
	//  или пользователь может редактировать свои сообщения без ограничения по времени
	// TODO Неясно, почему редактировать сообщение можно в течение 2000с, а удалить - 1800с.
	canUserEditMessage := timeAgo <= 1800 || userCanEditOwnForumMessages

	isMessageEditable := dbMessage.IsCensored == 0 && dbMessage.IsRed == 0

	// Из логики кода получается, что, в отличие от редактирования, удалять сообщения в админских форумах могут
	// только модераторы этих форумов.
	if !(user.UserId == dbMessage.UserID && canUserEditMessage && isMessageEditable) && !userIsForumModerator {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Вы не можете удалить данное сообщение",
		}
	}

	err = api.services.DB().DeleteForumMessage(r.Context(), dbMessage.MessageID, dbMessage.TopicId, dbMessage.ForumId,
		dbMessage.DateOfAdd, api.config.ForumMessagesInPage)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// TODO (FLGO-215):
	//  - удалить кеш текста сообщения
	//  - удалить директорию с аттачами сообщения

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
