package helpers

import (
	"docker-air-echo/constants"
	"docker-air-echo/structs"

	"github.com/gookit/validate"
)

func FormError(err validate.Errors) structs.Response {
	errors := make(map[string][]string)

	for key, errorVal := range err {
		var errorsTemp []string
		for _, value := range errorVal {
			errorsTemp = append(errorsTemp, value)
		}
		errors[key] = errorsTemp
	}

	return structs.Response{
		Message: constants.SM_UNPROCESSABLE_ENTITY,
		Data: map[string]interface{}{
			"errors": errors,
		},
	}
}
