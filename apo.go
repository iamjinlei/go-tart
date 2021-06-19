package tart

// The Absolute Price Oscillator displays the difference
// between two exponential moving averages of a security's
// price and is expressed as an absolute value.
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/apo
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
	a.sz++

	fast := a.fastMa.Update(v)
	slow := a.slowMa.Update(v)

	if a.sz < a.slowN {
		return 0
	}

	return fast - slow
}

func (a *Apo) InitPeriod() int64 {
	return a.slowN - 1
}

func (a *Apo) Valid() bool {
	return a.sz > a.InitPeriod()
}

// The Absolute Price Oscillator displays the difference
// between two exponential moving averages of a security's
// price and is expressed as an absolute value.
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/apo
func ApoArr(t MaType, in []float64, fastN, slowN int64) []float64 {
	out := make([]float64, len(in))

	a := NewApo(t, fastN, slowN)
	for i, v := range in {
		out[i] = a.Update(v)
	}

	return out
}
