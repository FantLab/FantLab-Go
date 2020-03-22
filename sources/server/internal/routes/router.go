package routes

import (
	"fantlab/base/httprouter"
	"fantlab/base/httputils"
	"fantlab/base/logs"
	"fantlab/base/logs/logger"
	"fantlab/base/protobuf"
	"fantlab/base/routing"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"fmt"
	"net/http"
	"regexp"
	"runtime/debug"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/proto"
)

func fill(x *httprouter.Group, y *routing.Group) {
	for _, mw := range y.Middlewares() {
		x.Middlewares = append(x.Middlewares, mw)
	}

	for _, ep := range y.Endpoints() {
		x.Endpoints = append(x.Endpoints, &httprouter.Endpoint{
			Method:  ep.Method(),
			Path:    ep.Path(),
			Handler: protobuf.Handle(ep.Handler()),
		})
	}

	for _, yy := range y.Subgroups() {
		x.Subgroup(func(xx *httprouter.Group) {
			fill(xx, yy)
		})
	}
}

type contextKey string

const (
	paramsKey       = contextKey("path_params")
	BasePath        = "v1"
	HealthcheckPath = "/ping"
)

func MakeHandler(appConfig *config.AppConfig, services *app.Services, logFunc func(*logger.Request)) (http.Handler, func()) {
	healthy := int32(1)

	routerConfig := &httprouter.Config{
		RootGroup: new(httprouter.Group),
		NotFoundHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if HealthcheckPath == r.URL.Path {
				if atomic.LoadInt32(&healthy) == 1 {
					w.WriteHeader(http.StatusNoContent)
				} else {
					w.WriteHeader(http.StatusServiceUnavailable)
				}
			} else {
				protobuf.Handle(func(r *http.Request) (int, proto.Message) {
					return http.StatusNotFound, &pb.Error_Response{
						Status: pb.Error_NOT_FOUND,
					}
				}).ServeHTTP(w, r)
			}
		}),
		RequestContextParamsKey: paramsKey,
		CommonPrefix:            BasePath,
		PathSegmentValidator:    regexp.MustCompile(`^\w+$`).MatchString,
		GlobalMiddlewares: []httprouter.Middleware{
			httputils.WrapResponseWriter,
			logs.HTTP(logFunc),
			httputils.CatchPanic(func(w http.ResponseWriter, r *http.Request, err interface{}) {
				logs.GetBuffer(r.Context()).Append(logger.Entry{
					Message: string(debug.Stack()),
					Err:     fmt.Errorf("Panic: %v", err),
					Time:    time.Now(),
				})

				protobuf.Handle(func(r *http.Request) (int, proto.Message) {
					return http.StatusInternalServerError, &pb.Error_Response{
						Status: pb.Error_SOMETHING_WENT_WRONG,
					}
				}).ServeHTTP(w, r)
			}),
		},
	}

	fill(routerConfig.RootGroup, Tree(appConfig, services, func(r *http.Request, valueKey string) string {
		if values, ok := r.Context().Value(paramsKey).(map[string]string); ok {
			if values != nil {
				return values[valueKey]
			}
		}
		return ""
	}))

	router, _ := httprouter.NewRouter(routerConfig)

	return router, func() {
		atomic.StoreInt32(&healthy, 0)
	}
}
