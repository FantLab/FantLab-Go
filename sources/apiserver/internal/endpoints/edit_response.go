package endpoints

import (
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"fmt"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
	"strings"
)

func (api *API) EditResponse(r *http.Request) (int, proto.Message) {
	var params struct {
		// id отзыва
		ResponseId uint64 `http:"id,path"`
		// новый текст отзыва
		Response string `http:"response,form"`
	}

	api.bindParams(&params, r)

	if params.ResponseId == 0 {
		return api.badParam("id")
	}

	dbResponse, err := api.services.DB().FetchResponse(r.Context(), params.ResponseId)

	if err != nil {
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.ResponseId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	formattedResponse := api.services.AppConfig().Smiles.RemoveFromString(strings.TrimSpace(params.Response))

	if uint64(len(formattedResponse)) < api.services.AppConfig().MinResponseLength {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: fmt.Sprintf("Текст сообщения слишком короткий (меньше %d символов после удаления смайлов)", api.services.AppConfig().MinResponseLength),
		}
	}

	userId := api.getUserId(r)

	userCanEditAnyResponses := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_EDIT_ANY_RESPONSES)

	if !(userId == dbResponse.UserId || userCanEditAnyResponses) {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Вы не можете отредактировать данный отзыв",
		}
	}

	err = api.services.DB().UpdateResponse(r.Context(), dbResponse.ResponseId, formattedResponse, dbResponse.UserId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	_ = api.services.DeleteUserCache(r.Context(), dbResponse.UserId)
	_ = api.services.DeleteHomepageResponsesCache(r.Context())

	helpers.DeleteResponseTextCache(dbResponse.ResponseId)

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
