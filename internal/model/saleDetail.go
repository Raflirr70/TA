package models

type SaleDetail struct {
	SaleDetailID uint `gorm:"primaryKey"`
	SaleID       uint
	ItemID       uint
	TotalItem    int
	TotalPrice   int

	Sale Sale
	Item Item
}
