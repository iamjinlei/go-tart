package tart

type Atr struct {
	n      int64
	k1     float64
	k2     float64
	tr     *TRange
	sz     int64
	prevTr float64
}

func NewAtr(n int64) *Atr {
	k := 1.0 / float64(n)
	return &Atr{
		n:  n,
		k1: k,
		k2: 1.0 - k,
		tr: NewTRange(),
		sz: 0,
	}
}

func (a *Atr) Update(h, l, c float64) float64 {
	tr := a.tr.Update(h, l, c)
	a.sz++

	if a.sz <= a.n+1 {
		a.prevTr += tr
		if a.sz == a.n+1 {
			a.prevTr /= float64(a.n)
			return a.prevTr
		}
		return 0
	}

	a.prevTr = tr*a.k1 + a.prevTr*a.k2
	return a.prevTr
}

func AtrArr(h, l, c []float64, n int64) []float64 {
	out := make([]float64, len(c))

	a := NewAtr(n)
	for i := 0; i < len(c); i++ {
		out[i] = a.Update(h[i], l[i], c[i])
	}

	return out
}
