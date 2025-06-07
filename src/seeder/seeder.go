package main

import (
	_ "docker-air-echo/autoload"
	"docker-air-echo/seeder/seeders"
)

func main() {
	seeders.SeedPermission()
}
