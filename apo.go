package tart

type Apo struct {
	slowN  int64
	fastMa *Ma
	slowMa *Ma
	sz     int64
}

func NewApo(t MaType, fastN, slowN int64) *Apo {
	if fastN > slowN {
		fastN, slowN = slowN, fastN
	}

	return &Apo{
		slowN:  slowN,
		fastMa: NewMa(t, fastN),
		slowMa: NewMa(t, slowN),
		sz:     0,
	}
}

func (a *Apo) Update(v float64) float64 {
	fast := a.fastMa.Update(v)
	slow := a.slowMa.Update(v)

	a.sz++
	if a.sz < a.slowN {
		return 0
	}

	return fast - slow
}

func ApoArr(t MaType, in []float64, fastN, slowN int64) []float64 {
	out := make([]float64, len(in))

	a := NewApo(t, fastN, slowN)
	for i, v := range in {
		out[i] = a.Update(v)
	}

	return out
}
