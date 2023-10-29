package status

import (
	"database/sql"
	"eos-layout/pkg/verifier"
	"errors"
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

func (e Error) Error() string {
	return e.message
}

func (e Error) Code() int {
	return e.code
}

func (e Error) Response(errs ...error) (int, *Response) {
	msg := e.Error()
	if len(errs) > 0 {
		msg = verifier.Translate(errs[0])
	}
	return e.StatusCode(), &Response{Code: e.code, Message: msg}
}

func (e Error) Message(msg string) *Error {
	err := &Error{code: e.code, message: e.message}
	if len(msg) != 0 {
		err.message = msg
	}
	return err
}

func (e Error) StatusCode() int {
	switch e.Code() {
	case ErrorLogin.Code():
		return http.StatusForbidden
	case ErrorAuthorize.Code():
		return http.StatusUnauthorized
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
	ErrorAuthorize     = NewError(401, "authentication failed")
	ErrorLogin         = NewError(403, "please log in again")
	ErrorServer        = NewError(1000, "service internal error")
	ErrorInvalidParams = NewError(1001, "error passing parameter")
	ErrorMobile        = NewError(1002, "wrong phone number")
	ErrorDataQuery     = NewError(1003, "data query error")
	ErrorUserNotFound  = NewError(1004, "the user does not exist")
	ErrorPassword      = NewError(1005, "the password is incorrect")
)

func IsRecordNotFound(err error) bool {
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true
		}
	}
	return false
}
