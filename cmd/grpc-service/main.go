package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"highness-grpc-gateway/internal/pkg/logging"

	"highness-grpc-gateway/config"
	"highness-grpc-gateway/internal/app/hello"
	"highness-grpc-gateway/proto/api"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-07

var logger = log.New(os.Stdout, "[SERVICE] ", log.Flags()|log.Lmicroseconds|log.Lshortfile)

func main() {
	// init zap logger
	logging.InitLogger(zapcore.DebugLevel)

	// create tcp listener
	addr := fmt.Sprintf("0.0.0.0:%d", *config.GRPC_PORT)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		zap.L().Fatal("Failed to listen:", zap.Error(err))
	}

	// create grpc server
	server := grpc.NewServer()
	helloServer := hello.NewHelloServer()
	api.RegisterHelloServiceServer(server, helloServer)

	// start grpc server
	zap.L().Info("GRPC service is serving at " + addr)
	if err = server.Serve(listener); err != nil {
		zap.L().Fatal("Failed to start grpc service", zap.Error(err))
	}
}
