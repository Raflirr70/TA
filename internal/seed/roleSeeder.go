package seeders

import (
	"log"

	model "tipes/internal/model"

	"gorm.io/gorm"
)

func RoleSeeder(db *gorm.DB) {
	roles := []model.Role{
		{Name: "Owner"},
		{Name: "Manager"},
		{Name: "Kasir"},
		{Name: "Admin Gudang"},
	}

	if err := db.Create(&roles).Error; err != nil {
		log.Fatal("Failed to seed roles:", err)
	}

	log.Println("Role seeder executed successfully")
}
