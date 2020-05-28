package endpoints

import (
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) AddBookcaseItem(r *http.Request) (int, proto.Message) {
	var params struct {
		// id полки
		BookcaseId uint64 `http:"id,path"`
		// группа, в которую входит полка (edition - издания, work - произведения, film - фильмы)
		Group string `http:"group,form"`
		// id item-а, который необходимо добавить на полку (editionId для изданий etc)
		ItemId uint64 `http:"item_id,form"`
	}

	api.bindParams(&params, r)

	if params.BookcaseId == 0 {
		return api.badParam("id")
	}
	if _, ok := converters.BookcaseGroupTitleMap[params.Group]; !ok {
		return api.badParam("group")
	}
	if params.ItemId == 0 {
		return api.badParam("item_id")
	}

	dbBookcase, err := api.services.DB().FetchTypedBookcase(r.Context(), params.Group, params.BookcaseId)

	if err != nil {
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.BookcaseId, 10),
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

	var dbItemId uint64
	itemIds := []uint64{params.ItemId}
	switch params.Group {
	case db.BookcaseEditionType:
		var dbEditions []db.Edition
		dbEditions, err = api.services.DB().FetchEditions(r.Context(), itemIds)
		if len(dbEditions) != 0 {
			dbItemId = dbEditions[0].EditionId
		}
	case db.BookcaseWorkType:
		var dbWorks []db.Work
		dbWorks, err = api.services.DB().FetchWorks(r.Context(), itemIds)
		if len(dbWorks) != 0 {
			dbItemId = dbWorks[0].WorkId
		}
	case db.BookcaseFilmType:
		var dbFilms []db.Film
		dbFilms, err = api.services.DB().FetchFilms(r.Context(), itemIds)
		if len(dbFilms) != 0 {
			dbItemId = dbFilms[0].FilmId
		}
	}

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if dbItemId == 0 {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: strconv.FormatUint(params.ItemId, 10),
		}
	}

	err = api.services.DB().InsertBookcaseItem(r.Context(), dbBookcase.BookcaseId, params.Group, dbItemId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
