package hello

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
	fullName := fmt.Sprintf("%s·%s", in.GetFirstName(), in.GetLastName())
	s.logger.Println("Greet to", fullName)
	return &api.HelloResponse{ReplyMessage: fmt.Sprintf("%s好, %s", s.GetPeriod(in.GetTimestamp()), fullName)}, nil
}

func (s *helloServer) SayHelloV2(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	fullName := fmt.Sprintf("%s %s", in.GetFirstName(), in.GetLastName())
	s.logger.Println("Greet to", fullName)
	return &api.HelloResponse{ReplyMessage: fmt.Sprintf("%s好, %s", s.GetPeriod(in.GetTimestamp()), fullName)}, nil
}

func (s *helloServer) GetPeriod(timestamp int64) string {
	if timestamp < 1e9 || timestamp > 1e10 {
		return "错误的时间"
	}
	hour := time.Unix(timestamp, 0).Hour()
	period := ""
	if hour >= 0 && hour <= 6 {
		period = "凌晨"
	} else if hour >= 6 && hour <= 12 {
		period = "上午"
	} else if hour > 12 && hour <= 18 {
		period = "下午"
	} else {
		period = "晚上"
	}
	return period
}
