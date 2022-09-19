package kmux

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"

	"google.golang.org/grpc/metadata"

	"highness-grpc-gateway/internal/pkg/kctx"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-08

// RequestMetaHandler sets meta data into ctx.
func RequestMetaHandler(ctx context.Context, request *http.Request) metadata.MD {
	log.Println("Host:", request.Host)
	log.Println("IP:", request.RemoteAddr)
	return metadata.New(map[string]string{
		kctx.MetaTraceID:    kctx.GetTraceID(request.Context()),
		kctx.MetaHttpHost:   request.Host,
		kctx.MetaHttpMethod: request.Method,
		kctx.MetaHttpUrl:    request.URL.String(),
		kctx.MetaHttpParam:  request.URL.Query().Encode(),
	})
}

// GetTraceID gets kctx.MetaTraceID from metadata.MD.
func GetTraceID(md metadata.MD) string {
	if val := md.Get(kctx.MetaTraceID); len(val) > 0 {
		return val[0]
	}
	return ""
}

// GetHttpUrl gets kctx.MetaHttpUrl from metadata.MD.
func GetHttpUrl(md metadata.MD) string {
	if val := md.Get(kctx.MetaHttpUrl); len(val) > 0 {
		return val[0]
	}
	return ""
}

// GetHttpParam gets kctx.MetaHttpParam from metadata.MD.
func GetHttpParam(md metadata.MD) string {
	if val := md.Get(kctx.MetaHttpParam); len(val) > 0 {
		return val[0]
	}
	return ""
}

// getLogger gets zap.Logger with kctx.MetaTraceID from ctx.
func getLogger(ctx context.Context) (*zap.Logger, metadata.MD) {
	metaData, _ := metadata.FromOutgoingContext(ctx)
	traceID := GetTraceID(metaData)
	return zap.L().With(zap.Field{
		Key:    kctx.MetaTraceID,
		Type:   zapcore.StringType,
		String: traceID,
	}), metaData
}
