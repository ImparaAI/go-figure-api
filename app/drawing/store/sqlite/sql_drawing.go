package sqlite

import (
	"time"
	"api/app/formatting"
)

type SqlDrawing struct {
	Id int `db:"id"`
	Featured bool `db:"featured"`
	OriginalPoints string `db:"originalPoints"`
	Image string `db:"image"`
	DrawVectors string `db:"drawVectors"`
	CalculatedDrawVectorCount int `db:"calculatedDrawVectorCount"`
	CreatedAt time.Time `db:"createdAt"`
	LastDrawVectorCalculatedAt formatting.SQLNullTime `db:"lastDrawVectorCalculatedAt"`
}