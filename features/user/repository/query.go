package repository

import (
	"database/sql"
	"errors"
	"fmt"

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
    stmt, errPrepare := repo.db.Prepare("INSERT INTO user (public_id, fullname, email, password, phone, gender) VALUES (?, ?, ?, ?, ?, ?)")
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
	stmt, err := repo.db.Prepare("UPDATE user SET fullname=?, email=?, password=?, phone=?, gender=? WHERE public_id=?")
	if err != nil {
		return user.User{}, fmt.Errorf("failed to prepare update statement: %v", err)
	}
	defer stmt.Close()

	result, errExec := stmt.Exec(input.FullName, input.Email, input.Password, input.Phone, input.Gender, id)
	if errExec != nil {
		return user.User{}, errExec
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return user.User{}, fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return user.User{}, fmt.Errorf("no user found with id %s", id)
	}
	userData := user.User{}
	row := repo.db.QueryRow("SELECT public_id, fullname, email, phone, gender FROM user WHERE public_id=?", id)

	err = row.Scan(&userData.PublicId, &userData.FullName, &userData.Email, &userData.Phone, &userData.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return user.User{}, fmt.Errorf("no user found with id %s", id)
		}
		return user.User{}, fmt.Errorf("failed to scan row: %v", err)
	}

	return userData, nil
}

// Delete implements user.RepositoryInterface.
func (repo *userRepository) Delete(id string) error {
    stmt, errPrepare := repo.db.Prepare("DELETE FROM user WHERE public_id = ?")
    if errPrepare != nil {
        return errPrepare
    }
    defer stmt.Close()

    _, errExec := stmt.Exec(id)
    if errExec != nil {
        return errExec
    }

    return nil
}

// GetById implements user.RepositoryInterface.
func (repo *userRepository) GetById(id string) (user.User, error) {
	userData := user.User{}
	row := repo.db.QueryRow("select public_id, fullname, email, phone, gender from user where public_id=?", id)

	err := row.Scan(&userData.PublicId, &userData.FullName, &userData.Email, &userData.Phone, &userData.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return user.User{}, fmt.Errorf("no user found with id %s", id)
		}
		return user.User{}, fmt.Errorf("failed to scan row: %v", err)
	}

	return userData, nil
}