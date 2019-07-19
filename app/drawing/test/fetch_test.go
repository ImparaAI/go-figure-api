package test

import (
	"strconv"
	"testing"
	"github.com/stretchr/testify/assert"

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

func TestFetchSuccess(t *testing.T) {
	points := []types.OriginalPoint{
		types.OriginalPoint{X: 4, Y: 5, Time: 0},
		types.OriginalPoint{X: 5, Y: 1, Time: 0.5},
		types.OriginalPoint{X: 2, Y: 3, Time: 1.5},
		types.OriginalPoint{X: 6, Y: 3, Time: 2.1},
	}

	store := store.New()
	id := store.Create(points)
	//drawing := store.Get(id)

	response := requester.Get("/drawing/" + strconv.Itoa(id))
	assert.True(t, response.Ok())
	//assert.Equal(t, `{"id":1,"featured":false,"originalPoints":"[{"x": 4, "y": 5, "time": 0}, {"x": 5, "y": 1, "time": 0}, {"x": 2, "y": 3, "time": 1.5}, {"x": 6, "y": 3, "time": 2.1}]","drawVectors":"[]","calculatedDrawVectorCount":0}`, response.Body())
}