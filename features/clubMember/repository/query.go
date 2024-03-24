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
	Update(input model.ClubMember, id string) error
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
	memberData := model.ClubMember{}
	row := u.db.QueryRow("SELECT m.public_id, u.public_id as user_id, c.public_id AS club_id, m.status "+
		"FROM club_member m "+
		"JOIN user u ON m.user_id = u.id "+
		"JOIN club c ON m.club_id = c.id "+
		"WHERE m.public_id=?", id)
	err := row.Scan(&memberData.PublicId, &memberData.UserId, &memberData.ClubId, &memberData.Status)
	if err != nil {
		fmt.Println(err)
		return model.ClubMember{}, errors.New("error parsing data to model")
	}
	return memberData, nil
}

// Update implements RepositoryInterface.
func (u *memberRepository) Update(input model.ClubMember, id string) error {
	stmt, err := u.db.Prepare("UPDATE club_member SET status=? WHERE public_id=?")
	if err != nil {
		return errors.New("error query prepare statement")
	}
	defer stmt.Close()

	result, errExec := stmt.Exec(input.Status, id)
	if errExec != nil {
		return errors.New("error query execution")
	}
	if row, err := result.RowsAffected(); row == 0 || err != nil {
		return errors.New("update failed, no rows affected")
	}

	return nil
}
