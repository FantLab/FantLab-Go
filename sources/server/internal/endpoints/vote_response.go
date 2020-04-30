package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"fantlab/server/internal/db"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) VoteResponse(r *http.Request) (int, proto.Message) {
	var params struct {
		// id отзыва
		ResponseId uint64 `http:"id,path"`
		// голос (плюс - plus, минус - minus)
		Vote string `http:"vote,form"`
	}

	api.bindParams(&params, r)

	if params.ResponseId == 0 {
		return api.badParam("id")
	}
	if _, ok := db.ResponseVoteMap[params.Vote]; !ok {
		return api.badParam("vote")
	}

	dbResponse, err := api.services.DB().FetchResponse(r.Context(), params.ResponseId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.ResponseId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	user := api.getUser(r)

	// TODO Пропущен весь кусок логики относительно Profile->no_vote_minus, поскольку этот запрет на выставление минусов
	//  задан хардкодом в Auth.pm
	if db.ResponseVoteMap[params.Vote] == -1 && !(user.OwnResponsesRating >= api.config.MinUserOwnResponsesRatingForMinusAbility) {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Вы не можете ставить минусы отзывам",
		}
	}

	if user.UserId == dbResponse.UserId {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Нельзя оценить собственный отзыв",
		}
	}

	responseUserVoteCount, err := api.services.DB().FetchResponseUserVoteCount(r.Context(), user.UserId, dbResponse.ResponseId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if responseUserVoteCount > 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Вы уже оценивали данный отзыв",
		}
	}

	dbResponse, err = api.services.DB().UpdateResponseVotes(r.Context(), dbResponse.ResponseId, user.UserId, params.Vote)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	response := converters.GetResponseRating(dbResponse)

	return http.StatusOK, response
}
