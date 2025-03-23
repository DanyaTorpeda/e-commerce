package services

import (
	"e-commerce/internal/domains/user"
	"e-commerce/internal/repository"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var JWTSecret = []byte("aisjdbnf8uaywbvc807q34fcoyb8q3c487yv")

type Claims struct {
	UserID    int
	UserEmail string
}

type AuthService struct {
	logger *logrus.Logger
	repo   repository.Authorization
}

func NewAuthService(repo repository.Authorization, logger *logrus.Logger) *AuthService {
	return &AuthService{
		repo:   repo,
		logger: logger,
	}
}

func (s *AuthService) CreateUser(input user.User) (int, error) {
	if err := user.Validate(input.Email, input.Password); err != nil {
		s.logger.Warnf("invalid input data: %s", err.Error())
		return 0, err
	}

	if err := input.Role.Validate(); err != nil {
		s.logger.Warnf("invalid input data: %s", err.Error())
		return 0, err
	}

	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		s.logger.Warnf("error hashing password: %s", err.Error())
		return 0, err
	}

	input.Password = hashedPassword

	return s.repo.CreateUser(input)
}

func (s *AuthService) CheckUser(input user.UserLogin) (*user.User, error) {
	if err := user.Validate(input.Email, input.Password); err != nil {
		s.logger.Warnf("invalid input data: %s", err.Error())
		return nil, err
	}

	usr, err := s.repo.GetUserByEmail(input.Email)
	if err != nil {
		s.logger.Warnf("no such user found: %s", err.Error())
		return nil, err
	}

	if err := checkPassword(input.Password, usr.Password); err != nil {
		s.logger.Warnf("passwords dont match: %s", err.Error())
		return nil, err
	}

	return usr, nil
}

func (s *AuthService) CreateToken(id int, email string) {

}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func checkPassword(password, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
