package endpoints

import (
	"fantlab/base/bindr"
	"fantlab/base/uuid"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"net/http"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
)

type PathParamGetter = func(r *http.Request, key string) string

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

func (api *API) getSession(r *http.Request) string {
	return r.Header.Get(app.SessionHeader)
}

func (api *API) getUserId(r *http.Request) uint64 {
	return app.GetUserId(r.Context())
}

func (api *API) generateSessionId() string {
	return uuid.GenerateNow()
}

func (api *API) badParam(name string) (int, proto.Message) {
	return http.StatusBadRequest, &pb.Error_Response{
		Status:  pb.Error_INVALID_PARAMETER,
		Context: "Некорректный параметр: " + name,
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
