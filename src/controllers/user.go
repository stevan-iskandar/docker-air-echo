package controllers

import (
	"net/http"
	"strconv"

	"docker-air-echo/constants"
	"docker-air-echo/database"
	"docker-air-echo/helpers"
	"docker-air-echo/models"
	"docker-air-echo/structs"
	"docker-air-echo/validations"

	"github.com/labstack/echo/v4"
)

func UserList(c echo.Context) error {
	var users []*models.User
	usernames := c.QueryParams().Get("username")
	email := c.QueryParam("email")
	data, err := helpers.Pagination(c, users, map[string]interface{}{
		"username": usernames,
		"email":    email,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, structs.Response{
		Message: constants.SM_RETRIEVE_SUCCESS,
		Data:    data,
	})
}

func UserStore(c echo.Context) error {
	db := database.DB()
	userForm := c.Get(validations.STORE_VALIDATION).(*validations.UserStoreForm)

	user := &models.User{
		Username:  userForm.Username,
		Email:     userForm.Email,
		FirstName: userForm.FirstName,
		LastName:  userForm.LastName,
	}

	if err := db.Create(user).Error; err != nil {
		return c.JSON(http.StatusUnprocessableEntity, structs.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, structs.Response{
		Message: constants.SM_STORE_SUCCESS,
		Data:    user,
	})
}

func UserUpdate(c echo.Context) error {
	strVal, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, structs.Response{Message: err.Error()})
	}

	encryptedData, err := helpers.Encrypt(strVal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{Message: err.Error()})
	}
	decryptedData, err := helpers.Decrypt(encryptedData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, structs.Response{
		Data: map[string]interface{}{
			"string":         strVal,
			"encrypted_data": encryptedData,
			"decrypted_data": decryptedData,
		},
	})
}
