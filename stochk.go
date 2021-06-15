package tart

type StochasticK struct {
	kN  int64
	min *Min
	max *Max
	sz  int64
}

func NewStochasticK(kN int64) *StochasticK {
	return &StochasticK{
		kN:  kN,
		min: NewMin(kN),
		max: NewMax(kN),
		sz:  0,
	}
}

func (s *StochasticK) Update(h, l, c float64) float64 {
	_, min := s.min.Update(l)
	_, max := s.max.Update(h)
	s.sz++

	if s.sz < s.kN {
		return 0
	}
	return (c - min) / (max - min) * 100.0
}

func StochasticKArr(h, l, c []float64, kN int64) []float64 {
	out := make([]float64, len(c))

	s := NewStochasticK(kN)
	for i := 0; i < len(c); i++ {
		out[i] = s.Update(h[i], l[i], c[i])
	}

	return out
}
