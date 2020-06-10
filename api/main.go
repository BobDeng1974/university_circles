package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"go.uber.org/zap"
	"university_circles/api/config"
	"university_circles/api/handler/v1"
	"university_circles/api/utils/logger"
)

func main() {

	webSrv := web.NewService(
		web.Name(config.SRV_NAME),
		web.Address(":20050"),
	)

	routerV1 := v1.ClientEngine()
	routerV1.Use(gin.Logger())
	webSrv.Handle("/", routerV1)
	if err := webSrv.Init(); err != nil {
		logger.Logger.Fatal("server init ", zap.Error(err))
	}

	if err := webSrv.Run(); err != nil {
		logger.Logger.Fatal("run err ", zap.Error(err))
	}

}
