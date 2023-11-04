package view

import (
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
	return nil
}
func (s *studentClassView) UnsignStudentFromAClass(c echo.Context) error {
	return nil
}
func (s *studentClassView) GetStudentClassList(c echo.Context) error {
	return nil
}