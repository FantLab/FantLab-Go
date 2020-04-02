package routes

import (
	"fantlab/base/httprouter"
	"fantlab/base/protobuf"
	"fantlab/base/sharedconfig"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"fantlab/server/internal/logs"
	"fantlab/server/routing"
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

func catchPanicMiddleware(panicHandler func(interface{})) httprouter.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if p := recover(); p != nil {
					if panicHandler != nil {
						panicHandler(p)
					}
					protobuf.Handle(func(r *http.Request) (int, proto.Message) {
						return http.StatusInternalServerError, &pb.Error_Response{
							Status: pb.Error_SOMETHING_WENT_WRONG,
						}
					}).ServeHTTP(w, r)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func globalMiddlewares() []httprouter.Middleware {
	if sharedconfig.IsDebug() {
		return []httprouter.Middleware{
			catchPanicMiddleware(func(e interface{}) {
				logs.Logger().Error(fmt.Sprintf("%v", e))
			}),
		}
	}
	return []httprouter.Middleware{
		catchPanicMiddleware(nil),
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				apmhttp.Wrap(next, apmhttp.WithServerRequestName(func(r *http.Request) string {
					return r.Context().Value(httprouter.PathKey).(string)
				}), apmhttp.WithPanicPropagation()).ServeHTTP(w, r)
			})
		},
	}
}
