package health

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"api/test/requester"
)

func TestHealth(t *testing.T) {
	response := requester.Get("/health")
	assert.True(t, response.Ok())
	assert.Equal(t, "OK", response.Body())
}
