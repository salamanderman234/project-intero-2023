package route

import (
	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	middleware "github.com/salamanderman234/project-intro-2023/layanan-kelas/middlewares"
)

func registerStudentClassRoute(group *echo.Group, studentClassView domain.StudentClassView) {
	group.POST("/assign/", studentClassView.AssignStudentToAClass, middleware.OnlyOperator)
	group.POST("/unassign/", studentClassView.UnsignStudentFromAClass, middleware.OnlyOperator)
	group.GET("/", studentClassView.GetStudentClassList)
}
