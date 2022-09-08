package hello

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"

	"highness-grpc-gateway/proto/api"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-07

type helloServer struct {
	api.HelloServiceServer
}

func NewHelloServer() api.HelloServiceServer {
	return &helloServer{}
}

func (s *helloServer) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	fullName := fmt.Sprintf("%s·%s", in.GetFirstName(), in.GetLastName())
	zap.L().Info("[SayHello]", zap.String("name", fullName))
	return &api.HelloResponse{ReplyMessage: fmt.Sprintf("%s好, %s", s.GetPeriod(in.GetTimestamp()), fullName)}, nil
}

func (s *helloServer) SayHelloV2(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	fullName := fmt.Sprintf("%s·%s", in.GetFirstName(), in.GetLastName())
	zap.L().Info("[SayHelloV2]", zap.String("name", fullName))
	return &api.HelloResponse{ReplyMessage: fmt.Sprintf("%s好, %s", s.GetPeriod(in.GetTimestamp()), fullName)}, nil
}

func (s *helloServer) SayGoodBye(ctx context.Context, in *api.ByeRequest) (*api.ByeResponse, error) {
	fullName := fmt.Sprintf("%s·%s", in.GetFirstName(), in.GetLastName())
	zap.L().Info("[SayGoodBye]", zap.String("name", fullName))
	return &api.ByeResponse{
		ReplyMessage: fmt.Sprintf("再见, %s。 于%s", fullName, time.Unix(in.GetTimestamp(), 0)),
	}, nil
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
