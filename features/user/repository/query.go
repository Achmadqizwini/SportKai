package repository

import (
	"database/sql"
	"github.com/Achmadqizwini/SportKai/features/user"
)

type userRepository struct {
	db *sql.DB
}

func New(db *sql.DB) user.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *userRepository) Create(input user.Core) (err error) {
	_, errExec := repo.db.Exec(("Insert into user (public_id, fullname, email, password, phone, gender) Values (1, ?, ?, ?, ?, ?)"), input.FullName, input.Email, input.Password, input.Phone, input.Gender)
	if errExec != nil {
		return errExec
	}
	return nil
}
