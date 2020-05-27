package endpoints

import (
	"encoding/json"
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (api *API) AddBookcase(r *http.Request) (int, proto.Message) {
	var params struct {
		// название
		Title string `http:"title,form"`
		// группа, в которую входит полка (edition - издания, work - произведения, film - фильмы)
		Group string `http:"group,form"`
		// тип полки (sale - на продажу, buy - купить, read - читать, wait - ожидаю, free - прочее)
		Type string `http:"type,form"`
		// описание, до 50 символов (иначе будет обрезано), может быть пустым
		Description string `http:"description,form"`
		// приватная?
		IsPrivate bool `http:"is_private,form"`
		// item-ы в формате [{"id1":"comment1"},...,{"idN":"commentN"}], id - это editionId для изданий etc, commentN может быть пустым
		Items string `http:"items,form"`
	}

	api.bindParams(&params, r)

	title := strings.Join(strings.Fields(params.Title), " ")

	if len(title) == 0 {
		return api.badParam("title")
	}
	if _, ok := converters.BookcaseGroupTitleMap[params.Group]; !ok {
		return api.badParam("group")
	}
	if _, ok := helpers.BookcaseTypeMap[params.Type]; !ok {
		return api.badParam("type")
	}

	var itemsInfo []map[uint64]string

	err := json.Unmarshal([]byte(params.Items), &itemsInfo)

	if err != nil {
		return api.badParam("items")
	}

	itemIds := make([]uint64, 0, len(itemsInfo))

	for _, itemInfo := range itemsInfo {
		for itemId := range itemInfo {
			itemIds = append(itemIds, itemId)
		}
	}

	var dbItemCount int
	switch params.Group {
	case db.BookcaseEditionType:
		var dbEditions []db.Edition
		dbEditions, err = api.services.DB().FetchEditions(r.Context(), itemIds)
		dbItemCount = len(dbEditions)
	case db.BookcaseWorkType:
		var dbWorks []db.Work
		dbWorks, err = api.services.DB().FetchWorks(r.Context(), itemIds)
		dbItemCount = len(dbWorks)
	case db.BookcaseFilmType:
		var dbFilms []db.Film
		dbFilms, err = api.services.DB().FetchFilms(r.Context(), itemIds)
		dbItemCount = len(dbFilms)
	}

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if dbItemCount != len(itemIds) {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: "Не все id item-ов указаны верно",
		}
	}

	description := strings.TrimSpace(params.Description)
	maxDescriptionLength := 50
	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength]
	}

	userId := api.getUserId(r)

	err = api.services.DB().InsertBookcase(r.Context(), userId, params.Group, params.Type, title, description, params.IsPrivate, itemsInfo)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
