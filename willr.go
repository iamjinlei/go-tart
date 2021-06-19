package tart

// Developed by Larry Williams, Williams %R is a momentum indicator that is
// the inverse of the Fast Stochastic Oscillator. Also referred to as %R,
// Williams %R reflects the level of the close relative to the highest high
// for the look-back period. In contrast, the Stochastic Oscillator reflects
// the level of the close relative to the lowest low. %R corrects for the
// inversion by multiplying the raw value by -100. As a result, the Fast
// Stochastic Oscillator and Williams %R produce the exact same lines, but
// with different scaling. Williams %R oscillates from 0 to -100; readings
// from 0 to -20 are considered overbought, while readings from -80 to -100
// are considered oversold. Unsurprisingly, signals derived from the Stochastic
// Oscillator are also applicable to Williams %R.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:williams_r
//  https://www.investopedia.com/terms/w/williamsr.asp
type WillR struct {
	n     int64
	stoch *StochasticK
	sz    int64
}

func NewWillR(n int64) *WillR {
	return &WillR{
		n:     n,
		stoch: NewStochasticK(n),
		sz:    0,
	}
}

func (w *WillR) Update(h, l, c float64) float64 {
	w.sz++

	k := w.stoch.Update(h, l, c)

	if w.sz < w.n {
		return 0
	}
	return k - 100.0
}

func (w *WillR) InitPeriod() int64 {
	return w.n - 1
}

func (w *WillR) Valid() bool {
	return w.sz > w.InitPeriod()
}

// Developed by Larry Williams, Williams %R is a momentum indicator that is
// the inverse of the Fast Stochastic Oscillator. Also referred to as %R,
// Williams %R reflects the level of the close relative to the highest high
// for the look-back period. In contrast, the Stochastic Oscillator reflects
// the level of the close relative to the lowest low. %R corrects for the
// inversion by multiplying the raw value by -100. As a result, the Fast
// Stochastic Oscillator and Williams %R produce the exact same lines, but
// with different scaling. Williams %R oscillates from 0 to -100; readings
// from 0 to -20 are considered overbought, while readings from -80 to -100
// are considered oversold. Unsurprisingly, signals derived from the Stochastic
// Oscillator are also applicable to Williams %R.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:williams_r
//  https://www.investopedia.com/terms/w/williamsr.asp
func WillRArr(h, l, c []float64, n int64) []float64 {
	out := make([]float64, len(c))

	w := NewWillR(n)
	for i := 0; i < len(c); i++ {
		out[i] = w.Update(h[i], l[i], c[i])
	}

	return out
}
