package errcode

import "fmt"

// StandardError struct
type StandardError struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func (err StandardError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", err.Code, err.Msg)
}

// WithMsg is method to set error msg
func (err StandardError) WithMsg(msg string) StandardError {
	err.Msg = err.Msg + ": " + msg
	return err
}

// New StarandError
func New(code int64, msg string) StandardError {
	return StandardError{code, msg}
}

var (
	Success            = StandardError{0, "success"}
	ErrUnknown         = StandardError{-1, "unknown error"}
	ErrReqForbidden    = StandardError{11001, "request forbidden"}
	ErrAuthExpired     = StandardError{11002, "signature is expired"}
	ErrParam           = StandardError{11003, "params error"}
	ErrMethodIncorrect = StandardError{11004, "method incorrect"}
	ErrTimeout         = StandardError{11005, "timeout"}
	ErrDbQuery         = StandardError{11006, "db query error"}
	ErrHttpResponse    = StandardError{11007, "http response not content"}

	ErrLoginSessionInvalid        = StandardError{20000, "token失效，请重新登录"}
	ErrRegisterStuNoExist         = StandardError{20001, "学号/编号已存在，请重新注册"}
	ErrRegisterIDCardNumberExist  = StandardError{20002, "身份证号已存在，请重新注册"}
	ErrRegisterEmailExist         = StandardError{20003, "邮箱已存在，请重新注册"}
	ErrRegisterPhoneExist         = StandardError{20004, "手机号已存在，请重新注册"}
	ErrRegisterNicknameExist      = StandardError{20005, "用户名已存在，请重新操作"}
	ErrLoginUserNameFailed        = StandardError{20006, "账号不正确，请重新登录"}
	ErrLoginUserPwdFailed         = StandardError{20007, "密码不正确，请重新登录"}
	ErrLoginUserNameOrPwdFailed   = StandardError{20008, "账号或密码不正确，请重新登录"}
	ErrRegisterFailed             = StandardError{20009, "注册失败，请重新注册"}
	ErrSaveFailed                 = StandardError{200010, "保存信息失败，请重试"}
	ErrVerifyCodeFailed           = StandardError{200011, "获取验证码失败，请重试"}
	ErrLogoutFailed               = StandardError{200012, "退出失败，请重试"}
	ErrUserNotExist               = StandardError{200013, "账号不存在，请重试"}
	ErrUserNotLogin               = StandardError{200014, "账号未登录，请登录"}
	ErrVerifyCodeNotExist         = StandardError{200015, "验证码失效，请重新获取"}
	ErrVerifyCodeCompareFailed    = StandardError{200016, "验证码错误，请重新登录"}
	ErrLoginFailed                = StandardError{200017, "登录失败，请重新登录"}
	ErrModifyUserInfoFailed       = StandardError{200018, "修改失败，请重新修改"}
	ErrGetUniversityListFailed    = StandardError{200019, "获取大学列表失败，请重新获取"}
	ErrGetCollegeListFailed       = StandardError{200020, "获取学院列表失败，请重新获取"}
	ErrGetScienceListFailed       = StandardError{200020, "获取科系列表失败，请重新获取"}
	ErrGetClassListFailed         = StandardError{200020, "获取班级列表失败，请重新获取"}
	ErrUserFollowFailed           = StandardError{200021, "关注失败"}
	ErrCancelUserFollowFailed     = StandardError{200022, "取消关注失败"}
	ErrRegisterPhoneNotExist      = StandardError{200023, "手机号不存在，请先注册"}
	ErrUserIsNotVerify            = StandardError{200024, "用户未验证，请提交验证信息"}
	ErrUserNameVerifyInvalid      = StandardError{200025, "校验的姓名为空或非法，请重新提交验证信息"}
	ErrUserIdCardVerifyInvalid    = StandardError{200026, "校验的身份证为空或非法，请重新提交验证信息"}
	ErrUserPhoneVerifyInvalid     = StandardError{200027, "校验的手机号为空或非法，请重新提交验证信息"}
	ErrRegisterUniversityNotExist = StandardError{200028, "学校不存在，请重新填写"}
	ErrRegisterProfessionNotExist = StandardError{200029, "专业不存在，请重新填写"}
	ErrUploadFileToOssFailed      = StandardError{200030, "上传文件失败"}
	ErrQueryUserFailed            = StandardError{200031, "查询用户失败"}

	// 短信
	ErrBusinessMinuteLimitControlFailed = StandardError{300001, "验证码请求过快，请一分钟后再请求"}
	ErrSendVerifyCodeFailed             = StandardError{300002, "验证码请求错误，请重新请求"}
	ErrBusinessDayLimitControlFailed    = StandardError{300003, "验证码请求次数过多，请稍后再获取"}
)
