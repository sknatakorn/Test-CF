package infrastructure

import (
	"fmt"
	"net/http"
)

type Error struct {
	HttpStatusCode int
	Message        string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d:%s: syntax error", e.HttpStatusCode, e.Message)
}

func NewError(httpStatusCode int, message string) error {
	return &Error{
		HttpStatusCode: httpStatusCode,
		Message:        message,
	}
}
func GetCodeError(err error) int {
	switch e := err.(type) {
	case *Error:
		return e.HttpStatusCode
	default:
		return http.StatusInternalServerError
	}

}
