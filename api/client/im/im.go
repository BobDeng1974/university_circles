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
	pb "university_circles/api/pb/im"

	"university_circles/api/utils/logger"
)

type ImClient struct {
	client      pb.ImService
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

func NewImClient() *ImClient {

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
		micro.Name(cf.CLIENT_IM_NAME),
		micro.Registry(etcdRegisty),
		//micro.Transport(grpc.NewTransport()),
		micro.WrapClient(NewMyClientWrapper()),
	)

	c := pb.NewImService(cf.SRV_IM_NAME, service.Client())
	return &ImClient{
		client:      c,
		serviceName: cf.SRV_IM_NAME,
	}
}

func (ic *ImClient) SendPeerMsg(ctx context.Context, req *pb.SendPeerMsgReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.SendPeerMsg(ctx, req)

	if err != nil {
		logger.Logger.Warn("client publish msg failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client SendAddFriendMsg success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (ic *ImClient) AddFriends(ctx context.Context, req *pb.AddFriendReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.AddFriends(ctx, req)

	if err != nil {
		logger.Logger.Warn("client publish msg failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client AddFriends success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (ic *ImClient) DelFriends(ctx context.Context, req *pb.DelFriendReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.DelFriends(ctx, req)
	fmt.Println("DelFriends client", resp, err)
	if err != nil {
		logger.Logger.Warn("client DelFriends failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client DelFriends success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) UpdateFriendRemark(ctx context.Context, req *pb.UpdateFriendRemarkReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.UpdateFriendRemark(ctx, req)
	if err != nil {
		logger.Logger.Warn("client UpdateFriendRemark failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client UpdateFriendRemark success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}


func (ic *ImClient) AddFriendBlackList(ctx context.Context, req *pb.AddFriendBlackListReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.AddFriendBlackList(ctx, req)
	fmt.Println("AddFriendBlackList client", resp, err)
	if err != nil {
		logger.Logger.Warn("client AddFriendBlackList failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client AddFriendBlackList success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) DelFriendBlackList(ctx context.Context, req *pb.DelFriendBlackListReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.DelFriendBlackList(ctx, req)
	fmt.Println("DelFriendBlackList client", resp, err)
	if err != nil {
		logger.Logger.Warn("client DelFriendBlackList failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client DelFriendBlackList success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) GetAllFriends(ctx context.Context, req *pb.FriendListReq, opts ...client.CallOption) (resp *pb.FriendListResp, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("GetAllFriends srv", err)
			return
		}
	}()
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.GetAllFriends(ctx, req)
	fmt.Println("GetAllFriends client", resp, err)
	if err != nil {
		logger.Logger.Warn("client GetAllFriends failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetAllFriends success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) GetBlackList(ctx context.Context, req *pb.BlackListReq, opts ...client.CallOption) (resp *pb.BlackListResp, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.GetBlackList(ctx, req)
	fmt.Println("GetBlackList client", resp, err)
	if err != nil {
		logger.Logger.Warn("client GetBlackList failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetBlackList success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) CreateGroup(ctx context.Context, req *pb.Group, opts ...client.CallOption) (resp *pb.CreateGroupResp, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.CreateGroup(ctx, req)
	fmt.Println("CreateGroup client", resp, err)
	if err != nil {
		logger.Logger.Warn("client CreateGroup failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client reqdentRegister success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) DelGroup(ctx context.Context, req *pb.DelGroupReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.DelGroup(ctx, req)
	fmt.Println("DelGroup client", resp, err)
	if err != nil {
		logger.Logger.Warn("client DelGroup failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client DelGroup success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) JoinGroup(ctx context.Context, req *pb.JoinGroupReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.JoinGroup(ctx, req)
	fmt.Println("JoinGroup client", resp, err)
	if err != nil {
		logger.Logger.Warn("client JoinGroup failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client JoinGroup success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) DelGroupMember(ctx context.Context, req *pb.DelGroupMemberReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.DelGroupMember(ctx, req)

	if err != nil {
		logger.Logger.Warn("client del group member failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client DelGroupMember success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) GetGroupMemberList(ctx context.Context, req *pb.GroupMemberListReq, opts ...client.CallOption) (resp *pb.GroupMemberListResp, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.GetGroupMemberList(ctx, req)

	if err != nil {
		logger.Logger.Warn("client publish msg failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetGroupMember success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) GetGroupList(ctx context.Context, req *pb.GroupListReq, opts ...client.CallOption) (resp *pb.GroupListResp, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.GetGroupList(ctx, req)
	fmt.Println("GetGroupList client", resp, err)
	if err != nil {
		logger.Logger.Warn("client GetGroupList failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetGroupList success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) GetUserUnReadCount(ctx context.Context, req *pb.GetUserUnReadCountReq, opts ...client.CallOption) (resp *pb.GetUserUnReadCountResp, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.GetUserUnReadCount(ctx, req)
	fmt.Println("GetUserUnReadCount client", resp, err)
	if err != nil {
		logger.Logger.Warn("client GetUserUnReadCount failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetUserUnReadCount success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) UpdateGroupAvatar(ctx context.Context, req *pb.UpdateGroupAvatarReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.UpdateGroupAvatar(ctx, req)
	fmt.Println("UpdateGroupAvatar client", resp, err)
	if err != nil {
		logger.Logger.Warn("client UpdateGroupAvatar failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client UpdateGroupAvatar success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) UpdateGroupNotice(ctx context.Context, req *pb.UpdateGroupNoticeReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.UpdateGroupNotice(ctx, req)
	fmt.Println("UpdateGroupNotice client", resp, err)
	if err != nil {
		logger.Logger.Warn("client UpdateGroupNotice failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client UpdateGroupNotice success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) UpdateGroupName(ctx context.Context, req *pb.UpdateGroupNameReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.UpdateGroupName(ctx, req)
	fmt.Println("UpdateGroupName client", resp, err)
	if err != nil {
		logger.Logger.Warn("client UpdateGroupName failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client UpdateGroupName success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}

func (ic *ImClient) UpdateGroupJoinAuth(ctx context.Context, req *pb.UpdateGroupJoinAuthReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = ic.client.UpdateGroupJoinAuth(ctx, req)
	fmt.Println("UpdateGroupJoinAuth client", resp, err)
	if err != nil {
		logger.Logger.Warn("client UpdateGroupJoinAuth failed", zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client UpdateGroupJoinAuth success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return resp, nil
}
