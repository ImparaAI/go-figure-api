package google_datastore

import (
	"time"
	"context"
	"encoding/json"
	"cloud.google.com/go/datastore"

	"api/app/drawing/types"
)

type GoogleDatastoreStore struct {
	Ctx    context.Context
	Client *datastore.Client
}

func (store *GoogleDatastoreStore) Exists(id int64) bool {
	var drawing DatastoreDrawing
	key := datastore.IDKey("Drawing", id, nil)

	if err := store.Client.Get(store.Ctx, key, &drawing); err != nil {
		return false
	}

	return true
}

func (store *GoogleDatastoreStore) Get(id int64) types.Drawing {
	var drawing DatastoreDrawing
	key := datastore.IDKey("Drawing", id, nil)

	if err := store.Client.Get(store.Ctx, key, &drawing); err != nil {
		panic(err)
	}

	drawing.Id = id

	return formatDatastoreDrawing(drawing)
}

func (store *GoogleDatastoreStore) GetRecent() []types.DrawingPreview {
	var datastoreDrawings []*DatastoreDrawing

	query := datastore.NewQuery("Drawing").Order("-created_at").Limit(20)
	keys, err := store.Client.GetAll(store.Ctx, query, &datastoreDrawings)
	if err != nil {
		panic(err)
	}

	// Set the id field on each Task from the corresponding key.
	for i, key := range keys {
		datastoreDrawings[i].Id = key.ID
	}

	return formatDatastoreDrawingPreviews(datastoreDrawings)
}

func (store *GoogleDatastoreStore) Create(points []types.OriginalPoint) int64 {
	json, _ := json.Marshal(points)
	drawing := &DatastoreDrawing{
		OriginalPoints: string(json[:]),
		DrawVectors:    "[]",
		CreatedAt:      time.Now(),
	}

  key := datastore.IncompleteKey("Drawing", nil)
  resultKey, err := store.Client.Put(store.Ctx, key, drawing)

  if err != nil {
		panic(err)
  }

  return resultKey.ID
}

func (store *GoogleDatastoreStore) AddVectors(drawingId int64, vectors []types.DrawVector) {
	key := datastore.IDKey("Drawing", drawingId, nil)

	_, err := store.Client.RunInTransaction(store.Ctx, func(tx *datastore.Transaction) error {
		var drawing DatastoreDrawing
		if err := tx.Get(key, &drawing); err != nil {
			return err
		}
		json, _ := json.Marshal(vectors)
		drawing.DrawVectors = string(json)
		_, err := tx.Put(key, &drawing)
		return err
	})


	if err != nil {
		panic(err)
	}
}
