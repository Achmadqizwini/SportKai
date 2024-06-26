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
	row := a.db.QueryRow(`
	SELECT id, public_id, fullname, email, password, phone, gender 
	FROM "user" WHERE email=$1`, input.Email)

	err := row.Scan(&userData.ID, &userData.PublicId, &userData.FullName, &userData.Email, &userData.Password, &userData.Phone, &userData.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Auth{}, errors.New("no user found")
		}
		return model.Auth{}, err
	}

	return userData, nil
}
