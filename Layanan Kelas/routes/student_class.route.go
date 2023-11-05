package route

import (
	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
)

func registerStudentClassRoute(group *echo.Group, studentClassView domain.StudentClassView) {
	group.POST("/assign_student/", studentClassView.AssignStudentToAClass)
	group.POST("/unassign_student/", studentClassView.UnsignStudentFromAClass)
	group.GET("/student_class_list/", studentClassView.GetStudentClassList)
}