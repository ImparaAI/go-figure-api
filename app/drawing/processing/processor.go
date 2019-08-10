package processing

import (
	"api/app/drawing/processing/draw_vector"
	"api/app/drawing/store"
	"api/app/drawing/types"
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
