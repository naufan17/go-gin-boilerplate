package main

import (
	"github.com/naufan17/go-gin-boilerplate/config"
)

func main() {
	db := config.ConnectDB()

	config.SeedAll(db)
}
