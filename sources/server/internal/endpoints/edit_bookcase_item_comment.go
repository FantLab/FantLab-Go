package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
	"strings"
)

func (api *API) EditBookcaseItemComment(r *http.Request) (int, proto.Message) {
	var params struct {
		// id item-а книжной полки
		BookcaseItemId uint64 `http:"id,path"`
		// текст комментария
		Comment string `http:"comment,form"`
	}

	api.bindParams(&params, r)

	if params.BookcaseItemId == 0 {
		return api.badParam("id")
	}

	dbBookcase, err := api.services.DB().FetchItemBookcase(r.Context(), params.BookcaseItemId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.BookcaseItemId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	userId := api.getUserId(r)

	if userId != dbBookcase.UserId {
		if dbBookcase.BookcaseShared == 1 {
			return http.StatusForbidden, &pb.Error_Response{
				Status:  pb.Error_ACTION_PERMITTED,
				Context: "Невозможно отредактировать чужую книжную полку",
			}
		} else {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.BookcaseItemId, 10),
			}
		}
	}

	text := strings.TrimSpace(params.Comment)

	err = api.services.DB().UpdateBookcaseItemComment(r.Context(), params.BookcaseItemId, text)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	response := converters.GetItemComment(text)

	return http.StatusOK, response
}
