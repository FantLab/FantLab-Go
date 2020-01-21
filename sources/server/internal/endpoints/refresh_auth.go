package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"fantlab/server/internal/db"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/bcrypt"
)

// Продлевает сессию с помощью рефреш-токена
func (api *API) RefreshAuth(r *http.Request) (int, proto.Message) {
	var params struct {
		// рефреш-токен, выданный при логине или предыдущем продлении сессии
		RefreshToken string `http:"refresh_token,form"`
	}

	api.bindParams(&params, r)

	auth := app.GetUserAuth(r.Context())

	// ищем токен в базе

	authToken, err := api.services.DB().FetchAuthToken(r.Context(), auth.TokenId)

	if dbtools.IsNotFoundError(err) {
		return http.StatusNotFound, &pb.Error_Response{
			Status: pb.Error_NOT_FOUND,
		}
	} else if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// проверяем рефреш-токен

	err = bcrypt.CompareHashAndPassword([]byte(authToken.RefreshHash), []byte(params.RefreshToken))

	if err != nil {
		return http.StatusUnauthorized, &pb.Error_Response{
			Status: pb.Error_INVALID_REFRESH_TOKEN,
		}
	}

	// проверяем валидность рефреш-токена

	if time.Until(authToken.IssuedAt.Add(config.RefreshTokenTimeout)) < 0 {
		return http.StatusUnauthorized, &pb.Error_Response{
			Status: pb.Error_REFRESH_TOKEN_EXPIRED,
		}
	}

	// выпускаем новый токен

	response, err := api.makeAuthResponse(r, time.Now(), auth.User.UserId, func(entry *db.AuthTokenEntry) error {
		return api.services.DB().ReplaceAuthToken(r.Context(), entry, auth.TokenId)
	})

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// успех

	return http.StatusOK, response
}
