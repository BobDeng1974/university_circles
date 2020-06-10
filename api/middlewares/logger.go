package middlewares

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"university_circles/api/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// RequestLogger is the logrus logger handler
func RequestLogger() gin.HandlerFunc {

	return func(c *gin.Context) {
		traceID := uuid.New().String()

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		start := time.Now()
		/*
		   c.Next()后就执行真实的路由函数，路由函数执行完成之后继续执行后续的代码
		*/
		c.Next()
		latency := time.Since(start)

		var code int
		var response interface{}
		if respMap, err := formatResponse(blw.body.Bytes()); err == nil {
			response = respMap
			code, _ = respMap["code"].(int)
		} else {
			response = blw.body.Bytes()
		}

		// HEAD 请求为探测
		if c.Request.Method == "HEAD" {
			return
		}

		logger.Logger.Info(
			"access_log",
			zap.String("id", traceID),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
			zap.Any("body", c.Request.Form),
			zap.Any("cookies", formatCookies(c.Request.Cookies())),
			zap.Any("response", response),
			zap.Any("code", code),
		)

		if len(c.Errors) > 0 {
			logger.Logger.Warn(c.Errors.ByType(gin.ErrorTypeAny).String())
		}

		if c.Request.Header.Get("HTTP_TENCENT_LEAKSCAN") != "" {
			return
		}
	}
}

func formatCookies(cookies []*http.Cookie) (cookiesMap map[string]string) {
	cookiesMap = make(map[string]string)
	for _, cookie := range cookies {
		cookiesMap[cookie.Name] = cookie.Value
	}
	return
}

func formatResponse(data []byte) (map[string]interface{}, error) {
	var responseMap = make(map[string]interface{})
	err := json.Unmarshal(data, &responseMap)
	if err != nil {
		return responseMap, err
	}
	return responseMap, nil
}
