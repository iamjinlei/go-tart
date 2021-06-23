package tart

// Sum of the last N values.
type Sum struct {
	n    int64
	hist *CBuf
	sum  float64
}

func NewSum(n int64) *Sum {
	return &Sum{
		n:    n,
		hist: NewCBuf(n),
		sum:  0,
	}
}

func (s *Sum) Update(v float64) float64 {
	old := s.hist.Append(v)
	s.sum += v - old

	if s.hist.Size() < s.n {
		return 0
	}

	return s.sum
}

func (s *Sum) InitPeriod() int64 {
	return s.n - 1
}

func (s *Sum) Valid() bool {
	return s.hist.Size() > s.InitPeriod()
}

func SumArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	s := NewSum(n)
	for i, v := range in {
		out[i] = s.Update(v)
	}

	return out
}
