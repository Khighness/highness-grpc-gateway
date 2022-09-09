package kgrpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-09

func MetaServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	zap.L().Sugar().Infof("[MetaServerInterceptor] %v", md)

	return handler(ctx, req)
}
