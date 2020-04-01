package routes

import (
	"fantlab/base/httprouter"
	"fantlab/base/httputils"
	"fantlab/base/protobuf"
	"fantlab/base/routing"
	"fantlab/base/sharedconfig"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"fantlab/server/internal/logs"
	"fmt"
	"net/http"
	"regexp"
	"sync/atomic"

	"go.elastic.co/apm/module/apmhttp"
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

const (
	BasePath        = "v1"
	HealthcheckPath = "/ping"
)

func MakeHandler(appConfig *config.AppConfig, services *app.Services) (http.Handler, func()) {
	healthy := int32(1)

	notFoundHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})

	routerConfig := &httprouter.Config{
		RootGroup:            new(httprouter.Group),
		NotFoundHandler:      notFoundHandler,
		CommonPrefix:         BasePath,
		PathSegmentValidator: regexp.MustCompile(`^\w+$`).MatchString,
		GlobalMiddlewares:    globalMiddlewares(),
	}

	fill(routerConfig.RootGroup, Tree(appConfig, services, func(r *http.Request, valueKey string) string {
		if values, ok := r.Context().Value(httprouter.ParamsKey).(map[string]string); ok {
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

func globalMiddlewares() []httprouter.Middleware {
	if sharedconfig.IsDebug() {
		return []httprouter.Middleware{
			httputils.CatchPanic(func(w http.ResponseWriter, r *http.Request, err interface{}) {
				logs.Logger().Error(fmt.Sprintf("%v", err))

				protobuf.Handle(func(r *http.Request) (int, proto.Message) {
					return http.StatusInternalServerError, &pb.Error_Response{
						Status: pb.Error_SOMETHING_WENT_WRONG,
					}
				}).ServeHTTP(w, r)
			}),
		}
	} else {
		return []httprouter.Middleware{
			httputils.CatchPanic(func(w http.ResponseWriter, r *http.Request, err interface{}) {
				protobuf.Handle(func(r *http.Request) (int, proto.Message) {
					return http.StatusInternalServerError, &pb.Error_Response{
						Status: pb.Error_SOMETHING_WENT_WRONG,
					}
				}).ServeHTTP(w, r)
			}),
			func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					apmhttp.Wrap(next, apmhttp.WithServerRequestName(func(r *http.Request) string {
						return r.Context().Value(httprouter.PathKey).(string)
					}), apmhttp.WithPanicPropagation()).ServeHTTP(w, r)
				})
			},
		}
	}
}
