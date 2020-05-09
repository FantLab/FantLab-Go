package logs

import (
	"context"

	"github.com/FantLab/go-kit/env"

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
	if env.IsDebug() {
		l, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		logger = l
	} else {
		logger = zap.NewExample(zap.WrapCore((&apmzap.Core{}).WrapCore))
	}
}
