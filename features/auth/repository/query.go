package repository

import (
	"database/sql"
	"errors"
	"github.com/Achmadqizwini/SportKai/features/auth/model"
)

type RepositoryInterface interface {
	FindUser(input model.Auth) (model.Auth, error)
}

type authRepository struct {
	db *sql.DB
}

func New(db *sql.DB) RepositoryInterface {
	return &authRepository{
		db: db,
	}
}

// FindUser implements RepositoryInterface.
func (a *authRepository) FindUser(input model.Auth) (model.Auth, error) {
	userData := model.Auth{}
	row := a.db.QueryRow("select public_id, fullname, email, password, phone, gender from user where email=? and password=?", input.Email, input.Password)

	err := row.Scan(&userData.PublicId, &userData.FullName, &userData.Email, &userData.Password, &userData.Phone, &userData.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Auth{}, errors.New("no user found")
		}
		return model.Auth{}, errors.New("error parsing to model")
	}

	return userData, nil
}
