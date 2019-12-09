package endpoints

import (
	"fantlab/base/bindr"
	"fantlab/base/uuid"
	"fantlab/server/internal/keys"
	"fantlab/server/internal/pb"
	"fantlab/server/internal/shared"
	"net/http"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
)

type PathParamGetter = func(r *http.Request, key string) string

type API struct {
	config          *shared.AppConfig
	services        *shared.Services
	pathParamGetter PathParamGetter
}

func MakeAPI(config *shared.AppConfig, services *shared.Services, pathParamGetter PathParamGetter) *API {
	return &API{
		config:          config,
		services:        services,
		pathParamGetter: pathParamGetter,
	}
}

func (api *API) getSession(r *http.Request) string {
	return r.Header.Get(keys.HeaderSessionId)
}

func (api *API) getUserId(r *http.Request) uint64 {
	return keys.GetUserId(r.Context())
}

func (api *API) generateSessionId() string {
	return uuid.GenerateNow()
}

func (api *API) badParam(name string) (int, proto.Message) {
	return http.StatusBadRequest, &pb.Error_Response{
		Status:  pb.Error_INVALID_PARAMETER,
		Context: name,
	}
}

func (api *API) bindParams(output interface{}, r *http.Request) {
	_ = bindr.BindStruct(output, func(f reflect.StructField) string {
		return getParamValue(r, f, api.pathParamGetter)
	})
}

// *******************************************************

func getParamValue(r *http.Request, f reflect.StructField, pathParamGetter PathParamGetter) string {
	s := strings.Split(f.Tag.Get("http"), ",")
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
