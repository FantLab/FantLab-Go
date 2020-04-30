package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
)

func (api *API) DeleteResponse(r *http.Request) (int, proto.Message) {
	var params struct {
		// id отзыва
		ResponseId uint64 `http:"id,path"`
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

	userId := /*api.getUserId(r)*/ uint64(58246)

	userCanEditAnyResponses := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_EDIT_ANY_RESPONSES)

	if !(userId == dbResponse.UserId || userCanEditAnyResponses) {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Вы не можете удалить данный отзыв",
		}
	}

	err = api.services.DB().DeleteResponse(r.Context(), dbResponse.ResponseId, dbResponse.WorkId, dbResponse.UserId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	_ = api.services.DeleteWorkStatCache(r.Context(), dbResponse.WorkId)
	_ = api.services.DeleteUserResponseCache(r.Context(), dbResponse.UserId, dbResponse.WorkId)
	_ = api.services.DeleteUserCache(r.Context(), dbResponse.UserId)
	_ = api.services.DeleteHomepageResponsesCache(r.Context())

	// TODO Удалить файловый кеш отзыва

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
