package models

type Stock struct {
	StockID  uint `gorm:"primaryKey"`
	ItemID   uint
	BranchID uint
	Stock    int

	Item           Item
	Branch         Branch
	StockHistories []StockHistory
}
