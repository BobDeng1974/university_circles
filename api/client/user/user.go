package userclient

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
	pb "university_circles/api/pb/user"

	"university_circles/api/utils/logger"
)

type UserClient struct {
	client      pb.UserService
	serviceName string
}

type ClientWrapper struct {
	client.Client
}

const (
	RPC_TIME_OUT_MS = 4000
)

func (c *ClientWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		return c.Client.Call(ctx, req, resp, opts...)
	}, func(e error) error {
		fmt.Println(e)
		fmt.Println("这是一个备用的服务")
		return e
	})
}

// NewClientWrapper returns a hystrix client Wrapper.
func NewUserClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &ClientWrapper{c}
	}
}

func NewUserClient() *UserClient {

	var configFile *cf.Config
	conf := new(cf.Config)

	if err := config.LoadFile(configFile.GetConfigFile()); err != nil {
		logger.Logger.Warn(" etcd init error", zap.Error(err))
	}
	if err := config.Scan(conf); err != nil {
		logger.Logger.Warn(" etcd init error", zap.Error(err))
	}

	etcdRegisty := etcdv3.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = conf.Etcd.Addr
			//etcdv3.Auth("root","1234")(options)
		})

	// 初始化服务
	service := micro.NewService(
		micro.Name(cf.CLIENT_USER_NAME),
		micro.Registry(etcdRegisty),
		//micro.Transport(grpc.NewTransport()),
		micro.WrapClient(NewUserClientWrapper()),
	)

	c := pb.NewUserService(cf.SRV_USER_NAME, service.Client())
	return &UserClient{
		client:      c,
		serviceName: cf.SRV_USER_NAME,
	}
}

