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

type person struct {
	name string
}

func (p person) String() string {
	return p.name
}

func TestJoinT(t *testing.T) {
	assert.Equal(t, "", morestrings.JoinT[person](nil, ",", "(", ")"))
	assert.Equal(t, "(1)", morestrings.JoinT([]person{{name: "1"}}, ",", "(", ")"))
	assert.Equal(t, "(1,2)", morestrings.JoinT([]person{{name: "1"}, {name: "2"}}, ",", "(", ")"))
}

func TestContainsAll(t *testing.T) {
	assert.True(t, morestrings.ContainsAll("", ""))
	assert.True(t, morestrings.ContainsAll("1", "1"))
	assert.True(t, morestrings.ContainsAll("12", "1", "2"))
	assert.False(t, morestrings.ContainsAll("", "1", "2"))
	assert.False(t, morestrings.ContainsAll("1", "1", "2"))
	assert.False(t, morestrings.ContainsAll("2", "1", "2"))
}

func TestContainsAllCaseInsensitive(t *testing.T) {
	assert.True(t, morestrings.ContainsAllCaseInsensitive("", ""))
	assert.True(t, morestrings.ContainsAllCaseInsensitive("1", "1"))
	assert.True(t, morestrings.ContainsAllCaseInsensitive("12", "1", "2"))
	assert.False(t, morestrings.ContainsAllCaseInsensitive("", "1", "2"))
	assert.False(t, morestrings.ContainsAllCaseInsensitive("1", "1", "2"))
	assert.False(t, morestrings.ContainsAllCaseInsensitive("2", "1", "2"))
	assert.True(t, morestrings.ContainsAllCaseInsensitive("AB", "A", "B"))
	assert.True(t, morestrings.ContainsAllCaseInsensitive("AB", "a", "b"))
	assert.False(t, morestrings.ContainsAllCaseInsensitive("", "1", "2"))
	assert.False(t, morestrings.ContainsAllCaseInsensitive("A", "A", "B"))
	assert.False(t, morestrings.ContainsAllCaseInsensitive("A", "a", "B"))
	assert.False(t, morestrings.ContainsAllCaseInsensitive("A", "A", "B"))
	assert.False(t, morestrings.ContainsAllCaseInsensitive("A", "a", "b"))
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

func TestContainAllCaseInsensitive(t *testing.T) {
	assert.True(t, morestrings.ContainAllCaseInsensitive([]string{""}, ""))
	assert.True(t, morestrings.ContainAllCaseInsensitive([]string{"1"}, "1"))
	assert.True(t, morestrings.ContainAllCaseInsensitive([]string{"1", "1"}, "1"))
	assert.True(t, morestrings.ContainAllCaseInsensitive([]string{"12"}, "1", "2"))
	assert.False(t, morestrings.ContainAllCaseInsensitive([]string{""}, "1", "2"))
	assert.False(t, morestrings.ContainAllCaseInsensitive([]string{"1"}, "1", "2"))
	assert.False(t, morestrings.ContainAllCaseInsensitive([]string{"2"}, "1", "2"))
	assert.False(t, morestrings.ContainAllCaseInsensitive([]string{"1", "2"}, "1", "2"))
	assert.False(t, morestrings.ContainAllCaseInsensitive([]string{"1", "2"}, "1", "2"))
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

func TestContainsAnyCaseInsensitive(t *testing.T) {
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("", ""))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("1", "1"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("12", "1"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("12", "2"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("12", "1", "2"))
	assert.False(t, morestrings.ContainsAnyCaseInsensitive("", "1", "2"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("1", "1", "2"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("2", "1", "2"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("AB", "A", "B"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("AB", "a", "b"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("AB", "A"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("AB", "a"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("AB", "B"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("AB", "b"))
	assert.False(t, morestrings.ContainsAnyCaseInsensitive("", "1", "2"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("A", "A", "B"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("A", "a", "B"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("A", "A", "B"))
	assert.True(t, morestrings.ContainsAnyCaseInsensitive("A", "a", "b"))
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

func TestContainAnyCaseInsensitive(t *testing.T) {
	assert.True(t, morestrings.ContainAnyCaseInsensitive([]string{""}, ""))
	assert.True(t, morestrings.ContainAnyCaseInsensitive([]string{"1"}, "1"))
	assert.True(t, morestrings.ContainAnyCaseInsensitive([]string{"12"}, "1"))
	assert.True(t, morestrings.ContainAnyCaseInsensitive([]string{"12"}, "2"))
	assert.True(t, morestrings.ContainAnyCaseInsensitive([]string{"12"}, "1", "2"))
	assert.False(t, morestrings.ContainAnyCaseInsensitive([]string{""}, "1", "2"))
	assert.True(t, morestrings.ContainAnyCaseInsensitive([]string{"1"}, "1", "2"))
	assert.True(t, morestrings.ContainAnyCaseInsensitive([]string{"2"}, "1", "2"))
}
