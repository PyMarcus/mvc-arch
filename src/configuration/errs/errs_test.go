package errs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrs(t *testing.T) {
	causes := []*Causes{
		{Field: "field1", Message: "error message 1"},
		{Field: "field2", Message: "error message 2"},
	}
	err := NewErrs("error message", "some_error", 123, causes)

	assert.NotNil(t, err)
	assert.Equal(t, err.Code, int(123))
	assert.Equal(t, len(causes), 2)
	assert.Equal(t, err.Err, "some_error")

	message := err.Error()
	assert.Equal(t, message, "error message")
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("bad request")
	assert.Equal(t, err.Code, int(400))
	assert.Equal(t, err.Message, "bad request")
}
