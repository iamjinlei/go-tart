package tart

// The Slow Stochastic Oscillator is a momentum indicator that shows the location
// of the close relative to the high-low range over a set number of periods. The
// indicator can range from 0 to 100. The difference between the Slow and Fast
// Stochastic Oscillator is the Slow %K incorporates a %K slowing period of 3 that
// controls the internal smoothing of %K. Setting the smoothing period to 1 is
// equivalent to plotting the Fast Stochastic Oscillator.
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/slow-stochastic
type StochSlow struct {
	fastKN int64
	slowKN int64
	slowDN int64
	util   int64
	stochK *StochasticK
	slowK  *Ma
	slowD  *Ma
	sz     int64
}

func NewStochSlow(fastKN int64, kt MaType, slowKN int64, dt MaType, slowDN int64) *StochSlow {
	return &StochSlow{
		fastKN: fastKN,
		slowKN: slowKN,
		slowDN: slowDN,
		util:   fastKN + slowKN + slowDN - 2,
		stochK: NewStochasticK(fastKN),
		slowK:  NewMa(kt, slowKN),
		slowD:  NewMa(dt, slowDN),
		sz:     0,
	}
}

func (s *StochSlow) Update(h, l, c float64) (float64, float64) {
	s.sz++

	fastK := s.stochK.Update(h, l, c)

	if s.sz < s.fastKN {
		return 0, 0
	}
	slowK := s.slowK.Update(fastK)
	slowD := s.slowD.Update(slowK)

	if s.sz < s.util {
		return 0, 0
	}

	return slowK, slowD
}

func (s *StochSlow) InitPeriod() int64 {
	return s.util - 1
}

func (s *StochSlow) Valid() bool {
	return s.sz > s.InitPeriod()
}

// The Slow Stochastic Oscillator is a momentum indicator that shows the location
// of the close relative to the high-low range over a set number of periods. The
// indicator can range from 0 to 100. The difference between the Slow and Fast
// Stochastic Oscillator is the Slow %K incorporates a %K slowing period of 3 that
// controls the internal smoothing of %K. Setting the smoothing period to 1 is
// equivalent to plotting the Fast Stochastic Oscillator.
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/slow-stochastic
func StochSlowArr(h, l, c []float64, fastKN int64, kt MaType, slowKN int64, dt MaType, slowDN int64) ([]float64, []float64) {
	k := make([]float64, len(c))
	d := make([]float64, len(c))

	s := NewStochSlow(fastKN, kt, slowKN, dt, slowDN)
	for i := 0; i < len(c); i++ {
		k[i], d[i] = s.Update(h[i], l[i], c[i])
	}

	return k, d
}
