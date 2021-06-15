package tart

type StochFast struct {
	kN  int64
	dN  int64
	min *Min
	max *Max
	d   *Ma
	sz  int64
}

func NewStochFast(kN int64, dt MaType, dN int64) *StochFast {
	return &StochFast{
		kN:  kN,
		dN:  dN,
		min: NewMin(kN),
		max: NewMax(kN),
		d:   NewMa(dt, dN),
		sz:  0,
	}
}

func (s *StochFast) Update(h, l, c float64) (float64, float64) {
	_, min := s.min.Update(l)
	_, max := s.max.Update(h)
	s.sz++

	if s.sz < s.kN {
		return 0, 0
	}
	k := (c - min) / (max - min) * 100.0
	d := s.d.Update(k)

	if s.sz <= s.kN+s.dN-2 {
		return 0, 0
	}

	return k, d
}

func StochFastArr(h, l, c []float64, kN int64, dt MaType, dN int64) ([]float64, []float64) {
	k := make([]float64, len(c))
	d := make([]float64, len(c))

	s := NewStochFast(kN, dt, dN)
	for i := 0; i < len(c); i++ {
		k[i], d[i] = s.Update(h[i], l[i], c[i])
	}

	return k, d
}
