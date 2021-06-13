package tart

type Atr struct {
	n   int64
	tr  *TRange
	ema *Ema
	sz  int64
}

func NewAtr(n int64) *Atr {
	return &Atr{
		n:   n,
		tr:  NewTRange(),
		ema: NewEma(n, 1.0/float64(n)),
		sz:  0,
	}
}

func (a *Atr) Update(h, l, c float64) float64 {
	tr := a.tr.Update(h, l, c)
	a.sz++
	if a.sz == 1 {
		return 0
	}

	return a.ema.Update(tr)
}

func AtrArr(h, l, c []float64, n int64) []float64 {
	out := make([]float64, len(c))

	a := NewAtr(n)
	for i := 0; i < len(c); i++ {
		out[i] = a.Update(h[i], l[i], c[i])
	}

	return out
}
