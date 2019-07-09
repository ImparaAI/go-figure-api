package submission

import (
	//"net/http"
	//"net/http/httptest"
	"testing"

	"api/test"
	"github.com/stretchr/testify/assert"

)

func TestFetch(t *testing.T) {
	request := test.Get("/submission/great")
	assert.Equal(t, request.Body(), "Here is your submission name: great")
}