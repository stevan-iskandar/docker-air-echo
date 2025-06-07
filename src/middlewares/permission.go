package middlewares

import (
	"fmt"
	"net/http"
	"docker-air-echo/database"
	"docker-air-echo/helpers"
	"docker-air-echo/models"
	"docker-air-echo/structs"

	"github.com/labstack/echo/v4"
)

func Permission(permission string) echo.MiddlewareFunc {
	db := database.DB()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userAuth := c.Get(USER).(*structs.JWTClaims)

			user := &models.User{}
			if err := db.Find(user, fmt.Sprintf("%s = ?", models.USER_ID), userAuth.ID).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, structs.Response{Message: err.Error()})
			}

			if !helpers.StringExistsInArray(user.Permissions, permission) {
				return c.JSON(http.StatusUnauthorized, structs.Response{
					Message: "Unauthorized",
				})
			}
			return next(c)
		}
	}
}
