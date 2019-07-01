package authapi

import (
	"net/http"

	"fantlab/protobuf/generated/fantlab/pb"
	"fantlab/shared"
	"fantlab/utils"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	services *shared.Services
}

func NewController(services *shared.Services) *Controller {
	return &Controller{services: services}
}

func (c *Controller) Login(ctx *gin.Context) {
	uid := ctx.Keys[gin.AuthUserKey].(int)

	if uid > 0 {
		utils.ShowError(ctx, http.StatusMethodNotAllowed, "log out first")
		return
	}

	userName := ctx.PostForm("login")
	password := ctx.PostForm("password")

	userData := fetchUserPasswordHash(c.services.DB, userName)

	if userData == nil {
		utils.ShowError(ctx, http.StatusNotFound, "user not found")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userData.OldHash), []byte(password))

	if err != nil {
		err = bcrypt.CompareHashAndPassword([]byte(userData.NewHash), []byte(password))
	}

	if err != nil {
		utils.ShowError(ctx, http.StatusUnauthorized, "incorrect password")
		return
	}

	sid := ksuid.New().String()

	if !insertNewSession(c.services.DB, sid, userData.UserID, ctx.ClientIP(), ctx.Request.UserAgent()) {
		utils.ShowError(ctx, http.StatusInternalServerError, "failed to create session")
		return
	}

	session := &pb.UserSessionResponse{
		UserId:       userData.UserID,
		SessionToken: sid,
	}

	utils.ShowProto(ctx, http.StatusOK, session)
}
