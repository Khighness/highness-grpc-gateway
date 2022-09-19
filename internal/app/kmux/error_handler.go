package kmux

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"highness-grpc-gateway/internal/pkg/gerror"
	"highness-grpc-gateway/internal/pkg/kctx"
)

// @Author Chen Zikang
// @Email  parakovo@gmail.com
// @Since  2022-09-08

// ErrorRoutingHandler will response status.Error directly by runtime.DefaultRoutingErrorHandler.
func ErrorRoutingHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
	logger, _ := getLogger(ctx)
	logger.Warn("[GRPC-ErrorRoutingHandler]",
		zap.String(kctx.MetaHttpMethod, r.Method),
		zap.String(kctx.MetaHttpUrl, r.URL.String()))
	runtime.DefaultRoutingErrorHandler(ctx, mux, marshaler, w, r, httpStatus)
}

// CustomErrorHandler will response status.Error directly and gateway filters will be ignored.
func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter,
	r *http.Request, err error) {
	logger, metaData := getLogger(ctx)
	logger.Warn("[GRPC-CustomErrorHandler]", zap.Error(err))

	w.Header().Set(kctx.MetaTraceID, GetTraceID(metaData))
	w.Header().Set("Content-Type", "application/json")
	statusErr := status.Convert(err)
	if statusErr.Code() == codes.InvalidArgument || statusErr.Code() == codes.Unknown {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(gerror.From(statusErr))
}
