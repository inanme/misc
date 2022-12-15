package maps_test

import (
	"github.com/inanme/misc/ints"
	"github.com/inanme/misc/ptr"
	"testing"

	"github.com/inanme/misc/maps"
	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	aMap := make(map[int]int)
	numbers := ints.FiniteRange(ints.Range{End: ptr.Int(10)})
	for _, key := range numbers {
		aMap[key] = key
	}

	extractedKeys := maps.Keys(aMap)
	extractedValues := maps.Values(aMap)
	for _, n := range numbers {
		assert.Contains(t, extractedKeys, n)
		assert.Contains(t, extractedValues, n)
	}
}
