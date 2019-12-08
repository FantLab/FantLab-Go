package routing

import (
	"fantlab/server/internal/endpoints"
	"fantlab/server/internal/middlewares"
	"fantlab/server/internal/shared"
)

const BasePath = "/v1"

func Routes(config *shared.AppConfig, services *shared.Services, pathParamGetter endpoints.PathParamGetter) *Group {
	api := endpoints.MakeAPI(config, services, pathParamGetter)

	g := new(Group)

	g.middleware(middlewares.DetectUser(services))

	g.subgroup("Общедоступные", func(g *Group) {
		g.endpoint("GET", "/forums", api.ShowForums, "Список форумов")
		g.endpoint("GET", "/forums/{id}", api.ShowForumTopics, "Список тем форума")
		g.endpoint("GET", "/topics/{id}", api.ShowTopicMessages, "Сообщения в теме форума")
		g.endpoint("GET", "/communities", api.ShowCommunities, "Список сообществ")
		g.endpoint("GET", "/communities/{id}", api.ShowCommunity, "Информация о сообществе")
		g.endpoint("GET", "/blogs", api.ShowBlogs, "Список блогов")
		g.endpoint("GET", "/blogs/{id}", api.ShowBlog, "Список статей в блоге")
		g.endpoint("GET", "/blog_articles/{id}", api.ShowArticle, "Статья в блоге")
		g.endpoint("GET", "/allgenres", api.ShowGenres, "Список жанров")
	})

	g.subgroup("Для анонимов", func(g *Group) {
		g.middleware(middlewares.RequireAnon)

		g.endpoint("POST", "/login", api.Login, "Логин")
	})

	g.subgroup("Для авторизованных пользователей", func(g *Group) {
		g.middleware(middlewares.RequireAuth)

		g.endpoint("DELETE", "/logout", api.Logout, "Разлогин")
	})

	g.subgroup("Для авторизованных незабаненных пользователей", func(g *Group) {
		g.middleware(middlewares.RequireAuth)
		g.middleware(middlewares.CheckBan(services))

		g.endpoint("POST", "/topics/{id}/subscription", api.SubscribeForumTopic, "Подписка на тему форума")
		g.endpoint("DELETE", "/topics/{id}/subscription", api.UnsubscribeForumTopic, "Отписка от темы форума")
		g.endpoint("POST", "/communities/{id}/subscription", api.SubscribeCommunity, "Вступление в сообщество")
		g.endpoint("DELETE", "/communities/{id}/subscription", api.UnsubscribeCommunity, "Выход из сообщества")
		g.endpoint("POST", "/blogs/{id}/subscription", api.SubscribeBlog, "Подписка на блог")
		g.endpoint("DELETE", "/blogs/{id}/subscription", api.UnsubscribeBlog, "Отписка от блога")
		g.endpoint("POST", "/blog_articles/{id}/subscription", api.SubscribeArticle, "Подписка на статью в блоге")
		g.endpoint("DELETE", "/blog_articles/{id}/subscription", api.UnsubscribeArticle, "Отписка от статьи в блоге")
		g.endpoint("POST", "/blog_articles/{id}/like", api.LikeArticle, "Лайк статьи в блоге")
		g.endpoint("DELETE", "/blog_articles/{id}/like", api.DislikeArticle, "Дизлайк статьи в блоге")
		g.endpoint("PUT", "/work/{id}/genres", api.SetWorkGenres, "Классификация произведения")
	})

	return g
}
