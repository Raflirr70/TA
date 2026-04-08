package models

type Owner struct {
	OwnerID uint `gorm:"primaryKey"`
	UserID  uint

	User User
}
