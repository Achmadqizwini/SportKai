package repository

import (
	"database/sql"
	"errors"

	"github.com/Achmadqizwini/SportKai/features/club/model"
)

type RepositoryInterface interface {
	Create(input model.Club) (uint, error)
	Get() ([]model.Club, error)
	GetById(id string) (model.Club, error)
	Update(input model.Club, id string) (model.Club, error)
	Delete(id string) error
}

type clubRepository struct {
	db *sql.DB
}

func New(db *sql.DB) RepositoryInterface {
	return &clubRepository{
		db: db,
	}
}

// Create implements RepositoryInterface.
func (c *clubRepository) Create(input model.Club) (uint, error) {
	stmt, err := c.db.Prepare("INSERT INTO club (public_id, name, address, city, description, joined_member, member_total, rules, requirements ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, errors.New("error prepare query statement")
	}
	defer stmt.Close()

	rows, err := stmt.Exec(input.PublicId, input.Name, input.Address, input.City, input.Description, input.JoinedMember, input.MemberTotal, input.Rules, input.Requirements)
	if err != nil {
		return 0, errors.New("error query execution")
	}

	if row, _ := rows.RowsAffected(); row > 0 {
		lastId, err := rows.LastInsertId()
		if err == nil {
			return uint(lastId), nil
		}
	}

	return 0, err
}

// Delete implements RepositoryInterface.
func (c *clubRepository) Delete(id string) error {
	panic("unimplemented")
}

// Get implements RepositoryInterface.
func (c *clubRepository) Get() ([]model.Club, error) {
	clubs := []model.Club{}
	rows, err := c.db.Query("select public_id, name, address, city, description, joined_member, member_total, rules, requirements from club")
	if err != nil {
		return nil, errors.New("error query")
	}
	defer rows.Close()

	for rows.Next() {
		var u model.Club
		err := rows.Scan(&u.PublicId, &u.Name, &u.Address, &u.City, &u.Description, &u.JoinedMember, &u.MemberTotal, &u.Rules, &u.Requirements)
		if err != nil {
			return nil, errors.New("error parsing data to model")
		}
		clubs = append(clubs, u)
	}

	return clubs, nil
}

// GetById implements RepositoryInterface.
func (c *clubRepository) GetById(id string) (model.Club, error) {
	panic("unimplemented")
}

// Update implements RepositoryInterface.
func (c *clubRepository) Update(input model.Club, id string) (model.Club, error) {
	panic("unimplemented")
}
