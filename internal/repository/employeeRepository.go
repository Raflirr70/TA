package repository

import (
	models "tipes/internal/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	GetByBranch(branchID uint) ([]models.User, error)
	Create(employee models.User) error
	Update(employee models.User) error
	ChangeStatus(employeeID uint, status string) error
	FindByID(employeeID uint) (models.User, error)
	DB() *gorm.DB
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) GetByBranch(branchID uint) ([]models.User, error) {
	var employees []models.User
	err := r.db.Where("branch_id = ?", branchID).Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) Create(employee models.User) error {
	return r.db.Create(&employee).Error
}

func (r *employeeRepository) Update(employee models.User) error {
	return r.db.Save(&employee).Error
}

func (r *employeeRepository) ChangeStatus(employeeID uint, status string) error {
	return r.db.Model(&models.User{}).
		Where("employee_id = ?", employeeID).
		Update("status", status).Error
}

func (r *employeeRepository) FindByID(employeeID uint) (models.User, error) {
	var employee models.User
	err := r.db.First(&employee, employeeID).Error
	return employee, err
}

func (r *employeeRepository) DB() *gorm.DB {
	return r.db
}
