package processing

import (
	"context"

	"api/app/drawing/processing/draw_vector"
	"api/app/drawing/store"
	"api/app/drawing/types"
)

func Process(drawingId int64) {
	var originalPointsFactory OriginalPointsFactory = OriginalPointsFactory{}

	originalPoints := originalPointsFactory.Build(drawingId)
	vectors := draw_vector.BuildSeries(originalPoints)

	saveDrawVectors(drawingId, vectors)
}

func saveDrawVectors(drawingId int64, vectors []types.DrawVector) {
	ctx := context.Background()
	store := store.New(ctx)
	store.AddVectors(drawingId, vectors)
}
