package repository

import (
	"AnonimousChat/internal/data"
	"database/sql"
)

type InterlocutorRepository interface {
	MigrationInterlocutorTable() error

	Registration(user data.Interlocutor) error
	GetInterlocutor(tag data.TagID) (data.Interlocutor, error)

	ChangeSelfGender(tag data.TagID, gender data.GenderID) error
	ChangeTargetGender(tag data.TagID, gender data.GenderID) error
	AddCountConnections(tag data.TagID) error
	AddSumDonation(tag data.TagID, sum int) error
}

type interlocutorRepository struct {
	db *sql.DB
}

func NewInterlocutorRepository(db *sql.DB) InterlocutorRepository {
	return &interlocutorRepository{
		db: db,
	}
}

func (r *interlocutorRepository) MigrationInterlocutorTable() error {
	_, err := r.db.Exec(`
	CREATE TABLE IF NOT EXISTS interlocutor(
		tag VARCHAR PRIMARY KEY,
		source VARCHAR NOT NULL,
		id INTEGER NOT NULL,
		self_gender INTEGER NOT NULL,
		target_gender INTEGER NOT NULL,
		count_connections INTEGER NOT NULL,
		sum_donation INTEGER NOT NULL
	);
	`)

	return err
}

func (r *interlocutorRepository) Registration(user data.Interlocutor) error {
	_, err := r.db.Exec(
		`INSERT INTO interlocutor VALUES($1, $2, $3, $4, $5, $6, $7);`,
		user.Tag,
		user.Source,
		user.ID,
		user.SelfGender,
		user.TargetGender,
		user.CountConnections,
		user.SumDonation,
	)

	return err
}

func (r *interlocutorRepository) GetInterlocutor(tag data.TagID) (data.Interlocutor, error) {
	var user data.Interlocutor

	row := r.db.QueryRow(`SELECT * FROM interlocutor WHERE tag=$1`, tag)

	err := row.Err()
	if err != nil {
		return user, err
	}

	err = row.Scan(
		&user.Tag,
		&user.Source,
		&user.ID,
		&user.SelfGender,
		&user.TargetGender,
		&user.CountConnections,
		&user.SumDonation,
	)

	return user, err
}
func (r *interlocutorRepository) ChangeSelfGender(tag data.TagID, gender data.GenderID) error {
	_, err := r.db.Exec(`UPDATE interlocutor SET self_gender = $2 WHERE tag = $1`, tag, gender)
	return err
}
func (r *interlocutorRepository) ChangeTargetGender(tag data.TagID, gender data.GenderID) error {
	_, err := r.db.Exec(`UPDATE interlocutor SET target_gender = $2 WHERE tag = $1`, tag, gender)
	return err
}
func (r *interlocutorRepository) AddCountConnections(tag data.TagID) error {
	_, err := r.db.Exec(`UPDATE interlocutor SET count_connections = count_connections+1 WHERE tag = $1`, tag)
	return err
}
func (r *interlocutorRepository) AddSumDonation(tag data.TagID, sum int) error {
	_, err := r.db.Exec(`UPDATE interlocutor SET sum_donation = sum_donation + $2 WHERE tag = $1`, tag, sum)
	return err
}
