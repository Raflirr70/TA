package models

type QR struct {
	QRID  uint `gorm:"primaryKey"`
	Key   string
	token string
}
