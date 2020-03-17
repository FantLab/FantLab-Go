package routes

import (
	"fantlab/base/httprouter"
	"fantlab/base/logs"
	"fantlab/base/logs/logger"
	"fantlab/base/protobuf"
	"fantlab/base/routing"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"net/http"
	"regexp"

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
	paramsKey = contextKey("path_params")
	BasePath  = "v1"
)

func MakeHandler(appConfig *config.AppConfig, services *app.Services, logFunc func(*logger.Request)) http.Handler {
	routerConfig := &httprouter.Config{
		RootGroup: new(httprouter.Group),
		NotFoundHandler: protobuf.Handle(func(r *http.Request) (int, proto.Message) {
			return http.StatusNotFound, &pb.Error_Response{
				Status: pb.Error_NOT_FOUND,
			}
		}),
		RequestContextParamsKey: paramsKey,
		CommonPrefix:            BasePath,
		PathSegmentValidator:    regexp.MustCompile(`^\w+$`).MatchString,
		GlobalMiddlewares: []httprouter.Middleware{
			logs.HTTP(logFunc, protobuf.Handle(func(r *http.Request) (int, proto.Message) {
				return http.StatusInternalServerError, &pb.Error_Response{
					Status: pb.Error_SOMETHING_WENT_WRONG,
				}
			})),
		},
	}

	routerConfig.RootGroup.Endpoint(http.MethodGet, "ping", protobuf.Handle(func(r *http.Request) (int, proto.Message) {
		return http.StatusOK, &pb.Common_SuccessResponse{}
	}))

	fill(routerConfig.RootGroup, Tree(appConfig, services, func(r *http.Request, valueKey string) string {
		value, _ := httprouter.GetValueFromContext(r.Context(), paramsKey, valueKey)
		return value
	}))

	router, _ := httprouter.NewRouter(routerConfig)

	return router
}
