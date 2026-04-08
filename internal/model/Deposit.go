package models

import "time"

type Deposit struct {
	DepositID  uint `gorm:"primaryKey"`
	FinanceID  uint
	TotalMoney int
	Time       time.Time

	Finance Finance
}
