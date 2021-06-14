package tart

type Wma struct {
	n    int64
	d    float64
	hist *cBuf
	sum  float64
	wsum float64
}

func NewWma(n int64) *Wma {
	return &Wma{
		n:    n,
		d:    float64(n*(n+1)) / 2,
		hist: newCBuf(n),
		sum:  0,
		wsum: 0,
	}
}

func (w *Wma) Update(v float64) float64 {
	if w.n == 1 {
		return v
	}

	old := w.hist.append(v)
	w.sum += v - old

	sz := w.hist.size()
	if sz < w.n {
		w.wsum += v * float64(sz)
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
