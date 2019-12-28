package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"fantlab/server/internal/helpers"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowCommunity(r *http.Request) (int, proto.Message) {
	params := struct {
		// айди сообщества
		CommunityId uint64 `http:"id,path"`
		// номер страницы (по умолчанию - 1)
		Page uint64 `http:"page,query"`
		// кол-во записей на странице (по умолчанию - 5)
		Limit uint64 `http:"limit,query"`
	}{
		Page:  1,
		Limit: api.config.BlogTopicsInPage,
	}

	api.bindParams(&params, r)

	if params.CommunityId == 0 {
		return api.badParam("id")
	}
	if params.Page == 0 {
		return api.badParam("page")
	}
	if !helpers.IsValidLimit(params.Limit) {
		return api.badParam("limit")
	}

	dbResponse, err := api.services.DB().FetchCommunityTopics(
		r.Context(),
		params.CommunityId,
		params.Limit,
		params.Limit*(params.Page-1),
	)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.CommunityId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	community := converters.GetCommunity(dbResponse, params.Page, params.Limit, api.config)
	return http.StatusOK, community
}
