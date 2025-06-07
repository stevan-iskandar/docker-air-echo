package validations

import (
	"docker-air-echo/helpers"
	"docker-air-echo/structs"
	"net/http"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const LOGIN_VALIDATION = "LoginValidation"

func LoginValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := validate.FromRequest(c.Request())
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, structs.Response{
				Message: err.Error(),
			})
		}

		v := data.Create()
		v.StringRules(map[string]string{
			"username": "required",
			"password": "required",
		})

		if !v.Validate() {
			return c.JSON(http.StatusUnprocessableEntity, helpers.FormError(v.Errors))
		}

		credentials := &Credentials{}
		v.BindSafeData(credentials)
		c.Set(LOGIN_VALIDATION, credentials)

		return next(c)
	}
}