func (u *UserClient) StudentRegister(ctx context.Context, stu *pb.UserRegisterReq) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.StudentRegister(ctx, stu)
	fmt.Println("UserRegister client", resp, err)
	if err != nil {
		logger.Logger.Warn(" student register failed", zap.Any("student", stu), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client StudentRegister success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", stu), zap.Any("resp", resp))

	return
}

//func (u *UserClient) AddStudentInfo(ctx context.Context, stu *pb.StudentRegInfo) (resp *pb.Response, err error) {
//
//	resp, err = u.client.AddStudentInfo(ctx, stu)
//
//	if err != nil {
//		logger.Logger.Warn("check student is exist failed", zap.Any("student", stu), zap.Error(err))
//		return nil, err
//	}
//	return
//}

func (u *UserClient) GetStudentInfoById(ctx context.Context, req *pb.GetStudentByIdReq, opts ...client.CallOption) (resp *pb.StudentInfoDetail, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.GetStudentInfoById(ctx, req)
	if err != nil {
		logger.Logger.Warn("get student info failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetStudentInfoById success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) GetStudentInfoByUsername(ctx context.Context, req *pb.GetStudentByUsernameReq, opts ...client.CallOption) (resp *pb.StudentInfoDetail, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.GetStudentInfoByUsername(ctx, req)

	if err != nil {
		logger.Logger.Warn("get student info by username failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetStudentInfoByUsername success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) UpdateStudentInfo(ctx context.Context, req *pb.UpdateStudentInfoReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.UpdateStudentInfo(ctx, req)

	if err != nil {
		logger.Logger.Warn("update student info failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client UpdateStudentInfo success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) UpdateUserAvatar(ctx context.Context, req *pb.UpdateUserAvatarReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.UpdateUserAvatar(ctx, req)

	if err != nil {
		logger.Logger.Warn("update user avatar failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client UpdateUserAvatar success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) UpdateUserPhone(ctx context.Context, req *pb.UpdateUserPhoneReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.UpdateUserPhone(ctx, req)

	if err != nil {
		logger.Logger.Warn("update user phone failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client UpdateUserPhone success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.UpdateUserPassword(ctx, req)

	if err != nil {
		logger.Logger.Warn("update user password failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client UpdateUserPassword success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) PwdLogin(ctx context.Context, req *pb.PwdLoginReq, opts ...client.CallOption) (resp *pb.LoginResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.PwdLogin(ctx, req)
	if err != nil {
		logger.Logger.Warn("user pwd login failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client PwdLogin success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}
func (u *UserClient) VerifyCodeLogin(ctx context.Context, req *pb.VerifyCodeLoginReq, opts ...client.CallOption) (resp *pb.LoginResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.VerifyCodeLogin(ctx, req)
	if err != nil {
		logger.Logger.Warn("user verify code login failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client VerifyCodeLogin success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) GetVerifyCode(ctx context.Context, req *pb.VerifyCodeReq, opts ...client.CallOption) (resp *pb.VerifyCodeResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.GetVerifyCode(ctx, req)
	if err != nil {
		logger.Logger.Warn("Verify Code Login failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetVerifyCode success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) CheckVerifyCode(ctx context.Context, req *pb.VerifyCodeRegReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.CheckVerifyCode(ctx, req)
	if err != nil {
		logger.Logger.Warn("Verify Code Login failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client CheckVerifyCode success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) Logout(ctx context.Context, req *pb.LogoutReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.Logout(ctx, req)
	if err != nil {
		logger.Logger.Warn("Verify Code Login failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client Logout success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) GetUniversity(ctx context.Context, req *pb.GetUniversityReq, opts ...client.CallOption) (resp *pb.GetUniversityListResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.GetUniversity(ctx, req)
	if err != nil {
		logger.Logger.Warn("GetUniversityList failed", zap.Any("req", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetUniversity success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
	return nil, nil
}

func (u *UserClient) GetUniversityList(ctx context.Context, req *pb.GetUniversityListReq, opts ...client.CallOption) (resp *pb.GetUniversityListResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.GetUniversityList(ctx, req)
	if err != nil {
		logger.Logger.Warn("GetUniversityList failed", zap.Any("req", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetUniversityList success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) GetCollegeList(ctx context.Context, req *pb.GetCollegeListReq, opts ...client.CallOption) (resp *pb.GetCollegeListResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.GetCollegeList(ctx, req)
	if err != nil {
		logger.Logger.Warn("GetCollegeList failed", zap.Any("req", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetCollegeList success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) GetProfessionList(ctx context.Context, req *pb.GetProfessionListReq, opts ...client.CallOption) (resp *pb.GetProfessionListResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.GetProfessionList(ctx, req)
	if err != nil {
		logger.Logger.Warn("GetScienceList failed", zap.Any("req", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetScienceList success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) GetClassList(ctx context.Context, req *pb.GetClassListReq, opts ...client.CallOption) (resp *pb.GetClassListResponse, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.GetClassList(ctx, req)
	if err != nil {
		logger.Logger.Warn("GetClassList failed", zap.Any("req", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetClassList success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

// 关注某个用户
func (u *UserClient) SaveUserFollow(ctx context.Context, req *pb.UserFollowOperateReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.SaveUserFollow(ctx, req)
	if err != nil {
		logger.Logger.Warn("SaveUserFollow failed", zap.Any("req", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client SaveUserFollow success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

// 取消关注某个用户
func (u *UserClient) CancelUserFollow(ctx context.Context, req *pb.UserFollowOperateReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.CancelUserFollow(ctx, req)
	if err != nil {
		logger.Logger.Warn("SaveUserFollow failed", zap.Any("req", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client CancelUserFollow success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) TeacherRegister(ctx context.Context, stu *pb.UserRegisterReq) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.TeacherRegister(ctx, stu)
	fmt.Println("UserRegister client", resp, err)
	if err != nil {
		logger.Logger.Warn(" teacher register failed", zap.Any("student", stu), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client TeacherRegister success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", stu), zap.Any("resp", resp))

	return
}

func (u *UserClient) GetTeacherInfoById(ctx context.Context, req *pb.GetTeacherByIdReq, opts ...client.CallOption) (resp *pb.TeacherInfoDetail, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.GetTeacherInfoById(ctx, req)
	if err != nil {
		logger.Logger.Warn("get teacher info failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetTeacherInfoById success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) UpdateTeacherInfo(ctx context.Context, req *pb.UpdateTeacherInfoReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.UpdateTeacherInfo(ctx, req)

	if err != nil {
		logger.Logger.Warn("update teacher info failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client UpdateTeacherInfo success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) GetTeacherListByUniversityId(ctx context.Context, req *pb.GetTeacherListByUniIdReq, opts ...client.CallOption) (resp *pb.TeacherListResp, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.GetTeacherListByUniversityId(ctx, req)

	if err != nil {
		logger.Logger.Warn("GetTeacherListByUniversityId failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetTeacherListByUniversityId success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) GetTeacherListByCollegeId(ctx context.Context, req *pb.GetTeacherListByCollegeIdReq, opts ...client.CallOption) (resp *pb.TeacherListResp, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.GetTeacherListByCollegeId(ctx, req)

	if err != nil {
		logger.Logger.Warn("GetTeacherListByCollegeId failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client GetTeacherListByCollegeId success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) QueryUser(ctx context.Context, req *pb.QueryUserReq, opts ...client.CallOption) (resp *pb.QueryUserResp, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.QueryUser(ctx, req)

	if err != nil {
		logger.Logger.Warn("QueryUser failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client QueryUser success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return
}

func (u *UserClient) CheckUserIsExist(ctx context.Context, req *pb.CheckUserIsExistReq, opts ...client.CallOption) (resp *pb.Response, err error) {
	startTm := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*RPC_TIME_OUT_MS)
	defer cancel()

	resp, err = u.client.CheckUserIsExist(ctx, req)

	if err != nil {
		logger.Logger.Warn("CheckUserIsExist failed", zap.Any("user", req), zap.Error(err))
		return nil, err
	}

	logger.Logger.Warn("Client CheckUserIsExist success ", zap.Int64("cost_time_ms", int64(time.Since(startTm)/time.Millisecond)), zap.Any("req", req), zap.Any("resp", resp))

	return 
}

