package main

import (
	"context"
	"fmt"
	"go.uber.org/zap/zapcore"
	"highness-grpc-gateway/internal/app/kgrpc"
	"highness-grpc-gateway/internal/app/kmux"
	"highness-grpc-gateway/internal/app/middleware"
	"highness-grpc-gateway/internal/pkg/logging"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"highness-grpc-gateway/config"
	"highness-grpc-gateway/proto/api"
)

// @Author Chen Zikang
// @Email  parakovo@gmail.com
// @Since  2022-09-07

var logger = log.New(os.Stdout, "[GATEWAY] ", log.Flags()|log.Lmicroseconds|log.Lshortfile)

func main() {
	// init zap logger
	logging.InitLogger(zapcore.DebugLevel)

	// dial to grpc service
	grpcAddr := fmt.Sprintf("0.0.0.0:%d", *config.GRPC_PORT)
	conn, err := grpc.DialContext(
		context.Background(),
		grpcAddr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(kgrpc.MetaClientInterceptor),
		grpc.WithStreamInterceptor(kgrpc.MetaClientStreamInterceptor),
	)
	if err != nil {
		zap.L().Fatal("Failed to dial grpc service ", zap.Error(err))
	}

	// create http mux
	gatewayMux := runtime.NewServeMux(kmux.DefaultServeMuxOptions()...)
	handler := middleware.WithMiddleWares(gatewayMux, middleware.DefaultMiddlewares()...)

	// register service into gateway mux
	err = api.RegisterHelloServiceHandler(context.Background(), gatewayMux, conn)
	if err != nil {
		zap.L().Fatal("Failed to register service into gateway", zap.Error(err))
	}

	// start grpc gateway by http service
	gatewayAddr := fmt.Sprintf("0.0.0.0:%d", *config.GATEWAY_PORT)
	httpServer := &http.Server{
		Addr:    gatewayAddr,
		Handler: handler,
	}
	zap.L().Info("GRPC gateway is serving at " + gatewayAddr)
	if err = httpServer.ListenAndServe(); err != nil {
		zap.L().Fatal("Failed to start grpc gateway ", zap.Error(err))
	}
}
