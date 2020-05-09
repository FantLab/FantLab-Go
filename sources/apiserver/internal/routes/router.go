package routes

import (
	"fantlab/apiserver/routing"
	"fantlab/base/protobuf"
	"fantlab/core/app"
	"fantlab/core/logs"
	"fantlab/pb"
	"fmt"
	"net/http"
	"regexp"
	"sync/atomic"

	"github.com/FantLab/go-kit/env"
	"github.com/FantLab/go-kit/http/mux"

	"go.elastic.co/apm/module/apmhttp"
	"google.golang.org/protobuf/proto"
)

func fill(x *mux.Group, y *routing.Group) {
	for _, mw := range y.Middlewares() {
		x.Middlewares = append(x.Middlewares, mw)
	}

	for _, ep := range y.Endpoints() {
		x.Endpoints = append(x.Endpoints, &mux.Endpoint{
			Method:  ep.Method(),
			Path:    ep.Path(),
			Handler: protobuf.Handle(ep.Handler()),
		})
	}

	for _, yy := range y.Subgroups() {
		x.Subgroup(func(xx *mux.Group) {
			fill(xx, yy)
		})
	}
}

const (
	BasePath        = "v1"
	HealthcheckPath = "/ping"
)

func MakeHandler(services *app.Services) (http.Handler, func()) {
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

	routerConfig := &mux.Config{
		RootGroup:            new(mux.Group),
		NotFoundHandler:      notFoundHandler,
		CommonPrefix:         BasePath,
		PathSegmentValidator: regexp.MustCompile(`^\w+$`).MatchString,
		GlobalMiddlewares:    globalMiddlewares(),
	}

	fill(routerConfig.RootGroup, Tree(services, func(r *http.Request, valueKey string) string {
		if values, ok := r.Context().Value(mux.ParamsKey).(map[string]string); ok {
			if values != nil {
				return values[valueKey]
			}
		}
		return ""
	}))

	router, _ := mux.NewRouter(routerConfig)

	return router, func() {
		atomic.StoreInt32(&healthy, 0)
	}
}

func catchPanicMiddleware(panicHandler func(interface{})) mux.Middleware {
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

func globalMiddlewares() []mux.Middleware {
	if env.IsDebug() {
		return []mux.Middleware{
			catchPanicMiddleware(func(e interface{}) {
				logs.Logger().Error(fmt.Sprintf("%v", e))
			}),
		}
	}
	return []mux.Middleware{
		catchPanicMiddleware(nil),
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				apmhttp.Wrap(next, apmhttp.WithServerRequestName(func(r *http.Request) string {
					return r.Context().Value(mux.PathKey).(string)
				}), apmhttp.WithPanicPropagation()).ServeHTTP(w, r)
			})
		},
	}
}
