package grpc_handler

import (
	"context"
	"go.uber.org/zap"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-08

func ErrorRoutingHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
	zap.L().Warn("[GRPC-ErrorRoutingHandler]",
		zap.String("http_method", r.Method),
		zap.String("http_path", r.URL.String()))
	if httpStatus >= 400 {
		zap.L().Warn("[ErrorHandler]", zap.Int("http_status", httpStatus))
	} else {
		zap.L().Info("[ErrorHandler]", zap.Int("http_status", httpStatus))
	}
	runtime.DefaultRoutingErrorHandler(ctx, mux, marshaler, w, r, httpStatus)
}
