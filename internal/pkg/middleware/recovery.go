package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-08

func Recovery(writer http.ResponseWriter, request *http.Request, next func(http.ResponseWriter, *http.Request)) {
	defer func() {
		zap.L().Info("[MIDDLEWARE-Recovery]")
		if p := recover(); p != nil {
			ctxzap.Extract(request.Context()).Error("recovery", zap.Error(fmt.Errorf("%v", p)))
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusInternalServerError)
			b, _ := json.Marshal(map[string]interface{}{
				"message": "Server Error",
			})
			_, _ = writer.Write(b)
		}
	}()

	next(writer, request)
}
