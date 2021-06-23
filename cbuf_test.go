package tart

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCBuf(t *testing.T) {
	c := NewCBuf(7)
	for v := 1.0; v < 10; v++ {
		c.Append(v)
		assert.EqualValues(t, v, c.NthNewest(0))
		assert.EqualValues(t, v, c.Size())
		assert.EqualValues(t, int(v-1)%7, c.NewestIndex())
		assert.EqualValues(t, int(v)%7, c.OldestIndex())
	}

	// 8,9,3,4,5,6,7
	//   ^
	// newest
	expected := []float64{9, 8, 7, 6, 5, 4, 3}
	for idx, ev := range expected {
		assert.EqualValues(t, ev, c.NthNewest(int64(idx)))
	}
	expected = []float64{3, 4, 5, 6, 7, 8, 9}
	for idx, ev := range expected {
		assert.EqualValues(t, ev, c.NthOldest(int64(idx)))
	}

	assert.EqualValues(t, 2, c.IndexToSeq(2))
	assert.EqualValues(t, 6, c.IndexToSeq(6))
	assert.EqualValues(t, 7, c.IndexToSeq(0))
	assert.EqualValues(t, 8, c.IndexToSeq(1))

	sum := float64(0)
	c.Iter(func(v float64) {
		sum += v
	})
	assert.EqualValues(t, 42, sum)
}
