package google_datastore

import (
	"time"
)

type DatastoreDrawing struct {
	Id                         int64
	Featured                   bool      `datastore:"featured`
	OriginalPoints             string    `datastore:"original_points,noindex`
	DrawVectors                string    `datastore:"draw_vectors,noindex`
	CreatedAt                  time.Time `datastore:"created_at`
	LastDrawVectorCalculatedAt time.Time `datastore:"last_draw_vector_calculated_at,noindex`
}
