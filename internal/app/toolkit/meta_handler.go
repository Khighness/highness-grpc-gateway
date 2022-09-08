package toolkit

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	"highness-grpc-gateway/internal/pkg/kctx"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-08

func RequestMetaHandler(ctx context.Context, request *http.Request) metadata.MD {
	return metadata.New(map[string]string{
		"http-method": request.Method,
		"http-url":    request.URL.String(),
		"http-param":  request.URL.Query().Encode(),
		"trace-id":    kctx.GetTraceID(request.Context()),
	})
}
