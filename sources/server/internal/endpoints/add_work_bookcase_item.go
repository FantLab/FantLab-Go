package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/db"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
)

func (api *API) AddWorkBookcaseItem(r *http.Request) (int, proto.Message) {
	var params struct {
		// id полки с произведениями
		BookcaseId uint64 `http:"id,path"`
		// id произведения, которое необходимо добавить на полку
		WorkId uint64 `http:"work_id,form"`
	}

	api.bindParams(&params, r)

	if params.BookcaseId == 0 {
		return api.badParam("id")
	}
	if params.WorkId == 0 {
		return api.badParam("work_id")
	}

	dbBookcase, err := api.services.DB().FetchTypedBookcase(r.Context(), db.BookcaseWorkType, params.BookcaseId)

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

	// TODO В Perl-бэке этой проверки нет, так что там можно добавить фантомное произведение на полку. Оно будет учитываться
	//  в счетчике содержимого в списке полок, но на самой полке отображаться не будет.
	dbWork, err := api.services.DB().FetchWork(r.Context(), params.WorkId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.WorkId, 10),
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

	err = api.services.DB().InsertWorkBookcaseItem(r.Context(), dbBookcase.BookcaseId, dbWork.WorkId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
