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
		kctx.HttpMethod: request.Method,
		kctx.HttpUrl:    request.URL.String(),
		kctx.HttpParam:  request.URL.Query().Encode(),
		kctx.TraceID:    kctx.GetTraceID(request.Context()),
	})
}
