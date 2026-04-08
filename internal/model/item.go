package models

type Item struct {
	ItemID uint `gorm:"primaryKey"`
	Code   string
	Name   string
	Price  int

	Stocks       []Stock
	SaleDetails  []SaleDetail
	ItemRequests []ItemRequest
}
