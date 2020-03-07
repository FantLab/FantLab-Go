package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/db"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/proto"
)

// Создаёт новый аутентификационный токен для пользователя на основе пары логин/пароль
func (api *API) Login(r *http.Request) (int, proto.Message) {
	var params struct {
		// никнейм пользователя
		Login string `http:"login,form"`
		// пароль
		Password string `http:"password,form"`
	}

	api.bindParams(&params, r)

	// ищем юзера в базе

	userLoginInfo, err := api.services.DB().FetchUserLoginInfo(r.Context(), params.Login)

	if dbtools.IsNotFoundError(err) {
		return http.StatusNotFound, &pb.Error_Response{
			Status: pb.Error_NOT_FOUND,
		}
	} else if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// проверяем пароль

	err = bcrypt.CompareHashAndPassword([]byte(userLoginInfo.OldHash), []byte(params.Password))

	if err != nil {
		err = bcrypt.CompareHashAndPassword([]byte(userLoginInfo.NewHash), []byte(params.Password))
	}

	if err != nil {
		return http.StatusUnauthorized, &pb.Error_Response{
			Status: pb.Error_INVALID_PASSWORD,
		}
	}

	// выпускаем новый токен

	response, err := api.makeAuthResponse(r, time.Now(), userLoginInfo.UserId, func(entry *db.AuthTokenEntry) error {
		return api.services.DB().InsertAuthToken(r.Context(), entry)
	})

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// успех

	return http.StatusOK, response
}
