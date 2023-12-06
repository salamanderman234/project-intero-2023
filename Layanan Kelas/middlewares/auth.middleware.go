package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/project-intro-2023/layanan-kelas/config"
	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
)

func OnlyOperator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		resp := domain.BasicResponse{
			Message: "success",
			Data:    nil,
			Errors:  nil,
		}
		if token == "" {
			resp.Message = "unauthorize error"
			return c.JSON(http.StatusUnauthorized, resp)
		}
		verifyAuthUrl := config.AuthServiceHost() + "/api/auth/me"
		client := http.Client{}
		req, _ := http.NewRequest("POST", verifyAuthUrl, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		res, _ := client.Do(req)
		if res.StatusCode != 200 {
			resp.Message = "forbidden error"
			return c.JSON(http.StatusForbidden, resp)
		}
		return next(c)
	}
}
