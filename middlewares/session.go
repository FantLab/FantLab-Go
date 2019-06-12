package middlewares

import (
	"fantlab/shared"

	"github.com/gin-gonic/gin"
)

func Session(services *shared.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sid := ctx.GetHeader("X-Session")

		if len(sid) > 0 {
			type userID struct {
				Value int `gorm:"Column:user_id"`
			}

			var uid userID

			services.DB.Table("sessions2").Where("code = ?", sid).First(&uid)

			ctx.Set(gin.AuthUserKey, uid.Value)
		} else {
			ctx.Set(gin.AuthUserKey, 0)
		}

		ctx.Next()
	}
}
