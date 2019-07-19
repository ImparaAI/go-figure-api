package store

import (
	"time"
	"encoding/json"
	"github.com/jmoiron/sqlx"

	"api/app/formatting"
	"api/app/drawing/types"
)

type SqliteStore struct {
	db *sqlx.DB
}

func (store *SqliteStore) Exists(id int) (bool) {
	var count int
	err := store.db.Get(&count, "SELECT COUNT(id) FROM drawings WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	return count > 0
}

func (store *SqliteStore) Get(id int) (types.Drawing) {
	var sqlDrawing SqlDrawing

	err := store.db.Get(&sqlDrawing, "SELECT * FROM drawings WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	drawing := types.Drawing{
		Id: sqlDrawing.Id,
		Featured: sqlDrawing.Featured,
		OriginalPoints: sqlDrawing.OriginalPoints,
		DrawVectors: sqlDrawing.DrawVectors,
		CalculatedDrawVectorCount: sqlDrawing.CalculatedDrawVectorCount,
		CreatedAt: formatting.JSONTime(sqlDrawing.CreatedAt),
	}

	if sqlDrawing.LastDrawVectorCalculatedAt.Valid {
		drawing.LastDrawVectorCalculatedAt = formatting.JSONTime(sqlDrawing.LastDrawVectorCalculatedAt.Time)
	} else {
		drawing.LastDrawVectorCalculatedAt = formatting.JSONTime(time.Time{})
	}

	return drawing
}

func (store *SqliteStore) Create(points []types.OriginalPoint) (int) {
	json, _ := json.Marshal(points)

	result := store.db.MustExec(`INSERT INTO drawings (originalPoints) VALUES (?)`, string(json[:]))
	id, _ := result.LastInsertId()

	return int(id)
}

func (store *SqliteStore) AddVectors(drawingId int, vectors string) {
	store.db.MustExec("UPDATE drawings SET vectors = ? WHERE drawingId = ?", vectors, drawingId)
}