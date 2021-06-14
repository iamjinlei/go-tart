package tart

type Var struct {
	n    int64
	hist *cBuf
	sum  float64
}

func NewVar(n int64) *Var {
	return &Var{
		n:    n,
		hist: newCBuf(n),
		sum:  0,
	}
}

func (r *Var) Update(v float64) float64 {
	old := r.hist.append(v)
	r.sum += v - old

	if r.hist.size() < r.n {
		return 0
	}

	mean := r.sum / float64(r.n)
	sum := float64(0)
	r.hist.iter(func(v float64) {
		diff := (v - mean)
		sum += diff * diff
	})

	return sum / float64(r.n)
}

func VarArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	s := NewVar(n)
	for i, v := range in {
		out[i] = s.Update(v)
	}

	return out
}
