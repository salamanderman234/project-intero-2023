package route

import (
	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
)

func registerClassSubjectRoute(group *echo.Group, classView domain.ClassSubjectView) {
	group.POST("/", classView.Create)
	group.GET("/", classView.Get)
	group.PATCH("/", classView.Update)
	group.DELETE("/", classView.Delete)
}