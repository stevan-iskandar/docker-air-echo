package controllers

import (
	"net/http"
	"time"

	"docker-air-echo/database"
	"docker-air-echo/helpers"
	"docker-air-echo/middlewares"
	"docker-air-echo/models"
	"docker-air-echo/structs"
	"docker-air-echo/validations"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	return c.JSON(http.StatusOK, structs.Response{
		Message: "User created",
	})
}

func getUserByUsername(username string) (*models.User, error) {
	db := database.DB()
	user := &models.User{}

	// Find the user document in the collection
	if err := db.Model(user).Where(&models.User{Username: username}).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func Login(c echo.Context) error {
	credentials := c.Get(validations.LOGIN_VALIDATION).(*validations.Credentials)

	user, err := getUserByUsername(credentials.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, structs.Response{
			Message: err.Error(),
		})
	}

	if validated := helpers.VerifyPassword(user.Password, credentials.Password); !validated {
		return echo.NewHTTPError(http.StatusUnauthorized, structs.Response{
			Message: "Wrong password",
		})
	}

	expirationTime := time.Now().Add(time.Hour * 24 * 7)
	token, err := helpers.GenerateToken(user, expirationTime)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, structs.Response{
		Message: "Successfully logged in",
		Data: map[string]interface{}{
			"token":      token,
			"user":       user,
			"expires_at": expirationTime,
		},
	})
}

func User(c echo.Context) *structs.JWTClaims {
	return c.Get(middlewares.USER).(*structs.JWTClaims)
}
