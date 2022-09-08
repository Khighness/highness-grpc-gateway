package middleware

import (
	"net/http"
	"sync"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-07

type Middleware func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))

func Middlewares() []Middleware {
	return []Middleware{
		Logging,
		Recovery,
	}
}

func WithMiddleWares(base http.Handler, middlewares ...Middleware) (handler http.Handler) {
	handler = base
	for i := len(middlewares) - 1; i >= 0; i-- {
		middleware := middlewares[i]
		handler = &handlerWithMiddleware{
			base:       handler,
			middleware: middleware,
		}
	}
	return handler
}

type handlerWithMiddleware struct {
	base       http.Handler
	middleware Middleware
}

func (h handlerWithMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	wrappedWriter, ok := writer.(*ResponseWriter)
	if !ok {
		wrappedWriter = &ResponseWriter{ResponseWriter: writer}
	}

	h.middleware(wrappedWriter, request, func(rw http.ResponseWriter, r *http.Request) {
		h.base.ServeHTTP(rw, r)
	})
}

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
	once       sync.Once
}

func (wr *ResponseWriter) Write(data []byte) (int, error) {
	wr.once.Do(func() {
		wr.statusCode = http.StatusOK
	})
	return wr.ResponseWriter.Write(data)
}

func (wr *ResponseWriter) WriteHeader(statusCode int) {
	wr.once.Do(func() {
		wr.statusCode = statusCode
	})
	wr.ResponseWriter.WriteHeader(statusCode)
}

func (wr *ResponseWriter) StatusCode() int {
	if wr.statusCode != 0 {
		return wr.statusCode
	}
	return wr.statusCode
}
