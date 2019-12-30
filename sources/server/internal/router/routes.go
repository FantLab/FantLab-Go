package router

import (
	"fantlab/base/routing"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"fantlab/server/internal/endpoints"
	"fantlab/server/internal/middlewares"
)

const BasePath = "/v1"

func Routes(config *config.AppConfig, services *app.Services, pathParamGetter endpoints.PathParamGetter) *routing.Group {
	api := endpoints.MakeAPI(config, services, pathParamGetter)

	g := new(routing.Group)

	g.Middleware(middlewares.DetectUser(services))

	g.Subgroup("Общедоступные", func(g *routing.Group) {
		g.Endpoint("GET", "/forums", api.ShowForums, "Список форумов")
		g.Endpoint("GET", "/forums/{id}", api.ShowForumTopics, "Список тем форума")
		g.Endpoint("GET", "/topics/{id}", api.ShowTopicMessages, "Сообщения в теме форума")
		g.Endpoint("GET", "/communities", api.ShowCommunities, "Список сообществ")
		g.Endpoint("GET", "/communities/{id}", api.ShowCommunity, "Информация о сообществе")
		g.Endpoint("GET", "/blogs", api.ShowBlogs, "Список блогов")
		g.Endpoint("GET", "/blogs/{id}", api.ShowBlog, "Список статей в блоге")
		g.Endpoint("GET", "/blog_articles/{id}", api.ShowArticle, "Статья в блоге")
		g.Endpoint("GET", "/allgenres", api.ShowGenres, "Список жанров")
		g.Endpoint("GET", "/work/{id}/classification", api.GetWorkClassification, "Классификация произведения")
	})

	g.Subgroup("Для анонимов", func(g *routing.Group) {
		g.Middleware(middlewares.RequireAnon)

		g.Endpoint("POST", "/login", api.Login, "Логин")
	})

	g.Subgroup("Для авторизованных пользователей", func(g *routing.Group) {
		g.Middleware(middlewares.RequireAuth)

		g.Endpoint("DELETE", "/logout", api.Logout, "Разлогин")
		g.Endpoint("GET", "/work/{id}/userclassification", api.GetUserWorkGenres, "Классификация произведения пользователем")

		g.Subgroup("Для авторизованных незабаненных пользователей", func(g *routing.Group) {
			g.Middleware(middlewares.CheckBan(services))

			g.Endpoint("POST", "/topics/{id}/subscription", api.SubscribeForumTopic, "Подписка на тему форума")
			g.Endpoint("DELETE", "/topics/{id}/subscription", api.UnsubscribeForumTopic, "Отписка от темы форума")
			g.Endpoint("POST", "/communities/{id}/subscription", api.SubscribeCommunity, "Вступление в сообщество")
			g.Endpoint("DELETE", "/communities/{id}/subscription", api.UnsubscribeCommunity, "Выход из сообщества")
			g.Endpoint("POST", "/blogs/{id}/subscription", api.SubscribeBlog, "Подписка на блог")
			g.Endpoint("DELETE", "/blogs/{id}/subscription", api.UnsubscribeBlog, "Отписка от блога")
			g.Endpoint("POST", "/blog_articles/{id}/subscription", api.SubscribeArticle, "Подписка на статью в блоге")
			g.Endpoint("DELETE", "/blog_articles/{id}/subscription", api.UnsubscribeArticle, "Отписка от статьи в блоге")
			g.Endpoint("POST", "/blog_articles/{id}/like", api.LikeArticle, "Лайк статьи в блоге")
			g.Endpoint("DELETE", "/blog_articles/{id}/like", api.DislikeArticle, "Дизлайк статьи в блоге")

			g.Subgroup("Для незабаненных философов", func(g *routing.Group) {
				g.Middleware(middlewares.CheckMinLevel(services, middlewares.UserClass_Philosopher))

				g.Endpoint("PUT", "/work/{id}/userclassification", api.SetWorkGenres, "Классификация произведения пользователем")
			})
		})
	})

	return g
}
