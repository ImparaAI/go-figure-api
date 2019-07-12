package store

import (
	"database/sql"

	"api/database"
)

func New() *SubmissionStore {
	store := &SubmissionStore{database.GetDb()}

	return store
}

type Submission struct {
	Id int `json:"id"`
}

type SubmissionStore struct {
	db *sql.DB
}

func (store *SubmissionStore) Get(id int) string {
	rows, _ := store.db.Query("select name from foo where id = " + string(id))

	var originalPoints string

	rows.Scan(&originalPoints)

	return originalPoints
}

func (store *SubmissionStore) Submit() {
	store.db.Exec("INSERT INTO submissions (requestedDrawVectorCount, originalPoints) VALUES (20, 'foodicks')")
}