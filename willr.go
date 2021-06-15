package tart

type WillR struct {
	n     int64
	stoch *StochasticK
	sz    int64
}

func NewWillR(n int64) *WillR {
	return &WillR{
		n:     n,
		stoch: NewStochasticK(n),
		sz:    0,
	}
}

func (w *WillR) Update(h, l, c float64) float64 {
	w.sz++
	k := w.stoch.Update(h, l, c)

	if w.sz < w.n {
		return 0
	}
	return k - 100.0
}

func WillRArr(h, l, c []float64, n int64) []float64 {
	out := make([]float64, len(c))

	w := NewWillR(n)
	for i := 0; i < len(c); i++ {
		out[i] = w.Update(h[i], l[i], c[i])
	}

	return out
}
