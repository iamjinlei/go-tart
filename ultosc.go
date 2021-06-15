package tart

import (
	"sort"
)

type UltOsc struct {
	n1    int64
	n2    int64
	n3    int64
	bp1   *Sum
	bp2   *Sum
	bp3   *Sum
	tr1   *Sum
	tr2   *Sum
	tr3   *Sum
	prevC float64
	sz    int64
}

func NewUltOsc(n1, n2, n3 int64) *UltOsc {
	arr := []int64{n1, n2, n3}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	n1, n2, n3 = arr[0], arr[1], arr[2]
	return &UltOsc{
		n1:    n1,
		n2:    n2,
		n3:    n3,
		bp1:   NewSum(n1),
		bp2:   NewSum(n2),
		bp3:   NewSum(n3),
		tr1:   NewSum(n1),
		tr2:   NewSum(n2),
		tr3:   NewSum(n3),
		prevC: 0,
		sz:    0,
	}
}

func (u *UltOsc) Update(h, l, c float64) float64 {
	u.sz++

	low := min(l, u.prevC)
	bp := c - low
	tr := max(h, u.prevC) - low
	u.prevC = c

	if u.sz == 1 {
		return 0
	}

	bp1 := u.bp1.Update(bp)
	bp2 := u.bp2.Update(bp)
	bp3 := u.bp3.Update(bp)
	tr1 := u.tr1.Update(tr)
	tr2 := u.tr2.Update(tr)
	tr3 := u.tr3.Update(tr)

	if u.sz <= u.n3 {
		return 0
	}

	d1 := bp1 / tr1
	d2 := bp2 / tr2
	d3 := bp3 / tr3

	return (d1*4.0 + d2*2.0 + d3) / 7.0 * 100.0
}

func UltOscArr(h, l, c []float64, n1, n2, n3 int64) []float64 {
	out := make([]float64, len(c))

	u := NewUltOsc(n1, n2, n3)
	for i := 0; i < len(c); i++ {
		out[i] = u.Update(h[i], l[i], c[i])
	}

	return out
}
