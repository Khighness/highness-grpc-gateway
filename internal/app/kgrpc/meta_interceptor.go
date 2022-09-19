package kgrpc

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-09

func MetaServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	logger := zap.L().Sugar()

	logger.Infof("[MetaServerInterceptor] %v", md)

	return handler(ctx, req)
}

func MetaClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	logger := zap.L().Sugar()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		logger.Error("[MetaServerInterceptor] %v", err)
		return err
	}
	//meta, _ := metadata.FromIncomingContext(ctx)
	meta, _ := runtime.ServerMetadataFromContext(ctx)
	logger.Infof("[MetaServerInterceptor] %v", meta)

	return nil
}

func MetaClientStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	zap.L().Sugar().Infof("[MetaClientStreamInterceptor] method: %v", method)
	stream, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	//meta, _ := runtime.ServerMetadataFromContext(ctx)
	zap.L().Sugar().Infof("[MetaClientStreamInterceptor] %v", stream.Trailer())
	return nil, nil
}
