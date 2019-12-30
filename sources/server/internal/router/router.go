package router

import (
	"fantlab/base/logs"
	"fantlab/base/logs/logger"
	"fantlab/base/protobuf"
	"fantlab/base/routing"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/golang/protobuf/proto"
)

func walk(r chi.Router, g *routing.Group) {
	r.Group(func(r chi.Router) {
		r.Use(g.Middlewares()...)

		for _, endpoint := range g.Endpoints() {
			r.Method(endpoint.Method(), endpoint.Path(), protobuf.Handle(endpoint.Handler()))
		}

		for _, sg := range g.Subgroups() {
			walk(r, sg)
		}
	})
}

func MakeRouter(config *config.AppConfig, services *app.Services, logFunc logger.ToString, isDebug bool) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(logs.HTTP(logs.Config{
		NeedsRecover: !isDebug,
		ToString:     logFunc,
		PanicHandler: protobuf.Handle(func(r *http.Request) (int, proto.Message) {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}),
	}))

	r.Route(BasePath, func(r chi.Router) {
		walk(r, Routes(config, services, chi.URLParam))
	})

	return r
}
