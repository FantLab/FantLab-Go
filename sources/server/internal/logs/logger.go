package logs

import (
	"context"
	"fantlab/base/sharedconfig"

	"go.uber.org/zap"

	"go.elastic.co/apm/module/apmzap"
)

var logger *zap.Logger

func Logger() *zap.Logger {
	return logger
}

func WithAPM(ctx context.Context) *zap.Logger {
	return logger.With(apmzap.TraceContext(ctx)...)
}

func init() {
	if sharedconfig.IsDebug() {
		l, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		logger = l
	} else {
		logger = zap.NewExample(zap.WrapCore((&apmzap.Core{}).WrapCore))
	}
}
