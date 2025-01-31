package main

import (
	"github.com/naufan17/go-gin-boilerplate/internal/configs"
)

func main() {
	db := configs.ConnectDB()

	configs.MigrateDB(db)
	configs.SeedAll(db)
}
