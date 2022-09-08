package toolkit

import (
	"context"
	"fmt"
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
		cookieStr := fmt.Sprintf("highness-%d", time.Now().Unix())
		http.SetCookie(w, NewHighnessCookie(cookieStr))
		zap.L().Info("[GRPC-CookieFilter] Set",
			zap.String("cookie_key", cookieKey),
			zap.String("cookie_value", cookieStr))
	}
	return nil
}
