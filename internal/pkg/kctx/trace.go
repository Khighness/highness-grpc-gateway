package kctx

import (
	"context"

	"highness-grpc-gateway/internal/pkg/random"
)

// @Author Chen Zikang
// @Email  parakovo@gmail.com
// @Since  2022-09-08

// SetTraceID sets traceID into ctx.
func SetTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, MetaTraceID, traceID)
}

// GetTraceID gets traceID from ctx.
func GetTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value(MetaTraceID).(string)
	if !ok {
		return DefaultValue
	}
	return traceID
}

// GenerateTraceId generates traceID.
func GenerateTraceId() string {
	return random.RandString(10)
}
