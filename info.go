package fail

type FailInfo interface {
	Code() *Code
	Kind() *Type
	RawMessage() *string
}

type Type string
type Code string

const (
	BadRequest    Type = "BAD_REQUEST"
	NotFound      Type = "NOT_FOUND"
	Unauthorized  Type = "UNAUTHORIZED"
	Forbidden     Type = "FORBIDDEN"
	Timeout       Type = "TIMEOUT"
	InternalError Type = "INTERNAL_ERROR"
	Unavailable   Type = "UNAVAILABLE"
	Business      Type = "BUSINESS"
)

type DefaultFailInfo struct {
	code    Code
	kind    Type
	message string
}

func NewFailInfo(code Code, kind Type, message string) FailInfo {
	return &DefaultFailInfo{
		code:    code,
		kind:    kind,
		message: message,
	}
}

func (failCode *DefaultFailInfo) Code() *Code {
	return &failCode.code
}

func (failCode *DefaultFailInfo) Kind() *Type {
	return &failCode.kind
}

func (failCode *DefaultFailInfo) RawMessage() *string {
	return &failCode.message
}
