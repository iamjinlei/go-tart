package tart

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCBuf(t *testing.T) {
	c := newCBuf(7)
	for v := 1.0; v < 10; v++ {
		c.append(v)
		assert.EqualValues(t, v, c.nthNewest(0))
		assert.EqualValues(t, v, c.size())
		assert.EqualValues(t, int(v-1)%7, c.newestIndex())
		assert.EqualValues(t, int(v)%7, c.oldestIndex())
	}

	// 8,9,3,4,5,6,7
	//   ^
	// newest
	expected := []float64{9, 8, 7, 6, 5, 4, 3}
	for idx, ev := range expected {
		assert.EqualValues(t, ev, c.nthNewest(int64(idx)))
	}
	expected = []float64{3, 4, 5, 6, 7, 8, 9}
	for idx, ev := range expected {
		assert.EqualValues(t, ev, c.nthOldest(int64(idx)))
	}

	assert.EqualValues(t, 2, c.indexToSeq(2))
	assert.EqualValues(t, 6, c.indexToSeq(6))
	assert.EqualValues(t, 7, c.indexToSeq(0))
	assert.EqualValues(t, 8, c.indexToSeq(1))
}
