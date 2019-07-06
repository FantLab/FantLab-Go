package authapi

import (
	"net/http"

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
	userId := ctx.Keys[gin.AuthUserKey].(int)

	if userId > 0 {
		utils.ShowError(ctx, http.StatusMethodNotAllowed, "log out first")
		return
	}

	userName := ctx.PostForm("login")
	password := ctx.PostForm("password")

	userData, err := fetchUserPasswordHash(c.services.DB, userName)

	if err != nil {
		if utils.IsRecordNotFoundError(err) {
			utils.ShowError(ctx, http.StatusNotFound, "user not found")
		} else {
			utils.ShowError(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.OldHash), []byte(password))

	if err != nil {
		err = bcrypt.CompareHashAndPassword([]byte(userData.NewHash), []byte(password))
	}

	if err != nil {
		utils.ShowError(ctx, http.StatusUnauthorized, "incorrect password")
		return
	}

	sid := ksuid.New().String()

	err = insertNewSession(c.services.DB, sid, userData.UserID, ctx.ClientIP(), ctx.Request.UserAgent())

	if err != nil {
		utils.ShowError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	session := createSession(userData, sid)

	utils.ShowProto(ctx, http.StatusOK, session)
}
