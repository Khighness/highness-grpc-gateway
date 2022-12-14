package middleware

import (
	"net"
	"net/http"
	"time"

	"highness-grpc-gateway/internal/pkg/kctx"
)

// @Author Chen Zikang
// @Email  parakovo@gmail.com
// @Since  2022-09-08

func Logging(writer http.ResponseWriter, request *http.Request, next func(http.ResponseWriter, *http.Request)) {
	logger := kctx.GetLogger(request.Context())
	start := time.Now()
	method := request.Method
	url := request.URL.String()
	clientIP := request.Header.Get("X-Remote-Addr")
	if clientIP == "" {
		clientIP = request.RemoteAddr
	}
	clientIP, _, _ = net.SplitHostPort(clientIP)
	headers := request.Header

	next(writer, request)

	latency := time.Since(start)
	var statusCode int
	if rw, ok := writer.(*ResponseWriter); ok {
		statusCode = rw.StatusCode()
	}
	logger.Sugar().With("headers", headers).Infof(
		"[MIDDLEWARE-Logging] %3d | %13v | %15s | %-7s %s",
		statusCode,
		latency,
		clientIP,
		method,
		url,
	)
}
