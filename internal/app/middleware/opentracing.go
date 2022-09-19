package middleware

import (
	"net/http"

	"go.uber.org/zap"

	"highness-grpc-gateway/internal/pkg/kctx"
)

// @Author Chen Zikang
// @Email  parakovo@gmail.com
// @Since  2022-09-08

func Opentracing(writer http.ResponseWriter, request *http.Request, next func(http.ResponseWriter, *http.Request)) {
	traceID := kctx.GenerateTraceId()
	request = request.WithContext(kctx.SetTraceID(request.Context(), traceID))
	zap.L().Info("[MIDDLEWARE-Tracing]",
		zap.String(kctx.MetaTraceID, kctx.GetTraceID(request.Context())))
	next(writer, request)
}
