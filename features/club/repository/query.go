package repository

import (
	"database/sql"
	"errors"

	"github.com/Achmadqizwini/SportKai/features/club/model"
)

type RepositoryInterface interface {
	Create(input model.Club) (string, error)
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
func (c *clubRepository) Create(input model.Club) (string, error) {
	stmt, err := c.db.Prepare("INSERT INTO club (public_id, name, address, city, description, joined_member, member_total, rules, requirements ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return "", errors.New("error prepare query statement")
	}
	defer stmt.Close()

	rows, err := stmt.Exec(input.PublicId, input.Name, input.Address, input.City, input.Description, input.JoinedMember, input.MemberTotal, input.Rules, input.Requirements)
	if err != nil {
		return "", errors.New("error query execution")
	}

	if row, _ := rows.RowsAffected(); row > 0 {
		lastId, err := rows.LastInsertId()
		if err == nil {
			var id string
			err := c.db.QueryRow("select public_id from club where id = ?", lastId).Scan(&id)
			if err != nil {
				return "", errors.New("error query statement")
			}
			return id, nil
		}
	}

	return "", err
}

// Delete implements RepositoryInterface.
func (c *clubRepository) Delete(id string) error {
	panic("unimplemented")
}

// Get implements RepositoryInterface.
func (c *clubRepository) Get() ([]model.Club, error) {
	clubs := []model.Club{}
	rows, err := c.db.Query("select public_id, name, address, city, description, joined_member, member_total, rules, requirements, created_at from club")
	if err != nil {
		return nil, errors.New("error query")
	}
	defer rows.Close()

	for rows.Next() {
		var u model.Club
		err := rows.Scan(&u.PublicId, &u.Name, &u.Address, &u.City, &u.Description, &u.JoinedMember, &u.MemberTotal, &u.Rules, &u.Requirements, &u.CreatedAt)
		if err != nil {
			return nil, errors.New("error parsing data to model")
		}
		clubs = append(clubs, u)
	}

	return clubs, nil
}

// GetById implements RepositoryInterface.
func (c *clubRepository) GetById(id string) (model.Club, error) {
	clubData := model.Club{}
	row := c.db.QueryRow("select public_id, name, address, city, description, joined_member, member_total, rules, requirements, created_at from club where public_id=?", id)

	err := row.Scan(&clubData.PublicId, &clubData.Name, &clubData.Address, &clubData.City, &clubData.Description, &clubData.JoinedMember, &clubData.MemberTotal, &clubData.Rules, &clubData.Requirements, &clubData.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Club{}, errors.New("no club found")
		}
		return model.Club{}, errors.New("error parsing to model")
	}

	return clubData, nil
}

// Update implements RepositoryInterface.
func (c *clubRepository) Update(input model.Club, id string) (model.Club, error) {
	panic("unimplemented")
}
