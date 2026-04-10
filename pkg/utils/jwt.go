package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = []byte(os.Getenv("KEY")) // ganti dengan env nanti

type Claims struct {
	UserID   uint
	Email    string
	RoleID   uint
	BranchID uint
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, email string, roleID uint, branchID uint) (string, error) {
	claims := Claims{
		UserID:   userID,
		Email:    email,
		RoleID:   roleID,
		BranchID: branchID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(SECRET_KEY)
}
