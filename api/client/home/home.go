package homeclient

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"time"

	"github.com/afex/hystrix-go/hystrix"

	"go.uber.org/zap"
	cf "university_circles/api/config"
	pb "university_circles/api/pb/home"

	"university_circles/api/utils/logger"
)

type HomeClient struct {
	client      pb.HomeService
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

func NewHomeClient() *HomeClient {

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
		micro.Name(cf.CLIENT_HOME_NAME),
		micro.Registry(etcdRegisty),
		//micro.Transport(grpc.NewTransport()),
		micro.WrapClient(NewMyClientWrapper()),
	)

	c := pb.NewHomeService(cf.SRV_HOME_NAME, service.Client())
	return &HomeClient{
		client:      c,
		serviceName: cf.SRV_HOME_NAME,
	}
}

func (h *HomeClient) SavePublishMsg(ctx context.Context, req *pb.PublishMsg, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = h.client.SavePublishMsg(ctx, req)

	if err != nil {
		logger.Logger.Warn("client publish msg failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client SavePublishMsg success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (h *HomeClient) DeletePublishMsg(ctx context.Context, req *pb.DeleteMsgRequest, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = h.client.DeletePublishMsg(ctx, req)
	if err != nil {
		logger.Logger.Warn("client delete user msg failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client DeletePublishMsg success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (h *HomeClient) GetHomeMsgList(ctx context.Context, req *pb.AllMsgListRequest, opts ...client.CallOption) (resp *pb.HomeMsgListResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = h.client.GetHomeMsgList(ctx, req)
	if err != nil {
		logger.Logger.Warn("client Get home Msg List failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetHomeMsgList success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (h *HomeClient) GetUserMsgList(ctx context.Context, req *pb.UserMsgListRequest, opts ...client.CallOption) (resp *pb.OtherMsgListResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = h.client.GetUserMsgList(ctx, req)
	if err != nil {
		logger.Logger.Warn("client Get user Msg List failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetUserMsgList success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (h *HomeClient) GetPublishMsgDetail(ctx context.Context, req *pb.OneMsgRequest, opts ...client.CallOption) (resp *pb.PublishMsgResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	res, err := h.client.GetMsgDetail(ctx, req)
	fmt.Println("client ", resp, err)
	if err != nil {
		logger.Logger.Warn("client Get user Msg detail failed", zap.Error(err))
		return nil, err
	}
	resp = res.PublishMsgResponseList[0]

	logger.Logger.Warn("Client GetMsgDetail success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (h *HomeClient) SaveMsgComment(ctx context.Context, req *pb.PublishMsgComment, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = h.client.SaveMsgComment(ctx, req)
	if err != nil {
		logger.Logger.Warn("client save msg comment failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client SaveMsgComment success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (h *HomeClient) GetMsgCommentList(ctx context.Context, req *pb.MsgCommentListRequest, opts ...client.CallOption) (resp *pb.MsgCommentListResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = h.client.GetMsgCommentList(ctx, req)
	if err != nil {
		logger.Logger.Warn("client get msg comment list failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetMsgCommentList success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (h *HomeClient) DeletePublishMsgComment(ctx context.Context, req *pb.DeleteMsgCommentRequest, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = h.client.DeletePublishMsgComment(ctx, req)
	if err != nil {
		logger.Logger.Warn("client del msg comment failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client DeletePublishMsgComment success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (h *HomeClient) SaveUserOperateMsgCount(ctx context.Context, req *pb.UserOperateCountRequest, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = h.client.SaveUserOperateMsgCount(ctx, req)
	if err != nil {
		logger.Logger.Warn("client save user operate publish msg count failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client SaveUserOperateMsgCount success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (h *HomeClient) GetUserOperateMsgCount(ctx context.Context, req *pb.UserOperateCountRequest, opts ...client.CallOption) (resp *pb.UserOperateCountResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = h.client.GetUserOperateMsgCount(ctx, req)
	if err != nil {
		logger.Logger.Warn("client get user operate publish msg record list failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetUserOperateMsgCount success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (h *HomeClient) GetUserOperateMsgRecodeList(ctx context.Context, req *pb.UserOperateRecodeListRequest, opts ...client.CallOption) (resp *pb.OtherMsgListResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = h.client.GetUserOperateMsgRecodeList(ctx, req)
	if err != nil {
		logger.Logger.Warn("client get user operate publish msg record list failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetUserOperateMsgRecodeList success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}
