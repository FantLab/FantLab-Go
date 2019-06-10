package routing

import (
	"fantlab/config"
	forumapi "fantlab/modules/forum"

	"github.com/gin-gonic/gin"
)

func SetupWith(db *config.FLDB) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	{
		v1 := router.Group("v1")

		// форум
		{
			controller := forumapi.NewController(db)

			v1.GET("/forums", controller.ShowForums)
			v1.GET("/forums/:id", controller.ShowForumTopics)
			v1.GET("/topics/:id", controller.ShowTopicMessages)
		}
	}

	return router
}
