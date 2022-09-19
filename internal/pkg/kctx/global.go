package kctx

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// @Author Chen Zikang
// @Email  parakovo@gmail.com
// @Since  2022-09-08

func GetLogger(ctx context.Context) *zap.Logger {
	return zap.L().With(zap.Field{
		Key:    MetaTraceID,
		Type:   zapcore.StringType,
		String: GetTraceID(ctx),
	})
}
