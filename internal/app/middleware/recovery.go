package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
	"highness-grpc-gateway/internal/pkg/kctx"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-08

func Recovery(writer http.ResponseWriter, request *http.Request, next func(http.ResponseWriter, *http.Request)) {
	defer func() {
		logger := kctx.GetLogger(request.Context())
		logger.Info("[MIDDLEWARE-Recovery]")
		if p := recover(); p != nil {
			logger.Error("[MIDDLEWARE-Recovery]", zap.Error(fmt.Errorf("%v", p)))
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
