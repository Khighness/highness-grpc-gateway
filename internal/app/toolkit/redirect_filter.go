package toolkit

import (
	"context"
	"net/http"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"

	"highness-grpc-gateway/internal/pkg/httputil"
	"highness-grpc-gateway/internal/pkg/kctx"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-09

// RedirectFilter checks if url needs redirect.
func RedirectFilter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	metaData, _ := metadata.FromOutgoingContext(ctx)
	logger := zap.L().With(zap.Field{
		Key:    kctx.TraceID,
		Type:   zapcore.StringType,
		String: getTraceID(metaData),
	})

	url := getHttpUrl(metaData)
	if strings.HasPrefix(url, "/v1/hello") {
		url = strings.ReplaceAll(url, "hello", "bye")
		logger.Info("Redirect:", zap.String("url", url))
		httputil.Redirect(w, url, 307)
	}
	return nil
}

func getHttpUrl(md metadata.MD) string {
	if val := md.Get(kctx.HttpUrl); len(val) > 0 {
		return val[0]
	}
	return ""
}
