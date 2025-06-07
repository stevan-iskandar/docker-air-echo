package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Generate a salt and hash the password with it.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
