package view

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	helper "github.com/salamanderman234/project-intro-2023/layanan-kelas/helpers"
)

type studentClassView struct {
	classService domain.StudentClassService
}

func NewStudentClassView(s domain.StudentClassService) domain.StudentClassView {
	return &studentClassView{
		classService: s,
	}
}

func (s *studentClassView) AssignStudentToAClass(c echo.Context) error {
	requestContext := c.Request().Context()
	form := domain.AssignStudentForm{}
	servFunc := func(createdForm domain.Form) (any, error) {
		data := createdForm.(*domain.AssignStudentForm)
		classID, studentID, err := s.classService.AssignStudent(requestContext, *data)
		result := map[string]uint{
			"student_id": studentID,
			"class_id":   classID,
		}
		return result, err
	}
	bindFunc := func() []error {
		errs := echo.FormFieldBinder(c).
			Uint("class_id", &form.ClassID).
			Uint("student_id", &form.StudentID).
			Uint("year", &form.Year).
			BindErrors()
		return errs
	}
	statusCode, resp := basicCreateView(c, servFunc, &form, bindFunc)
	return c.JSON(statusCode, resp)
}
func (s *studentClassView) UnsignStudentFromAClass(c echo.Context) error {
	requestContext := c.Request().Context()
	respStatusCode := http.StatusOK
	resp := domain.BasicResponse{
		Message: "success",
		Data:    nil,
		Errors:  nil,
	}
	form := domain.AssignStudentForm{}
	if errs := echo.FormFieldBinder(c).
		Uint("class_id", &form.ClassID).
		Uint("student_id", &form.StudentID).
		Uint("year", &form.Year).
		BindErrors(); len(errs) != 0 {
		resp.Message = "request error"
		resp.Errors = helper.GenerateBindingErrorDetail(errs)
		return c.JSON(http.StatusBadRequest, resp)
	}

	_, err := s.classService.UnasssignStudent(requestContext, form)
	if err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return c.JSON(respStatusCode, resp)
	}
	resp.Data = map[string]any{
		"deleted_data": map[string]uint{
			"student_id": form.StudentID,
			"class_id":   form.ClassID,
			"year":       form.Year,
		},
	}
	return c.JSON(respStatusCode, resp)
}
func (s *studentClassView) GetStudentClassList(c echo.Context) error {
	requestContext := c.Request().Context()
	respStatusCode := http.StatusOK
	resp := domain.BasicResponse{
		Message: "ok",
		Data:    nil,
		Errors:  nil,
	}

	var studentID uint
	var classID uint
	var year uint
	if errs := echo.QueryParamsBinder(c).
		Uint("class_id", &classID).
		Uint("student_id", &studentID).
		Uint("year", &year).
		BindErrors(); len(errs) > 0 {

		resp.Message = "request error"
		resp.Errors = helper.GenerateBindingErrorDetail(errs)
		return c.JSON(http.StatusBadRequest, resp)
	}

	results, err := s.classService.GetStudentClassList(requestContext, studentID, classID, year)
	if err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return c.JSON(respStatusCode, resp)
	}
	resp.Data = map[string]any{
		"results": results,
		"length":  len(results),
	}
	return c.JSON(respStatusCode, resp)
}
