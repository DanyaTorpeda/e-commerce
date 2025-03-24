package services

import (
	"e-commerce/internal/domains/user"
	"e-commerce/internal/repository"

	"github.com/sirupsen/logrus"
)

type Authorization interface {
	CreateUser(input user.User) (int, error)
	CheckUser(input user.UserLogin) (*user.User, error)
	CreateToken(id int, email string) (string, error)
	ParseToken(tokenString string) (*Claims, error)
}

type Service struct {
	logger *logrus.Logger
	Authorization
}

func New(repo *repository.Repository, logger *logrus.Logger) *Service {
	return &Service{
		logger:        logger,
		Authorization: NewAuthService(repo.Authorization, logger),
	}
}
