package tart

type Aroon struct {
	n   int64
	min *Min
	max *Max
	sz  int64
}

func NewAroon(n int64) *Aroon {
	return &Aroon{
		n:   n,
		min: NewMin(n + 1),
		max: NewMax(n + 1),
		sz:  0,
	}
}

func (a *Aroon) aroonValue(idx int64) float64 {
	return float64(a.n-(a.sz-1-idx)) / float64(a.n) * 100.0
}

func (a *Aroon) Update(h, l float64) (float64, float64) {
	minIdx, _ := a.min.Update(l)
	maxIdx, _ := a.max.Update(h)
	a.sz++

	if a.sz <= a.n {
		return 0, 0
	}

	return a.aroonValue(minIdx), a.aroonValue(maxIdx)
}

func AroonArr(h, l []float64, n int64) ([]float64, []float64) {
	dn := make([]float64, len(h))
	up := make([]float64, len(h))

	a := NewAroon(n)
	for i := 0; i < len(h); i++ {
		dn[i], up[i] = a.Update(h[i], l[i])
	}

	return dn, up
}
