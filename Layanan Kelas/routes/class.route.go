package route

import (
	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
)

func registerClassRoute(group *echo.Group, classView domain.ClassView) {
	group.POST("/", classView.Create)
	group.GET("/", classView.Read)
	group.PATCH("/", classView.Update)
	group.DELETE("/", classView.Delete)
}