package tart

// Developed by George C. Lane in the late 1950s, the Stochastic Oscillator is
// a momentum indicator that shows the location of the close relative to the
// high-low range over a set number of periods. According to an interview with
// Lane, the Stochastic Oscillator “doesn't follow price, it doesn't follow
// volume or anything like that. It follows the speed or the momentum of price.
// As a rule, the momentum changes direction before price.” As such, bullish and
// bearish divergences in the Stochastic Oscillator can be used to foreshadow
// reversals. This was the first, and most important, signal that Lane identified.
// Lane also used this oscillator to identify bull and bear set-ups to anticipate
// a future reversal. As the Stochastic Oscillator is range-bound, it is also
// useful for identifying overbought and oversold levels.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:stochastic_oscillator_fast_slow_and_full
//  https://www.investopedia.com/terms/s/stochasticoscillator.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/fast-stochastic
type StochFast struct {
	kN     int64
	dN     int64
	stochK *StochasticK
	d      *Ma
	sz     int64
}

func NewStochFast(kN int64, dt MaType, dN int64) *StochFast {
	return &StochFast{
		kN:     kN,
		dN:     dN,
		stochK: NewStochasticK(kN),
		d:      NewMa(dt, dN),
		sz:     0,
	}
}

func (s *StochFast) Update(h, l, c float64) (float64, float64) {
	s.sz++

	k := s.stochK.Update(h, l, c)

	if s.sz < s.kN {
		return 0, 0
	}
	d := s.d.Update(k)

	if s.sz <= s.kN+s.dN-2 {
		return 0, 0
	}

	return k, d
}

func (s *StochFast) InitPeriod() int64 {
	return s.kN + s.dN - 2
}

func (s *StochFast) Valid() bool {
	return s.sz > s.InitPeriod()
}

// Developed by George C. Lane in the late 1950s, the Stochastic Oscillator is
// a momentum indicator that shows the location of the close relative to the
// high-low range over a set number of periods. According to an interview with
// Lane, the Stochastic Oscillator “doesn't follow price, it doesn't follow
// volume or anything like that. It follows the speed or the momentum of price.
// As a rule, the momentum changes direction before price.” As such, bullish and
// bearish divergences in the Stochastic Oscillator can be used to foreshadow
// reversals. This was the first, and most important, signal that Lane identified.
// Lane also used this oscillator to identify bull and bear set-ups to anticipate
// a future reversal. As the Stochastic Oscillator is range-bound, it is also
// useful for identifying overbought and oversold levels.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:stochastic_oscillator_fast_slow_and_full
//  https://www.investopedia.com/terms/s/stochasticoscillator.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/fast-stochastic
func StochFastArr(h, l, c []float64, kN int64, dt MaType, dN int64) ([]float64, []float64) {
	k := make([]float64, len(c))
	d := make([]float64, len(c))

	s := NewStochFast(kN, dt, dN)
	for i := 0; i < len(c); i++ {
		k[i], d[i] = s.Update(h[i], l[i], c[i])
	}

	return k, d
}
