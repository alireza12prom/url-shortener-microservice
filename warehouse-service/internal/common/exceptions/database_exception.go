package exceptions

import (
	"fmt"
	"time"
)

type DatabaseException struct {
	Code     ErrorCode
	Extra    any
	DateTime string
}

func (e *DatabaseException) Error() string {
	return fmt.Sprintf("{Code=%s, DataTime=%s, Extra=%v}", e.Code, e.DateTime, e.Extra)
}

func NewDatabaseException(err error) *DatabaseException {
	return &DatabaseException{
		Code:     DatabaseFailer,
		Extra:    map[string]any{"reason": err.Error()},
		DateTime: time.Now().UTC().Format(time.RFC3339),
	}
}
