package errcode

import (
	"fmt"
	"net/http"
)

// Error 声明结构体用于表示错误的响应结果
type Error struct {
	code int 			`json:"code"`
	msg string 			`json:"msg"`
	details []string 	`json:"details"`
}

//作为全局错误码的存储载体、便于查看当前的注册情况
var codes = map[int]string{}

// NewError 创建一个新的error时进行一个排重的校验
func NewError(code int, msg string) *Error  {
	if _ , ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在了、请更换一个",code))
	}

	codes[code] = msg
	return &Error{code: code,msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d、错误信息: %s",e.Code(),e.Msg())
}

func (e *Error)Code() int {
	return e.code
}

func (e *Error)Msg() string {
	return e.msg
}

func (e *Error) Msgf (args []interface{}) string  {
	return fmt.Sprintf(e.msg,args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error)WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

func (e *Error)StatusCode() int  {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.code:
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}

