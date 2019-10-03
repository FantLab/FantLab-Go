package endpoints

import (
	"fantlab/api/internal/endpoints/internal/datahelpers"
	"fantlab/helpers"
	"fantlab/pb"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowBlogs(r *http.Request) (int, proto.Message) {
	page, err := uintQueryParam(r, "page", 1)

	if err != nil {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "page",
		}
	}

	limit, err := uintQueryParam(r, "limit", api.config.BlogsInPage)

	if err != nil || !helpers.IsValidLimit(limit) {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "limit",
		}
	}

	sort := strings.ToLower(queryParam(r, "sort", "update"))
	offset := limit * (page - 1)

	dbResponse, err := api.services.DB().FetchBlogs(uint32(limit), uint32(offset), sort)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	blogs := datahelpers.GetBlogs(dbResponse, uint32(page), uint32(limit), api.config)
	return http.StatusOK, blogs
}
