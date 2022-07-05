package google_datastore

import (
	"encoding/json"
	"fmt"
	"strings"

	"api/app/drawing/types"
	"api/app/formatting"
)

func formatDatastoreDrawing(datastoreDrawing DatastoreDrawing) types.Drawing {
	drawing := types.Drawing{
		Id:                         datastoreDrawing.Id,
		Featured:                   datastoreDrawing.Featured,
		CalculatedDrawVectorCount:  datastoreDrawing.CalculatedDrawVectorCount,
		CreatedAt:                  formatting.JSONTime(datastoreDrawing.CreatedAt),
		LastDrawVectorCalculatedAt: formatting.JSONTime(datastoreDrawing.LastDrawVectorCalculatedAt),
	}

	json.Unmarshal([]byte(datastoreDrawing.OriginalPoints), &drawing.OriginalPoints)
	json.Unmarshal([]byte(datastoreDrawing.DrawVectors), &drawing.DrawVectors)

	return drawing
}

func formatDatastoreDrawingPreviews(datastoreDrawings []*DatastoreDrawing) []types.DrawingPreview {
	drawings := []types.DrawingPreview{}

	for _, datastoreDrawing := range datastoreDrawings {
		drawings = append(drawings, formatDatastoreDrawingPreview(datastoreDrawing))
	}

	return drawings
}

func formatDatastoreDrawingPreview(datastoreDrawing *DatastoreDrawing) types.DrawingPreview {
	drawingPreview := types.DrawingPreview{
		Id:      datastoreDrawing.Id,
		SvgPath: buildSvgPath(datastoreDrawing),
	}

	return drawingPreview
}

func buildSvgPath(datastoreDrawing *DatastoreDrawing) string {
	var path strings.Builder
	var originalPoints []types.OriginalPoint

	path.Grow(len(originalPoints) * 10)
	json.Unmarshal([]byte(datastoreDrawing.OriginalPoints), &originalPoints)

	for i, point := range originalPoints {
		if i == 0 {
			fmt.Fprintf(&path, "M %d %d ", point.X, point.Y)
		} else {
			fmt.Fprintf(&path, "L %d %d ", point.X, point.Y)
		}
	}

	return path.String()
}
