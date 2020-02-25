package router

import (
	"fantlab/base/routing"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"fantlab/server/internal/endpoints"
	"fantlab/server/internal/middlewares"
)

const BasePath = "/v1"

func Routes(config *config.AppConfig, services *app.Services, pathParamGetter endpoints.PathParamGetter) *routing.Group {
	api := endpoints.MakeAPI(config, services, pathParamGetter)

	g := new(routing.Group)

	g.Middleware(middlewares.CheckSession(services))

	g.Subgroup("Общедоступные", func(g *routing.Group) {
		g.Endpoint("POST", "/auth/login", api.Login, "Логин")
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
		g.Endpoint("GET", "/work/{id}/subworks", api.GetWorkSubWorks, "Иерархия произведений, входящих в запрашиваемое")
		g.Endpoint("GET", "/blog_articles/{id}/comments", api.BlogArticleComments, "Комментарии к статье в блоге")
	})

	g.Subgroup("Для зарегистрированных пользователей", func(g *routing.Group) {
		g.Middleware(middlewares.CheckAuth)

		g.Endpoint("POST", "/auth/refresh", api.RefreshAuth, "Продление сессии")

		g.Subgroup("Для пользователей с валидной сессией", func(g *routing.Group) {
			g.Middleware(middlewares.CheckAuthExpiration)

			g.Endpoint("GET", "/work/{id}/userclassification", api.GetUserWorkGenres, "Классификация произведения пользователем")

			g.Subgroup("С проверкой на бан", func(g *routing.Group) {
				g.Middleware(middlewares.CheckBan(services))

				g.Endpoint("POST", "/topics/{id}/message", api.AddForumMessage, "Создание нового сообщения в форуме")
				g.Endpoint("PUT", "/topics/{id}/subscription", api.ToggleForumTopicSubscription, "Подписка/отписка от темы форума")
				g.Endpoint("PUT", "/communities/{id}/subscription", api.ToggleCommunitySubscription, "Вступление/выход из сообщества")
				g.Endpoint("PUT", "/blogs/{id}/subscription", api.ToggleBlogSubscription, "Подписка/отписка от блога")
				g.Endpoint("PUT", "/blog_articles/{id}/subscription", api.ToggleArticleSubscription, "Подписка/отписка от статьи в блоге")
				g.Endpoint("PUT", "/blog_articles/{id}/like", api.ToggleArticleLike, "Лайк/дизлайк статьи в блоге")

				g.Subgroup("Для философов", func(g *routing.Group) {
					g.Middleware(middlewares.CheckMinLevel(pb.Common_USERCLASS_PHILOSOPHER))

					g.Endpoint("PUT", "/work/{id}/userclassification", api.SetWorkGenres, "Классификация произведения пользователем")
					g.Endpoint("PUT", "/forum_messages/{id}/voting", api.SetForumMessageVoting, "Плюс/минус посту в форуме")
				})
			})
		})
	})

	return g
}
