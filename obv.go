package tart

type Obv struct {
	prev float64
	obv  float64
	sz   int64
}

func NewObv() *Obv {
	return &Obv{
		prev: 0,
		obv:  0,
		sz:   0,
	}
}

func (o *Obv) Update(c, v float64) float64 {
	o.sz++

	prev := o.prev
	o.prev = c
	if o.sz == 1 {
		o.obv = v
		return o.obv
	}

	if c > prev {
		o.obv += v
	} else if c < prev {
		o.obv -= v
	}

	return o.obv
}

func ObvArr(c, v []float64) []float64 {
	out := make([]float64, len(c))

	o := NewObv()
	for i := 0; i < len(c); i++ {
		out[i] = o.Update(c[i], v[i])
	}

	return out
}
