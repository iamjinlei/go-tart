package tart

type StochRsi struct {
	n     int64
	kN    int64
	dN    int64
	util  int64
	rsi   *Rsi
	stoch *StochFast
	sz    int64
}

func NewStochRsi(n int64, kN int64, dt MaType, dN int64) *StochRsi {
	return &StochRsi{
		n:     n,
		kN:    kN,
		dN:    dN,
		util:  n + kN + dN - 1,
		rsi:   NewRsi(n),
		stoch: NewStochFast(kN, dt, dN),
		sz:    0,
	}
}

func (s *StochRsi) Update(v float64) (float64, float64) {
	rsi := s.rsi.Update(v)
	s.sz++

	if s.sz <= s.n {
		return 0, 0
	}
	k, d := s.stoch.Update(rsi, rsi, rsi)
	if s.sz < s.util {
		return 0, 0
	}

	return k, d
}

func StochRsiArr(in []float64, n, kN int64, dt MaType, dN int64) ([]float64, []float64) {
	k := make([]float64, len(in))
	d := make([]float64, len(in))

	s := NewStochRsi(n, kN, dt, dN)
	for i, v := range in {
		k[i], d[i] = s.Update(v)
	}

	return k, d
}
