package tart

// The Percentage Price Oscillator (PPO) is a momentum oscillator
// that measures the difference between two moving averages as a
// percentage of the larger moving average. As with its cousin,
// MACD, the Percentage Price Oscillator is shown with a signal line,
// a histogram and a centerline. Signals are generated with signal
// line crossovers, centerline crossovers, and divergences. These
// signals are no different than those associated with MACD, with a
// few differences between the two: first, PPO readings are not
// subject to the price level of the security. Second, PPO readings
// for different securities can be compared, even when there are
// large differences in the price.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:price_oscillators_ppo
//  https://www.investopedia.com/terms/p/ppo.asp
type Ppo struct {
	fastN int64
	slowN int64
	fast  *Ma
	slow  *Ma
	sz    int64
}

func NewPpo(t MaType, fastN, slowN int64) *Ppo {
	if slowN < fastN {
		fastN, slowN = slowN, fastN
	}
	return &Ppo{
		fastN: fastN,
		slowN: slowN,
		fast:  NewMa(t, fastN),
		slow:  NewMa(t, slowN),
		sz:    0,
	}
}

func (p *Ppo) Update(v float64) float64 {
	p.sz++

	fast := p.fast.Update(v)
	slow := p.slow.Update(v)

	if p.sz < p.slowN {
		return 0
	}

	if almostZero(slow) {
		return 0
	}
	return (fast - slow) / slow * 100.0
}

func (p *Ppo) InitPeriod() int64 {
	return p.slowN - 1
}

func (p *Ppo) Valid() bool {
	return p.sz > p.InitPeriod()
}

// The Percentage Price Oscillator (PPO) is a momentum oscillator
// that measures the difference between two moving averages as a
// percentage of the larger moving average. As with its cousin,
// MACD, the Percentage Price Oscillator is shown with a signal line,
// a histogram and a centerline. Signals are generated with signal
// line crossovers, centerline crossovers, and divergences. These
// signals are no different than those associated with MACD, with a
// few differences between the two: first, PPO readings are not
// subject to the price level of the security. Second, PPO readings
// for different securities can be compared, even when there are
// large differences in the price.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:price_oscillators_ppo
//  https://www.investopedia.com/terms/p/ppo.asp
func PpoArr(in []float64, t MaType, fastN, slowN int64) []float64 {
	out := make([]float64, len(in))

	p := NewPpo(t, fastN, slowN)
	for i, v := range in {
		out[i] = p.Update(v)
	}

	return out
}
