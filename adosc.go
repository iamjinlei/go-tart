package tart

type AdOsc struct {
	fastN  int64
	slowN  int64
	fastK1 float64
	fastK2 float64
	slowK1 float64
	slowK2 float64
	ad     *Ad
	fast   float64
	slow   float64
	sz     int64
}

func NewAdOsc(fastN, slowN int64) *AdOsc {
	if slowN < fastN {
		fastN, slowN = slowN, fastN
	}
	return &AdOsc{
		fastN:  fastN,
		slowN:  slowN,
		fastK1: 2.0 / float64(fastN+1),
		fastK2: 1.0 - 2.0/float64(fastN+1),
		slowK1: 2.0 / float64(slowN+1),
		slowK2: 1 - 2.0/float64(slowN+1),
		ad:     NewAd(),
		fast:   0,
		slow:   0,
		sz:     0,
	}
}

func (a *AdOsc) Update(h, l, c, v float64) float64 {
	a.sz++
	ad := a.ad.Update(h, l, c, v)
	if a.sz == 1 {
		a.fast = ad
		a.slow = ad
		return 0
	}
	a.fast = ad*a.fastK1 + a.fast*a.fastK2
	a.slow = ad*a.slowK1 + a.slow*a.slowK2

	if a.sz < a.slowN {
		return 0
	}
	return a.fast - a.slow
}

func AdOscArr(h, l, c, v []float64, fastN, slowN int64) []float64 {
	out := make([]float64, len(c))

	a := NewAdOsc(fastN, slowN)
	for i := 0; i < len(c); i++ {
		out[i] = a.Update(h[i], l[i], c[i], v[i])
	}
	return out
}
