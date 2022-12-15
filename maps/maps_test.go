package maps_test

import (
	"fmt"
	"github.com/inanme/misc/maps"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestKeys(t *testing.T) {
	aMap := make(map[int]string)
	keys := []int{1, 2, 3, 4}
	for _, key := range keys {
		aMap[key] = fmt.Sprintf("%d", key)
	}

	extractedKeys := maps.Keys(aMap)
	sort.Ints(extractedKeys)
	assert.Equal(t, extractedKeys, keys)
}
