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

func TestNonIntegerId(t *testing.T) {
	nonInts := []string{"foo", "[]", "4.2", "{}"}

	for _, val := range nonInts {
		response := requester.Get("/drawing/" + val)
		assert.True(t, response.IsBadRequest())

		message := "The string '" + val + "' should produce a non-int error."
		assert.Equal(t, `{"message":"The request is not properly formatted."}`, response.Body(), message)
	}
}

func TestUnknownId(t *testing.T) {
	database.ClearTestingDb()

	response := requester.Get("/drawing/1")
	assert.True(t, response.IsNotFound())
	assert.Equal(t, `{"message":"This drawing doesn't exist."}`, response.Body())
}

func TestFetchOneSuccess(t *testing.T) {
	points := []types.OriginalPoint{
		types.OriginalPoint{X: 4, Y: 5, Time: 0},
		types.OriginalPoint{X: 5, Y: 1, Time: 0.5},
		types.OriginalPoint{X: 2, Y: 3, Time: 1.5},
		types.OriginalPoint{X: 6, Y: 3, Time: 2.1},
	}

	store := store.New()
	id := store.Create(points)

	response := requester.Get("/drawing/" + strconv.Itoa(id))
	assert.True(t, response.Ok())

	drawing := store.Get(id)
	createdAt := string(time.Time(drawing.CreatedAt).Format("2006-01-02T15:04:05-0700"))

	expectedJson := json.Compact(`{
		"id": 1,
		"featured": false,
		"originalPoints": [
			{"x": 4, "y": 5, "time": 0},
			{"x": 5, "y": 1, "time": 0.5},
			{"x": 2, "y": 3, "time": 1.5},
			{"x": 6, "y": 3, "time": 2.1}
		],
		"drawVectors": [],
		"calculatedDrawVectorCount": 0,
		"createdAt": "` + createdAt + `",
		"lastDrawVectorCalculatedAt": null
	}`)

	assert.Equal(t, expectedJson, response.Body())
}