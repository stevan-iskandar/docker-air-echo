package seeders

import (
	"docker-air-echo/constants"
	"docker-air-echo/database"
	"docker-air-echo/models"
)

func SeedPermission() error {
	db := database.DB()

	permissions := []models.Permission{
		{Code: constants.PER_PERMISSION_CREATE},
		{Code: constants.PER_PERMISSION_DELETE},
		{Code: constants.PER_PERMISSION_UPDATE},
		{Code: constants.PER_PERMISSION_VIEW},

		{Code: constants.PER_USER_CREATE},
		{Code: constants.PER_USER_DELETE},
		{Code: constants.PER_USER_UPDATE},
		{Code: constants.PER_USER_VIEW},
	}

	for _, permission := range permissions {
		if err := db.FirstOrCreate(&models.Permission{}, &models.Permission{Code: permission.Code}).Error; err != nil {
			return err
		}
	}

	return nil
}
