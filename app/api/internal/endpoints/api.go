package endpoints

import (
	"fantlab/api/internal/consts"
	"fantlab/shared"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/segmentio/ksuid"
)

type API struct {
	config   *shared.AppConfig
	services *shared.Services
}

func MakeAPI(config *shared.AppConfig, services *shared.Services) *API {
	return &API{config: config, services: services}
}

// *******************************************************

func (api *API) getSession(r *http.Request) string {
	return r.Header.Get(consts.SessionHeader)
}

func (api *API) getUserId(r *http.Request) uint64 {
	return r.Context().Value(consts.UserKey).(uint64)
}

func (api *API) generateSessionId() string {
	return ksuid.New().String()
}

// *******************************************************

func urlParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func uintURLParam(r *http.Request, key string) (uint64, error) {
	return strconv.ParseUint(urlParam(r, key), 10, 32)
}

func queryParam(r *http.Request, key string, defaultValue string) string {
	value := r.URL.Query().Get(key)

	if len(value) > 0 {
		return value
	}

	return defaultValue
}

func uintQueryParam(r *http.Request, key string, defaultValue uint64) (uint64, error) {
	return strconv.ParseUint(queryParam(r, key, strconv.FormatUint(defaultValue, 10)), 10, 32)
}
