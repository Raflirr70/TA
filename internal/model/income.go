package models

type Income struct {
	IncomeID    uint `gorm:"primaryKey"`
	FinanceID   uint
	BranchID    uint
	ReportID    uint
	TotalIncome int

	Finance Finance
	Branch  Branch
	Report  Report
}
