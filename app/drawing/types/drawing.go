package types

import (
	//"time"
	"api/app/formatting"
)

type Drawing struct {
	Id int `db:"id" json:"id"`
	Featured bool `db:"featured" json:"featured"`
	OriginalPoints string `db:"originalPoints" json:"originalPoints"`
	DrawVectors string `db:"drawVectors" json:"drawVectors"`
	CalculatedDrawVectorCount int `db:"calculatedDrawVectorCount" json:"calculatedDrawVectorCount"`
	CreatedAt formatting.JSONTime `db:"createdAt" json:"createdAt"`
	LastDrawVectorCalculatedAt formatting.JSONTime `db:"lastDrawVectorCalculatedAt" json:"lastDrawVectorCalculatedAt"`
}