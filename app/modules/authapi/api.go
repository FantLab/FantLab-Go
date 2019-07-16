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
	userId := ctx.Keys[gin.AuthUserKey].(int)

	if userId > 0 {
		utils.ShowProto(ctx, http.StatusMethodNotAllowed, &pb.Error_Response{
			Status: pb.Error_LOG_OUT_FIRST,
		})
		return
	}

	userName := ctx.PostForm("login")
	password := ctx.PostForm("password")

	userData, err := c.services.DB.FetchUserPasswordHash(userName)

	if err != nil {
		if utils.IsRecordNotFoundError(err) {
			utils.ShowProto(ctx, http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: userName,
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

	err = c.services.DB.InsertNewSession(sid, userData.UserID, ctx.ClientIP(), ctx.Request.UserAgent())

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	utils.ShowProto(ctx, http.StatusOK, &pb.Auth_LoginResponse{
		UserId:       userData.UserID,
		SessionToken: sid,
	})
}
