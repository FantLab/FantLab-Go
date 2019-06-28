// +build debug

package utils

import (
	"github.com/gin-gonic/gin"
)

func SetupGinMode() {
	gin.SetMode(gin.DebugMode)
}
