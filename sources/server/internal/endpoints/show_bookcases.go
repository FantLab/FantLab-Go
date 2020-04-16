package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) ShowBookcases(r *http.Request) (int, proto.Message) {
	var params struct {
		// id пользователя
		UserId uint64 `http:"id,path"`
	}

	api.bindParams(&params, r)

	if params.UserId == 0 {
		return api.badParam("id")
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

	isOwner := user != nil && user.UserId == dbUser.UserId

	dbBookcases, err := api.services.DB().FetchBookcases(r.Context(), dbUser.UserId, isOwner)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	bookcasesResponse := converters.GetBookcases(dbBookcases)

	return http.StatusOK, bookcasesResponse
}
