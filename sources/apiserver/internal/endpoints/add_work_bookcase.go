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

func (api *API) AddWorkBookcase(r *http.Request) (int, proto.Message) {
	var params struct {
		// название
		Title string `http:"title,form"`
		// тип полки (sale - на продажу, buy - купить, read - читать, wait - ожидаю, free - прочее)
		Type string `http:"type,form"`
		// описание, может быть пустым
		Description string `http:"description,form"`
		// приватная?
		IsPrivate bool `http:"is_private,form"`
		// произведения в формате [{"workId1":"comment1"},...,{"workIdN":"commentN"}], commentN может быть пустым
		Works string `http:"works,form"`
	}

	api.bindParams(&params, r)

	title := strings.Join(strings.Fields(params.Title), " ")

	if len(title) == 0 {
		return api.badParam("title")
	}
	if _, ok := helpers.BookcaseTypeMap[params.Type]; !ok {
		return api.badParam("type")
	}

	var worksInfo []map[uint64]string

	err := json.Unmarshal([]byte(params.Works), &worksInfo)

	if err != nil {
		return api.badParam("works")
	}

	workIds := make([]uint64, 0, len(worksInfo))

	for _, workInfo := range worksInfo {
		for workId := range workInfo {
			workIds = append(workIds, workId)
		}
	}

	dbWorks, err := api.services.DB().FetchWorks(r.Context(), workIds)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if len(dbWorks) != len(workIds) {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: "Не все id произведений указаны верно",
		}
	}

	description := strings.TrimSpace(params.Description)

	userId := api.getUserId(r)

	err = api.services.DB().InsertBookcase(r.Context(), userId, db.BookcaseWorkType, params.Type, title, description, params.IsPrivate, worksInfo)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
