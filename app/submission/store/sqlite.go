package store

import (
	"github.com/jmoiron/sqlx"

	"api/database"
	"api/app/submission/types"
)

func New() *SubmissionStore {
	store := &SubmissionStore{database.GetDb()}

	return store
}

type SubmissionStore struct {
	db *sqlx.DB
}

func (store *SubmissionStore) Exists(id int) (bool, error) {
	var count int
	err := store.db.Get(&count, "SELECT COUNT(id) FROM submissions WHERE id = ?", id)

	return count > 0, err
}

func (store *SubmissionStore) Get(id int) (types.Submission, error) {
	var submission types.Submission
	err := store.db.Get(&submission, "SELECT * FROM submissions WHERE id = ?", id)

	return submission, err
}

func (store *SubmissionStore) Submit() (int, error) {
	result := store.db.MustExec("INSERT INTO submissions (requestedDrawVectorCount, originalPoints) VALUES (?, ?)", 20, "foopicks")
	id, err := result.LastInsertId()

	return int(id), err
}