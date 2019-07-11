package store

import (
	"api/database"
)

func New() *SubmissionStore {
	store := &SubmissionStore{database.GetConnection()}

	return store
}

type Submission struct {
	Id int `json:"id"`
}

type SubmissionStore struct {
	connection *database.Connection
}

func (store *SubmissionStore) Get() string {
	return store.connection.Val
}