package endpoints

import (
	"encoding/json"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func (api *API) ChangeBookcasesOrder(r *http.Request) (int, proto.Message) {
	var params struct {
		// новый порядок сортировки в формате {"bookcaseId1":index1,...,"bookcaseIdN":indexN}, indexN > 0
		Order string `http:"order,form"`
	}

	api.bindParams(&params, r)

	var newOrder map[uint64]uint64

	err := json.Unmarshal([]byte(params.Order), &newOrder)

	if err != nil {
		return api.badParam("order")
	}

	userId := api.getUserId(r)

	var bookcaseIds []uint64
	for bookcaseId, newIndex := range newOrder {
		if newIndex == 0 {
			return api.badParam("order")
		}
		bookcaseIds = append(bookcaseIds, bookcaseId)
	}

	dbBookcases, err := api.services.DB().FetchUserBookcases(r.Context(), userId, bookcaseIds)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if len(dbBookcases) != len(bookcaseIds) {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: "Не все id книжных полок указаны верно",
		}
	}

	oldOrder := make(map[uint64]uint64, len(dbBookcases))
	for _, dbBookcase := range dbBookcases {
		oldOrder[dbBookcase.BookcaseId] = dbBookcase.Sort
	}

	// Оставляем только полки, у которых действительно изменился порядок сортировки, чтобы в дальнейшем не делать
	// лишних запросов к базе
	finalOrder := map[uint64]uint64{}
	for bookcaseId, newIndex := range newOrder {
		if oldOrder[bookcaseId] != newIndex {
			finalOrder[bookcaseId] = newIndex
		}
	}

	dbBookcases, err = api.services.DB().UpdateBookcasesOrder(r.Context(), userId, finalOrder)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	bookcasesResponse := converters.GetBookcases(dbBookcases)

	return http.StatusOK, bookcasesResponse
}
