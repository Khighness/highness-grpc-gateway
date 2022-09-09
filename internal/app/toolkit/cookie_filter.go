package toolkit

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"

	"highness-grpc-gateway/internal/pkg/kctx"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-07

type CookieOption struct {
	Path     string
	Domain   string
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite http.SameSite
}

var DefaultCookieOption = &CookieOption{
	Path:     "/",
	Domain:   "127.0.0.1",
	MaxAge:   86400 * 30,
	Secure:   true,
	HttpOnly: true,
	SameSite: http.SameSiteDefaultMode,
}

var cookieKey = "_highness_grpc_gateway_session"

func NewHighnessCookie(cookie string) *http.Cookie {
	return &http.Cookie{
		Name:     cookieKey,
		Value:    cookie,
		Path:     DefaultCookieOption.Path,
		Domain:   DefaultCookieOption.Domain,
		Expires:  time.Now().Add(time.Duration(DefaultCookieOption.MaxAge) * time.Second),
		Secure:   DefaultCookieOption.Secure,
		HttpOnly: DefaultCookieOption.HttpOnly,
		SameSite: DefaultCookieOption.SameSite,
	}
}

// CookieFilter sets cookie before sending http response.
func CookieFilter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	if resp.ProtoReflect().Type().Descriptor().FullName() == "HelloResponse" {
		metaData, _ := metadata.FromOutgoingContext(ctx)
		logger := zap.L().With(zap.Field{
			Key:    kctx.TraceID,
			Type:   zapcore.StringType,
			String: getTraceID(metaData),
		})

		cookieStr := fmt.Sprintf("highness-%d", time.Now().Unix())
		http.SetCookie(w, NewHighnessCookie(cookieStr))

		logger.Info("[GRPC-CookieFilter] Set",
			zap.String(kctx.TraceID, getTraceID(metaData)),
			zap.String("cookie-key", cookieKey),
			zap.String("cookie-value", cookieStr))
	}
	return nil
}
