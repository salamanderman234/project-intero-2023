package route

import (
	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
)

func RegisterAllRoutes(router *echo.Echo, viewRegistry domain.ViewRegistry) {
	// groups
	classGroup := router.Group("/class")
	// register
	registerClassRoute(classGroup, viewRegistry.ClassVw)
}