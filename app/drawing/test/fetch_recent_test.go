package test

import (
	"time"
	"strconv"
	"testing"
	"github.com/stretchr/testify/assert"

	"api/database"
	"api/test/json"
	"api/test/requester"
	"api/app/drawing/types"
	"api/app/drawing/store"
)

func TestFetchRecent(t *testing.T) {
	database.ClearTestingDb()

	firstDrawingJson := createFetchRecentDrawingJson()
	secondDrawingJson := createFetchRecentDrawingJson()

	response := requester.Get("/drawings/recent")
	assert.True(t, response.Ok())

	expectedJson := json.Compact(`[` + secondDrawingJson + `,` + firstDrawingJson + `]`)

	assert.Equal(t, expectedJson, response.Body())
}

func createFetchRecentDrawingJson() string {
	store := store.New()
	points := []types.OriginalPoint{
		types.OriginalPoint{X: 4, Y: 5, Time: 0},
		types.OriginalPoint{X: 5, Y: 1, Time: 1},
	}
	id := store.Create(points, "data:image/png;image")
	drawing := store.Get(id)
	createdAt := string(time.Time(drawing.CreatedAt).Format("2006-01-02T15:04:05-0700"))

	return `{
		"id": ` + strconv.Itoa(id) + `,
		"featured": false,
		"originalPoints": [
			{"x": 4, "y": 5, "time": 0},
			{"x": 5, "y": 1, "time": 1}
		],
		"image": "data:image/png;image",
		"drawVectors": [],
		"calculatedDrawVectorCount": 0,
		"createdAt": "` + createdAt + `",
		"lastDrawVectorCalculatedAt": null
	}`
}