package endpoints

import (
	"fantlab/keys"
	"fantlab/shared"
	"fantlab/uuid"
	"net/http"
	"strconv"
)

type URLParamGetter = func(r *http.Request, key string) string

type API struct {
	config         *shared.AppConfig
	services       *shared.Services
	urlParamGetter URLParamGetter
}

func MakeAPI(config *shared.AppConfig, services *shared.Services, urlParamGetter URLParamGetter) *API {
	return &API{
		config:         config,
		services:       services,
		urlParamGetter: urlParamGetter,
	}
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

func (api *API) urlParam(r *http.Request, key string) string {
	return api.urlParamGetter(r, key)
}

func (api *API) uintURLParam(r *http.Request, key string) (uint64, error) {
	return strconv.ParseUint(api.urlParam(r, key), 10, 32)
}

func (api *API) queryParam(r *http.Request, key string, defaultValue string) string {
	value := r.URL.Query().Get(key)

	if len(value) > 0 {
		return value
	}

	return defaultValue
}

func (api *API) uintQueryParam(r *http.Request, key string, defaultValue uint64) (uint64, error) {
	return strconv.ParseUint(api.queryParam(r, key, strconv.FormatUint(defaultValue, 10)), 10, 32)
}
