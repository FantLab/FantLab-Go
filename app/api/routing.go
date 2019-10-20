package api

import (
	"fantlab/api/internal/endpoints"
	"fantlab/api/internal/middlewares"
	"fantlab/logs"
	"fantlab/logs/logger"
	"fantlab/protobuf"
	"fantlab/shared"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func MakeRouter(config *shared.AppConfig, services *shared.Services, logFunc logger.ToString) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(logs.HTTP(logFunc))

	api := endpoints.MakeAPI(config, services)

	r.Route("/v1", func(r chi.Router) {
		r.Use(middlewares.DetectUser(services))

		// Доступно всем
		r.Group(func(r chi.Router) {
			r.Get("/forums", protobuf.Handle(api.ShowForums))
			r.Get("/forums/{id}", protobuf.Handle(api.ShowForumTopics))
			r.Get("/topics/{id}", protobuf.Handle(api.ShowTopicMessages))
			r.Get("/communities", protobuf.Handle(api.ShowCommunities))
			r.Get("/communities/{id}", protobuf.Handle(api.ShowCommunity))
			r.Get("/blogs", protobuf.Handle(api.ShowBlogs))
			r.Get("/blogs/{id}", protobuf.Handle(api.ShowBlog))
			r.Get("/blog_articles/{id}", protobuf.Handle(api.ShowArticle))
			r.Get("/allgenres", protobuf.Handle(api.ShowGenres))
		})

		// Доступно только анонимам
		r.Group(func(r chi.Router) {
			r.Use(middlewares.RequireAnon)

			r.Post("/login", protobuf.Handle(api.Login))
		})

		// Требуется авторизация
		r.Group(func(r chi.Router) {
			r.Use(middlewares.RequireAuth)

			r.Delete("/logout", protobuf.Handle(api.Logout))
		})

		// Требуется авторизация и проверка на бан
		r.Group(func(r chi.Router) {
			r.Use(middlewares.RequireAuth)
			r.Use(middlewares.CheckBan(services))

			r.Post("/topics/{id}/subscription", protobuf.Handle(api.SubscribeForumTopic))
			r.Delete("/topics/{id}/subscription", protobuf.Handle(api.UnsubscribeForumTopic))
			r.Post("/communities/{id}/subscription", protobuf.Handle(api.SubscribeCommunity))
			r.Delete("/communities/{id}/subscription", protobuf.Handle(api.UnsubscribeCommunity))
			r.Post("/blogs/{id}/subscription", protobuf.Handle(api.SubscribeBlog))
			r.Delete("/blogs/{id}/subscription", protobuf.Handle(api.UnsubscribeBlog))
			r.Post("/blog_articles/{id}/subscription", protobuf.Handle(api.SubscribeArticle))
			r.Delete("/blog_articles/{id}/subscription", protobuf.Handle(api.UnsubscribeArticle))
			r.Post("/blog_articles/{id}/like", protobuf.Handle(api.LikeArticle))
			r.Delete("/blog_articles/{id}/like", protobuf.Handle(api.DislikeArticle))
			r.Put("/work/{id}/genres", protobuf.Handle(api.SetWorkGenres))
		})
	})

	return r
}
