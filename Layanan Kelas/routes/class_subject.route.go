package route

import (
	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	middleware "github.com/salamanderman234/project-intro-2023/layanan-kelas/middlewares"
)

func registerClassSubjectRoute(group *echo.Group, classView domain.ClassSubjectView) {
	group.POST("/", classView.Create, middleware.OnlyOperator)
	group.GET("/", classView.Get)
	group.PATCH("/", classView.Update, middleware.OnlyOperator)
	group.DELETE("/", classView.Delete, middleware.OnlyOperator)
}
