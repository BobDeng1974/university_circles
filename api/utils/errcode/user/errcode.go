package user

import (
	"university_circles/api/utils/errcode"
)

var (
	Success            = errcode.StandardError{0, "success"}
	ErrUnknown         = errcode.StandardError{-1, "unknown error"}
	ErrReqForbidden    = errcode.StandardError{11001, "request forbidden"}
	ErrAuthExpired     = errcode.StandardError{11002, "signature is expired"}
	ErrParam           = errcode.StandardError{11003, "params error"}
	ErrMethodIncorrect = errcode.StandardError{11004, "method incorrect"}
	ErrTimeout         = errcode.StandardError{11005, "timeout"}
	ErrDbQuery         = errcode.StandardError{11006, "db query error"}
	ErrHttpResponse    = errcode.StandardError{11007, "http response not content"}

	ErrLoginSessionInvalid        = errcode.StandardError{20000, "token失效，请重新登录"}
	ErrRegisterStuNoExist         = errcode.StandardError{20001, "学号/编号已存在，请重新注册"}
	ErrRegisterIDCardNumberExist  = errcode.StandardError{20002, "身份证号已存在，请重新注册"}
	ErrRegisterEmailExist         = errcode.StandardError{20003, "邮箱已存在，请重新注册"}
	ErrRegisterPhoneExist         = errcode.StandardError{20004, "手机号已存在，请重新注册"}
	ErrRegisterNicknameExist      = errcode.StandardError{20005, "用户名已存在，请重新操作"}
	ErrLoginUserNameFailed        = errcode.StandardError{20006, "账号不正确，请重新登录"}
	ErrLoginUserPwdFailed         = errcode.StandardError{20007, "密码不正确，请重新登录"}
	ErrLoginUserNameOrPwdFailed   = errcode.StandardError{20008, "账号或密码不正确，请重新登录"}
	ErrRegisterFailed             = errcode.StandardError{20009, "注册失败，请重新注册"}
	ErrSaveFailed                 = errcode.StandardError{200010, "保存信息失败，请重试"}
	ErrVerifyCodeFailed           = errcode.StandardError{200011, "获取验证码失败，请重试"}
	ErrLogoutFailed               = errcode.StandardError{200012, "退出失败，请重试"}
	ErrUserNotExist               = errcode.StandardError{200013, "账号不存在，请重试"}
	ErrUserNotLogin               = errcode.StandardError{200014, "账号未登录，请登录"}
	ErrVerifyCodeNotExist         = errcode.StandardError{200015, "验证码失效，请重新获取"}
	ErrVerifyCodeCompareFailed    = errcode.StandardError{200016, "验证码错误，请重新登录"}
	ErrLoginFailed                = errcode.StandardError{200017, "登录失败，请重新登录"}
	ErrModifyUserInfoFailed       = errcode.StandardError{200018, "修改失败，请重新修改"}
	ErrGetUniversityListFailed    = errcode.StandardError{200019, "获取大学列表失败，请重新获取"}
	ErrGetCollegeListFailed       = errcode.StandardError{200020, "获取学院列表失败，请重新获取"}
	ErrGetScienceListFailed       = errcode.StandardError{200020, "获取科系列表失败，请重新获取"}
	ErrGetClassListFailed         = errcode.StandardError{200020, "获取班级列表失败，请重新获取"}
	ErrUserFollowFailed           = errcode.StandardError{200021, "关注失败"}
	ErrCancelUserFollowFailed     = errcode.StandardError{200022, "取消关注失败"}
	ErrRegisterPhoneNotExist      = errcode.StandardError{200023, "手机号不存在，请先注册"}
	ErrUserIsNotVerify            = errcode.StandardError{200024, "用户未验证，请提交验证信息"}
	ErrUserNameVerifyInvalid      = errcode.StandardError{200025, "校验的姓名为空或非法，请重新提交验证信息"}
	ErrUserIdCardVerifyInvalid    = errcode.StandardError{200026, "校验的身份证为空或非法，请重新提交验证信息"}
	ErrUserPhoneVerifyInvalid     = errcode.StandardError{200027, "校验的手机号为空或非法，请重新提交验证信息"}
	ErrRegisterUniversityNotExist = errcode.StandardError{200028, "学校不存在，请重新填写"}
	ErrRegisterProfessionNotExist = errcode.StandardError{200029, "专业不存在，请重新填写"}
	ErrUploadFileToOssFailed      = errcode.StandardError{200030, "上传文件失败"}
	ErrQueryUserFailed            = errcode.StandardError{200031, "查询用户失败"}

	// 短信
	ErrBusinessMinuteLimitControlFailed = errcode.StandardError{300001, "验证码请求过快，请一分钟后再请求"}
	ErrSendVerifyCodeFailed             = errcode.StandardError{300002, "验证码请求错误，请重新请求"}
	ErrBusinessDayLimitControlFailed    = errcode.StandardError{300003, "验证码请求次数过多，请稍后再获取"}
)
