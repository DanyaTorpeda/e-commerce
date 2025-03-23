package repository

import (
	"e-commerce/internal/domains/user"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthPostgres struct {
	db     *sqlx.DB
	logger *logrus.Logger
}

func NewAuthPostgres(db *sqlx.DB, logger *logrus.Logger) *AuthPostgres {
	return &AuthPostgres{
		db:     db,
		logger: logger,
	}
}

func (r *AuthPostgres) CreateUser(input user.User) (int, error) {
	query := fmt.Sprintf(`
		insert into %s(first_name, last_name, password_hash, email, role)
		values($1, $2, $3, $4, $5)
		returning id;
	`, usersTable)

	var id int
	if err := r.db.QueryRow(query, input.FirstName, input.LastName,
		input.Password, input.Email, input.Role).Scan(&id); err != nil {
		r.logger.Warnf("error with query: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUserByEmail(email string) (*user.User, error) {
	query := fmt.Sprintf(`
		select *
		from %s
		where email = $1;
	`, usersTable)

	var usr user.User
	if err := r.db.Get(&usr, query, email); err != nil {
		r.logger.Warnf("no such user found: %s", err.Error())
		return nil, err
	}

	return &usr, nil
}
