package store

import (
	"github.com/jmoiron/sqlx"

	"api/database"
	"api/app/drawing/types"
)

func New() *DrawingStore {
	store := &DrawingStore{database.GetDb()}

	return store
}

type DrawingStore struct {
	db *sqlx.DB
}

func (store *DrawingStore) Exists(id int) (bool, error) {
	var count int
	err := store.db.Get(&count, "SELECT COUNT(id) FROM drawings WHERE id = ?", id)

	return count > 0, err
}

func (store *DrawingStore) Get(id int) (types.Drawing, error) {
	var drawing types.Drawing
	err := store.db.Get(&drawing, "SELECT * FROM drawings WHERE id = ?", id)

	return drawing, err
}

func (store *DrawingStore) Submit() (int, error) {
	result := store.db.MustExec("INSERT INTO drawings (requestedDrawVectorCount, originalPoints) VALUES (?, ?)", 20, "foopicks")
	id, err := result.LastInsertId()

	return int(id), err
}

func (store *DrawingStore) AddVectors(drawingId int, vectors string) {
	store.db.MustExec("UPDATE drawings SET vectors = ? WHERE drawingId = ?", vectors, drawingId)
}