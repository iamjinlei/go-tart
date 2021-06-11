package tart

type Sma struct {
	n    int64
	sz   int64
	hist []float64
	sum  float64
}

func NewSma(n int64) *Sma {
	return &Sma{
		n:    n,
		sz:   0,
		hist: make([]float64, n),
		sum:  0,
	}
}

func (s *Sma) Update(v float64) float64 {
	idx := s.sz % s.n
	s.sum += v - s.hist[idx]
	s.hist[idx] = v

	s.sz++

	if s.sz < s.n {
		return 0
	}

	return s.sum / float64(s.n)
}

func SmaArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	s := NewSma(n)
	for i, v := range in {
		out[i] = s.Update(v)
	}

	return out
}
