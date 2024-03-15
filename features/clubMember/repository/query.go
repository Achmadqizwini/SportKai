package repository

import (
	"database/sql"
	"errors"
	model "github.com/Achmadqizwini/SportKai/features/clubMember/model"
)

type RepositoryInterface interface {
	Create(input model.ClubMember) error
	Get() ([]model.ClubMember, error)
	GetById(id string) (model.ClubMember, error)
	Update(input model.ClubMember, id string) (model.ClubMember, error)
	Delete(id string) error
}

type memberRepository struct {
	db *sql.DB
}

func New(db *sql.DB) RepositoryInterface {
	return &memberRepository{
		db: db,
	}
}

// Create implements RepositoryInterface.
func (u *memberRepository) Create(input model.ClubMember) error {
	stmt, errPrepare := u.db.Prepare("INSERT INTO user (public_id, user_id, club_id, status) VALUES (?, ?, ?, ?)")
	if errPrepare != nil {
		return errors.New("error prepare query statement")
	}
	defer stmt.Close()

	_, errExec := stmt.Exec(input.PublicId, input.UserId, input.ClubId, input.Status)
	if errExec != nil {
		return errors.New("error query execution")
	}

	return nil
}

// Delete implements RepositoryInterface.
func (u *memberRepository) Delete(id string) error {
	panic("unimplemented")
}

// Get implements RepositoryInterface.
func (u *memberRepository) Get() ([]model.ClubMember, error) {
	panic("unimplemented")
}

// GetById implements RepositoryInterface.
func (u *memberRepository) GetById(id string) (model.ClubMember, error) {
	panic("unimplemented")
}

// Update implements RepositoryInterface.
func (u *memberRepository) Update(input model.ClubMember, id string) (model.ClubMember, error) {
	panic("unimplemented")
}
