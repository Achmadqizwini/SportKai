package repository

import (
	"database/sql"

	"github.com/Achmadqizwini/SportKai/features/club/model"
)

type RepositoryInterface interface {
	Create(input model.Club) (string, error)
	Get() ([]model.Club, error)
	GetById(id string) (model.Club, error)
	Update(input model.Club, id string) error
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
	stmt, err := c.db.Prepare(`
		INSERT INTO club (public_id, name, address, city, description, joined_member, member_total, rules, requirements ) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING public_id
	`)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var id string
	err = stmt.QueryRow(input.PublicId, input.Name, input.Address, input.City, input.Description, input.JoinedMember, input.MemberTotal, input.Rules, input.Requirements).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

// Delete implements RepositoryInterface.
func (c *clubRepository) Delete(id string) error {
	stmt, errPrepare := c.db.Prepare("DELETE FROM club WHERE public_id = $1")
	if errPrepare != nil {
		return errPrepare
	}
	defer stmt.Close()
	// wip delete child data also
	res, errExec := stmt.Exec(id)
	if errExec != nil {
		return errExec
	}
	if row, err := res.RowsAffected(); row == 0 || err != nil {
		return err
	}

	return nil
}

// Get implements RepositoryInterface.
func (c *clubRepository) Get() ([]model.Club, error) {
	clubs := []model.Club{}
	rows, err := c.db.Query("SELECT public_id, name, address, city, description, joined_member, member_total, rules, requirements, created_at FROM club")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u model.Club
		err := rows.Scan(&u.PublicId, &u.Name, &u.Address, &u.City, &u.Description, &u.JoinedMember, &u.MemberTotal, &u.Rules, &u.Requirements, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		clubs = append(clubs, u)
	}

	return clubs, nil
}

// GetById implements RepositoryInterface.
func (c *clubRepository) GetById(id string) (model.Club, error) {
	clubData := model.Club{}
	row := c.db.QueryRow("SELECT public_id, name, address, city, description, joined_member, member_total, rules, requirements, created_at FROM club WHERE public_id=$1", id)

	err := row.Scan(&clubData.PublicId, &clubData.Name, &clubData.Address, &clubData.City, &clubData.Description, &clubData.JoinedMember, &clubData.MemberTotal, &clubData.Rules, &clubData.Requirements, &clubData.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Club{}, err
		}
		return model.Club{}, err
	}

	return clubData, nil
}

// Update implements RepositoryInterface.
func (c *clubRepository) Update(input model.Club, id string) error {
	stmt, err := c.db.Prepare("UPDATE club SET name=$1, address=$2, city=$3, description=$4, rules=$5, requirements=$6  WHERE public_id=$7")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, errExec := stmt.Exec(input.Name, input.Address, input.City, input.Description, input.Rules, input.Requirements, id)
	if errExec != nil {
		return err
	}
	if row, err := result.RowsAffected(); row == 0 || err != nil {
		return err
	}

	return nil
}
