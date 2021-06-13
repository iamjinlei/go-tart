package tart

import (
	"math"
)

type smooth struct {
	n  int64
	k  float64
	sz int64
	sm float64
}

func newSmooth(n int64) *smooth {
	return &smooth{
		n:  n,
		k:  float64(n-1) / float64(n),
		sz: 0,
		sm: 0,
	}
}

func (s *smooth) update(v float64) float64 {
	s.sz++
	if s.sz < s.n {
		s.sm += v
		return 0
	} else {
		s.sm = s.sm*s.k + v
	}

	return s.sm
}

type Adx struct {
	n        int64
	prevH    float64
	prevL    float64
	tr       *TRange
	sTr      *smooth
	sPlusDm  *smooth
	sMinusDm *smooth
	sz       int64
	adx      *Ema
}

func NewAdx(n int64) *Adx {
	return &Adx{
		n:        n,
		tr:       NewTRange(),
		prevH:    0,
		prevL:    0,
		sTr:      newSmooth(n),
		sPlusDm:  newSmooth(n),
		sMinusDm: newSmooth(n),
		sz:       0,
		adx:      NewEma(n, 1.0/float64(n)),
	}
}

func (a *Adx) Update(h, l, c float64) float64 {
	tr := a.tr.Update(h, l, c)
	plusDm := h - a.prevH
	minusDm := a.prevL - l
	a.prevH = h
	a.prevL = l

	a.sz++
	if a.sz == 1 {
		return 0
	}

	if minusDm > plusDm || plusDm < 0 {
		plusDm = 0
	}
	if plusDm > minusDm || minusDm < 0 {
		minusDm = 0
	}

	str := a.sTr.update(tr)
	spDm := a.sPlusDm.update(plusDm)
	smDm := a.sMinusDm.update(minusDm)
	if a.sz <= a.n {
		return 0
	}
	dx := float64(0)
	if !almostZero(str) {
		// NOTE: when tr is non-zero, str * 100.0 seems cancelled out in the formula?
		pDi := spDm /// str * 100.0
		mDi := smDm /// str * 100.0
		sumDi := pDi + mDi
		if !almostZero(sumDi) {
			dx = math.Abs(pDi-mDi) / sumDi * 100.0
		}
	}

	return a.adx.Update(dx)
}

func AdxArr(h, l, c []float64, n int64) []float64 {
	out := make([]float64, len(c))

	a := NewAdx(n)
	for i := 0; i < len(c); i++ {
		out[i] = a.Update(h[i], l[i], c[i])
	}

	return out
}
