package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/db"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
)

func (api *API) AddEditionBookcaseItem(r *http.Request) (int, proto.Message) {
	var params struct {
		// id полки с изданиями
		BookcaseId uint64 `http:"id,path"`
		// id издания, которое необходимо добавить на полку
		EditionId uint64 `http:"edition_id,form"`
	}

	api.bindParams(&params, r)

	if params.BookcaseId == 0 {
		return api.badParam("id")
	}
	if params.EditionId == 0 {
		return api.badParam("edition_id")
	}

	dbBookcase, err := api.services.DB().FetchTypedBookcase(r.Context(), db.BookcaseEditionType, params.BookcaseId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.BookcaseId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// TODO В Perl-бэке этой проверки нет, так что там можно добавить фантомное издание на полку. Оно будет учитываться
	//  в счетчике содержимого в списке полок, но на самой полке отображаться не будет.
	dbEdition, err := api.services.DB().FetchEdition(r.Context(), params.EditionId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.EditionId, 10),
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
				Context: strconv.FormatUint(params.BookcaseId, 10),
			}
		}
	}

	err = api.services.DB().InsertEditionBookcaseItem(r.Context(), dbBookcase.BookcaseId, dbEdition.EditionId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
