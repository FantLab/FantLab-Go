package endpoints

import (
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"fmt"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) AddPrivateMessage(r *http.Request) (int, proto.Message) {
	var params struct {
		// id пользователя, которому отправляется сообщение
		UserId uint64 `http:"id,path"`
		// текст сообщения
		Message string `http:"message,form"`
		// отправить копию посредством Email?
		SendCopyViaEmail bool `http:"send_copy_via_email, form"`
	}

	api.bindParams(&params, r)

	if params.UserId == 0 {
		return api.badParam("id")
	}

	dbUser, err := api.services.DB().FetchUser(r.Context(), params.UserId)

	if err != nil {
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.UserId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	user := api.getUser(r)

	if user.UserId == dbUser.UserId {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Нельзя отправить сообщение самому себе",
		}
	}

	userCanModeratePrivateMessages := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_MODERATE_PRIVATE_MESSAGES)

	formattedMessage := helpers.FormatMessage(params.Message)

	messageContainsModerTags := helpers.ContainsModerTags(formattedMessage)

	if !userCanModeratePrivateMessages && messageContainsModerTags {
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
	if userCanModeratePrivateMessages && messageContainsModerTags {
		isRed = 1
	}

	var sendCopyViaEmail uint8
	if params.SendCopyViaEmail {
		sendCopyViaEmail = 1
	}

	dbMessage, err := api.services.DB().InsertPrivateMessage(r.Context(), user.UserId, dbUser.UserId, formattedMessage, isRed, sendCopyViaEmail)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// NOTE Пропущена вся логика с аттачами

	messageResponse := converters.GetPrivateMessage(dbMessage, api.services.AppConfig())

	if params.SendCopyViaEmail || user.AlwaysCopyPrivateMessageViaEmail {
		api.services.SendPrivateMessageMail(r.Context(), user.UserId, user.Login, []string{dbUser.Email}, formattedMessage)
	}

	return http.StatusOK, messageResponse
}
