package route

import (
	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	middleware "github.com/salamanderman234/project-intro-2023/layanan-kelas/middlewares"
)

func registerClassRoute(group *echo.Group, classView domain.ClassView) {
	group.POST("/", classView.Create, middleware.OnlyOperator)
	group.GET("/", classView.Read)
	group.PUT("/", classView.Update, middleware.OnlyOperator)
	group.DELETE("/", classView.Delete, middleware.OnlyOperator)
}
