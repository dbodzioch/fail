package fail

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Fail interface {
	GetId() *string
	GetTimestamp() *time.Time
	GetFailInfo() FailInfo
	GetCause() error
	GetParams() []any
	Error() string
}

type DefaultFail struct {
	id        string
	timestamp time.Time
	info      FailInfo
	cause     error
	params    []any
}

func NewFail(info FailInfo) *DefaultFail {
	return &DefaultFail{
		id:        uuid.NewString(),
		timestamp: time.Now(),
		info:      info,
	}
}

func (f *DefaultFail) WithCause(cause error) *DefaultFail {
	f.cause = cause
	return f
}

func (f *DefaultFail) WithParams(params ...any) *DefaultFail {
	f.params = append(f.params, params...)
	return f
}

func (f *DefaultFail) Error() string {
	if cause, ok := f.cause.(Fail); ok {
		return fmt.Sprint(f.info.RawMessage(), f.params) + "\n" + cause.Error()
	}
	return f.Error()
}

func (f *DefaultFail) GetFailInfo() FailInfo {
	return f.info
}

func (f *DefaultFail) GetId() *string {
	return &f.id
}

func (f *DefaultFail) GetTimestamp() *time.Time {
	return &f.timestamp
}

func (f *DefaultFail) GetCause() error {
	return f.cause
}

func (f *DefaultFail) GetParams() []any {
	return f.params
}

func (f *DefaultFail) StringParams() []string {
	stringParams := make([]string, len(f.params))

	for _, param := range f.params {
		switch typedParam := param.(type) {
		case string:
			stringParams = append(stringParams, typedParam)
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			stringParams = append(stringParams, fmt.Sprintf("%d", typedParam))
		case float32, float64, complex64, complex128:
			stringParams = append(stringParams, fmt.Sprintf("%.2f", typedParam))
		case bool:
			stringParams = append(stringParams, fmt.Sprintf("%t", typedParam))
		}
	}

	return stringParams
}
