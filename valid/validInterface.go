package valid

import (
	"fmt"
	"reflect"
)

// ValidateError 验证错误提示
type ValidateError struct {
	Field   string // 验证不通过的属性
	Valid   string // 不通过的条件
	ErrCode int    // 不通过的错误码
}

func (ve ValidateError) Error() string {
	return fmt.Sprintf("Unverified <%s>:%s", ve.Field, ve.Valid)
}

// ValidateInterface 校验封装
type ValidateInterface interface {
	validate() bool
}

type validateModel struct {
	required bool
	fieldT   reflect.StructField
	fieldV   interface{}
	error    int64
}
