package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"highness-grpc-gateway/config"
	"highness-grpc-gateway/internal/app/hello"
	"highness-grpc-gateway/proto/api"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-07

var logger = log.New(os.Stdout, "[SERVICE] ", log.Flags()|log.Lmicroseconds|log.Lshortfile)

func main() {
	// create tcp listener
	addr := fmt.Sprintf("0.0.0.0:%d", *config.GRPC_PORT)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatalln("Failed to listen:", err)
	}

	// create grpc server
	server := grpc.NewServer()
	helloServer := hello.NewHelloServer()
	api.RegisterHelloServiceServer(server, helloServer)

	// start grpc server
	logger.Println("GRPC service is serving at", addr)
	if err = server.Serve(listener); err != nil {
		logger.Fatalln("Failed to start:", err)
	}
}
