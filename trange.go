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

	ret := h - l
	ret = max(ret, d0)
	ret = max(ret, d1)
	return ret
}

func TRangeArr(h, l, c []float64) []float64 {
	out := make([]float64, len(c))

	t := NewTRange()
	for i := 0; i < len(c); i++ {
		out[i] = t.Update(h[i], l[i], c[i])
	}

	return out
}
