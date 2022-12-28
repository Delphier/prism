package httpx

import (
	"net/http"
)

// Error 定义一个包含状态码的自定义 Http 错误。参见：https://medium.com/rungo/error-handling-in-go-f0125de052f0
type Error interface {
	error
	Code() int
}

type errorObject struct {
	text string
	code int
}

func (e *errorObject) Error() string {
	return e.text
}

func (e *errorObject) Code() int {
	return e.code
}

// NewError 返回一个新的 Http 错误
func NewError(text string, code int) Error {
	return &errorObject{text, code}
}

// NewBadRequestError 返回一个 BadRequest 错误
func NewBadRequestError(text string) Error {
	return NewError(text, http.StatusBadRequest)
}

func NewUnauthorizedError(text string) Error {
	return NewError(text, http.StatusUnauthorized)
}

// 定义常用 HTTP 错误
var (
	ErrBadRequest   = NewBadRequestError("Invalid Request")
	ErrUnauthorized = NewUnauthorizedError(http.StatusText(http.StatusUnauthorized))
	ErrNotFound     = NewError(http.StatusText(http.StatusNotFound), http.StatusNotFound)
)
