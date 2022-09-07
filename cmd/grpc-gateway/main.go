package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"highness-grpc-gateway/config"
	"highness-grpc-gateway/proto/api"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-07

var logger = log.New(os.Stdout, "[GATEWAY] ", log.Flags()|log.Lmicroseconds|log.Lshortfile)

func main() {
	// dial to grpc service
	grpcAddr := fmt.Sprintf("0.0.0.0:%d", *config.GRPC_PORT)
	conn, err := grpc.DialContext(context.Background(), grpcAddr,
		grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("Failed to dial grpc service:", err)
	}

	// register service into gateway mux
	gatewayMux := runtime.NewServeMux()
	err = api.RegisterHelloServiceHandler(context.Background(), gatewayMux, conn)
	if err != nil {
		logger.Fatalln("Failed to register service into gateway:", err)
	}

	// start grpc gateway by http service
	gatewayAddr := fmt.Sprintf("0.0.0.0:%d", *config.GATEWAY_PORT)
	httpServer := &http.Server{
		Addr:    gatewayAddr,
		Handler: gatewayMux,
	}
	logger.Println("GRPC gateway is serving at", gatewayAddr)
	if err = httpServer.ListenAndServe(); err != nil {
		logger.Fatalln("Failed to start grpc gateway:", err)
	}
}
