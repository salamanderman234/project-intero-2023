package view

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	helper "github.com/salamanderman234/project-intro-2023/layanan-kelas/helpers"
)

type classSubjectView struct {
	service domain.ClassSubjectService
}

func NewClassSubjectView(s domain.ClassSubjectService) domain.ClassSubjectView {
	return &classSubjectView{
		service: s,
	}
}

func (cs *classSubjectView) Create(c echo.Context) error {
	requestContext := c.Request().Context()
	form := domain.ClassSubjectCreateForm{}
	fun := func(createdForm domain.Form) (any, error) {
		return cs.service.CreateClassSubject(requestContext, *createdForm.(*domain.ClassSubjectCreateForm))
	}
	bindFunc := func() []error {
		return echo.FormFieldBinder(c).
			Uint("class_id", &form.ClassID).
			Uint("subject_id", &form.SubjectID).
			Uint("teacher_id", &form.TeacherID).
			Uint("year", &form.Year).
			BindErrors()
	}
	statusCode, resp := basicCreateView(c, fun, &form, bindFunc)
	return c.JSON(statusCode, resp)
}
func (cs *classSubjectView) Get(c echo.Context) error {
	requestContext := c.Request().Context()
	respStatusCode := http.StatusOK
	resp := domain.BasicResponse{
		Message: "ok",
		Data:    nil,
		Errors:  nil,
	}

	var id uint
	var studentID uint
	var classID uint
	var teacherID uint
	var year uint
	var page uint
	var sortBy string
	var sort string
	if errs := echo.QueryParamsBinder(c).
		Uint("class_id", &classID).
		Uint("id", &id).
		Uint("student_id", &studentID).
		Uint("teacher_id", &teacherID).
		Uint("year", &year).
		Uint("page", &page).
		String("order_by", &sortBy).
		String("order", &sort).
		BindErrors(); len(errs) > 0 {

		resp.Message = "request error"
		resp.Errors = helper.GenerateBindingErrorDetail(errs)
		return c.JSON(http.StatusBadRequest, resp)
	}

	results, pagination, err := cs.service.GetClassSubject(requestContext, id, classID, studentID, teacherID, year, page, sortBy, sort)
	if err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return c.JSON(respStatusCode, resp)
	}
	if id != 0 {
		resp.Data = results[0]
	} else {
		resp.Data = map[string]any{
			"results":    results,
			"length":     len(results),
			"pagination": pagination,
		}
	}

	return c.JSON(respStatusCode, resp)
}
func (cs *classSubjectView) Update(c echo.Context) error {
	requestContext := c.Request().Context()
	form := domain.ClassSubjectUpdateForm{}
	servFunc := func(id uint, updateForm domain.Form) (int, any, error) {
		data := updateForm.(*domain.ClassSubjectUpdateForm)
		aff, result, err := cs.service.UpdateClassSubject(requestContext, id, *data)
		return int(aff), result, err
	}
	bindFunc := func() []error {
		return echo.FormFieldBinder(c).
			Uint("subject_id", &form.SubjectID).
			Uint("teacher_id", &form.TeacherID).
			Uint("year", &form.Year).
			BindErrors()
	}
	statusCode, resp := basicUpdateView(c, servFunc, &form, bindFunc)
	return c.JSON(statusCode, resp)
}
func (cs *classSubjectView) Delete(c echo.Context) error {
	requestContext := c.Request().Context()
	deleteFunc := func(id uint) error {
		err := cs.service.DeleteClassSubject(requestContext, id)
		return err
	}
	statusCode, resp := basicDeleteView(c, deleteFunc)
	return c.JSON(statusCode, resp)
}
