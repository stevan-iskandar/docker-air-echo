package helpers

import (
	"docker-air-echo/constants"
	"os"

	"github.com/speps/go-hashids"
)

func Encrypt(data int) (string, error) {
	hd := hashids.NewData()
	hd.Salt = os.Getenv(constants.ENV_HASHIDS_SALT)
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)

	return h.Encode([]int{45, data})
}

func Decrypt(encryptedData string) ([]int, error) {
	hd := hashids.NewData()
	hd.Salt = os.Getenv(constants.ENV_HASHIDS_SALT)
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)

	return h.DecodeWithError(encryptedData)
}
