package routing

import (
	"fantlab/logger"
	"fantlab/middlewares"
	"fantlab/modules/authapi"
	"fantlab/modules/blogsapi"
	"fantlab/modules/forumapi"
	"fantlab/modules/genresapi"
	"fantlab/shared"

	"github.com/gin-gonic/gin"
)

func SetupWith(services *shared.Services) *gin.Engine {
	router := gin.New()
	router.Use(logger.GinLogger, gin.Recovery(), middlewares.Session(services))

	{
		v1 := router.Group("v1")

		// Пользователь
		{
			controller := authapi.NewController(services)

			v1.POST("/login", controller.Login)
			v1.DELETE("/logout", controller.Logout)
		}

		// Форум
		{
			controller := forumapi.NewController(services)

			v1.GET("/forums", controller.ShowForums)
			v1.GET("/forums/:id", controller.ShowForumTopics)
			v1.GET("/topics/:id", controller.ShowTopicMessages)
		}

		// Блоги
		{
			controller := blogsapi.NewController(services)

			v1.GET("/communities", controller.ShowCommunities)
			v1.GET("/communities/:id", controller.ShowCommunity)
			v1.GET("/blogs", controller.ShowBlogs)
			v1.GET("/blogs/:id", controller.ShowBlog)
			v1.GET("/blog_articles/:id", controller.ShowArticle)
		}

		// Жанры
		{
			controller := genresapi.NewController(services)

			v1.GET("/genres", controller.ShowGenres)
		}
	}

	return router
}
