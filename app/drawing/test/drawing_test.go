package test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"api/test/requester"
)

func TestFetchBadInput(t *testing.T) {
	nonInts := []string{"foo", "[]", "4.2", "{}"}

	for _, val := range nonInts {
		response := requester.Get("/drawing/" + val)
		assert.True(t, response.IsBadRequest())
		assert.Equal(t, "{\"message\":\"The request is not properly formatted.\"}", response.Body())
	}
}

/*func TestFetchBadInput(t *testing.T) {
	response := requester.Get("/drawing/great")
	assert.True(t, response.Ok())
	assert.Equal(t, response.Body(), "Here is your drawing name: great")
}*/