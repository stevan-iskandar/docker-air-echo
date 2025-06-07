package autoload

import (
	"github.com/gookit/validate"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
	})
}
