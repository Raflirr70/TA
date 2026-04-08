package service

import (
	"errors"
	models "tipes/internal/model"
	"tipes/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	CheckNewUser(email, noTelephone string) (bool, error)
	Register(user models.User) error
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) CheckNewUser(email, noTelephone string) (bool, error) {
	_, err := s.repo.FindUser(email, noTelephone)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, nil // user belum ada
	}
	if err == nil {
		return false, nil // user sudah ada
	}
	return false, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	return string(bytes), err
}

func (s *authService) Validate(user models.User) (bool, []error) {
	var errs []error

	if user.FirstName == "" {
		errs = append(errs, errors.New("Firstname wajib diisi"))
	}
	if user.LastName == "" {
		errs = append(errs, errors.New("Lastname wajib diisi"))
	}
	if user.Email == "" {
		errs = append(errs, errors.New("Email wajib diisi"))
	}
	if user.NoTelephone == "" {
		errs = append(errs, errors.New("No Telephone wajib diisi"))
	}
	if user.Password == "" {
		errs = append(errs, errors.New("Password wajib diisi"))
	}

	if len(errs) > 0 {
		return false, errs
	}
	return true, nil
}

func (s *authService) Register(user models.User) error {
	// validasi
	valid, errs := s.Validate(user)
	if !valid {
		return errors.Join(errs...)
	}

	// cek user sudah ada atau belum
	ok, err := s.CheckNewUser(user.Email, user.NoTelephone)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("email atau no telephone sudah digunakan")
	}

	// hash password
	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash

	// simpan ke DB
	return s.repo.Create(user)
}
