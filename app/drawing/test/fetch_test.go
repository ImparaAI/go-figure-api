package test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"api/test/requester"
)

func TestNonIntegerId(t *testing.T) {
	nonInts := []string{"foo", "[]", "4.2", "{}"}

	for _, val := range nonInts {
		response := requester.Get("/drawing/" + val)
		assert.True(t, response.IsBadRequest())

		message := "The string '" + val + "' should produce a non-int error."
		assert.Equal(t, "{\"message\":\"The request is not properly formatted.\"}", response.Body(), message)
	}
}