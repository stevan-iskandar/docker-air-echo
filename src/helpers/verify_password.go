package helpers

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(hashedPassword, password string) bool {
	// Compare the stored hashed password with the entered password.
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
