package test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"api/database"
	"api/test/requester"
)

func TestMain(m *testing.M) {
	database.SetTestingEnvironment()
	err := database.Initialize()

	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestNonJsonInput(t *testing.T) {
	nonInts := []string{"foo", "[]", "4.2", ""}

	for _, val := range nonInts {
		response := requester.Post("/drawing", val)
		assert.True(t, response.IsBadRequest())
		assert.Equal(t, `{"message":"The request is not properly formatted."}`, response.Body())
	}
}

func TestMissingPointsField(t *testing.T) {
	response := requester.Post("/drawing", "{}")
	assert.True(t, response.IsBadRequest())
	assert.Equal(t, `{"message":"The request is not properly formatted."}`, response.Body())
}

func TestEmptyPointsField(t *testing.T) {
	response := requester.Post("/drawing", `{"points":[]}`)
	assert.True(t, response.IsBadRequest())
	assert.Equal(t, `{"message":"There needs to be at least 1 point."}`, response.Body())
}

func TestInvalidPoints(t *testing.T) {
	nonInts := []string{`"foo"`, "[]", "4.2"}

	for _, val := range nonInts {
		response := requester.Post("/drawing", `{"points":[`+val+`]}`)
		assert.True(t, response.IsBadRequest())
		assert.Equal(t, `{"message":"The request is not properly formatted."}`, response.Body())
	}
}

func TestFirstPointTimeNonZero(t *testing.T) {
	json := `{"points":[{"x": 4, "y": 5, "time": 0.5}, {"x": 5, "y": 1, "time": 1}]}`
	response := requester.Post("/drawing", json)
	assert.True(t, response.IsBadRequest())
	assert.Equal(t, `{"message":"The first point's time must be zero."}`, response.Body())
}

func TestNonSequentialTimes(t *testing.T) {
	json := `{"points":[{"x": 4, "y": 5, "time": 0}, {"x": 5, "y": 1, "time": 0}, {"x": 2, "y": 3, "time": 1.5}, {"x": 6, "y": 3, "time": 1.1}]}`
	response := requester.Post("/drawing", json)
	assert.True(t, response.IsBadRequest())
	assert.Equal(t, `{"message":"Each point's time should be equal to or greater than the previous point."}`, response.Body())
}

func TestSubmitSuccess(t *testing.T) {
	database.ClearTestingDb()
	json := `{"points":[{"x": 4, "y": 5, "time": 0}, {"x": 5, "y": 1, "time": 0}, {"x": 2, "y": 3, "time": 1.5}, {"x": 6, "y": 3, "time": 2.1}]}`
	response := requester.Post("/drawing", json)
	assert.True(t, response.Ok())
	assert.Equal(t, `{"id":1}`, response.Body())
}
