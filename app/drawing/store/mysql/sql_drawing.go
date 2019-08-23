package mysql

import (
	"api/app/formatting"
	"time"
)

type SqlDrawing struct {
	Id                         int                    `db:"id"`
	Featured                   bool                   `db:"featured"`
	OriginalPoints             string                 `db:"originalPoints"`
	DrawVectors                string                 `db:"drawVectors"`
	CreatedAt                  time.Time              `db:"createdAt"`
	LastDrawVectorCalculatedAt formatting.SQLNullTime `db:"lastDrawVectorCalculatedAt"`
}
