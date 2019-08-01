package middlewares

import (
	"fantlab/pb"
	"fantlab/shared"
	"fantlab/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DetectUser(services *shared.Services) gin.HandlerFunc {
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

func AuthorizedUserIsRequired(ctx *gin.Context) {
	userId := ctx.GetInt64(gin.AuthUserKey)

	if userId > 0 {
		ctx.Next()
	} else {
		utils.ShowProto(ctx, http.StatusUnauthorized, &pb.Error_Response{
			Status: pb.Error_INVALID_SESSION,
		})

		ctx.Abort()
	}
}

func AnonymousIsRequired(ctx *gin.Context) {
	userId := ctx.GetInt64(gin.AuthUserKey)

	if userId > 0 {
		utils.ShowProto(ctx, http.StatusMethodNotAllowed, &pb.Error_Response{
			Status: pb.Error_LOG_OUT_FIRST,
		})

		ctx.Abort()
	} else {
		ctx.Next()
	}
}
