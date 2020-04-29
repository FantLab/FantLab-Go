package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/db"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
)

func (api *API) DeleteBookcaseItem(r *http.Request) (int, proto.Message) {
	var params struct {
		// id item-а книжной полки
		BookcaseItemId uint64 `http:"id,path"`
	}

	api.bindParams(&params, r)

	if params.BookcaseItemId == 0 {
		return api.badParam("id")
	}

	dbBookcaseItem, err := api.services.DB().FetchBookcaseItem(r.Context(), params.BookcaseItemId)

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

	dbBookcase, err := api.services.DB().FetchBookcase(r.Context(), dbBookcaseItem.BookcaseId)

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

	switch dbBookcase.BookcaseType {
	case db.BookcaseEditionType:
		err = api.services.DB().DeleteEditionBookcaseItem(r.Context(), dbBookcaseItem.BookcaseItemId, dbBookcaseItem.ItemId)
	case db.BookcaseWorkType:
		err = api.services.DB().DeleteWorkBookcaseItem(r.Context(), dbBookcaseItem.BookcaseItemId, dbBookcaseItem.ItemId)
	case db.BookcaseFilmType:
		err = api.services.DB().DeleteFilmBookcaseItem(r.Context(), dbBookcaseItem.BookcaseItemId)
	}

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
