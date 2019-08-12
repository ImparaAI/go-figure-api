package test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"

	"api/app/drawing/store"
	"api/app/drawing/types"
	"api/database"
	"api/test/json"
	"api/test/requester"
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
		{X: 4, Y: 5, Time: 0},
		{X: 5, Y: 1, Time: 1},
	}
	id := store.Create(points)

	return `{
		"id": ` + strconv.Itoa(id) + `,
		"svgPath": "M 4 5 L 5 1 "
	}`
}
