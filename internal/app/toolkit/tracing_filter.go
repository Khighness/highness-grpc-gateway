package toolkit

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-08

// TracingFilter sets cookie before sending http response.
func TracingFilter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	metaData, _ := metadata.FromOutgoingContext(ctx)
	w.Header().Set("X-Trace-Id", getTraceID(metaData))
	return nil
}

func getTraceID(md metadata.MD) string {
	if val := md.Get("trace-id"); len(val) > 0 {
		return val[0]
	}
	return ""
}
