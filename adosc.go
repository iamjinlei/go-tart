package tart

// Developed by Marc Chaikin, the Chaikin Oscillator measures the
// momentum of the Accumulation Distribution Line using the MACD
// formula. (This makes it an indicator of an indicator.) The
// Chaikin Oscillator is the difference between the 3-day and
// 10-day EMAs of the Accumulation Distribution Line. Like other
// momentum indicators, this indicator is designed to anticipate
// directional changes in the Accumulation Distribution Line by
// measuring the momentum behind the movements. A momentum change
// is the first step to a trend change. Anticipating trend changes
// in the Accumulation Distribution Line can help chartists
// anticipate trend changes in the underlying security. The
// Chaikin Oscillator generates signals with crosses above/below
// the zero line or with bullish/bearish divergences.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:chaikin_oscillator
//  https://www.investopedia.com/terms/c/chaikinoscillator.asp
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

func (a *AdOsc) InitPeriod() int64 {
	return a.slowN - 1
}

func (a *AdOsc) Valid() bool {
	return a.sz > a.InitPeriod()
}

// Developed by Marc Chaikin, the Chaikin Oscillator measures the
// momentum of the Accumulation Distribution Line using the MACD
// formula. (This makes it an indicator of an indicator.) The
// Chaikin Oscillator is the difference between the 3-day and
// 10-day EMAs of the Accumulation Distribution Line. Like other
// momentum indicators, this indicator is designed to anticipate
// directional changes in the Accumulation Distribution Line by
// measuring the momentum behind the movements. A momentum change
// is the first step to a trend change. Anticipating trend changes
// in the Accumulation Distribution Line can help chartists
// anticipate trend changes in the underlying security. The
// Chaikin Oscillator generates signals with crosses above/below
// the zero line or with bullish/bearish divergences.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:chaikin_oscillator
//  https://www.investopedia.com/terms/c/chaikinoscillator.asp
func AdOscArr(h, l, c, v []float64, fastN, slowN int64) []float64 {
	out := make([]float64, len(c))

	a := NewAdOsc(fastN, slowN)
	for i := 0; i < len(c); i++ {
		out[i] = a.Update(h[i], l[i], c[i], v[i])
	}
	return out
}
