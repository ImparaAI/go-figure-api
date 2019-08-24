package mysql

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"api/app/drawing/types"
	"api/app/formatting"
)

func formatSqlDrawing(sqlDrawing SqlDrawing) types.Drawing {
	drawing := types.Drawing{
		Id:        sqlDrawing.Id,
		Featured:  sqlDrawing.Featured,
		CreatedAt: formatting.JSONTime(sqlDrawing.CreatedAt),
	}

	if sqlDrawing.LastDrawVectorCalculatedAt.Valid {
		drawing.LastDrawVectorCalculatedAt = formatting.JSONTime(sqlDrawing.LastDrawVectorCalculatedAt.Time)
	} else {
		drawing.LastDrawVectorCalculatedAt = formatting.JSONTime(time.Time{})
	}

	json.Unmarshal([]byte(sqlDrawing.OriginalPoints), &drawing.OriginalPoints)
	json.Unmarshal([]byte(sqlDrawing.DrawVectors), &drawing.DrawVectors)

	return drawing
}

func formatSqlDrawingPreviews(sqlDrawings []SqlDrawing) []types.DrawingPreview {
	drawings := []types.DrawingPreview{}

	for _, sqlDrawing := range sqlDrawings {
		drawings = append(drawings, formatSqlDrawingPreview(sqlDrawing))
	}

	return drawings
}

func formatSqlDrawingPreview(sqlDrawing SqlDrawing) types.DrawingPreview {
	drawingPreview := types.DrawingPreview{
		Id:      sqlDrawing.Id,
		SvgPath: buildSvgPath(sqlDrawing),
	}

	return drawingPreview
}

func buildSvgPath(sqlDrawing SqlDrawing) string {
	var path strings.Builder
	var originalPoints []types.OriginalPoint

	path.Grow(len(originalPoints) * 10)
	json.Unmarshal([]byte(sqlDrawing.OriginalPoints), &originalPoints)

	for i, point := range originalPoints {
		if i == 0 {
			fmt.Fprintf(&path, "M %d %d ", point.X, point.Y)
		} else {
			fmt.Fprintf(&path, "L %d %d ", point.X, point.Y)
		}
	}

	return path.String()
}
