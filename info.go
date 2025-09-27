package fail

type FailInfo interface {
	Code() *FailCode
	Kind() *FailType
	RawMessage() *string
}

type FailType string
type FailCode string

const (
	BadRequest    FailType = "BAD_REQUEST"
	NotFound      FailType = "NOT_FOUND"
	Unauthorized  FailType = "UNAUTHORIZED"
	Forbidden     FailType = "FORBIDDEN"
	Timeout       FailType = "TIMEOUT"
	InternalError FailType = "INTERNAL_ERROR"
	Unavailable   FailType = "UNAVAILABLE"
	Business      FailType = "BUSINESS"
)

type DefaultFailInfo struct {
	code    FailCode
	kind    FailType
	message string
}

func NewFailInfo(code FailCode, kind FailType, message string) FailInfo {
	return &DefaultFailInfo{
		code:    code,
		kind:    kind,
		message: message,
	}
}

func (failCode *DefaultFailInfo) Code() *FailCode {
	return &failCode.code
}

func (failCode *DefaultFailInfo) Kind() *FailType {
	return &failCode.kind
}

func (failCode *DefaultFailInfo) RawMessage() *string {
	return &failCode.message
}
