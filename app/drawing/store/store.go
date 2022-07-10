package store

import (
	"context"
	"os"

	"api/app/drawing/store/mysql"
	"api/app/drawing/store/google_datastore"
	"api/app/drawing/types"
	"api/database"
)

func New(ctx context.Context) Store {
	var store Store

	if os.Getenv("DB_STORE") == "google_datastore" {
		store = &google_datastore.GoogleDatastoreStore{
			Ctx:    ctx,
			Client: database.GetDatastore(),
		}
	} else {
		store = &mysql.MySqlStore{database.GetDb()}
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
