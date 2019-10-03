package endpoints

import (
	"fantlab/api/internal/endpoints/internal/datahelpers"
	"fantlab/dbtools"
	"fantlab/helpers"
	"fantlab/pb"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowCommunity(r *http.Request) (int, proto.Message) {
	communityId, err := uintURLParam(r, "id")

	if err != nil {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "id",
		}
	}

	page, err := uintQueryParam(r, "page", 1)

	if err != nil {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "page",
		}
	}

	limit, err := uintQueryParam(r, "limit", api.config.BlogTopicsInPage)

	if err != nil || !helpers.IsValidLimit(limit) {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "limit",
		}
	}

	offset := limit * (page - 1)

	dbResponse, err := api.services.DB().FetchCommunityTopics(uint32(communityId), uint32(limit), uint32(offset))

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(communityId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	community := datahelpers.GetCommunity(dbResponse, uint32(page), uint32(limit), api.config)
	return http.StatusOK, community
}
