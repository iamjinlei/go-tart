package tart

import (
	"math"
)

// Developed by Perry Kaufman, Kaufman's Adaptive Moving
// Average (KAMA) is a moving average designed to account
// for market noise or volatility. KAMA will closely follow
// prices when the price swings are relatively small and
// the noise is low. KAMA will adjust when the price swings
// widen and follow prices from a greater distance. This
// trend-following indicator can be used to identify the
// overall trend, time turning points and filter price
// movements.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:kaufman_s_adaptive_moving_average
type Kama struct {
	n          int64
	constA     float64
	constB     float64
	sz         int64
	hist       []float64
	absChgHist []float64
	sumAbsChg  float64
	kama       float64
}

func NewKama(n int64) *Kama {
	return &Kama{
		n:          n,
		constA:     2.0/(2.0+1.0) - 2.0/(30.0+1.0),
		constB:     2.0 / (30.0 + 1.0),
		sz:         0,
		hist:       make([]float64, n+1),
		absChgHist: make([]float64, n),
		sumAbsChg:  0,
		kama:       0,
	}
}

func (k *Kama) Update(v float64) float64 {
	idx := k.sz % (k.n + 1)
	prevIdx := (idx + k.n) % (k.n + 1)
	nextIdx := (idx + 1) % (k.n + 1)
	k.hist[idx] = v

	absChgIdx := (k.sz + k.n - 1) % k.n
	chg := math.Abs(v - k.hist[prevIdx])
	k.sumAbsChg += chg - k.absChgHist[absChgIdx]
	k.absChgHist[absChgIdx] = chg

	k.sz++

	if k.sz <= k.n {
		k.kama = v
		return 0
	}

	// er = change / volatility
	//    = abs(Nth value - 1st value) / (sum of N period abs chg)
	totalChg := math.Abs(v - k.hist[nextIdx])
	var er float64
	if (totalChg >= k.sumAbsChg) || (k.sumAbsChg < 0.00000000000001 && k.sumAbsChg > -0.00000000000001) {
		er = 1.0
	} else {
		er = totalChg / k.sumAbsChg
	}

	sc := er*k.constA + k.constB
	sc *= sc

	k.kama = sc*v + (1-sc)*k.kama
	return k.kama
}

func (k *Kama) InitPeriod() int64 {
	return k.n
}

func (k *Kama) Valid() bool {
	return k.sz > k.InitPeriod()
}

// Developed by Perry Kaufman, Kaufman's Adaptive Moving
// Average (KAMA) is a moving average designed to account
// for market noise or volatility. KAMA will closely follow
// prices when the price swings are relatively small and
// the noise is low. KAMA will adjust when the price swings
// widen and follow prices from a greater distance. This
// trend-following indicator can be used to identify the
// overall trend, time turning points and filter price
// movements.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:kaufman_s_adaptive_moving_average
func KamaArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	k := NewKama(n)
	for i, v := range in {
		out[i] = k.Update(v)
	}

	return out
}
