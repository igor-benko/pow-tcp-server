package domain

import "errors"

type StatusCode int

const (
	CodeNotFound   = 404
	CodeBadRequest = 400
	CodeConflict   = 409
	CodeUnknown    = 500
)

type DomainError struct {
	Err        error
	Code       string
	StatusCode StatusCode
}

func New(text string, code string, statusCode StatusCode) error {
	return NewFromError(errors.New(text), code, statusCode)
}

func NewFromError(err error, code string, statusCode StatusCode) error {
	return &DomainError{
		Err:        err,
		Code:       code,
		StatusCode: statusCode,
	}
}

func (e *DomainError) Error() string {
	return e.Err.Error()
}

var ErrTariffNotFound = New("tariff not found", "tariff-svc.tariff_not_found", CodeNotFound)
