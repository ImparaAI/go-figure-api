package store

import (
	"api/database"
	"api/app/drawing/types"
)

func New() Store {
	var store Store = &SqliteStore{database.GetDb()}

	return store
}

type Store interface {
	Exists(id int) (bool)
	Get(id int) (types.Drawing)
	Create(points []types.OriginalPoint) (int)
	AddVectors(drawingId int, vectors string)
}