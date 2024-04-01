package repository

import (
	"database/sql"
	"errors"
	"github.com/Achmadqizwini/SportKai/features/user/model"
)

type RepositoryInterface interface {
	Create(input model.User) error
	Get() ([]model.User, error)
	GetById(id string) (model.User, error)
	Update(input model.User, id string) (model.User, error)
	Delete(id string) error
}

type userRepository struct {
	db *sql.DB
}

func New(db *sql.DB) RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements Repository
func (repo *userRepository) Create(input model.User) (err error) {
	stmt, errPrepare := repo.db.Prepare(`
	INSERT INTO "user" (public_id, fullname, email, password, phone, gender) 
	VALUES ($1, $2, $3, $4, $5, $6)
	`)

	if errPrepare != nil {
		return errPrepare
	}
	defer stmt.Close()

	_, errExec := stmt.Exec(input.PublicId, input.FullName, input.Email, input.Password, input.Phone, input.Gender)
	if errExec != nil {
		return errExec
	}

	return nil
}

// Get implements RepositoryInterface.
func (repo *userRepository) Get() ([]model.User, error) {
	userData := []model.User{}
	rows, err := repo.db.Query(`SELECT public_id, fullname, email, phone, gender, created_at, updated_at FROM "user"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.PublicId, &u.FullName, &u.Email, &u.Phone, &u.Gender, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		userData = append(userData, u)
	}

	return userData, nil
}

// Update implements RepositoryInterface.
func (repo *userRepository) Update(input model.User, id string) (model.User, error) {
	stmt, err := repo.db.Prepare(`
	UPDATE "user" SET fullname=$1, email=$2, password=$3, phone=$4, gender=$5 WHERE public_id=$6`)
	if err != nil {
		return model.User{}, err
	}
	defer stmt.Close()

	result, errExec := stmt.Exec(input.FullName, input.Email, input.Password, input.Phone, input.Gender, id)
	if errExec != nil {
		return model.User{}, errExec
	}
	if row, err := result.RowsAffected(); row == 0 || err != nil {
		return model.User{}, err
	}

	userData := model.User{}
	row := repo.db.QueryRow(`SELECT public_id, fullname, email, phone, gender FROM "user" WHERE public_id=$1`, id)

	err = row.Scan(&userData.PublicId, &userData.FullName, &userData.Email, &userData.Phone, &userData.Gender)
	if err != nil {
		return model.User{}, err
	}

	return userData, nil
}

// Delete implements RepositoryInterface.
func (repo *userRepository) Delete(id string) error {
	stmt, errPrepare := repo.db.Prepare(`DELETE FROM "user" WHERE public_id = $1`)
	if errPrepare != nil {
		return errPrepare
	}
	defer stmt.Close()

	res, errExec := stmt.Exec(id)
	if row, err := res.RowsAffected(); row == 0 || err != nil {
		return err
	}
	if errExec != nil {
		return errExec
	}

	return nil
}

// GetById implements RepositoryInterface.
func (repo *userRepository) GetById(id string) (model.User, error) {
	userData := model.User{}
	row := repo.db.QueryRow(`SELECT public_id, fullname, email, phone, gender FROM "user" WHERE public_id=$1`, id)

	err := row.Scan(&userData.PublicId, &userData.FullName, &userData.Email, &userData.Phone, &userData.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, errors.New("no user found")
		}
		return model.User{}, err
	}

	return userData, nil
}
