package types

import (
	//"time"
	//"api/app/formatting"
)

type Submission struct {
	Id int `db:"id" json:"id"`
	RequestedDrawVectorCount int `db:"requestedDrawVectorCount" json:"requestedDrawVectorCount"`
	Featured bool `db:"featured" json:"featured"`
	OriginalPoints string `db:"originalPoints" json:"originalPoints"`
	DrawVectors string `db:"drawVectors" json:"drawVectors"`
	CalculatedDrawVectorCount int `db:"calculatedDrawVectorCount" json:"calculatedDrawVectorCount"`
	CreatedAt string `db:"createdAt" json:"createdAt"`
	LastDrawVectorCalculatedAt string `db:"lastDrawVectorCalculatedAt" json:"lastDrawVectorCalculatedAt"`
}