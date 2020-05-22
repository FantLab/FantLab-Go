package endpoints

import (
	"encoding/json"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strings"
)

func (api *API) AddFilmBookcase(r *http.Request) (int, proto.Message) {
	var params struct {
		// название
		Title string `http:"title,form"`
		// тип полки (sale - на продажу, buy - купить, read - читать, wait - ожидаю, free - прочее)
		Type string `http:"type,form"`
		// описание, может быть пустым
		Description string `http:"description,form"`
		// приватная?
		IsPrivate bool `http:"is_private,form"`
		// фильмы в формате [{"filmId1":"comment1"},...,{"filmIdN":"commentN"}], commentN может быть пустым
		Films string `http:"films,form"`
	}

	api.bindParams(&params, r)

	if len(params.Title) == 0 {
		return api.badParam("title")
	}
	if _, ok := helpers.BookcaseTypeMap[params.Type]; !ok {
		return api.badParam("type")
	}

	title := strings.TrimSpace(params.Title)
	title = whitespaceCharactersInRowRegex.ReplaceAllLiteralString(title, " ")

	var filmsInfo []map[uint64]string

	err := json.Unmarshal([]byte(params.Films), &filmsInfo)

	if err != nil {
		return api.badParam("films")
	}

	filmIds := make([]uint64, 0, len(filmsInfo))

	for _, filmInfo := range filmsInfo {
		for filmId := range filmInfo {
			filmIds = append(filmIds, filmId)
		}
	}

	dbFilms, err := api.services.DB().FetchFilms(r.Context(), filmIds)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if len(dbFilms) != len(filmIds) {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: "Не все id фильмов указаны верно",
		}
	}

	description := strings.TrimSpace(params.Description)

	userId := api.getUserId(r)

	err = api.services.DB().InsertBookcase(r.Context(), userId, db.BookcaseFilmType, params.Type, title, description, params.IsPrivate, filmsInfo)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
