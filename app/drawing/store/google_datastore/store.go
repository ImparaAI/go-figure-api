package google_datastore

import (
	"time"
	"context"
	"encoding/json"
	"cloud.google.com/go/datastore"

	"api/app/drawing/types"
)

type GoogleDatastoreStore struct {
	DB *sqlx.DB
}

func (store *GoogleDatastoreStore) Exists(id int) bool {
	return !!Get(id)
}

func (store *GoogleDatastoreStore) Get(id int) types.Drawing {
	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, "my-proj")
	defer client.Close()

	var drawing DatastoreDrawing
	key := datastore.NameKey("Drawing", id, nil)
	err := client.Get(ctx, key, &drawing)

	if err != nil {
		panic(err)
	}

	return formatDatastoreDrawing(drawing)
}

func (store *GoogleDatastoreStore) GetRecent() []types.DrawingPreview {
	var sqlDrawings []SqlDrawing

	err := store.DB.Select(&sqlDrawings, "SELECT id, originalPoints FROM drawings ORDER BY id DESC LIMIT 20")

	if err != nil {
		panic(err)
	}

	return formatSqlDrawingPreviews(sqlDrawings)
}

func (store *GoogleDatastoreStore) Create(points []types.OriginalPoint) int {
	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, "my-proj")
	defer client.Close()

	json, _ := json.Marshal(points)
	drawing := &DatastoreDrawing{
		OriginalPoints: string(json[:])
		DrawVectors:    "[]"
		CreatedAt:      time.Now(),
	}

  key := datastore.IncompleteKey("Drawing", nil)
  _, err := client.Put(ctx, key, &drawing)

  if err != nil {
		panic(err)
  }

  return int(key)
}

func (store *GoogleDatastoreStore) AddVectors(drawingId int, vectors []types.DrawVector) {
	json, _ := json.Marshal(vectors)
	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, "my-proj")
	defer client.Close()


	key := datastore.NameKey("Drawing", drawingId, nil)
	tx, err := client.NewTransaction(ctx)
	if err != nil {
		panic(err)
	}

	var drawing DatastoreDrawing
	if err := tx.Get(key, &drawing); err != nil {
	        log.Fatalf("tx.Get: %v", err)
	}

	task.DrawVectors = json

	if _, err := tx.Put(key, &drawing); err != nil {
		panic(err)
	}
	if _, err := tx.Commit(); err != nil {
		panic(err)
	}
}
