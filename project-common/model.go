package common

import "net/http"

type BusinessCode int
type Result struct {
	Code BusinessCode `json:"code"`
	Msg  string       `json:"msg"`
	Data any          `json:"data"`
}

// Success 成功返回
func (r *Result) Success(data any) *Result {
	r.Code = http.StatusOK
	r.Msg = "success"
	r.Data = data
	return r
}

// Fail 失败返回
func (r *Result) Fail(code BusinessCode, msg string) *Result {
	r.Code = code
	r.Msg = msg
	return r
}
