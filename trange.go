package tart

import (
	"math"
)

// Welles Wilder described these calculations to determine the trading range
// for a stock or commodity. True Range is defined as the largest of the
// following: (1) The distance from today's high to today's low. (2) The
// distance from yesterday's close to today's high. (3) The distance from
// yesterday's close to today's low. Wilder included price comparisons among
// subsequent bars in order to account for gaps in his range calculation.
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
	t.sz++

	d0 := math.Abs(h - t.prevC)
	d1 := math.Abs(l - t.prevC)
	t.prevC = c

	if t.sz == 1 {
		return 0
	}

	ret := h - l
	ret = max(ret, d0)
	ret = max(ret, d1)
	return ret
}

func (t *TRange) InitPeriod() int64 {
	return 1
}

func (t *TRange) Valid() bool {
	return t.sz > t.InitPeriod()
}

// Welles Wilder described these calculations to determine the trading range
// for a stock or commodity. True Range is defined as the largest of the
// following: (1) The distance from today's high to today's low. (2) The
// distance from yesterday's close to today's high. (3) The distance from
// yesterday's close to today's low. Wilder included price comparisons among
// subsequent bars in order to account for gaps in his range calculation.
func TRangeArr(h, l, c []float64) []float64 {
	out := make([]float64, len(c))

	t := NewTRange()
	for i := 0; i < len(c); i++ {
		out[i] = t.Update(h[i], l[i], c[i])
	}

	return out
}
