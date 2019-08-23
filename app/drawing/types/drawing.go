package types

import (
	"api/app/formatting"
)

type Drawing struct {
	Id                         int                 `json:"id"`
	Featured                   bool                `json:"featured"`
	OriginalPoints             []OriginalPoint     `json:"originalPoints"`
	DrawVectors                []DrawVector        `json:"drawVectors"`
	CreatedAt                  formatting.JSONTime `json:"createdAt"`
	LastDrawVectorCalculatedAt formatting.JSONTime `json:"lastDrawVectorCalculatedAt"`
}
