package models

type User struct {
	UserID      uint `gorm:"primaryKey"`
	RoleID      uint
	BranchID    uint
	Status      string
	FirstName   string
	LastName    string
	Email       string
	NoTelephone string
	Password    string

	Role   Role
	Branch Branch
}
