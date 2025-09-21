package fail

import (
	"fmt"

	"github.com/google/uuid"
)

type Fail struct {
	id       string
	failCode FailCode
	cause    *error
	params   []any
}

func NewFail(failCode FailCode, cause *error, params []any) *Fail {
	fail := Fail{
		id:       uuid.NewString(),
		failCode: failCode,
		cause:    cause,
		params:   params,
	}

	if len(failCode.getParams()) > 0 && len(params) == 0 {
		fail.params = failCode.getParams()
	}

	return &fail
}

func (f *Fail) Error() string {
	if f.cause != nil {
		return fmt.Errorf("%s: %w", f.failCode.String(), *f.cause).Error()
	}

	return f.failCode.String()
}
