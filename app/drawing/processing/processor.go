package processing

import (
	"api/app/drawing/store"
	"api/app/drawing/types"
)

func Process(drawingId int) {
	originalPoints := BuildOriginalPoints(drawingId)
	vectors := buildDrawVectors(originalPoints)

	saveDrawVectors(drawingId, vectors)
}

func saveDrawVectors(drawingId int, vectors []types.DrawVector) {
	store := store.New()
	store.AddVectors(drawingId, vectors)
}