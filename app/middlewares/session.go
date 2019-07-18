package middlewares

import (
	"fantlab/shared"
	"fantlab/utils"

	"github.com/gin-gonic/gin"
)

func Session(services *shared.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sid := ctx.GetHeader(utils.SessionHeader)

		if len(sid) > 0 {
			uid, err := utils.GetUserIdBySessionFromCache(services.Cache, sid)

			if err != nil || uid == 0 {
				dbSession := services.DB.FetchUserSessionInfo(sid)

				if utils.PutSessionInCache(services.Cache, sid, dbSession.UserID, dbSession.DateOfCreate) {
					uid = dbSession.UserID
				}
			}

			ctx.Set(gin.AuthUserKey, uid)
		} else {
			ctx.Set(gin.AuthUserKey, uint64(0))
		}

		ctx.Next()
	}
}
