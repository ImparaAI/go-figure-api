package store

import (
	"os"

	"api/app/drawing/store/mysql"
	"api/app/drawing/store/google_datastore"
	"api/app/drawing/types"
	"api/database"
)

func New() Store {
	if os.Getenv("DB_STORE") == "google_datastore" {
		var store Store = &google_datastore.GoogleDatastoreStore{database.GetDb()}
	} else {
		var store Store = &mysql.MySqlStore{database.GetDb()}
	}

	return store
}

type Store interface {
	Exists(id int64) bool
	Get(id int64) types.Drawing
	GetRecent() []types.DrawingPreview
	Create(points []types.OriginalPoint) int64
	AddVectors(drawingId int64, vectors []types.DrawVector)
}
