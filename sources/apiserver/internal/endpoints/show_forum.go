package endpoints

import (
	"fantlab/core/converters"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) ShowForumTopics(r *http.Request) (int, proto.Message) {
	params := struct {
		// id форума
		ForumId uint64 `http:"id,path"`
		// номер страницы (по умолчанию - 1)
		Page uint64 `http:"page,query"`
		// кол-во записей на странице (по умолчанию - 20)
		Limit uint64 `http:"limit,query"`
	}{
		Page:  1,
		Limit: api.services.AppConfig().ForumTopicsInPage,
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

	availableForums := api.getAvailableForums(r)

	isForumExists, err := api.services.DB().FetchForumExists(r.Context(), params.ForumId, availableForums)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if !isForumExists {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: strconv.FormatUint(params.ForumId, 10),
		}
	}

	userId := api.getUserId(r)

	dbResponse, err := api.services.DB().FetchForumTopics(r.Context(), userId, params.ForumId, params.Limit, params.Limit*(params.Page-1))

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// NOTE Пропущена следующая логика Perl-бэка:
	// 1. workgroup_referee - модераторы (поскольку задается хардкодом в Auth.pm)
	// 2. хардкод конкретных юзеров как модераторов
	// 3. списки тем, посвященных отдельным фильмам, режиссерам, авторам и тд.
	// 4. список пользователей, находящихся сейчас на этом форуме

	forumTopics := converters.GetForumTopics(dbResponse, params.Page, params.Limit, api.services.AppConfig())
	return http.StatusOK, forumTopics
}
