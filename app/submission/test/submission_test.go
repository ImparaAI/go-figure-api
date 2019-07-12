package submission

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"api/test/requester"
)

func TestFetch(t *testing.T) {
	response := requester.Get("/submission/great")
	assert.True(t, response.Ok())
	assert.Equal(t, response.Body(), "Here is your submission name: great")
}