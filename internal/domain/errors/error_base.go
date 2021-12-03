package errors

import "fmt"

type ErrorBase struct {
	Message string
	error
}

func NewError(err error) *ErrorBase {
	return &ErrorBase{
		Message: err.Error(),
	}
}

func (e ErrorBase) String() string {
	return fmt.Sprintf("domain: %s", e.Message)
}

func (e ErrorBase) Error() string {
	return fmt.Sprintf("domain: %s", e.Message)
}
