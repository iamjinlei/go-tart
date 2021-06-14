package tart

type Sum struct {
	n    int64
	hist *cBuf
	sum  float64
}

func NewSum(n int64) *Sum {
	return &Sum{
		n:    n,
		hist: newCBuf(n),
		sum:  0,
	}
}

func (s *Sum) Update(v float64) float64 {
	old := s.hist.append(v)
	s.sum += v - old

	if s.hist.size() < s.n {
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
