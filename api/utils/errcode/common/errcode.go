package common

import "university_circles/api/utils/errcode"

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

	ErrUploadFileToOssFailed = errcode.StandardError{50001, "上传文件失败"}
	ErrReportMsgFailed       = errcode.StandardError{50002, "提交数据失败"}
)
