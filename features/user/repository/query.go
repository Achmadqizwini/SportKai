package repository

import (
	"database/sql"
	"errors"
	"github.com/Achmadqizwini/SportKai/features/user/model"
	"github.com/Achmadqizwini/SportKai/utils/logger"
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

var logService = logger.NewLogger().Logger.With().Logger()

// Create implements Repository
func (repo *userRepository) Create(input model.User) (err error) {
	stmt, errPrepare := repo.db.Prepare(`
	INSERT INTO "user" (public_id, fullname, email, password, phone, gender) 
	VALUES ($1, $2, $3, $4, $5, $6)
	`)

	if errPrepare != nil {
		logService.Error().Err(errPrepare).Msg("error prepare query statement")
		return errPrepare
	}
	defer stmt.Close()

	_, errExec := stmt.Exec(input.PublicId, input.FullName, input.Email, input.Password, input.Phone, input.Gender)
	if errExec != nil {
		logService.Error().Err(errExec).Msg("error query execution")
		return errExec
	}

	return nil
}

// Get implements RepositoryInterface.
func (repo *userRepository) Get() ([]model.User, error) {
	userData := []model.User{}
	rows, err := repo.db.Query("select public_id, fullname, email, phone, gender, created_at, updated_at from user")
	if err != nil {
		return nil, errors.New("error query")
	}
	defer rows.Close()

	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.PublicId, &u.FullName, &u.Email, &u.Phone, &u.Gender, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, errors.New("error parsing data to model")
		}
		userData = append(userData, u)
	}

	return userData, nil
}

// Update implements RepositoryInterface.
func (repo *userRepository) Update(input model.User, id string) (model.User, error) {
	stmt, err := repo.db.Prepare("UPDATE user SET fullname=?, email=?, password=?, phone=?, gender=? WHERE public_id=?")
	if err != nil {
		return model.User{}, errors.New("error query prepare statement")
	}
	defer stmt.Close()

	result, errExec := stmt.Exec(input.FullName, input.Email, input.Password, input.Phone, input.Gender, id)
	if errExec != nil {
		return model.User{}, errors.New("error query execution")
	}
	if row, err := result.RowsAffected(); row == 0 || err != nil {
		return model.User{}, errors.New("update failed, no rows affected")
	}

	userData := model.User{}
	row := repo.db.QueryRow("SELECT public_id, fullname, email, phone, gender FROM user WHERE public_id=?", id)

	err = row.Scan(&userData.PublicId, &userData.FullName, &userData.Email, &userData.Phone, &userData.Gender)
	if err != nil {
		return model.User{}, errors.New("error parsing data to model")
	}

	return userData, nil
}

// Delete implements RepositoryInterface.
func (repo *userRepository) Delete(id string) error {
	stmt, errPrepare := repo.db.Prepare("DELETE FROM user WHERE public_id = ?")
	if errPrepare != nil {
		return errors.New("error prepare query statement")
	}
	defer stmt.Close()

	res, errExec := stmt.Exec(id)
	if row, err := res.RowsAffected(); row == 0 || err != nil {
		return errors.New("no user found")
	}
	if errExec != nil {
		return errExec
	}

	return nil
}

// GetById implements RepositoryInterface.
func (repo *userRepository) GetById(id string) (model.User, error) {
	userData := model.User{}
	row := repo.db.QueryRow("select public_id, fullname, email, phone, gender from user where public_id=?", id)

	err := row.Scan(&userData.PublicId, &userData.FullName, &userData.Email, &userData.Phone, &userData.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, errors.New("no user found")
		}
		return model.User{}, errors.New("error parsing to model")
	}

	return userData, nil
}
