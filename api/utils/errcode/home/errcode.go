package home

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

	ErrOssStsTokenFailed           = errcode.StandardError{40001, "获取STS token失败"}
	ErrPublishMsgFailed            = errcode.StandardError{40002, "发布动态失败，请重新发布"}
	ErrGetPublishMsgListFailed     = errcode.StandardError{40003, "获取动态失败"}
	ErrGetPublishMsgDetailFailed   = errcode.StandardError{40004, "获取动态详情失败"}
	ErrDelPublishMsgFailed         = errcode.StandardError{40005, "删除动态失败"}
	ErrPublishMsgCommentFailed     = errcode.StandardError{40006, "评论动态失败"}
	ErrDelPublishMsgCommentFailed  = errcode.StandardError{40007, "删除动态评论失败"}
	ErrUserOperatePublishMsgFailed = errcode.StandardError{40008, "保存用户操作动态失败"}
)
