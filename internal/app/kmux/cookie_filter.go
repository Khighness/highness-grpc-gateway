package kmux

import (
	"context"
	"fmt"
	"highness-grpc-gateway/proto/api"
	"net/http"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
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
	MaxAge:   10,
	Secure:   false,
	HttpOnly: true,
	SameSite: http.SameSiteDefaultMode,
}

var cookieKey = "_highness_grpc_gateway_session"

func NewHighnessCookie(cookie string) *http.Cookie {
	return &http.Cookie{
		Name:     cookieKey,
		Value:    cookie,
		Path:     DefaultCookieOption.Path,
		Expires:  time.Now().Add(time.Duration(DefaultCookieOption.MaxAge) * time.Second),
		Secure:   DefaultCookieOption.Secure,
		HttpOnly: DefaultCookieOption.HttpOnly,
	}
}

// CookieFilter sets cookie before sending http response.
func CookieFilter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	logger, _ := getLogger(ctx)
	descriptor := resp.ProtoReflect().Type().Descriptor()
	logger.Sugar().Infof("name: %v, fullname: %v", descriptor.Name(), descriptor.FullName())
	if descriptor.FullName() == "HelloResponse" {
		response := resp.(*api.HelloResponse)
		logger.Info("[CookieFilter]", zap.Any("resp", response))

		cookieStr := fmt.Sprintf("highness-%d", time.Now().Unix())
		//http.SetCookie(w, &http.Cookie{
		//	Name:    cookieKey,
		//	Value:   cookieStr,
		//	Expires: time.Now().Add(1 * time.Minute),
		//})
		http.SetCookie(w, NewHighnessCookie(cookieStr))

		logger.Info("[GRPC-CookieFilter] Set",
			zap.String("cookie-key", cookieKey),
			zap.String("cookie-value", cookieStr))
	}
	return nil
}
