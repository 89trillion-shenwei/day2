package internal

import "net/http"

// GlobalError 全局异常结构体
type GlobalError struct {
	Status  int    `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//获取err的内容
func (err GlobalError) Error() string {
	return err.Message
}

const (
	IsEmpty           = 1001 //表达式不能为空
	IllegalExpression = 1002 //表达式非法
)

//IsEmptyError  表达式不能为空
func IsEmptyError(message string) GlobalError {
	return GlobalError{
		Status:  http.StatusForbidden,
		Code:    IsEmpty,
		Message: message,
	}
}

//IllegalExpressionError     表达式非法
func IllegalExpressionError(message string) GlobalError {
	return GlobalError{
		Status:  http.StatusForbidden,
		Code:    IllegalExpression,
		Message: message,
	}
}
