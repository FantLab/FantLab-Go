package endpoints

import (
	"fantlab/dbtools"
	"fantlab/pb"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/bcrypt"
)

func (api *API) Login(r *http.Request) (int, proto.Message) {
	login := r.PostFormValue("login")
	password := r.PostFormValue("password")

	userData, err := api.services.DB().FetchUserPasswordHash(r.Context(), login)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: login,
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.OldHash), []byte(password))

	if err != nil {
		err = bcrypt.CompareHashAndPassword([]byte(userData.NewHash), []byte(password))
	}

	if err != nil {
		return http.StatusUnauthorized, &pb.Error_Response{
			Status: pb.Error_INVALID_PASSWORD,
		}
	}

	sid := api.generateSessionId()

	dateOfCreate := time.Now()

	ok, err := api.services.DB().InsertNewSession(r.Context(), dateOfCreate, sid, userData.UserID, r.RemoteAddr, r.UserAgent())

	if !ok || err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	_ = api.services.Cache().PutSession(r.Context(), sid, uint64(userData.UserID), dateOfCreate)

	return http.StatusOK, &pb.Auth_LoginResponse{
		UserId:       userData.UserID,
		SessionToken: sid,
	}
}
