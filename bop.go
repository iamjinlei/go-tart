package tart

type Bop struct {
}

func NewBop() *Bop {
	return &Bop{}
}

func (b *Bop) Update(o, h, l, c float64) float64 {
	d := h - l
	if almostZero(d) {
		return 0
	}
	return (c - o) / d
}

func BopArr(o, h, l, c []float64) []float64 {
	out := make([]float64, len(o))

	b := NewBop()
	for i := 0; i < len(o); i++ {
		out[i] = b.Update(o[i], h[i], l[i], c[i])
	}

	return out
}
