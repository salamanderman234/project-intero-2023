package view

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
)

type studentClassView struct {
	classService domain.StudentClassService
}

func NewStudentClassView(s domain.StudentClassService) domain.StudentClassView{
	return &studentClassView{
		classService: s,
	}
}

func (s *studentClassView) AssignStudentToAClass(c echo.Context) error {
	requestContext := c.Request().Context()
	respStatusCode := http.StatusOK
	resp := domain.BasicResponse{
		Message: "siswa berhasil diassign ke kelas",
		Datas: nil,
		Errors: nil,
	}
	
	form := domain.AssignStudentForm{}
	if err := c.Bind(&form); err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(domain.ErrBind)
		return c.JSON(respStatusCode, resp)
	}

	classId, studentId, err := s.classService.AssignStudent(requestContext, form)
	if err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return c.JSON(respStatusCode, resp)
	}
	resp.Datas = map[string]uint{
		"student_id" : studentId,
		"class_id" : classId,
	}
	return c.JSON(respStatusCode, resp)
}
func (s *studentClassView) UnsignStudentFromAClass(c echo.Context) error {
	requestContext := c.Request().Context()
	respStatusCode := http.StatusOK
	resp := domain.BasicResponse{
		Message: "siswa berhasil diunassign ke kelas",
		Datas: nil,
		Errors: nil,
	}
	
	form := domain.AssignStudentForm{}
	if err := c.Bind(&form); err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(domain.ErrBind)
		return c.JSON(respStatusCode, resp)
	}

	ok, err := s.classService.UnasssignStudent(requestContext, form)
	if !ok {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return c.JSON(respStatusCode, resp)
	}
	resp.Datas = map[string]int{
		"student_id" : int(*form.SiswaID),
		"class_id" : int(*form.KelasID),
	}
	return c.JSON(respStatusCode, resp)
}
func (s *studentClassView) GetStudentClassList(c echo.Context) error {
	requestContext := c.Request().Context()
	respStatusCode := http.StatusOK
	resp := domain.BasicResponse{
		Message: "ok",
		Datas: nil,
		Errors: nil,
	}
	
	studentIdString := c.QueryParam("student_id")
	studentId, _ := strconv.Atoi(studentIdString)

	results, err := s.classService.GetStudentClassList(requestContext, uint(studentId))
	if err != nil {
		respStatusCode, resp.Message, resp.Errors = handleErrorResponse(err)
		return c.JSON(respStatusCode, resp)
	}
	resp.Datas = map[string]int{
		"datas" : studentId,
		"length" : len(results),
	}
	return c.JSON(respStatusCode, resp)
}