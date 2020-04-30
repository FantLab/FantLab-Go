package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
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
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.ResponseId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	userId := api.getUserId(r)

	userCanEditAnyResponses := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_EDIT_ANY_RESPONSES)

	if !(userId == dbResponse.UserId || userCanEditAnyResponses) {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Вы не можете отредактировать данный отзыв",
		}
	}

	if len(params.Response) == 0 || len(strings.TrimSpace(params.Response)) == 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Текст отзыва пустой",
		}
	}

	err = api.services.DB().UpdateResponse(r.Context(), dbResponse.ResponseId, params.Response, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	_ = api.services.DeleteUserCache(r.Context(), userId)
	_ = api.services.DeleteHomepageResponsesCache(r.Context())

	// TODO Удалить файловый кеш отзыва

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
