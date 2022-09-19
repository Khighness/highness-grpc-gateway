package kmux

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"net/http"
)

// @Author Chen Zikang
// @Email  parakovo@gmail.com
// @Since  2022-09-09

// RedirectFilter checks if url needs redirect.
func RedirectFilter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	meta, _ := runtime.ServerMetadataFromContext(ctx)
	zap.L().Info("[GRPC-RedirectFilter]", zap.Any("meta", meta))

	return nil
}
