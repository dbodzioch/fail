package fail

import (
	"fmt"
)

type FailCode interface {
	Code() *Code
	Kind() *Kind
	Message() *string
	getParams() []any
	String() string
}

type Kind string
type Code string

const (
	BadRequest    Kind = "BAD_REQUEST"
	NotFound      Kind = "NOT_FOUND"
	Unauthorized  Kind = "UNAUTHORIZED"
	Forbidden     Kind = "FORBIDDEN"
	Timeout       Kind = "TIMEOUT"
	InternalError Kind = "INTERNAL_ERROR"
	Unavailable   Kind = "UNAVAILABLE"
	Business      Kind = "BUSINESS"
)

type DefaultFailCode struct {
	code    Code
	kind    Kind
	message string
	params  []any
}

func NewFailCode(code Code, kind Kind, message string) FailCode {
	return &DefaultFailCode{
		code:    code,
		kind:    kind,
		message: message,
	}
}

func (failCode *DefaultFailCode) With(params ...any) FailCode {
	failCode.params = params
	return failCode
}

func (failCode *DefaultFailCode) Code() *Code {
	return &failCode.code
}

func (failCode *DefaultFailCode) Kind() *Kind {
	return &failCode.kind
}

func (failCode *DefaultFailCode) Message() *string {
	return &failCode.message
}

func (failCode *DefaultFailCode) String() string {
	return fmt.Sprintf("[%s] %s: %s", *failCode.Code(), *failCode.Kind(), failCode.formatMessage())
}

func (failCode *DefaultFailCode) formatMessage() string {
	return fmt.Sprintf(failCode.message, failCode.params...)
}

func (failCode *DefaultFailCode) getParams() []any {
	return failCode.params
}
