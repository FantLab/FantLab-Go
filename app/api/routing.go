package api

import (
	"fantlab/api/internal/endpoints"
	"fantlab/shared"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func MakeRouter(config *shared.AppConfig, services *shared.Services) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	api := endpoints.MakeAPI(config, services)
	m := &middlewares{services: services}

	r.Route("/v1", func(r chi.Router) {
		r.Use(m.detectUser)

		// Доступно всем
		r.Group(func(r chi.Router) {
			r.Get("/forums", httpHandler(api.ShowForums))
			r.Get("/forums/{id}", httpHandler(api.ShowForumTopics))
			r.Get("/topics/{id}", httpHandler(api.ShowTopicMessages))
			r.Get("/communities", httpHandler(api.ShowCommunities))
			r.Get("/communities/{id}", httpHandler(api.ShowCommunity))
			r.Get("/blogs", httpHandler(api.ShowBlogs))
			r.Get("/blogs/{id}", httpHandler(api.ShowBlog))
			r.Get("/blog_articles/{id}", httpHandler(api.ShowArticle))
			r.Get("/allgenres", httpHandler(api.ShowGenres))
		})

		// Доступно только анонимам
		r.Group(func(r chi.Router) {
			r.Use(m.anonymousIsRequired)

			r.Post("/login", httpHandler(api.Login))
		})

		r.Group(func(r chi.Router) {
			r.Use(m.authorizedUserIsRequired)

			r.Delete("/logout", httpHandler(api.Logout))
		})

		// Требуется авторизация и проверка на бан
		r.Group(func(r chi.Router) {
			r.Use(m.authorizedUserIsRequired)
			r.Use(m.checkUserIsBanned)

			r.Post("/communities/{id}/subscription", httpHandler(api.SubscribeCommunity))
			r.Delete("/communities/{id}/subscription", httpHandler(api.UnsubscribeCommunity))
			r.Post("/blogs/{id}/subscription", httpHandler(api.SubscribeBlog))
			r.Delete("/blogs/{id}/subscription", httpHandler(api.UnsubscribeBlog))
			r.Post("/blog_articles/{id}/subscription", httpHandler(api.SubscribeArticle))
			r.Delete("/blog_articles/{id}/subscription", httpHandler(api.UnsubscribeArticle))
			r.Post("/blog_articles/{id}/like", httpHandler(api.LikeArticle))
			r.Delete("/blog_articles/{id}/like", httpHandler(api.DislikeArticle))
			r.Put("/work/{id}/genres", httpHandler(api.SetWorkGenres))
		})
	})

	return r
}
