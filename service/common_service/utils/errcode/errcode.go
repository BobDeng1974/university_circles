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

	ErrUploadFileToOssFailed = StandardError{50001, "上传文件失败"}
)
