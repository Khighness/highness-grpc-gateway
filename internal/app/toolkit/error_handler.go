package toolkit

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"

	"highness-grpc-gateway/internal/pkg/kctx"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-08

func ErrorRoutingHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
	kctx.GetLogger(r.Context()).Warn("[GRPC-ErrorRoutingHandler]",
		zap.String(kctx.HttpMethod, r.Method),
		zap.String(kctx.HttpUrl, r.URL.String()))
	runtime.DefaultRoutingErrorHandler(ctx, mux, marshaler, w, r, httpStatus)
}
