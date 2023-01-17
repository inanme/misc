package morestrings_test

import (
	"testing"

	"github.com/inanme/misc/morestrings"
	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	assert.Equal(t, "", morestrings.Join(nil, ",", "(", ")"))
	assert.Equal(t, "(1)", morestrings.Join([]string{"1"}, ",", "(", ")"))
	assert.Equal(t, "(1,2)", morestrings.Join([]string{"1", "2"}, ",", "(", ")"))
}
