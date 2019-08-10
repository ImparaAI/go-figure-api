package processing

import (
	"api/app/drawing/store"
	"api/app/drawing/types"
	"api/app/drawing/processing/draw_vector"
)

func Process(drawingId int) {
	var originalPointsFactory OriginalPointsFactory = OriginalPointsFactory{}

	originalPoints := originalPointsFactory.Build(drawingId)
	vectors := draw_vector.BuildSeries(originalPoints)

	saveDrawVectors(drawingId, vectors)
}

func saveDrawVectors(drawingId int, vectors []types.DrawVector) {
	store := store.New()
	store.AddVectors(drawingId, vectors)
}
