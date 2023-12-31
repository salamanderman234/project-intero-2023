package helper

import (
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
)

func ValidateForm(obj any) (bool, govalidator.Errors) {
	_, errs := govalidator.ValidateStruct(obj)
	if errs != nil {
		return false, errs.(govalidator.Errors)
	}
	return true, nil
}	

func GenerateErrorDetail(errs govalidator.Errors) []domain.ErrorDetail {
	var errsResult []domain.ErrorDetail
	for _, err := range errs {
		errValidator := err.(govalidator.Error)
		errDetail := domain.ErrorDetail{
			Field: strings.ToLower(errValidator.Name),
			Type: errValidator.Validator,
			Message: err.Error(),
		}
		errsResult = append(errsResult, errDetail)
	}
	return errsResult
}

func GenerateBindingErrorDetail(errs []error) []domain.ErrorDetail {
	var errsResult []domain.ErrorDetail
	for _, errRaw := range errs {
		err := errRaw.(*echo.BindingError)
		errDetail := domain.ErrorDetail{
			Field: strings.ToLower(err.Field),
			Type: "binding",
			Message: err.Message.(string),
		}
		errsResult = append(errsResult, errDetail)
	}
	return errsResult
}