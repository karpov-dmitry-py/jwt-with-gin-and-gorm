package service

import (
	"context"
	"fmt"

	"github.com/karpov-dmitry-py/jwt-with-gin-and-gorm/internal/pkg/repository"
	"gorm.io/gorm"
)

type UserService interface {
	SignUp(ctx context.Context, user SignUpUser) (int64, error)
}

type userService struct {
	dbConn     *gorm.DB
	jwtService JwtService
}

// NewUserService inits a user service
func NewUserService(dbConn *gorm.DB, jwtService JwtService) UserService {
	return &userService{
		dbConn:     dbConn,
		jwtService: jwtService,
	}
}

func (s *userService) SignUp(_ context.Context, user SignUpUser) (int64, error) {
	var (
		dbUser = repository.User{Email: user.Email}
	)

	hash, err := s.jwtService.GeneratePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}

	dbUser.Password = string(hash)

	if err = s.dbConn.Create(&dbUser).Error; err != nil {
		return 0, fmt.Errorf("failed to create a user: %v", err)
	}

	return int64(dbUser.ID), nil
}
