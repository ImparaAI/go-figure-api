package submission

import (
	"testing"

	"api/test"
	"github.com/stretchr/testify/assert"

)

func TestFetch(t *testing.T) {
	response := test.Get("/submission/great")
	assert.True(t, response.Ok())
	assert.Equal(t, response.Body(), "Here is your submission name: great")
}