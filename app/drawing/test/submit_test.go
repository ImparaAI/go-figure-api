package test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"api/test/requester"
)

func TestNonJsonInput(t *testing.T) {
	nonInts := []string{"foo", "[]", "4.2", ""}

	for _, val := range nonInts {
		response := requester.Post("/drawing", val)
		assert.True(t, response.IsBadRequest())
		assert.Equal(t, "{\"message\":\"The request is not properly formatted.\"}", response.Body())
	}
}

func TestMissingJsonField(t *testing.T) {

}