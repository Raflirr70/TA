package models

type Role struct {
	RoleID uint `gorm:"primaryKey"`
	Name   string

	Users []User
}
