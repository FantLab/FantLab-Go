package routes

import (
	"fantlab/apiserver/internal/endpoints"
	"fantlab/apiserver/internal/middlewares"
	"fantlab/apiserver/routing"
	"fantlab/core/app"
	"fantlab/pb"
	"net/http"
)

func Tree(services *app.Services, pathParamGetter endpoints.PathParamGetter) *routing.Group {
	api := endpoints.MakeAPI(services, pathParamGetter)

	g := new(routing.Group)

	g.Middleware(middlewares.CheckSession(services))

	g.Subgroup("Общедоступные", func(g *routing.Group) {
		g.Endpoint(http.MethodPost, "/auth/login", api.Login, "Логин")
		g.Endpoint(http.MethodGet, "/work/:id/classification", api.GetWorkClassification, "Классификация произведения")
		g.Endpoint(http.MethodGet, "/work/:id/subworks", api.GetWorkSubWorks, "Иерархия произведений, входящих в запрашиваемое")
		g.Endpoint(http.MethodGet, "/forums", api.ShowForums, "Список форумов")
		g.Endpoint(http.MethodGet, "/forums/:id", api.ShowForumTopics, "Список тем форума")
		g.Endpoint(http.MethodGet, "/topics/:id", api.ShowTopicMessages, "Сообщения в теме форума")
		g.Endpoint(http.MethodGet, "/communities", api.ShowCommunities, "Список сообществ")
		g.Endpoint(http.MethodGet, "/communities/:id", api.ShowCommunity, "Информация о сообществе")
		g.Endpoint(http.MethodGet, "/blogs", api.ShowBlogs, "Список блогов")
		g.Endpoint(http.MethodGet, "/blogs/:id", api.ShowBlog, "Список статей в блоге")
		g.Endpoint(http.MethodGet, "/blog_articles/:id", api.ShowArticle, "Статья в блоге")
		g.Endpoint(http.MethodGet, "/allgenres", api.ShowGenres, "Список жанров")
		g.Endpoint(http.MethodGet, "/blog_articles/:id/comments", api.BlogArticleComments, "Комментарии к статье в блоге")
		g.Endpoint(http.MethodGet, "/users/:id/bookcases", api.ShowBookcases, "Список книжных полок пользователя")
		g.Endpoint(http.MethodGet, "/edition_bookcases/:id", api.ShowEditionBookcase, "Содержимое полки с изданиями")
		g.Endpoint(http.MethodGet, "/work_bookcases/:id", api.ShowWorkBookcase, "Содержимое полки с произведениями")
		g.Endpoint(http.MethodGet, "/film_bookcases/:id", api.ShowFilmBookcase, "Содержимое полки с фильмами")
	})

	g.Subgroup("Для зарегистрированных пользователей", func(g *routing.Group) {
		g.Middleware(middlewares.CheckAuth)

		g.Endpoint(http.MethodPost, "/auth/refresh", api.RefreshAuth, "Продление сессии")

		g.Subgroup("Для пользователей с валидной сессией", func(g *routing.Group) {
			g.Middleware(middlewares.CheckAuthExpiration)

			g.Subgroup("С проверкой на бан", func(g *routing.Group) {
				g.Middleware(middlewares.CheckBan(services))

				g.Endpoint(http.MethodPost, "/users/:id/private_message", api.AddPrivateMessage, "Создание нового сообщения в личной переписке")
				g.Endpoint(http.MethodGet, "/work/:id/userclassification", api.GetUserWorkGenres, "Классификация произведения пользователем")
				g.Endpoint(http.MethodPut, "/response/:id", api.EditResponse, "Редактирование отзыва на произведение")
				g.Endpoint(http.MethodPut, "/response/:id/voting", api.VoteResponse, "Плюс/минус отзыву на произведение")
				g.Endpoint(http.MethodDelete, "/response/:id", api.DeleteResponse, "Удаление отзыва на произведение")
				g.Endpoint(http.MethodPost, "/topics/:id/message", api.AddForumMessage, "Создание нового сообщения в форуме")
				g.Endpoint(http.MethodPut, "/forum_messages/:id", api.EditForumMessage, "Редактирование сообщения в форуме")
				g.Endpoint(http.MethodDelete, "/forum_messages/:id", api.DeleteForumMessage, "Удаление сообщения в форуме")
				g.Endpoint(http.MethodGet, "/forum_messages/:id/file_upload_url", api.GetForumMessageFileUploadUrl, "Получение URL для загрузки аттача к сообщению в форуме")
				g.Endpoint(http.MethodDelete, "/forum_messages/:id/file", api.DeleteForumMessageFile, "Удаление аттача сообщения в форуме")
				g.Endpoint(http.MethodPut, "/topics/:id/message_draft", api.SaveForumMessageDraft, "Сохранение черновика сообщения в форуме")
				g.Endpoint(http.MethodPost, "/topics/:id/message_draft", api.ConfirmForumMessageDraft, "Подтверждение черновика сообщения в форуме")
				g.Endpoint(http.MethodDelete, "/topics/:id/message_draft", api.CancelForumMessageDraft, "Отмена черновика сообщения в форуме")
				g.Endpoint(http.MethodGet, "/topics/:id/message_draft/file_upload_url", api.GetForumMessageDraftFileUploadUrl, "Получение URL для загрузки аттача к черновику сообщения в форуме")
				g.Endpoint(http.MethodDelete, "/topics/:id/message_draft/file", api.DeleteForumMessageDraftFile, "Удаление аттача черновика сообщения в форуме")
				g.Endpoint(http.MethodPut, "/topics/:id/subscription", api.ToggleForumTopicSubscription, "Подписка/отписка от темы форума")
				g.Endpoint(http.MethodPost, "/blog_articles/:id/comment", api.AddBlogArticleComment, "Создание нового комментария к статье в блоге")
				g.Endpoint(http.MethodPut, "/blog_article_comments/:id", api.EditBlogArticleComment, "Редактирование комментария к статье в блоге")
				g.Endpoint(http.MethodDelete, "/blog_article_comments/:id", api.DeleteBlogArticleComment, "Удаление комментария к статье в блоге")
				g.Endpoint(http.MethodPut, "/communities/:id/subscription", api.ToggleCommunitySubscription, "Вступление/выход из сообщества")
				g.Endpoint(http.MethodPut, "/blogs/:id/subscription", api.ToggleBlogSubscription, "Подписка/отписка от блога")
				g.Endpoint(http.MethodPut, "/blog_articles/:id/subscription", api.ToggleArticleSubscription, "Подписка/отписка от статьи в блоге")
				g.Endpoint(http.MethodPut, "/blog_articles/:id/like", api.ToggleArticleLike, "Лайк/дизлайк статьи в блоге")
				g.Endpoint(http.MethodPost, "/bookcases", api.CreateDefaultBookcases, "Создание первичных книжных полок")
				g.Endpoint(http.MethodPut, "/bookcases/order", api.ChangeBookcasesOrder, "Изменение порядка сортировки книжных полок внутри блоков")
				g.Endpoint(http.MethodPost, "/edition_bookcases/:id/items", api.AddEditionBookcaseItem, "Добавление item-а на полку изданий")
				g.Endpoint(http.MethodPost, "/work_bookcases/:id/items", api.AddWorkBookcaseItem, "Добавление item-а на полку произведений")
				g.Endpoint(http.MethodPost, "/film_bookcases/:id/items", api.AddFilmBookcaseItem, "Добавление item-а на полку фильмов")
				g.Endpoint(http.MethodPut, "/bookcase_items/:id/comment", api.EditBookcaseItemComment, "Редактирование комментария к item-у книжной полки")
				g.Endpoint(http.MethodDelete, "/bookcase_items/:id", api.DeleteBookcaseItem, "Удаление item-а с книжной полки")
				g.Endpoint(http.MethodDelete, "/bookcases/:id", api.DeleteBookcase, "Удаление книжной полки")

				g.Subgroup("Для философов", func(g *routing.Group) {
					g.Middleware(middlewares.CheckMinLevel(pb.Common_USERCLASS_PHILOSOPHER))

					g.Endpoint(http.MethodPut, "/work/:id/userclassification", api.SetWorkGenres, "Классификация произведения пользователем")
					g.Endpoint(http.MethodPut, "/forum_messages/:id/voting", api.VoteForumMessage, "Плюс/минус посту в форуме")
				})
			})
		})
	})

	return g
}
