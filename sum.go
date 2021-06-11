package tart

type Sum struct {
	n    int64
	sz   int64
	hist []float64
	sum  float64
}

func NewSum(n int64) *Sum {
	return &Sum{
		n:    n,
		sz:   0,
		hist: make([]float64, n),
		sum:  0,
	}
}

func (s *Sum) Update(v float64) float64 {
	idx := s.sz % s.n
	s.sum += v - s.hist[idx]
	s.hist[idx] = v

	s.sz++

	if s.sz < s.n {
		return 0
	}

	return s.sum
}

func SumArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	s := NewSum(n)
	for i, v := range in {
		out[i] = s.Update(v)
	}

	return out
}
