package models

type Owner struct {
	ownerID uint `gorm:"primaryKey"`
	UserID  uint

	User User
}
