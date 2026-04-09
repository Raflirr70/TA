package database

import (
	"log"
	"tipes/internal/config"
	model "tipes/internal/model"
)

func Migrate() {
	err := config.DB.AutoMigrate(
		&model.Role{},
		&model.Finance{},
		&model.Item{},
		&model.Report{},

		&model.FinanceHistory{},
		&model.Deposit{},
		&model.Withdraw{},
		&model.User{},

		&model.Owner{},
		&model.Branch{},
		&model.Stock{},
		&model.Income{},
		&model.StockHistory{},
		&model.Attendance{},
		&model.ItemRequest{},

		&model.Restock{},
		&model.Sale{},

		&model.SaleDetail{},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Database migrated")
}
