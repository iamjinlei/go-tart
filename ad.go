package tart

type Ad struct {
	ad float64
}

func NewAd() *Ad {
	return &Ad{
		ad: 0,
	}
}

func (a *Ad) Update(h, l, c, v float64) float64 {
	h2l := h - l
	if h2l > 0.0 {
		a.ad += (((c - l) - (h - c)) / h2l) * v
	}
	return a.ad
}

func AdArr(h, l, c, v []float64) []float64 {
	out := make([]float64, len(c))

	a := NewAd()
	for i := 0; i < len(c); i++ {
		out[i] = a.Update(h[i], l[i], c[i], v[i])
	}
	return out
}
