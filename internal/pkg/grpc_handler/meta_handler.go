package grpc_handler

import (
	"context"
	"net/http"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-08

func RequestMetaHandler(ctx context.Context, request *http.Request) metadata.MD {

	zap.L().Info("[GRPC-RequestMetaHandler]",
		zap.String("http-method", request.Method),
		zap.String("http-url", request.URL.String()),
		zap.String("http-param", request.URL.Query().Encode()),
	)
	return metadata.New(map[string]string{
		"http-method": request.Method,
		"http-url":    request.URL.String(),
		"http-param":  request.URL.Query().Encode(),
	})
}
