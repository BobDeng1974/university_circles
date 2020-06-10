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

	ErrAddFriendFailed          = StandardError{60001, "添加好友失败"}
	ErrDelFriendFailed          = StandardError{60002, "删除好友失败"}
	ErrAddFriendBlackListFailed = StandardError{60003, "添加好友黑名单失败"}
	ErrDelFriendBlackListFailed = StandardError{60004, "从黑名单删除好友失败"}
	ErrGetAllFriendFailed       = StandardError{60005, "获取好友列表失败"}
	ErrGetBlackListFailed       = StandardError{60006, "获取黑名单列表失败"}
	ErrCreteGroupFailed         = StandardError{60007, "创建群组失败"}
	ErrDelGroupFailed           = StandardError{60008, "解散群组失败"}
	ErrGetGroupMemberFailed     = StandardError{60009, "获取群组成员失败"}
	ErrGroupListFailed          = StandardError{600010, "获取群组列表失败"}
	ErrJoinGroupFailed          = StandardError{600011, "加入群组失败"}
	ErrIMAuthTokenFailed        = StandardError{600012, "获取im auth token失败"}
	ErrIsFriendFailed           = StandardError{600013, "已经是好友了"}
	ErrIsBlackListFailed        = StandardError{600014, "已经添加到黑名单了"}
	ErrGroupNameIsExistFailed   = StandardError{600015, "群组名已存在了，请重新命名群组"}
	ErrAuthFailed               = StandardError{600016, "无操作权限"}
	ErrUserNotExist             = StandardError{600017, "账号不存在，请重试"}
	ErrUpdateFriendRemarkFailed = StandardError{600018, "修改好友备注失败"}
	ErrUpdateGroupAvatarFailed  = StandardError{600019, "设置群组头像失败"}
	ErrUpdateGroupNoticeFailed  = StandardError{600020, "设置群组备注失败"}
	ErrUpdateGroupNameFailed  	= StandardError{600021, "设置群组名称失败"}
	ErrUpdateGroupJoinAuthFailed  	= StandardError{600022, "设置群组加入权限失败"}
	ErrDelGroupMemberFailed  	= StandardError{600023, "删除群组成员失败"}
	ErrGroupNotExist  	= StandardError{600024, "群组不存在"}
)
