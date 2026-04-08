package models

type Finance struct {
	FinanceID  uint `gorm:"primaryKey"`
	TotalMoney string

	Incomes   []Income
	Withdraws []Withdraw
	Deposits  []Deposit
}
