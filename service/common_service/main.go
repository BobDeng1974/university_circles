package main

import (
	"flag"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go.uber.org/zap"

	"github.com/micro/go-micro/config"
	"time"

	cf "university_circles/service/common_service/config"
	"university_circles/service/common_service/handler"
	pb "university_circles/service/common_service/pb/common"
	"university_circles/service/common_service/utils/logger"
	"university_circles/service/common_service/wrapper"
)

func main() {
	configFile := flag.String("f", "./config/config.json", "please use config.json")
	conf := new(cf.Config)

	flag.Parse()

	if err := config.LoadFile(*configFile); err != nil {
		logger.Logger.Fatal(" etcd init error", zap.Error(err))
		return
	}
	if err := config.Scan(conf); err != nil {
		logger.Logger.Fatal(" etcd init error", zap.Error(err))
		return
	}

	etcdRegisty := etcdv3.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = conf.Etcd.Addr
			//etcdv3.Auth("root","1234")(options)
		})

	// New Service
	service := micro.NewService(
		micro.Name(cf.SRV_NAME),
		micro.Version("latest"),
		micro.WrapClient(wrapper.LogClientWrap),
		micro.Registry(etcdRegisty),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		//micro.WrapHandler(ratelimit.NewHandlerWrapper(&rate.Bucket{}, false)),
	)
	// 必须提前初始化
	err := cmd.Init()
	if err != nil {
		logger.Logger.Fatal(" cmd init error", zap.Error(err))
	}
	// Initialise service
	service.Init()

	// Register Handler
	err = pb.RegisterCommonServiceHandler(service.Server(), new(handler.CommonHandler))
	if err != nil {
		logger.Logger.Fatal("register home handler error", zap.Error(err))
	}

	if err := broker.Init(); err != nil {
		logger.Logger.Fatal("Broker Init error", zap.Error(err))
	}

	if err := broker.Connect(); err != nil {
		logger.Logger.Fatal("Broker Connect error", zap.Error(err))
	}

	// Run service
	if err := service.Run(); err != nil {
		logger.Logger.Fatal("home service Run error", zap.Error(err))
	}
}
