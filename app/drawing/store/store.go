package store

import (
	"api/app/drawing/store/sqlite"
	"api/app/drawing/types"
	"api/database"
)

func New() Store {
	var store Store = &sqlite.SqliteStore{database.GetDb()}

	return store
}

type Store interface {
	Exists(id int) bool
	Get(id int) types.Drawing
	GetRecent() []types.DrawingPreview
	Create(points []types.OriginalPoint) int
	AddVectors(drawingId int, vectors []types.DrawVector)
}
