package homeclient

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go.uber.org/zap"
	"time"
	cf "university_circles/api/config"
	"university_circles/api/utils/logger"

	pb "university_circles/api/pb/common"
)

type CommonClient struct {
	client      pb.CommonService
	serviceName string
}

type ClientWrapper struct {
	client.Client
}

const (
	RPC_TIME_OUT_MS = 4000
)

func (c *ClientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(e error) error {
		fmt.Println(e)
		fmt.Println("这是一个备用的服务")
		return e
	})
}

// NewClientWrapper returns a hystrix client Wrapper.
func NewMyClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &ClientWrapper{c}
	}
}

func NewCommonClient() *CommonClient {

	var configFile *cf.Config
	conf := new(cf.Config)

	if err := config.LoadFile(configFile.GetConfigFile()); err != nil {
		fmt.Println(err)
		logger.Logger.Warn(" etcd init error", zap.Error(err))
	}
	if err := config.Scan(conf); err != nil {
		fmt.Println(err)
		logger.Logger.Warn(" etcd init error", zap.Error(err))
	}

	fmt.Println("etcd地址", conf.Etcd.Addr)
	etcdRegisty := etcdv3.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = conf.Etcd.Addr
			//etcdv3.Auth("root","1234")(options)
		})

	// 初始化服务
	service := micro.NewService(
		micro.Name(cf.CLIENT_COMMON_NAME),
		micro.Registry(etcdRegisty),
		//micro.Transport(grpc.NewTransport()),
		micro.WrapClient(NewMyClientWrapper()),
	)

	c := pb.NewCommonService(cf.SRV_COMMON_NAME, service.Client())
	return &CommonClient{
		client:      c,
		serviceName: cf.SRV_COMMON_NAME,
	}
}

func (c *CommonClient) Report(ctx context.Context, req *pb.ReportReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = c.client.Report(ctx, req)
	if err != nil {
		logger.Logger.Warn("client user report msg failed", zap.Any("report", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client Report success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}
