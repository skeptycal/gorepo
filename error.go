package gorepo

import (
	"fmt"
)

func NewShellError(msg string, err error) ShellError {
	return ShellError{
		msg,
		fmt.Errorf("http request error: %w", err),
	}
}

type ShellError struct {
	msg string
	err error
}

func (e *ShellError) Error() string {
	return fmt.Sprintf("%T%s: %v%T", redText, e.msg, e.err, reset)
}

func (e *ShellError) Unwrap() error {
	return e.err
}

type RequestError struct {
	statusCode int
	err        error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.statusCode, r.err)
}

func (r *RequestError) Unwrap() error {
	return r.err
}

func NewRequestError(statusCode int, err error) *RequestError {
	return &RequestError{
		statusCode: statusCode,
		err:        fmt.Errorf("http request error: %w", err),
	}
}
