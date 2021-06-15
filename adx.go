package tart

type Adx struct {
	n   int64
	dx  *Dx
	adx *Ema
	sz  int64
}

func NewAdx(n int64) *Adx {
	return &Adx{
		n:   n,
		dx:  NewDx(n),
		adx: NewEma(n, 1.0/float64(n)),
		sz:  0,
	}
}

func (a *Adx) Update(h, l, c float64) float64 {
	dx := a.dx.Update(h, l, c)
	a.sz++
	if a.sz <= a.n {
		return 0
	}

	return a.adx.Update(dx)
}

func AdxArr(h, l, c []float64, n int64) []float64 {
	out := make([]float64, len(c))

	a := NewAdx(n)
	for i := 0; i < len(c); i++ {
		out[i] = a.Update(h[i], l[i], c[i])
	}

	return out
}
