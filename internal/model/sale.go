package models

type Sale struct {
	SaleID     uint `gorm:"primaryKey"`
	UserID     uint
	IncomeID   uint
	TotalItem  int
	TotalPrice int

	User        User
	Income      Income
	SaleDetails []SaleDetail
}
