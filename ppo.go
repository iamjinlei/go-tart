package tart

type Ppo struct {
	fastN int64
	slowN int64
	fast  *Ma
	slow  *Ma
	sz    int64
}

func NewPpo(t MaType, fastN, slowN int64) *Ppo {
	if slowN < fastN {
		fastN, slowN = slowN, fastN
	}
	return &Ppo{
		fastN: fastN,
		slowN: slowN,
		fast:  NewMa(t, fastN),
		slow:  NewMa(t, slowN),
		sz:    0,
	}
}

func (p *Ppo) Update(v float64) float64 {
	fast := p.fast.Update(v)
	slow := p.slow.Update(v)

	p.sz++
	if p.sz < p.slowN {
		return 0
	}

	if almostZero(slow) {
		return 0
	}
	return (fast - slow) / slow * 100.0
}

func PpoArr(in []float64, t MaType, fastN, slowN int64) []float64 {
	out := make([]float64, len(in))

	p := NewPpo(t, fastN, slowN)
	for i, v := range in {
		out[i] = p.Update(v)
	}

	return out
}
