package endpoints

import (
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) VoteResponse(r *http.Request) (int, proto.Message) {
	var params struct {
		// id отзыва
		ResponseId uint64 `http:"id,path"`
		// голос (true - плюс, false - минус)
		VotePlus bool `http:"vote_plus,form"`
	}

	api.bindParams(&params, r)

	if params.ResponseId == 0 {
		return api.badParam("id")
	}

	dbResponse, err := api.services.DB().FetchResponse(r.Context(), params.ResponseId)

	if err != nil {
		if db.IsNotFoundError(err) {
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

	// NOTE Пропущен весь кусок логики относительно Profile->no_vote_minus, поскольку этот запрет на выставление минусов
	// задан хардкодом в Auth.pm
	if !params.VotePlus && !(user.OwnResponsesRating >= api.services.AppConfig().MinUserOwnResponsesRatingForMinusAbility) {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Вы не можете ставить минусы отзывам",
		}
	}

	if user.UserId == dbResponse.UserId {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
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
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Вы уже оценивали данный отзыв",
		}
	}

	dbResponse, err = api.services.DB().UpdateResponseVotes(r.Context(), dbResponse.ResponseId, user.UserId, params.VotePlus)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	response := converters.GetResponseRating(dbResponse)

	return http.StatusOK, response
}
