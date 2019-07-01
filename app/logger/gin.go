package logger

import (
	"fmt"

	"fantlab/utils"

	"github.com/gin-gonic/gin"
)

// Адаптировано из дефолтного Gin'овского Logger-а.
// Формат лога:
// [$CurrentTime]  [$Duration]  [$ClientIP]  $StatusCode $Method $Path

var GinLogger = gin.LoggerWithFormatter(ginFormatter)

var ginFormatter = func(params gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if params.IsOutputColor() {
		statusColor = params.StatusCodeColor()
		methodColor = params.MethodColor()
		resetColor = params.ResetColor()
	}

	return fmt.Sprintf("%v  %v  [%s]  %s %3d %s %s %s %s %s\n%s",
		utils.FormatLogTime(params.TimeStamp),
		utils.FormatLogDuration(params.Latency),
		params.ClientIP,
		statusColor, params.StatusCode, resetColor,
		methodColor, params.Method, resetColor,
		params.Path,
		params.ErrorMessage,
	)
}
