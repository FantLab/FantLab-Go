package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"fantlab/server/internal/helpers"
	"fmt"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) AddPrivateMessage(r *http.Request) (int, proto.Message) {
	params := struct {
		// id пользователя, которому отправляется сообщение
		UserId uint64 `http:"id,path"`
		// текст сообщения
		Message string `http:"message,form"`
		// отправить копию посредством Email? (да - 1, нет - 0)
		SendCopyViaEmail uint8 `http:"send_copy_via_email, form"`
	}{
		SendCopyViaEmail: 0,
	}

	api.bindParams(&params, r)

	if params.UserId == 0 {
		return api.badParam("id")
	}
	if !(params.SendCopyViaEmail == 0 || params.SendCopyViaEmail == 1) {
		return api.badParam("also_send_via_email")
	}

	dbUser, err := api.services.DB().FetchUser(r.Context(), params.UserId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
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
			Status:  pb.Error_ACTION_PERMITTED,
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
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Текст сообщения пустой (после форматирования)",
		}
	}

	// NOTE По логике Perl-бэка, в отличие от форума, в личке у ботов нет преимуществ в отношении максимального
	// размера сообщения
	if formattedMessageLength > api.config.MaxForumMessageLength {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: fmt.Sprintf("Текст сообщения слишком длинный (больше %d символов после форматирования)", api.config.MaxForumMessageLength),
		}
	}

	var isRed uint8
	if userCanModeratePrivateMessages && messageContainsModerTags {
		isRed = 1
	}

	dbMessage, err := api.services.DB().InsertPrivateMessage(r.Context(), user.UserId, dbUser.UserId, formattedMessage, isRed, params.SendCopyViaEmail)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// TODO Пропущена вся логика с аттачами, пока вроде не нужна

	messageResponse := converters.GetPrivateMessage(dbMessage, api.config)

	if params.SendCopyViaEmail == 1 || user.AlwaysCopyPrivateMessageViaEmail {
		_ = api.services.SendPrivateMessageMail(r.Context(), user.UserId, user.Login, []string{dbUser.Email}, formattedMessage, api.config)
	}

	return http.StatusOK, messageResponse
}
