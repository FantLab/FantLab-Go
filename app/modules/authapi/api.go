package authapi

import (
	"fantlab/pb"
	"net/http"

	"fantlab/shared"
	"fantlab/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	services *shared.Services
}

func NewController(services *shared.Services) *Controller {
	return &Controller{services: services}
}

func (c *Controller) Login(ctx *gin.Context) {
	userId := ctx.GetInt64(gin.AuthUserKey)

	if userId > 0 {
		utils.ShowProto(ctx, http.StatusMethodNotAllowed, &pb.Error_Response{
			Status: pb.Error_LOG_OUT_FIRST,
		})
		return
	}

	login := ctx.PostForm("login")
	password := ctx.PostForm("password")

	userData, err := c.services.DB.FetchUserPasswordHash(login)

	if err != nil {
		if utils.IsRecordNotFoundError(err) {
			utils.ShowProto(ctx, http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: login,
			})
		} else {
			utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			})
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.OldHash), []byte(password))

	if err != nil {
		err = bcrypt.CompareHashAndPassword([]byte(userData.NewHash), []byte(password))
	}

	if err != nil {
		utils.ShowProto(ctx, http.StatusUnauthorized, &pb.Error_Response{
			Status: pb.Error_INVALID_PASSWORD,
		})
		return
	}

	sid := utils.GenerateUniqueId()

	dateOfCreate, err := c.services.DB.InsertNewSession(sid, userData.UserID, ctx.ClientIP(), ctx.Request.UserAgent())

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	_ = utils.PutSessionInCache(c.services.Cache, sid, uint64(userData.UserID), dateOfCreate)

	utils.ShowProto(ctx, http.StatusOK, &pb.Auth_LoginResponse{
		UserId:       userData.UserID,
		SessionToken: sid,
	})
}

func (c *Controller) Logout(ctx *gin.Context) {
	userId := ctx.GetInt64(gin.AuthUserKey)

	if userId == 0 {
		utils.ShowProto(ctx, http.StatusUnauthorized, &pb.Error_Response{
			Status: pb.Error_INVALID_SESSION,
		})
		return
	}

	sid := ctx.GetHeader(utils.SessionHeader)

	err := c.services.DB.DeleteSession(sid)

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	utils.DeleteSessionFromCache(c.services.Cache, sid)

	utils.ShowProto(ctx, http.StatusOK, &pb.Common_SuccessResponse{})
}
