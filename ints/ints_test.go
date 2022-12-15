package ints_test

import (
	"github.com/inanme/misc/ints"
	"github.com/inanme/misc/ptr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFiniteRange(t *testing.T) {
	assert.Equal(t, []int{},
		ints.FiniteRange(ints.Range{}))
	assert.Equal(t, []int{0, 1, 2},
		ints.FiniteRange(ints.Range{End: ptr.Int(3)}))
	assert.Equal(t, []int{0, 1, 2, 3},
		ints.FiniteRange(ints.Range{End: ptr.Int(3), IsInclusive: true}))
	assert.Equal(t, []int{1, 2, 3},
		ints.FiniteRange(ints.Range{Start: ptr.Int(1), End: ptr.Int(3), IsInclusive: true}))
	assert.Equal(t, []int{1, 3, 5, 7, 9},
		ints.FiniteRange(ints.Range{Start: ptr.Int(1), End: ptr.Int(10), Step: ptr.Int(2), IsInclusive: true}))
	assert.Equal(t, []int{1, 4, 7, 10},
		ints.FiniteRange(ints.Range{Start: ptr.Int(1), End: ptr.Int(10), Step: ptr.Int(3), IsInclusive: true}))
}
