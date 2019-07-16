package middlewares

import (
	"fantlab/shared"

	"github.com/gin-gonic/gin"
)

func Session(services *shared.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sid := ctx.GetHeader("X-Session")

		if len(sid) > 0 {
			userId := services.DB.FetchUserIdBySession(sid)

			ctx.Set(gin.AuthUserKey, userId)
		} else {
			ctx.Set(gin.AuthUserKey, 0)
		}

		ctx.Next()
	}
}
