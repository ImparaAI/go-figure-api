package sqlite

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"

	"api/app/drawing/types"
)

type SqliteStore struct {
	DB *sqlx.DB
}

func (store *SqliteStore) Exists(id int) bool {
	var count int
	err := store.DB.Get(&count, "SELECT COUNT(id) FROM drawings WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	return count > 0
}

func (store *SqliteStore) Get(id int) types.Drawing {
	var sqlDrawing SqlDrawing

	err := store.DB.Get(&sqlDrawing, "SELECT * FROM drawings WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	return formatSqlDrawing(sqlDrawing)
}

func (store *SqliteStore) GetRecent() []types.DrawingPreview {
	var sqlDrawings []SqlDrawing

	err := store.DB.Select(&sqlDrawings, "SELECT id, originalPoints FROM drawings ORDER BY id DESC LIMIT 20")

	if err != nil {
		panic(err)
	}

	return formatSqlDrawingPreviews(sqlDrawings)
}

func (store *SqliteStore) Create(points []types.OriginalPoint) int {
	json, _ := json.Marshal(points)

	result := store.DB.MustExec(`INSERT INTO drawings (originalPoints) VALUES (?)`, string(json[:]))
	id, _ := result.LastInsertId()

	return int(id)
}

func (store *SqliteStore) AddVectors(drawingId int, vectors []types.DrawVector) {
	json, _ := json.Marshal(vectors)

	store.DB.MustExec("UPDATE drawings SET drawVectors = ? WHERE id = ?", string(json), drawingId)
}