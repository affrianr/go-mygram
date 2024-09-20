package errorcontext

import (
	"fmt"
	"strings"
)

type AppError struct {
	BaseError    error
	ResponseCode int
	Description  string
}

type ErrorContext interface {
	UpdateErrorContext(err error, responseCode int, description string)
	GetDescription() string
	GetResponseCode() int
	GetBaseError() error
}

func NewAppError(err error, responseCode int, description string) ErrorContext {
	return &AppError{
		BaseError:    err,
		ResponseCode: responseCode,
		Description:  strings.TrimSpace(description),
	}
}

func (e *AppError) UpdateErrorContext(err error, responseCode int, description string) {
	if err != nil {
		e.BaseError = fmt.Errorf("%w: %v", e.BaseError, err)
	}
	if responseCode != 0 {
		e.ResponseCode = responseCode
	}
	if description != "" {
		e.Description = strings.TrimSpace(description)
	}
}

func (e *AppError) GetDescription() string {
	return e.Description
}

func (e *AppError) GetResponseCode() int {
	return e.ResponseCode
}

func (e *AppError) GetBaseError() error {
	return e.BaseError
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %d: %s", e.ResponseCode, e.Description)
}
