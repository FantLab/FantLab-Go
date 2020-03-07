package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) SetForumMessageVoting(r *http.Request) (int, proto.Message) {
	var params struct {
		// id сообщения
		MessageId uint64 `http:"id,path"`
		// плюс посту - plus, минус посту - minus, удалить голос - none (для модераторов)
		Vote string `http:"vote,form"`
	}

	api.bindParams(&params, r)

	if params.MessageId == 0 {
		return api.badParam("id")
	}

	if params.Vote != "plus" && params.Vote != "minus" && params.Vote != "none" {
		return api.badParam("vote")
	}

	availableForums := api.getAvailableForums(r)

	message, err := api.services.DB().FetchForumMessage(r.Context(), params.MessageId, availableForums)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.MessageId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	userId := api.getUserId(r)

	if message.UserID == userId && params.Vote != "none" {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Нельзя оценить собственное сообщение",
		}
	}

	if message.IsRed == 1 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Нельзя оценить сообщение модератора",
		}
	}

	userIsForumModerator, err := api.services.DB().FetchUserIsForumModerator(r.Context(), userId, message.TopicId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if userIsForumModerator && params.Vote != "none" {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Модератор не может оценивать сообщения посетителей",
		}
	}

	if !userIsForumModerator && params.Vote == "none" {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Нельзя удалить оценку сообщения",
		}
	}

	messageUserVoteCount, err := api.services.DB().FetchForumMessageUserVoteCount(r.Context(), userId, params.MessageId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if messageUserVoteCount > 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Вы уже оценивали данное сообщение",
		}
	}

	switch params.Vote {
	case "plus":
		err = api.services.DB().UpdateForumMessageVotedPlus(r.Context(), params.MessageId, userId)
	case "minus":
		err = api.services.DB().UpdateForumMessageVotedMinus(r.Context(), params.MessageId, userId)
	case "none":
		err = api.services.DB().UpdateForumMessageVoteDeleted(r.Context(), params.MessageId)
	}

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
