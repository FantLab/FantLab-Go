package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/base/utils"
	"fantlab/server/internal/convers"
	"fantlab/server/internal/helpers"
	"fantlab/server/internal/pb"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowForumTopics(r *http.Request) (int, proto.Message) {
	params := struct {
		// айди форума
		ForumId uint64 `http:"id,path"`
		// номер страницы (по умолчанию - 1)
		Page uint64 `http:"page,query"`
		// кол-во записей на странице (по умолчанию - 20)
		Limit uint64 `http:"limit,query"`
	}{
		Page:  1,
		Limit: api.config.ForumTopicsInPage,
	}

	api.bindParams(&params, r)

	if params.ForumId == 0 {
		return api.badParam("id")
	}
	if params.Page == 0 {
		return api.badParam("page")
	}
	if !helpers.IsValidLimit(params.Limit) {
		return api.badParam("limit")
	}

	availableForums := api.config.DefaultAccessToForums

	userId := api.getUserId(r)

	if userId > 0 {
		availableForumsString, err := api.services.DB().FetchAvailableForums(r.Context(), userId)

		if err != nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}

		availableForums = utils.ParseUints(strings.Split(availableForumsString, ","))

		if availableForums == nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}
	}

	dbResponse, err := api.services.DB().FetchForumTopics(
		r.Context(),
		availableForums,
		params.ForumId,
		params.Limit,
		params.Limit*(params.Page-1),
	)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.ForumId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	forumTopics := convers.GetForumTopics(dbResponse, params.Page, params.Limit, api.config)
	return http.StatusOK, forumTopics
}
