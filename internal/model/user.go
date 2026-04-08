package models

type User struct {
	UserID    uint `gorm:"primaryKey"`
	RoleID    uint
	BranchID  uint
	Status    string
	FirstName string
	LasttName string
	Username  string
	Password  string

	Role   Role
	Branch Branch
}
