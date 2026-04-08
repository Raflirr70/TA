package models

type StockHistory struct {
	StockHistoryID uint `gorm:"primaryKey"`
	StockID        uint
	ReportID       uint
	StockItem      int

	Stock  Stock
	Report Report
}
