package database

import (
	"tipes/internal/config"
	seed "tipes/internal/seed"
)

func Seed() {
	seed.RoleSeeder(config.DB)
}
