package authapi

import (
	"fantlab/pb"
	"net/http"
	"time"

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

	dateOfCreate := time.Now()

	ok, err := c.services.DB.InsertNewSession(dateOfCreate, sid, userData.UserID, ctx.ClientIP(), ctx.Request.UserAgent())

	if !ok || err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	_ = c.services.Cache.PutSession(sid, uint64(userData.UserID), dateOfCreate)

	utils.ShowProto(ctx, http.StatusOK, &pb.Auth_LoginResponse{
		UserId:       userData.UserID,
		SessionToken: sid,
	})
}

func (c *Controller) Logout(ctx *gin.Context) {
	sid := ctx.GetHeader(utils.SessionHeader)

	_, err := c.services.DB.DeleteSession(sid)

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	_ = c.services.Cache.DeleteSession(sid)

	utils.ShowProto(ctx, http.StatusOK, &pb.Common_SuccessResponse{})
}
