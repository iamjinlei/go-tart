package tart

type Trix struct {
	n    int64
	ema1 *Ema
	ema2 *Ema
	ema3 *Ema
	prev float64
	sz   int64
}

func NewTrix(n int64) *Trix {
	k := 2.0 / float64(n+1)
	return &Trix{
		n:    n,
		ema1: NewEma(n, k),
		ema2: NewEma(n, k),
		ema3: NewEma(n, k),
		prev: 0,
		sz:   0,
	}
}

func (t *Trix) Update(v float64) float64 {
	t.sz++
	v = t.ema1.Update(v)
	if t.sz < t.n {
		return 0
	}

	v = t.ema2.Update(v)
	if t.sz < 2*t.n-1 {
		return 0
	}

	v = t.ema3.Update(v)
	roc := float64(0)
	if !almostZero(t.prev) {
		roc = (v - t.prev) / t.prev
	}
	t.prev = v
	if t.sz < 3*t.n-2 {
		return 0
	}

	return roc * 100.0
}

func TrixArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	t := NewTrix(n)
	for i, v := range in {
		out[i] = t.Update(v)
	}

	return out
}
