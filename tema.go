package tart

type Tema struct {
	n    int64
	sz   int64
	ema1 *Ema
	ema2 *Ema
	ema3 *Ema
}

func NewTema(n int64, k float64) *Tema {
	return &Tema{
		n:    n,
		sz:   0,
		ema1: NewEma(n, k),
		ema2: NewEma(n, k),
		ema3: NewEma(n, k),
	}
}

func (t *Tema) Update(v float64) float64 {
	e1 := t.ema1.Update(v)
	t.sz++

	if t.sz > t.n-1 {
		e2 := t.ema2.Update(e1)

		if t.sz > 2*t.n-2 {
			e3 := t.ema3.Update(e2)

			if t.sz > t.n*3-3 {
				return e3 + 3.0*(e1-e2)
			}
		}
	}

	return 0
}

func TemaArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	k := 2.0 / float64(n+1)
	t := NewTema(n, k)
	for i, v := range in {
		out[i] = t.Update(v)
	}

	return out
}
