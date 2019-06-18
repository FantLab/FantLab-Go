package authapi

import (
	"fantlab/shared"
	"fantlab/utils"
	"net/http"

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
		ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, utils.ErrorJSON("log out first"))
		return
	}

	userName := ctx.PostForm("login")
	password := ctx.PostForm("password")

	userData := fetchUserPasswordHash(c.services.DB, userName)

	if userData == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, utils.ErrorJSON("user not found"))
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userData.OldHash), []byte(password))

	if err != nil {
		err = bcrypt.CompareHashAndPassword([]byte(userData.NewHash), []byte(password))
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorJSON("incorrect password"))
		return
	}

	sid := ksuid.New().String()

	if !insertNewSession(c.services.DB, sid, userData.UserID, ctx.ClientIP(), ctx.Request.UserAgent()) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorJSON("failed to create session"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": userData.UserID,
		"session": sid,
	})
}
