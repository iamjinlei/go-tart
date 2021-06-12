package tart

import (
	"math"
)

type TRange struct {
	sz    int64
	prevC float64
}

func NewTRange() *TRange {
	return &TRange{
		sz:    0,
		prevC: 0,
	}
}

func (t *TRange) Update(h, l, c float64) float64 {
	d0 := math.Abs(h - t.prevC)
	d1 := math.Abs(l - t.prevC)
	t.prevC = c
	t.sz++

	if t.sz == 1 {
		return 0
	}

	// math.Max() does some extra we don'tc care (overhead)
	max := h - l
	if max < d0 {
		max = d0
	}
	if max < d1 {
		max = d1
	}

	return max
}

func TRangeArr(h, l, c []float64) []float64 {
	out := make([]float64, len(c))

	t := NewTRange()
	for i := 0; i < len(c); i++ {
		out[i] = t.Update(h[i], l[i], c[i])
	}

	return out
}
