package models

type Report struct {
	ReportID      uint `gorm:"primaryKey"`
	TotalIncome   int
	TotalMoney    int
	TotalWithdraw int
	TotalDeposit  int
	TotalProfit   int

	Restocks       []Restock
	Incomes        []Income
	StockHistories []StockHistory
	Attendances    []Attendance
}
