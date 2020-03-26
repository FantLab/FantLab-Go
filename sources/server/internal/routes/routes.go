package routes

import (
	"fantlab/base/routing"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"fantlab/server/internal/endpoints"
	"fantlab/server/internal/middlewares"
	"net/http"
)

func Tree(config *config.AppConfig, services *app.Services, pathParamGetter endpoints.PathParamGetter) *routing.Group {
	api := endpoints.MakeAPI(config, services, pathParamGetter)

	g := new(routing.Group)

	g.Middleware(middlewares.CheckSession(services))

	g.Subgroup("Общедоступные", func(g *routing.Group) {
		g.Endpoint(http.MethodPost, "/auth/login", api.Login, "Логин")
		g.Endpoint(http.MethodGet, "/forums", api.ShowForums, "Список форумов")
		g.Endpoint(http.MethodGet, "/forums/:id", api.ShowForumTopics, "Список тем форума")
		g.Endpoint(http.MethodGet, "/topics/:id", api.ShowTopicMessages, "Сообщения в теме форума")
		g.Endpoint(http.MethodGet, "/communities", api.ShowCommunities, "Список сообществ")
		g.Endpoint(http.MethodGet, "/communities/:id", api.ShowCommunity, "Информация о сообществе")
		g.Endpoint(http.MethodGet, "/blogs", api.ShowBlogs, "Список блогов")
		g.Endpoint(http.MethodGet, "/blogs/:id", api.ShowBlog, "Список статей в блоге")
		g.Endpoint(http.MethodGet, "/blog_articles/:id", api.ShowArticle, "Статья в блоге")
		g.Endpoint(http.MethodGet, "/allgenres", api.ShowGenres, "Список жанров")
		g.Endpoint(http.MethodGet, "/work/:id/classification", api.GetWorkClassification, "Классификация произведения")
		g.Endpoint(http.MethodGet, "/work/:id/subworks", api.GetWorkSubWorks, "Иерархия произведений, входящих в запрашиваемое")
		g.Endpoint(http.MethodGet, "/blog_articles/:id/comments", api.BlogArticleComments, "Комментарии к статье в блоге")
	})

	g.Subgroup("Для зарегистрированных пользователей", func(g *routing.Group) {
		g.Middleware(middlewares.CheckAuth)

		g.Endpoint(http.MethodPost, "/auth/refresh", api.RefreshAuth, "Продление сессии")

		g.Subgroup("Для пользователей с валидной сессией", func(g *routing.Group) {
			g.Middleware(middlewares.CheckAuthExpiration)

			g.Endpoint(http.MethodGet, "/work/:id/userclassification", api.GetUserWorkGenres, "Классификация произведения пользователем")

			g.Subgroup("С проверкой на бан", func(g *routing.Group) {
				g.Middleware(middlewares.CheckBan(services))

				g.Endpoint(http.MethodPost, "/topics/:id/message", api.AddForumMessage, "Создание нового сообщения в форуме")
				g.Endpoint(http.MethodPut, "/topics/:id/message_draft", api.SaveForumMessageDraft, "Сохранение черновика сообщения в форуме")
				g.Endpoint(http.MethodPost, "/topics/:id/message_draft", api.ConfirmForumMessageDraft, "Подтверждение черновика сообщения в форуме")
				g.Endpoint(http.MethodDelete, "/topics/:id/message_draft", api.CancelForumMessageDraft, "Отмена черновика сообщения в форуме")
				g.Endpoint(http.MethodPut, "/topics/:id/subscription", api.ToggleForumTopicSubscription, "Подписка/отписка от темы форума")
				g.Endpoint(http.MethodPut, "/forum_messages/:id", api.EditForumMessage, "Редактирование сообщения в форуме")
				g.Endpoint(http.MethodDelete, "/forum_messages/:id", api.DeleteForumMessage, "Удаление сообщения в форуме")
				g.Endpoint(http.MethodPut, "/communities/:id/subscription", api.ToggleCommunitySubscription, "Вступление/выход из сообщества")
				g.Endpoint(http.MethodPut, "/blogs/:id/subscription", api.ToggleBlogSubscription, "Подписка/отписка от блога")
				g.Endpoint(http.MethodPost, "/blog_articles/:id/comment", api.AddBlogArticleComment, "Создание нового комментария к статье в блоге")
				g.Endpoint(http.MethodPut, "/blog_articles/:id/subscription", api.ToggleArticleSubscription, "Подписка/отписка от статьи в блоге")
				g.Endpoint(http.MethodPut, "/blog_articles/:id/like", api.ToggleArticleLike, "Лайк/дизлайк статьи в блоге")

				g.Subgroup("Для философов", func(g *routing.Group) {
					g.Middleware(middlewares.CheckMinLevel(pb.Common_USERCLASS_PHILOSOPHER))

					g.Endpoint(http.MethodPut, "/work/:id/userclassification", api.SetWorkGenres, "Классификация произведения пользователем")
					g.Endpoint(http.MethodPut, "/forum_messages/:id/voting", api.SetForumMessageVoting, "Плюс/минус посту в форуме")
				})
			})
		})
	})

	return g
}
