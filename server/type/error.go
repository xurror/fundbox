package _type

import (
	"errors"
)

type AppError struct {
	Err     error  `json:"-"`       // low-level runtime error
	Code    int    `json:"code"`    // http response status code
	Message string `json:"message"` // user-level status message
}

func NewAppError(statusCode int, err string) *AppError {
	return &AppError{
		Err:     errors.New(err),
		Code:    statusCode,
		Message: err,
	}
}

func (e AppError) Error() string {
	return e.Message
}
