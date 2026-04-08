package models

type FinanceHistory struct {
	FinanceHistoryID uint `gorm:"primaryKey"`
	FinanceID        uint
	TotalMoney       int

	Finance Finance
}
