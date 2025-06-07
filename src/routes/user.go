package routes

import (
	"docker-air-echo/constants"
	"docker-air-echo/controllers"
	"docker-air-echo/middlewares"
	"docker-air-echo/validations"

	"github.com/labstack/echo/v4"
)

func RouteUser(api *echo.Group) {
	api.GET(
		"/user",
		controllers.UserList,
		middlewares.Permission(constants.PER_USER_VIEW),
	)
	api.POST(
		"/user",
		controllers.UserStore,
		middlewares.Permission(constants.PER_USER_CREATE),
		validations.StoreValidation,
	)
	api.GET(
		"/user/id",
		controllers.UserUpdate,
	)
}
