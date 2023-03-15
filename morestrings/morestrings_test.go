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

func TestContainsAll(t *testing.T) {
	assert.True(t, morestrings.ContainsAll("", ""))
	assert.True(t, morestrings.ContainsAll("1", "1"))
	assert.True(t, morestrings.ContainsAll("12", "1", "2"))
	assert.False(t, morestrings.ContainsAll("", "1", "2"))
	assert.False(t, morestrings.ContainsAll("1", "1", "2"))
	assert.False(t, morestrings.ContainsAll("2", "1", "2"))
}

func TestContainAll(t *testing.T) {
	assert.True(t, morestrings.ContainAll([]string{""}, ""))
	assert.True(t, morestrings.ContainAll([]string{"1"}, "1"))
	assert.True(t, morestrings.ContainAll([]string{"1", "1"}, "1"))
	assert.True(t, morestrings.ContainAll([]string{"12"}, "1", "2"))
	assert.False(t, morestrings.ContainAll([]string{""}, "1", "2"))
	assert.False(t, morestrings.ContainAll([]string{"1"}, "1", "2"))
	assert.False(t, morestrings.ContainAll([]string{"2"}, "1", "2"))
	assert.False(t, morestrings.ContainAll([]string{"1", "2"}, "1", "2"))
	assert.False(t, morestrings.ContainAll([]string{"1", "2"}, "1", "2"))
}

func TestContainsAny(t *testing.T) {
	assert.True(t, morestrings.ContainsAny("", ""))
	assert.True(t, morestrings.ContainsAny("1", "1"))
	assert.True(t, morestrings.ContainsAny("12", "1"))
	assert.True(t, morestrings.ContainsAny("12", "2"))
	assert.True(t, morestrings.ContainsAny("12", "1", "2"))
	assert.False(t, morestrings.ContainsAny("", "1", "2"))
	assert.True(t, morestrings.ContainsAny("1", "1", "2"))
	assert.True(t, morestrings.ContainsAny("2", "1", "2"))
}

func TestContainAny(t *testing.T) {
	assert.True(t, morestrings.ContainAny([]string{""}, ""))
	assert.True(t, morestrings.ContainAny([]string{"1"}, "1"))
	assert.True(t, morestrings.ContainAny([]string{"12"}, "1"))
	assert.True(t, morestrings.ContainAny([]string{"12"}, "2"))
	assert.True(t, morestrings.ContainAny([]string{"12"}, "1", "2"))
	assert.False(t, morestrings.ContainAny([]string{""}, "1", "2"))
	assert.True(t, morestrings.ContainAny([]string{"1"}, "1", "2"))
	assert.True(t, morestrings.ContainAny([]string{"2"}, "1", "2"))
}
