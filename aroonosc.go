package tart

type AroonOsc struct {
	n     int64
	aroon *Aroon
}

func NewAroonOsc(n int64) *AroonOsc {
	return &AroonOsc{
		n:     n,
		aroon: NewAroon(n),
	}
}

func (a *AroonOsc) Update(h, l float64) float64 {
	dn, up := a.aroon.Update(h, l)
	return up - dn
}

func AroonOscArr(h, l []float64, n int64) []float64 {
	out := make([]float64, len(h))

	a := NewAroonOsc(n)
	for i := 0; i < len(h); i++ {
		out[i] = a.Update(h[i], l[i])
	}

	return out
}
