package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Achmadqizwini/SportKai/utils/logger"

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

var logRepo = logger.NewLogger().Logger.With().Logger()

// Create implements RepositoryInterface.
func (u *memberRepository) Create(input model.ClubMember) error {
	fmt.Println(input)
	stmt, errPrepare := u.db.Prepare("INSERT INTO club_member (public_id, user_id, club_id, status) VALUES (?, ?, ?, ?)")
	if errPrepare != nil {
		return errors.New("error prepare query statement")
	}
	defer stmt.Close()

	_, errExec := stmt.Exec(input.PublicId, input.UserId, input.ClubId, input.Status)
	if errExec != nil {
		logRepo.Error().Err(errExec).Msg("error query execution")
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
	members := []model.ClubMember{}
	rows, err := u.db.Query("select id, public_id, user_id, club_id, status, joined_at, left_at from club_member")
	if err != nil {
		return nil, errors.New("error query")
	}
	defer rows.Close()

	for rows.Next() {
		var u model.ClubMember
		err := rows.Scan(&u.ID, &u.PublicId, &u.UserId, &u.ClubId, &u.Status, &u.JoinedAt, &u.LeftAt)
		if err != nil {
			return nil, errors.New("error parsing data to model")
		}
		members = append(members, u)
	}

	return members, nil
}

// GetById implements RepositoryInterface.
func (u *memberRepository) GetById(id string) (model.ClubMember, error) {
	panic("unimplemented")
}

// Update implements RepositoryInterface.
func (u *memberRepository) Update(input model.ClubMember, id string) (model.ClubMember, error) {
	panic("unimplemented")
}
