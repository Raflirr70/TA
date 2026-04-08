package models

type Restock struct {
	RestockID     uint `gorm:"primaryKey"`
	ItemRequestID uint
	UserID        uint
	ReportID      uint
	TotalIncome   int

	ItemRequest ItemRequest
	Report      Report
	User        User
}
