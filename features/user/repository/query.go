package repository

import (
	"database/sql"
	"errors"

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
func (repo *userRepository) Create(input user.User) (err error) {
	_, errExec := repo.db.Exec(("Insert into user (public_id, fullname, email, password, phone, gender) Values (?, ?, ?, ?, ?, ?)"), input.PublicId, input.FullName, input.Email, input.Password, input.Phone, input.Gender)
	if errExec != nil {
		return errExec
	}
	return nil
}

// Get implements user.RepositoryInterface.
func (repo *userRepository) Get() ([]user.User, error) {
	userData := []user.User{}
	rows, err := repo.db.Query("select public_id, fullname, email, phone, gender from user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u user.User
		err := rows.Scan(&u.PublicId, &u.FullName, &u.Email, &u.Phone, &u.Gender)
		if err != nil {
			return nil, errors.New("failed retrieve data, error query")
		}
		userData = append(userData, u)
	}

	return userData, nil
}

// Update implements user.RepositoryInterface.
func (repo *userRepository) Update(input user.User, id string) (user.User, error) {
	panic("unimplemented")
}


// Delete implements user.RepositoryInterface.
func (repo *userRepository) Delete(id string) error {
	_, errExec := repo.db.Query(("Delete from user where id = ?"), id)
	if errExec != nil {
		return errExec
	}
	return nil
}