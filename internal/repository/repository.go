package repository

import (
	"e-commerce/internal/domains/user"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Authorization interface {
	CreateUser(input user.User) (int, error)
	GetUserByEmail(email string) (*user.User, error)
}

type Repository struct {
	logger *logrus.Logger
	Authorization
}

func New(db *sqlx.DB, logger *logrus.Logger) *Repository {
	return &Repository{
		logger:        logger,
		Authorization: NewAuthPostgres(db, logger),
	}
}
