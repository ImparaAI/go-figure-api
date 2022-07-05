package types

import (
	"api/app/formatting"
)

type Drawing struct {
	Id                         int64               `json:"id"`
	Featured                   bool                `json:"featured"`
	OriginalPoints             []OriginalPoint     `json:"originalPoints"`
	DrawVectors                []DrawVector        `json:"drawVectors"`
	CalculatedDrawVectorCount  int                 `json:"calculatedDrawVectorCount"`
	CreatedAt                  formatting.JSONTime `json:"createdAt"`
	LastDrawVectorCalculatedAt formatting.JSONTime `json:"lastDrawVectorCalculatedAt"`
}
