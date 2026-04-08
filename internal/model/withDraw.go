package models

import "time"

type Withdraw struct {
	WithdrawID uint `gorm:"primaryKey"`
	FinanceID  uint
	TotalMoney int
	Time       time.Time

	Finance Finance
}
