package tart

type Dema struct {
	n    int64
	sz   int64
	ema1 *Ema
	ema2 *Ema
}

func NewDema(n int64, k float64) *Dema {
	return &Dema{
		n:    n,
		sz:   0,
		ema1: NewEma(n, k),
		ema2: NewEma(n, k),
	}
}

func (d *Dema) Update(v float64) float64 {
	e1 := d.ema1.Update(v)
	d.sz++

	if d.sz > d.n-1 {
		e2 := d.ema2.Update(e1)
		if d.sz > d.n*2-2 {
			return 2.0*e1 - e2
		}
	}

	return 0
}

func DemaArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	k := 2.0 / float64(n+1)
	d := NewDema(n, k)
	for i, v := range in {
		out[i] = d.Update(v)
	}

	return out
}
