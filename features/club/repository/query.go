package repository

import (
	"database/sql"
	"errors"
	"github.com/Achmadqizwini/SportKai/features/club/model"
)

type RepositoryInterface interface {
	Create(input model.Club) error
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
func (c *clubRepository) Create(input model.Club) error {
	stmt, errPrepare := c.db.Prepare("INSERT INTO club (public_id, name, address, city, description, joined_member, member_total, rules, requirements ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if errPrepare != nil {
		return errors.New("error prepare query statement")
	}
	defer stmt.Close()

	_, errExec := stmt.Exec(input.PublicId, input.Name, input.Address, input.City, input.Description, input.JoinedMember, input.MemberTotal, input.Rules, input.Requirements)
	if errExec != nil {
		return errors.New("error query execution")
	}

	return nil
}

// Delete implements RepositoryInterface.
func (c *clubRepository) Delete(id string) error {
	panic("unimplemented")
}

// Get implements RepositoryInterface.
func (c *clubRepository) Get() ([]model.Club, error) {
	panic("unimplemented")
}

// GetById implements RepositoryInterface.
func (c *clubRepository) GetById(id string) (model.Club, error) {
	panic("unimplemented")
}

// Update implements RepositoryInterface.
func (c *clubRepository) Update(input model.Club, id string) (model.Club, error) {
	panic("unimplemented")
}
