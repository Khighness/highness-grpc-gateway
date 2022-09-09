package toolkit

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"

	"highness-grpc-gateway/internal/pkg/kctx"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-08

// TracingFilter sets cookie before sending http response.
func TracingFilter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	metaData, _ := metadata.FromOutgoingContext(ctx)
	w.Header().Set(kctx.TraceID, getTraceID(metaData))
	return nil
}
