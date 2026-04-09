package service

import (
	"errors"
	models "tipes/internal/model"
	"tipes/internal/repository"
)

type EmployeeService interface {
	GetEmployees(branchID uint) ([]models.User, error)
	HireEmployee(emp models.User) error
	EditEmployee(emp models.User) error
	FireEmployee(employeeID uint) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo}
}

func (s *employeeService) GetEmployees(branchID uint) ([]models.User, error) {
	return s.repo.GetByBranch(branchID)
}

func (s *employeeService) HireEmployee(emp models.User) error {
	if emp.FirstName == "" || emp.LastName == "" || emp.Email == "" || emp.NoTelephone == "" {
		return errors.New("semua field wajib diisi")
	}
	emp.Status = "active"
	return s.repo.Create(emp)
}

func (s *employeeService) EditEmployee(emp models.User) error {
	existing, err := s.repo.FindByID(emp.UserID)
	if err != nil {
		return err
	}
	existing.FirstName = emp.FirstName
	existing.LastName = emp.LastName
	existing.Email = emp.Email
	existing.NoTelephone = emp.NoTelephone
	existing.BranchID = emp.BranchID
	return s.repo.Update(existing)
}

func (s *employeeService) FireEmployee(employeeID uint) error {
	return s.repo.ChangeStatus(employeeID, "inactive")
}
