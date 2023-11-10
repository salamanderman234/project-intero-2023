package view

import (
	echo "github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
)

type classSubjectView struct {
	service domain.ClassSubjectService
}

func NewClassSubjectView(s domain.ClassSubjectService) domain.ClassSubjectView {
	return &classSubjectView {
		service: s,
	}
}

func (cs *classSubjectView) Create(c echo.Context) error {
	requestContext := c.Request().Context()
	form := domain.ClassSubjectCreateForm{}
	fun := func (createdForm domain.Form) (any, error) {
		return cs.service.CreateClassSubject(requestContext, createdForm.(domain.ClassSubjectCreateForm))
	}
	statusCode, resp := basicCreateView(c, fun, &form)
	return c.JSON(statusCode, resp)
}
func (cs *classSubjectView) Get(c echo.Context) error {
	return nil
}
func (cs *classSubjectView) Update(c echo.Context) error {
	return nil
}
func (cs *classSubjectView) Delete(c echo.Context) error {
	return nil
}