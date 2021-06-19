package tart

import (
	"math"
)

// Refer to ADX.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:average_directional_index_adx
type Dx struct {
	n        int64
	prevH    float64
	prevL    float64
	tr       *TRange
	sTr      *smooth
	sPlusDm  *smooth
	sMinusDm *smooth
	sz       int64
}

func NewDx(n int64) *Dx {
	return &Dx{
		n:        n,
		tr:       NewTRange(),
		prevH:    0,
		prevL:    0,
		sTr:      newSmooth(n),
		sPlusDm:  newSmooth(n),
		sMinusDm: newSmooth(n),
		sz:       0,
	}
}

func (d *Dx) Update(h, l, c float64) float64 {
	d.sz++

	tr := d.tr.Update(h, l, c)
	plusDm := h - d.prevH
	minusDm := d.prevL - l
	d.prevH = h
	d.prevL = l

	if d.sz == 1 {
		return 0
	}

	if minusDm > plusDm || plusDm < 0 {
		plusDm = 0
	}
	if plusDm > minusDm || minusDm < 0 {
		minusDm = 0
	}

	str := d.sTr.update(tr)
	spDm := d.sPlusDm.update(plusDm)
	smDm := d.sMinusDm.update(minusDm)
	if d.sz <= d.n {
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

	return dx
}

func (d *Dx) InitPeriod() int64 {
	return d.n
}

func (d *Dx) Valid() bool {
	return d.sz > d.InitPeriod()
}

// Refer to ADX.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:average_directional_index_adx
func DxArr(h, l, c []float64, n int64) []float64 {
	out := make([]float64, len(c))

	d := NewDx(n)
	for i := 0; i < len(c); i++ {
		out[i] = d.Update(h[i], l[i], c[i])
	}

	return out
}

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

func (s *smooth) initPeriod() int64 {
	return s.n - 1
}
