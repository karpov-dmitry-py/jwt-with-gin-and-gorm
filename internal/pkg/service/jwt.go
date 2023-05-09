package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type JwtService interface {
	GeneratePasswordHash(password string) ([]byte, error)
}

type jwtService struct{}

// NewJwtService inits a jwt service
func NewJwtService() JwtService {
	return &jwtService{}
}

func (s *jwtService) GeneratePasswordHash(password string) ([]byte, error) {
	//nolint:gomnd
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, fmt.Errorf("failed to hash the password: %v", err)
	}

	return hash, nil
}
