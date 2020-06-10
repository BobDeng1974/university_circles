package logger

import (
	"go.uber.org/zap"
	"university_circles/service/common_service/utils/zaplog"
)

var (
	// Logger is a run log instance
	Logger *zap.Logger
	cfg    *zaplog.Config
)

func init() {
	cfg := zaplog.Config{
		EncodeLogsAsJson:   true,
		FileLoggingEnabled: true,
		Directory:          "/data/logs/go/",
		Filename:           "university_circles.log",
		MaxSize:            512,
		MaxBackups:         30,
		MaxAge:             7,
	}
	Logger = zaplog.GetLogger(cfg)
}
