package repository

import (
	models "tipes/internal/model"

	"gorm.io/gorm"
)

type AuthRepository interface {
	FindUser(email, noTelephone string) (models.User, error)
	Create(user models.User) error
	DB() *gorm.DB
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) FindUser(email, noTelephone string) (models.User, error) {
	var user models.User
	err := r.db.
		Where("email = ? OR no_telephone = ?", email, noTelephone).
		First(&user).Error

	return user, err
}

func (r *authRepository) Create(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *authRepository) DB() *gorm.DB {
	return r.db
}
