package processing

import (
	"api/app/drawing/store"
	"api/app/drawing/types"
)

func buildOriginalPoints(drawingId int) []types.OriginalPoint {
	drawing := getDrawing(drawingId)
	normalizeTime(drawing.OriginalPoints)

	return drawing.OriginalPoints
}

func getDrawing(drawingId int) types.Drawing {
	store := store.New()
	drawing := store.Get(drawingId)

	if drawing.Id == 0 {
		panic("The drawing could not be found in storage.")
	}

	return drawing
}

func normalizeTime(originalPoints []types.OriginalPoint) {
	finalPoint := originalPoints[len(originalPoints) - 1]

	if finalPoint.Time == 0 {
		return
	}

	for i := 0; i < len(originalPoints); i++ {
		originalPoints[i].Time = originalPoints[i].Time / finalPoint.Time
	}
}