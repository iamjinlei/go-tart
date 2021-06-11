package tart

type Var struct {
	n    int64
	sz   int64
	hist []float64
	sum  float64
}

func NewVar(n int64) *Var {
	return &Var{
		n:    n,
		sz:   0,
		hist: make([]float64, n),
		sum:  0,
	}
}

func (r *Var) Update(v float64) float64 {
	idx := r.sz % r.n
	r.sum += v - r.hist[idx]
	r.hist[idx] = v

	r.sz++

	if r.sz < r.n {
		return 0
	}

	mean := r.sum / float64(r.n)
	sum := float64(0)
	for _, h := range r.hist {
		diff := (h - mean)
		sum += diff * diff
	}

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
