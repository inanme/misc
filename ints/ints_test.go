package ints_test

import (
	"testing"

	"github.com/inanme/misc/ints"
	"github.com/inanme/misc/ptr"
	"github.com/stretchr/testify/assert"
)

func TestFiniteRange(t *testing.T) {
	equal := func(expected []int, range0 ints.Range, ranges ...ints.Range) {
		assert.Equal(t, expected, ints.FiniteRange(range0, ranges...))
	}
	equal([]int{}, ints.Range{})
	equal([]int{}, ints.Range{}, ints.Range{})
	equal([]int{0, 1, 2},
		ints.Range{End: ptr.Int(3)})
	equal([]int{0, 1, 2, 0, 1, 2},
		ints.Range{End: ptr.Int(3)},
		ints.Range{End: ptr.Int(3)})
	equal([]int{0, 1, 2, 3},
		ints.Range{End: ptr.Int(3), IsInclusive: true})
	equal([]int{0, 1, 2, 3, 0, 1, 2, 3},
		ints.Range{End: ptr.Int(3), IsInclusive: true},
		ints.Range{End: ptr.Int(3), IsInclusive: true})
	equal([]int{1, 2, 3},
		ints.Range{Start: ptr.Int(1), End: ptr.Int(3), IsInclusive: true})
	equal([]int{1, 3, 5, 7, 9},
		ints.Range{Start: ptr.Int(1), End: ptr.Int(10), Step: ptr.Int(2), IsInclusive: true})
	equal([]int{1, 4, 7, 10},
		ints.Range{Start: ptr.Int(1), End: ptr.Int(10), Step: ptr.Int(3), IsInclusive: true})
}
