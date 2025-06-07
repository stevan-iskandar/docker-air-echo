package structs

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	jwt.RegisteredClaims
}
