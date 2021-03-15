package endpoints

import (
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) VoteForumMessage(r *http.Request) (int, proto.Message) {
	var params struct {
		// id сообщения
		MessageId uint64 `http:"id,path"`
		// голос (true - плюс, false - минус)
		VotePlus bool `http:"vote_plus,form"`
	}

	api.bindParams(&params, r)

	if params.MessageId == 0 {
		return api.badParam("id")
	}

	availableForums := api.getAvailableForums(r)

	isMessageExists, err := api.services.DB().FetchForumMessageExists(r.Context(), params.MessageId, availableForums)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if !isMessageExists {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: strconv.FormatUint(params.MessageId, 10),
		}
	}

	dbMessage, err := api.services.DB().FetchForumMessage(r.Context(), params.MessageId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	var isForumWithEnabledRating bool
	for _, forumId := range api.services.AppConfig().ForumsWithEnabledRating {
		if dbMessage.ForumId == forumId {
			isForumWithEnabledRating = true
			break
		}
	}

	if isForumWithEnabledRating {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "В данном форуме у сообщений нет оценок",
		}
	}

	if !params.VotePlus {
		var isForumWithDisabledMinuses bool
		for _, forumId := range api.services.AppConfig().ForumsWithDisabledMinuses {
			if dbMessage.ForumId == forumId {
				isForumWithDisabledMinuses = true
				break
			}
		}

		if isForumWithDisabledMinuses {
			return http.StatusForbidden, &pb.Error_Response{
				Status:  pb.Error_ACTION_FORBIDDEN,
				Context: "В этом форуме запрещены минусы",
			}
		}
	}

	if dbMessage.IsCensored == 1 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Нельзя оценить зацензурированное сообщение",
		}
	}

	if dbMessage.IsRed == 1 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Нельзя оценить данное сообщение",
		}
	}

	userId := api.getUserId(r)

	var isReadOnlyUser bool
	for _, readOnlyUserId := range api.services.AppConfig().ReadOnlyForumUsers[dbMessage.ForumId] {
		if userId == readOnlyUserId {
			isReadOnlyUser = true
			break
		}
	}

	if isReadOnlyUser {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Вы не можете оценивать сообщения в данном форуме",
		}
	}

	if dbMessage.UserId == userId {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Нельзя оценить собственное сообщение",
		}
	}

	isMessageUserVoteExists, err := api.services.DB().FetchForumMessageUserVoteExists(r.Context(), userId, dbMessage.MessageId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if isMessageUserVoteExists {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Вы уже оценивали данное сообщение",
		}
	}

	err = api.services.DB().UpdateForumMessageVotes(r.Context(), dbMessage.MessageId, userId, params.VotePlus)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
