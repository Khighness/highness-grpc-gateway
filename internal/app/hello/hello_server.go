package hello

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"

	"go.uber.org/zap"

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
	period, err := s.GetPeriod(in.GetTimestamp())
	if err != nil {
		return nil, err
	}
	fullName := fmt.Sprintf("%s·%s", in.GetFirstName(), in.GetLastName())
	zap.L().Info("[SayHello]", zap.String("name", fullName))
	return &api.HelloResponse{ReplyMessage: fmt.Sprintf("%s好, %s", period, fullName)}, nil
}

func (s *helloServer) SayHelloV2(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	period, err := s.GetPeriod(in.GetTimestamp())
	if err != nil {
		return nil, err
	}
	fullName := fmt.Sprintf("%s·%s", in.GetFirstName(), in.GetLastName())
	zap.L().Info("[SayHelloV2]", zap.String("name", fullName))
	return &api.HelloResponse{ReplyMessage: fmt.Sprintf("%s好, %s", period, fullName)}, nil
}

func (s *helloServer) SayGoodBye(ctx context.Context, in *api.ByeRequest) (*api.ByeResponse, error) {
	meta := metadata.Pairs("K1", "V1")
	grpc.SetTrailer(ctx, meta)

	fullName := fmt.Sprintf("%s·%s", in.GetFirstName(), in.GetLastName())
	zap.L().Info("[SayGoodBye]", zap.String("name", fullName))
	return &api.ByeResponse{
		ReplyMessage: fmt.Sprintf("再见, %s。 于%s", fullName, time.Unix(in.GetTimestamp(), 0)),
	}, nil
}

func (s *helloServer) GetPeriod(timestamp int64) (string, error) {
	if timestamp < 1e9 || timestamp > 1e10 {
		return "", errors.New("invalid timestamp")
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
	return period, nil
}
