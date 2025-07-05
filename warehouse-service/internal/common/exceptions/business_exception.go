package exceptions

import (
	"fmt"
	"time"
)

type BusinessException struct {
	Code     ErrorCode
	Extra    any
	DateTime string
}

func (e *BusinessException) Error() string {
	return fmt.Sprintf("{Code=%s, DataTime=%s, Extra=%v}", e.Code, e.DateTime, e.Extra)
}

func NewBusinessException(code ErrorCode, extra map[string]any) *BusinessException {
	return &BusinessException{
		Code:     code,
		Extra:    extra,
		DateTime: time.Now().UTC().Format(time.RFC3339),
	}
}
