package tart

type Wma struct {
	n    int64
	d    float64
	sz   int64
	hist []float64
	sum  float64
	wsum float64
}

func NewWma(n int64) *Wma {
	return &Wma{
		n:    n,
		d:    float64(n*(n+1)) / 2,
		sz:   0,
		hist: make([]float64, n),
		sum:  0,
		wsum: 0,
	}
}

func (w *Wma) Update(v float64) float64 {
	if w.n == 1 {
		return v
	}

	idx := w.sz % w.n
	w.sum += v - w.hist[idx]
	w.hist[idx] = v

	w.sz++

	if w.sz < w.n {
		w.wsum += v * float64(w.sz)
		return 0
	} else {
		w.wsum += v * float64(w.n)
	}

	ret := w.wsum / w.d
	w.wsum -= w.sum

	return ret
}

func WmaArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	w := NewWma(n)
	for i, v := range in {
		out[i] = w.Update(v)
	}

	return out
}
