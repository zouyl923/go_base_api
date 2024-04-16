package errx

import "fmt"

type CodeError struct {
	Code    int
	Message string
	Data    interface{}
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.Code, e.Message)
}

func NewCodeError(code int) *CodeError {
	return &CodeError{Code: code, Message: GetCnMessage(code)}
}

func NewParamError(message string) *CodeError {
	return &CodeError{Code: ParamError, Message: message}
}
