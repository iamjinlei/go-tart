package tart

type Sma struct {
	n    int64
	hist *cBuf
	sz   int64
	sum  float64
}

func NewSma(n int64) *Sma {
	return &Sma{
		n:    n,
		hist: newCBuf(n),
		sz:   0,
		sum:  0,
	}
}

func (s *Sma) Update(v float64) float64 {
	old := s.hist.append(v)
	s.sum += v - old

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
