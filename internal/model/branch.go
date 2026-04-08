package models

type Branch struct {
	BranchID uint `gorm:"primaryKey"`
	Name     string
	Location string

	Users   []User
	Stocks  []Stock
	Incomes []Income
}
