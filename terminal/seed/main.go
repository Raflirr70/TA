package main

import (
	"tipes/internal/config"
	"tipes/internal/database"
)

func main() {
	config.Connection()
	database.Seed()
}
