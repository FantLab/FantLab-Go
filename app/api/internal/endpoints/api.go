package endpoints

import (
	"fantlab/keys"
	"fantlab/shared"
	"fantlab/uuid"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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
	return r.Header.Get(keys.HeaderSessionId)
}

func (api *API) getUserId(r *http.Request) uint64 {
	return keys.GetUserId(r.Context())
}

func (api *API) generateSessionId() string {
	return uuid.GenerateNow()
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
