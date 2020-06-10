package im

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

	ErrAddFriendFailed          = errcode.StandardError{60001, "添加好友失败"}
	ErrDelFriendFailed          = errcode.StandardError{60002, "删除好友失败"}
	ErrAddFriendBlackListFailed = errcode.StandardError{60003, "添加好友黑名单失败"}
	ErrDelFriendBlackListFailed = errcode.StandardError{60004, "从黑名单删除好友失败"}
	ErrGetAllFriendFailed       = errcode.StandardError{60005, "获取好友列表失败"}
	ErrGetBlackListFailed       = errcode.StandardError{60006, "获取黑名单列表失败"}
	ErrCreteGroupFailed         = errcode.StandardError{60007, "创建群组失败"}
	ErrDelGroupFailed           = errcode.StandardError{60008, "解散群组失败"}
	ErrGetGroupMemberFailed     = errcode.StandardError{60009, "获取群组成员失败"}
	ErrGroupListFailed          = errcode.StandardError{600010, "获取群组列表失败"}
	ErrJoinGroupFailed          = errcode.StandardError{600011, "加入群组失败"}
	ErrIMAuthTokenFailed        = errcode.StandardError{600012, "获取im auth token失败"}
	ErrIsFriendFailed           = errcode.StandardError{600013, "已经是好友了"}
	ErrIsBlackListFailed        = errcode.StandardError{600014, "已经添加到黑名单了"}
	ErrGroupNameIsExistFailed   = errcode.StandardError{600015, "群组名已存在了，请重新命名群组"}
	ErrAuthFailed               = errcode.StandardError{600016, "无操作权限"}
	ErrUserNotExist             = errcode.StandardError{600017, "账号不存在，请重试"}
	ErrUpdateFriendRemarkFailed = errcode.StandardError{600018, "修改好友备注失败"}
	ErrUpdateGroupAvatarFailed  = errcode.StandardError{600019, "设置群组头像失败"}
	ErrUpdateGroupNoticeFailed  = errcode.StandardError{600020, "设置群组备注失败"}
	ErrUpdateGroupNameFailed  	= errcode.StandardError{600021, "设置群组名称失败"}
	ErrUpdateGroupJoinAuthFailed  	= errcode.StandardError{600022, "设置群组加入权限失败"}
	ErrDelGroupMemberFailed  	= errcode.StandardError{600023, "删除群组成员失败"}
	ErrGroupNotExist  	= errcode.StandardError{600024, "群组不存在"}
)
