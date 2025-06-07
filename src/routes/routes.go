package routes

import (
	"docker-air-echo/controllers"
	"docker-air-echo/middlewares"
	"docker-air-echo/structs"
	"docker-air-echo/validations"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	api := e.Group("/api")

	api.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, structs.Response{
			Message: "Welcome to Podman Echo API",
			Data:    nil,
		})
	})

	// api.GET("/", controllers.Root)
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login, validations.LoginValidation)

	api.Use(middlewares.Auth)

	RouteUser(api)
}
