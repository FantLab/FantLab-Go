package auth

import (
	"fantlab/shared"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	services *shared.Services
}

func NewController(services *shared.Services) *Controller {
	return &Controller{
		services: services,
	}
}

func (c *Controller) Login(ctx *gin.Context) {
	uid := ctx.Keys[gin.AuthUserKey].(int)

	if uid > 0 {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "log out first",
		})

		return
	}

	userName := ctx.DefaultPostForm("login", "")
	password := ctx.DefaultPostForm("pass", "")

	userData := fetchUserPasswordHash(c.services.DB, userName)

	if userData == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userData.OldHash), []byte(password))

	if err != nil {
		err = bcrypt.CompareHashAndPassword([]byte(userData.NewHash), []byte(password))
	}

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "incorrect password",
		})

		return
	}

	sid := ksuid.New().String()

	if !insertNewSession(c.services.DB, sid, userData.UserID, ctx.ClientIP(), ctx.Request.UserAgent()) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create session",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": userData.UserID,
		"session": sid,
	})
}
