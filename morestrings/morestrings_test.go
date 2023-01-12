package morestrings_test

import (
	"github.com/inanme/misc/morestrings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoin(t *testing.T) {
	assert.Equal(t, "", morestrings.Join(nil, ",", "(", ")"))
	assert.Equal(t, "(1)", morestrings.Join([]string{"1"}, ",", "(", ")"))
	assert.Equal(t, "(1,2)", morestrings.Join([]string{"1", "2"}, ",", "(", ")"))
}
