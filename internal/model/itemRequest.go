package models

type ItemRequest struct {
	ItemRequestID uint `gorm:"primaryKey"`
	UserID        uint
	ItemID        uint
	Status        string
	total         string

	User     User
	Item     Item
	Restocks []Restock
}
