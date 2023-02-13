package errs

import "fmt"

type ErrorCode int

// BError 通用错误码
type BError struct {
	Code ErrorCode
	Msg  string
}

// Error 实现error接口
func (e *BError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

// NewError 创建错误
func NewError(code ErrorCode, msg string) *BError {
	return &BError{
		Code: code,
		Msg:  msg,
	}
}
