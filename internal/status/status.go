package status

import (
	"net/http"
	"sync"
)

var (
	errorMaps = map[int]string{}
	mu        sync.Mutex
)

type Error struct {
	code    int
	message string
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Response(errs ...error) (int, *Response) {
	msg := e.Error()
	if len(errs) > 0 {
		msg = errs[0].Error()
	}
	return e.StatusCode(), &Response{Code: e.code, Message: msg}
}

func (e *Error) Message(msg string) *Error {
	err := &Error{code: e.code, message: e.message}
	if len(msg) != 0 {
		err.message = msg
	}
	return err
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case ErrorServer.Code():
		return http.StatusInternalServerError
	default:
		return http.StatusOK
	}
}

func NewError(code int, message string) *Error {
	mu.Lock()
	defer mu.Unlock()
	errorMaps[code] = message
	return &Error{
		code:    code,
		message: message,
	}
}

var (
	Success            = NewError(0, "success")
	ErrorServer        = NewError(1000, "服务内部错误")
	ErrorInvalidParams = NewError(1001, "传入参数错误")
)

func IsRecordNotFound(err error) bool {
	if err != nil {
		if err.Error() == "record not found" {
			return true
		}
	}
	return false
}
