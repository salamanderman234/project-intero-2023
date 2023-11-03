package view

import (
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	helper "github.com/salamanderman234/project-intro-2023/layanan-kelas/helpers"
)

func handleErrorResponse(err error) (int, string, []domain.ErrorDetail) {
	if errs, ok := err.(govalidator.Errors); ok {
		return http.StatusBadRequest, domain.ErrFormValidation.Error(), helper.GenerateErrorDetail(errs)
	}
	if errors.Is(err, domain.ErrConversionType) || errors.Is(err, domain.ErrGormInGeneral) {
		return http.StatusInternalServerError, "internal server error", nil
	}
	if errors.Is(err, domain.ErrDuplicateEnties) {
		return http.StatusUnprocessableEntity, domain.ErrDuplicateEnties.Error(), nil
	}
	if errors.Is(err, domain.ErrResourceNotFound) {
		return http.StatusNotFound, domain.ErrResourceNotFound.Error(), nil
	}
	return 0, "", nil
}