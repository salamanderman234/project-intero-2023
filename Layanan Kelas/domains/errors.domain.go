package domain

import "errors"

var (
	ErrGormInGeneral = errors.New("gorm error")
	ErrConversionType = errors.New("conversion type error")
	ErrFormValidation = errors.New("validation error")
	ErrResourceNotFound = errors.New("resource not found error")
	ErrDuplicateEnties = errors.New("duplicate entries error")
	ErrForeignKeyViolated = errors.New("foreign key violated error")
)