package store

import (
	"api/database"
	"api/app/drawing/types"
	"api/app/drawing/store/sqlite"
)

func New() Store {
	var store Store = &sqlite.SqliteStore{database.GetDb()}

	return store
}

type Store interface {
	Exists(id int) (bool)
	Get(id int) (types.Drawing)
	Create(points []types.OriginalPoint, image string) (int)
	AddVectors(drawingId int, vectors []types.DrawVector)
}