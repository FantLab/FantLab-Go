package routing

import (
	"fantlab/base/protobuf"
	"fantlab/server/internal/logs"
	"fantlab/server/internal/logs/logger"
	"fantlab/server/internal/shared"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func walk(r chi.Router, g *Group) {
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

func MakeRouter(config *shared.AppConfig, services *shared.Services, logFunc logger.ToString) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(logs.HTTP(logFunc))

	r.Route(BasePath, func(r chi.Router) {
		walk(r, Routes(config, services, chi.URLParam))
	})

	return r
}
