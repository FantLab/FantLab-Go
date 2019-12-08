package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/server/internal/pb"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/bcrypt"
)

// Создает новую сессию пользователя
func (api *API) Login(r *http.Request) (int, proto.Message) {
	var params struct {
		// никнейм пользователя
		Login string `http:"login,form"`
		// пароль
		Password string `http:"password,form"`
	}

	api.bindParams(&params, r)

	userData, err := api.services.DB().FetchUserPasswordHash(r.Context(), params.Login)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: params.Login,
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.OldHash), []byte(params.Password))

	if err != nil {
		err = bcrypt.CompareHashAndPassword([]byte(userData.NewHash), []byte(params.Password))
	}

	if err != nil {
		return http.StatusUnauthorized, &pb.Error_Response{
			Status: pb.Error_INVALID_PASSWORD,
		}
	}

	sid := api.generateSessionId()

	dateOfCreate := time.Now()

	err = api.services.DB().InsertNewSession(r.Context(), dateOfCreate, sid, userData.UserID, r.RemoteAddr, r.UserAgent())

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	_ = api.services.Cache().PutSession(r.Context(), sid, userData.UserID, dateOfCreate)

	return http.StatusOK, &pb.Auth_LoginResponse{
		UserId:       userData.UserID,
		SessionToken: sid,
	}
}
