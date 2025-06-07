package helpers

import (
	"fmt"
	"os"
	"docker-air-echo/constants"
	"docker-air-echo/models"
	"docker-air-echo/structs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *models.User, expirationTime time.Time) (string, error) {
	claims := structs.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:     fmt.Sprintf("%d", user.ID),
			Issuer: user.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv(constants.ENV_JWT_KEY)))
}
