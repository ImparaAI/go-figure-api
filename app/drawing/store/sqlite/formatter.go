package sqlite

import (
	"time"
	"encoding/json"

	"api/app/formatting"
	"api/app/drawing/types"
)

func formatSqlDrawing(sqlDrawing SqlDrawing) types.Drawing {
	drawing := types.Drawing{
		Id: sqlDrawing.Id,
		Featured: sqlDrawing.Featured,
		Image: sqlDrawing.Image,
		CalculatedDrawVectorCount: sqlDrawing.CalculatedDrawVectorCount,
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