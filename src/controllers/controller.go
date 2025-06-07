package controllers

import (
	"fmt"
	"net/http"
	"docker-air-echo/database"
	"docker-air-echo/helpers"
	"docker-air-echo/models"
	"docker-air-echo/structs"
	"time"

	"github.com/labstack/echo/v4"
)

func Root(c echo.Context) error {
	// Test write speed.
	startTime := time.Now()
	db := database.DB()

	permissions := &[]models.Permission{}
	if err := db.Model(&models.Permission{}).Find(permissions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, structs.Response{Message: err.Error()})
	}

	var allowedPermissions []string
	for _, permission := range *permissions {
		allowedPermissions = append(allowedPermissions, permission.Code)
	}

	for i := 1; i <= 1000; i++ {
		password, _ := helpers.HashPassword(fmt.Sprintf("password*%d*", i))
		user := &models.User{
			Username:    fmt.Sprintf("user%d", i),
			Email:       fmt.Sprintf("user%d@email.com", i),
			Password:    password,
			FirstName:   fmt.Sprintf("first%d", i),
			LastName:    fmt.Sprintf("last%d", i),
			WrongPass:   i % 2,
			Permissions: allowedPermissions,
		}
		db.FirstOrCreate(user, models.User{Username: user.Username})
	}

	writeDuration := fmt.Sprintf("Write Time: %v\n", time.Since(startTime))

	// Test read speed.
	startTime = time.Now()
	var users []models.User
	db.Model(&models.User{}).Find(&users)

	readDuration := fmt.Sprintf("Read Time: %v\n", time.Since(startTime))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"write_duration": writeDuration,
		"read_duration":  readDuration,
		"users":          users,
	})
}
