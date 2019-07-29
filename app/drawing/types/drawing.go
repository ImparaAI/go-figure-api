package types

import (
	"api/app/formatting"
)

type Drawing struct {
	Id int `json:"id"`
	Featured bool `json:"featured"`
	OriginalPoints []OriginalPoint `json:"originalPoints"`
	Image string `json:"image"`
	DrawVectors []DrawVector `json:"drawVectors"`
	CalculatedDrawVectorCount int `json:"calculatedDrawVectorCount"`
	CreatedAt formatting.JSONTime `json:"createdAt"`
	LastDrawVectorCalculatedAt formatting.JSONTime `json:"lastDrawVectorCalculatedAt"`
}