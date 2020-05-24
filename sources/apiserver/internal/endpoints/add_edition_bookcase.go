package endpoints

import (
	"encoding/json"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (api *API) AddEditionBookcase(r *http.Request) (int, proto.Message) {
	var params struct {
		// название
		Title string `http:"title,form"`
		// тип полки (sale - на продажу, buy - купить, read - читать, wait - ожидаю, free - прочее)
		Type string `http:"type,form"`
		// описание, может быть пустым
		Description string `http:"description,form"`
		// приватная?
		IsPrivate bool `http:"is_private,form"`
		// издания в формате [{"editionId1":"comment1"},...,{"editionIdN":"commentN"}], commentN может быть пустым
		Editions string `http:"editions,form"`
	}

	api.bindParams(&params, r)

	title := strings.Join(strings.Fields(params.Title), " ")

	if len(title) == 0 {
		return api.badParam("title")
	}
	if _, ok := helpers.BookcaseTypeMap[params.Type]; !ok {
		return api.badParam("type")
	}

	var editionsInfo []map[uint64]string

	err := json.Unmarshal([]byte(params.Editions), &editionsInfo)

	if err != nil {
		return api.badParam("editions")
	}

	editionIds := make([]uint64, 0, len(editionsInfo))

	for _, editionInfo := range editionsInfo {
		for editionId := range editionInfo {
			editionIds = append(editionIds, editionId)
		}
	}

	dbEditions, err := api.services.DB().FetchEditions(r.Context(), editionIds)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if len(dbEditions) != len(editionIds) {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: "Не все id изданий указаны верно",
		}
	}

	description := strings.TrimSpace(params.Description)

	userId := api.getUserId(r)

	err = api.services.DB().InsertBookcase(r.Context(), userId, db.BookcaseEditionType, params.Type, title, description, params.IsPrivate, editionsInfo)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
