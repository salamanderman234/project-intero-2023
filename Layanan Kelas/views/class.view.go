package view

import (
	echo "github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
)

type classView struct {
	classService domain.ClassService
}

func NewClassView(c domain.ClassService) domain.ClassView {
	return &classView{
		classService: c,
	}
}
func (cv *classView) Create(c echo.Context) error {
	requestContext := c.Request().Context()
	form := domain.ClassCreateForm{}
	servFunc := func (createdForm domain.Form) (any, error) {
		data := createdForm.(*domain.ClassCreateForm)
		return cv.classService.CreateClass(requestContext, *data)
	}
	bindFunc := func () ([]error) {
		return echo.FormFieldBinder(c).
			Uint("focus_id", &form.FocusID).
			Uint("grade_id", &form.GradeID).
			String("group", &form.Group).
			BindErrors()
	}
	statusCode, resp := basicCreateView(c, servFunc, &form, bindFunc)
	return c.JSON(statusCode, resp)
}
func (cv *classView) Read(c echo.Context) error {
	requestContext := c.Request().Context()
	searchFunc := func(q string, page uint, orderBy string, order string) (any, domain.Pagination, error){
		return cv.classService.GetClassList(requestContext, q, page, orderBy, order)
	}
	findFunc := func(id uint) (any, error) {
		return cv.classService.GetClassInfo(requestContext, uint(id))
	}
	statusCode, resp := basicSearchView(c, searchFunc, findFunc)
	return c.JSON(statusCode, resp)
}

func (cv *classView) Update(c echo.Context) error {
	requestContext := c.Request().Context()
	form := domain.ClassUpdateForm{}
	servFunc := func (id uint, updateForm domain.Form) (int, any, error) {
		data := updateForm.(*domain.ClassUpdateForm)
		return cv.classService.UpdateClassInfo(requestContext, id, *data)
	}
	bindFunc := func () ([]error) {
		return echo.FormFieldBinder(c).
			Uint("focus_id", &form.FocusID).
			Uint("grade_id", &form.GradeID).
			String("group", &form.Group).
			BindErrors()
	}
	statusCode, resp := basicUpdateView(c, servFunc, &form, bindFunc)
	return c.JSON(statusCode, resp)
}
func (cv *classView) Delete(c echo.Context) error {
	requestContext := c.Request().Context()
	deleteFunc := func(id uint) (error){
		_, err := cv.classService.DeleteClass(requestContext, id)
		return err
	}
	statusCode, resp := basicDeleteView(c, deleteFunc)
	return c.JSON(statusCode, resp)
}