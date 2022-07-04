package google_datastore

import (
	"time"
)

type DatastoreDrawing struct {
	Id                         int
	Featured                   bool
	OriginalPoints             string    `datastore:",noindex`
	DrawVectors                string    `datastore:",noindex`
	CalculatedDrawVectorCount  int       `datastore:",noindex`
	CreatedAt                  time.Time
	LastDrawVectorCalculatedAt time.Time `datastore:",noindex`
}
