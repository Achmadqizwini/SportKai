package repository

import (
	"database/sql"
	model "github.com/Achmadqizwini/SportKai/features/clubMember/model"
)

type RepositoryInterface interface {
	Create(input model.MemberPayload) error
	Get(club_id string) ([]model.MemberPayload, error)
	GetById(club_id, id string) (model.MemberPayload, error)
	Update(input model.MemberPayload, id string) error
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
func (u *memberRepository) Create(input model.MemberPayload) error {
	type ID struct {
		UserId uint
		ClubId uint
	}
	var id ID
	err := u.db.QueryRow(`
		SELECT u.id AS user_id, c.id AS club_id 
		FROM "user" u 
		JOIN club c ON u.public_id = $1 AND c.public_id = $2
	`, input.UserId, input.ClubId).Scan(&id.UserId, &id.ClubId)
	if err != nil {
		return err
	}
	stmt, errPrepare := u.db.Prepare(`
		INSERT INTO club_member (public_id, user_id, club_id, status) 
		VALUES ($1, $2, $3, $4)
	`)
	if errPrepare != nil {
		return errPrepare
	}
	defer stmt.Close()

	_, errExec := stmt.Exec(input.PublicId, id.UserId, id.ClubId, input.Status)
	if errExec != nil {
		return errExec
	}

	return nil
}

// Delete implements RepositoryInterface.
func (u *memberRepository) Delete(id string) error {
	stmt, errPrepare := u.db.Prepare(`DELETE FROM club_member WHERE public_id = $1`)
	if errPrepare != nil {
		return errPrepare
	}
	defer stmt.Close()

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
func (u *memberRepository) Get(club_id string) ([]model.MemberPayload, error) {
	members := []model.MemberPayload{}
	rows, err := u.db.Query(`
	SELECT m.id, m.public_id, m.user_id, m.club_id, m.status, m.joined_at, m.left_at FROM club_member m
	LEFT JOIN club c ON m.club_id = c.public_id
	WHERE c.public_id = $1`, club_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u model.MemberPayload
		err := rows.Scan(&u.ID, &u.PublicId, &u.UserId, &u.ClubId, &u.Status, &u.JoinedAt, &u.LeftAt)
		if err != nil {
			return nil, err
		}
		members = append(members, u)
	}

	return members, nil
}

// GetById implements RepositoryInterface.
func (u *memberRepository) GetById(club_id, id string) (model.MemberPayload, error) {
	memberData := model.MemberPayload{}
	row := u.db.QueryRow(`SELECT m.public_id, u.public_id as user_id, c.public_id AS club_id, m.status, m.joined_at
		FROM club_member m
		JOIN "user" u ON m.user_id = u.id
		JOIN club c ON m.club_id = c.id
		WHERE m.public_id=$1 AND c.public_id=$2`, id, club_id)
	err := row.Scan(&memberData.PublicId, &memberData.UserId, &memberData.ClubId, &memberData.Status, &memberData.JoinedAt)
	if err != nil {
		return model.MemberPayload{}, err
	}
	return memberData, nil
}

// Update implements RepositoryInterface.
func (u *memberRepository) Update(input model.MemberPayload, id string) error {
	stmt, err := u.db.Prepare("UPDATE club_member SET status=$1 WHERE public_id=$2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, errExec := stmt.Exec(input.Status, id)
	if errExec != nil {
		return errExec
	}
	if row, err := result.RowsAffected(); row == 0 || err != nil {
		return err
	}

	return nil
}
