package endpoints

import (
	"fantlab/base/protobuf/pbutils"
	"fantlab/base/reflectutils"
	"fantlab/base/uuid"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"fantlab/server/internal/db"
	"fantlab/server/internal/helpers"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type PathParamGetter = func(r *http.Request, key string) string

func getParamValue(r *http.Request, tagValue string, pathParamGetter PathParamGetter) string {
	s := strings.Split(tagValue, ",")
	if len(s) != 2 {
		return ""
	}

	name, source := s[0], s[1]

	switch source {
	case "path":
		return pathParamGetter(r, name)
	case "form":
		return r.PostFormValue(name)
	case "query":
		return r.URL.Query().Get(name)
	}
	return ""
}

// *******************************************************

type API struct {
	config          *config.AppConfig
	services        *app.Services
	pathParamGetter PathParamGetter
}

func MakeAPI(config *config.AppConfig, services *app.Services, pathParamGetter PathParamGetter) *API {
	return &API{
		config:          config,
		services:        services,
		pathParamGetter: pathParamGetter,
	}
}

// *******************************************************

func (api *API) badParam(name string) (int, proto.Message) {
	return http.StatusBadRequest, &pb.Error_Response{
		Status:  pb.Error_INVALID_PARAMETER,
		Context: "Некорректный параметр: " + name,
	}
}

func (api *API) bindParams(output interface{}, r *http.Request) {
	reflectutils.SetStructValues(output, "http", func(s string) string {
		return getParamValue(r, s, api.pathParamGetter)
	})
}

// *******************************************************

func (api *API) getUser(r *http.Request) *pb.Auth_Claims_UserInfo {
	if auth := app.GetUserAuth(r.Context()); auth != nil {
		return auth.User
	}
	return nil
}

func (api *API) getUserId(r *http.Request) uint64 {
	if user := api.getUser(r); user != nil {
		return user.UserId
	}
	return 0
}

func (api *API) getAvailableForums(r *http.Request) []uint64 {
	if user := api.getUser(r); user != nil {
		if len(user.AvailableForumIds) > 0 {
			return user.AvailableForumIds
		}
	}
	return api.config.DefaultAccessToForums
}

func (api *API) isPermissionGranted(r *http.Request, destPermission pb.Auth_Claims_Permission) bool {
	if user := api.getUser(r); user != nil {
		for _, permission := range user.Permissions {
			if permission == destPermission {
				return true
			}
		}
		return false
	}
	return false
}

// *******************************************************

func (api *API) makeAuthResponse(r *http.Request, issuedAt time.Time, userId uint64, saveFn func(entry *db.AuthTokenEntry) error) (*pb.Auth_AuthResponse, error) {
	// получаем инфу о пользователе

	userInfo, err := api.services.DB().FetchUserInfo(r.Context(), userId)

	if err != nil {
		return nil, err
	}

	// формируем данные токена

	var permissions []pb.Auth_Claims_Permission

	if userInfo.CanEditForumMessages == "1" {
		permissions = append(permissions, pb.Auth_Claims_PERMISSION_CAN_EDIT_OWN_FORUM_MESSAGES)
	}
	if userInfo.CanEditResponses == "1" {
		permissions = append(permissions, pb.Auth_Claims_PERMISSION_CAN_EDIT_ANY_RESPONSES)
	}

	claims := &pb.Auth_Claims{
		TokenId: uuid.Generate(issuedAt),
		Issued:  pbutils.TimestampProto(issuedAt),
		User: &pb.Auth_Claims_UserInfo{
			UserId:             userId,
			Login:              userInfo.Login,
			Gender:             helpers.GetGender(userId, userInfo.Gender),
			Class:              helpers.UserClassMap[userInfo.Class],
			OwnResponsesRating: userInfo.VoteCount,
			AvailableForumIds:  helpers.ParseUints(strings.Split(userInfo.AvailableForums, ",")),
			Permissions:        permissions,
		},
	}

	claimsBytes, err := protojson.Marshal(claims)

	if err != nil {
		return nil, err
	}

	// генерируем новый рефреш токен

	refreshToken := uuid.GenerateNow()

	refreshTokenHash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	// сохраняем токен в базе

	err = saveFn(&db.AuthTokenEntry{
		TokenId:     claims.TokenId,
		UserId:      userId,
		RefreshHash: string(refreshTokenHash),
		IssuedAt:    issuedAt,
		RemoteAddr:  r.RemoteAddr,
		DeviceInfo:  "{}", // TODO:
	})

	if err != nil {
		return nil, err
	}

	// подписываем токен

	signedClaimsBytes := api.services.CryptoCoder().Encode(claimsBytes)

	// успех

	return &pb.Auth_AuthResponse{
		UserId:       userId,
		Token:        string(signedClaimsBytes),
		RefreshToken: refreshToken,
	}, nil
}
