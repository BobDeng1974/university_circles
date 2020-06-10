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

	ErrOssStsTokenFailed = StandardError{40001, "获取STS token失败"}
	ErrFailed            = StandardError{40002, "发布动态失败，请重新发布"}
	ErrGetListFailed     = StandardError{40003, "获取动态失败"}
	ErrGetDetailFailed   = StandardError{40004, "获取动态详情失败"}
	ErrDelFailed         = StandardError{40005, "删除动态失败"}
	ErrCommentFailed     = StandardError{40006, "评论动态失败"}
	ErrDelCommentFailed  = StandardError{40007, "删除动态评论失败"}
	ErrUserOperateFailed = StandardError{40008, "保存用户操作动态失败"}
)
