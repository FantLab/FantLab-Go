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

type controllers struct {
	auth   *authapi.Controller
	forum  *forumapi.Controller
	blogs  *blogsapi.Controller
	genres *genresapi.Controller
}

const (
	baseGroupName = "v1"
)

func SetupWith(services *shared.Services) *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger, gin.Recovery(), middlewares.DetectUser(services))

	c := controllers{
		auth:   authapi.NewController(services),
		forum:  forumapi.NewController(services),
		blogs:  blogsapi.NewController(services),
		genres: genresapi.NewController(services),
	}

	// Доступно всем
	{
		g := r.Group(baseGroupName)
		g.GET("/forums", c.forum.ShowForums)
		g.GET("/forums/:id", c.forum.ShowForumTopics)
		g.GET("/topics/:id", c.forum.ShowTopicMessages)
		g.GET("/communities", c.blogs.ShowCommunities)
		g.GET("/communities/:id", c.blogs.ShowCommunity)
		g.GET("/blogs", c.blogs.ShowBlogs)
		g.GET("/blogs/:id", c.blogs.ShowBlog)
		g.GET("/blog_articles/:id", c.blogs.ShowArticle)
		g.GET("/allgenres", c.genres.ShowGenres)
	}

	// Доступно только анонимам
	{
		g := r.Group(baseGroupName).Use(middlewares.AnonymousIsRequired)
		g.POST("/login", c.auth.Login)
	}

	// Требуется авторизация
	{
		g := r.Group(baseGroupName).Use(middlewares.AuthorizedUserIsRequired)
		g.DELETE("/logout", c.auth.Logout)
		g.POST("/blog_articles/:id/like", c.blogs.LikeArticle)
		g.DELETE("/blog_articles/:id/like", c.blogs.DislikeArticle)
		g.PUT("/work/:id/genres", c.genres.SetWorkGenres)
	}

	return r
}
