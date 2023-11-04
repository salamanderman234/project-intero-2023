package main

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/project-intro-2023/layanan-kelas/config"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	repository "github.com/salamanderman234/project-intro-2023/layanan-kelas/repositories"
	route "github.com/salamanderman234/project-intro-2023/layanan-kelas/routes"
	service "github.com/salamanderman234/project-intro-2023/layanan-kelas/services"
	view "github.com/salamanderman234/project-intro-2023/layanan-kelas/views"
)

func init() {
	config.SetConfig("./.env")
}

func main() {
	connection, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	router := echo.New()
	// declare repo
	classRepository := repository.NewClassRepository(connection)
	// declare service
	repoRegistry := domain.ServiceRegistry{}
	classService := service.NewClassService(classRepository, repoRegistry)
	// declare view
	classView := view.NewClassView(classService)
	viewRegistry := domain.ViewRegistry{
		ClassVw: classView,
	}
	// register route
	route.RegisterAllRoutes(router, viewRegistry)
	router.Logger.Fatal(router.Start(":1323"))
}