package errors

import (
	"fmt"
	"net/http"
)

type ErrorCode int

const (
	ErrCodeValidation   ErrorCode = 40000 // 验证错误
	ErrCodeUnauthorized ErrorCode = 40100 // 未授权
	ErrCodeForbidden    ErrorCode = 40300 // 禁止访问
	ErrCodeNotFound     ErrorCode = 40400 // 资源不存在
	ErrCodeConflict     ErrorCode = 40900 // 资源冲突
	ErrCodeInternal     ErrorCode = 50000 // 内部错误
)

type AppError struct {
	Code       ErrorCode
	Message    string
	Err        error
	HttpStatus int
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// NewValidationError 创建验证错误
func NewValidationError(msg string) *AppError {
	return &AppError{
		Code:       ErrCodeValidation,
		Message:    msg,
		HttpStatus: http.StatusBadRequest,
	}
}

// NewNotFoundError 创建资源不存在错误
func NewNotFoundError(resource string) *AppError {
	return &AppError{
		Code:       ErrCodeNotFound,
		Message:    fmt.Sprintf("%s not found", resource),
		HttpStatus: http.StatusNotFound,
	}
}

// NewConflictError 创建资源冲突错误
func NewConflictError(msg string) *AppError {
	return &AppError{
		Code:       ErrCodeConflict,
		Message:    msg,
		HttpStatus: http.StatusConflict,
	}
}

// NewInternalError 创建内部错误
func NewInternalError(err error) *AppError {
	return &AppError{
		Code:       ErrCodeInternal,
		Message:    "Internal server error",
		Err:        err,
		HttpStatus: http.StatusInternalServerError,
	}
}

// NewUnauthorizedError 创建未授权错误
func NewUnauthorizedError(msg string) *AppError {
	return &AppError{
		Code:       ErrCodeUnauthorized,
		Message:    msg,
		HttpStatus: http.StatusUnauthorized,
	}
}

// NewForbiddenError 创建禁止访问错误
func NewForbiddenError(msg string) *AppError {
	return &AppError{
		Code:       ErrCodeForbidden,
		Message:    msg,
		HttpStatus: http.StatusForbidden,
	}
}
