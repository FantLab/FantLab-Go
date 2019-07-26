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
				dbSession, err := services.DB.FetchUserSessionInfo(sid)

				if err == nil {
					err = utils.PutSessionInCache(services.Cache, sid, dbSession.UserID, dbSession.DateOfCreate)

					if err == nil {
						uid = dbSession.UserID
					}
				}
			}

			ctx.Set(gin.AuthUserKey, int64(uid))
		} else {
			ctx.Set(gin.AuthUserKey, int64(0))
		}

		ctx.Next()
	}
}
