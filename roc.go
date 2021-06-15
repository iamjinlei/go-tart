package tart

type Roc struct {
	n    int64
	hist *cBuf
	sz   int64
}

func NewRoc(n int64) *Roc {
	return &Roc{
		n:    n,
		hist: newCBuf(n),
		sz:   0,
	}
}

func (r *Roc) Update(v float64) float64 {
	old := r.hist.append(v)

	r.sz++
	if r.sz <= r.n {
		return 0
	}

	if almostZero(old) {
		return 0
	}
	return (v - old) / old * 100.0
}

func RocArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	r := NewRoc(n)
	for i, v := range in {
		out[i] = r.Update(v)
	}

	return out
}
