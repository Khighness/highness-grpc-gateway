package hello

import (
	"context"
	"log"
	"os"

	"highness-grpc-gateway/proto/api"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-07

type helloServer struct {
	logger *log.Logger
	api.HelloServiceServer
}

func NewHelloServer() api.HelloServiceServer {
	return &helloServer{
		logger: log.New(os.Stdout, "[GRPC] ", log.Flags()|log.Lmicroseconds|log.Lshortfile),
	}
}

func (s *helloServer) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	s.logger.Println("Greet to", in.Name)
	return &api.HelloResponse{Reply: "Hi " + in.Name}, nil
}
