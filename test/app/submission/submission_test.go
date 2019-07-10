package submission

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"api/test"
)

func TestFetch(t *testing.T) {
	response := test.Get("/submission/great")
	assert.True(t, response.Ok())
	assert.Equal(t, response.Body(), "Here is your submission name: great")
}