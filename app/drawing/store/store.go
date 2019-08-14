package store

import (
	"api/app/drawing/store/mysql"
	"api/app/drawing/types"
	"api/database"
)

func New() Store {
	var store Store = &mysql.MySqlStore{database.GetDb()}

	return store
}

type Store interface {
	Exists(id int) bool
	Get(id int) types.Drawing
	GetRecent() []types.DrawingPreview
	Create(points []types.OriginalPoint) int
	AddVectors(drawingId int, vectors []types.DrawVector)
}
