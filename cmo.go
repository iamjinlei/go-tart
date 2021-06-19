package tart

// The Chande momentum oscillator is a technical momentum
// indicator introduced by Tushar Chande in his 1994 book
// The New Technical Trader. The formula calculates the
// difference between the sum of recent gains and the sum
// of recent losses and then divides the result by the sum
// of all price movements over the same period.
//  https://www.investopedia.com/terms/c/chandemomentumoscillator.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/cmo
type Cmo struct {
	n          int64
	initPeriod int64
	su         *Ema
	sd         *Ema
	prevC      float64
	sz         int64
}

func NewCmo(n int64) *Cmo {
	k := 1.0 / float64(n)
	su := NewEma(n, k)
	sd := NewEma(n, k)
	a := su.InitPeriod()
	b := sd.InitPeriod()
	if a < b {
		a = b
	}
	return &Cmo{
		n:          n,
		initPeriod: a,
		su:         su,
		sd:         sd,
		prevC:      0,
		sz:         0,
	}
}

func (c *Cmo) Update(v float64) float64 {
	c.sz++

	d := v - c.prevC
	c.prevC = v

	if c.sz == 1 {
		return 0
	}

	var asu, asd float64
	if d > 0 {
		asu = c.su.Update(d)
		asd = c.sd.Update(0)
	} else {
		asu = c.su.Update(0)
		asd = c.sd.Update(-d)
	}

	sum := asu + asd
	if almostZero(sum) {
		return 0
	}

	return (asu - asd) / sum * 100.0
}

func (c *Cmo) InitPeriod() int64 {
	return c.initPeriod
}

func (c *Cmo) Valid() bool {
	return c.sz > c.initPeriod
}

// The Chande momentum oscillator is a technical momentum
// indicator introduced by Tushar Chande in his 1994 book
// The New Technical Trader. The formula calculates the
// difference between the sum of recent gains and the sum
// of recent losses and then divides the result by the sum
// of all price movements over the same period.
//  https://www.investopedia.com/terms/c/chandemomentumoscillator.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/cmo
func CmoArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	c := NewCmo(n)
	for i, v := range in {
		out[i] = c.Update(v)
	}

	return out
}
