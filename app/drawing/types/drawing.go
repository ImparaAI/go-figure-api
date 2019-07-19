package types

import (
	"api/app/formatting"
)

type Drawing struct {
	Id int `json:"id"`
	Featured bool `json:"featured"`
	OriginalPoints string `json:"originalPoints"`
	DrawVectors string `json:"drawVectors"`
	CalculatedDrawVectorCount int `json:"calculatedDrawVectorCount"`
	CreatedAt formatting.JSONTime `json:"createdAt"`
	LastDrawVectorCalculatedAt formatting.JSONTime `json:"lastDrawVectorCalculatedAt"`
}