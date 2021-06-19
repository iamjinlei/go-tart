package tart

// TRIX is a momentum oscillator that displays the percent rate of change of
// a triple exponentially smoothed moving average. It was developed in the
// early 1980's by Jack Hutson, an editor for Technical Analysis of Stocks
// and Commodities magazine. With its triple smoothing, TRIX is designed to
// filter out insignificant price movements. Chartists can use TRIX to
// generate signals similar to MACD. A signal line can be applied to look
// for signal line crossovers. A directional bias can be determined with the
// absolute level. Bullish and bearish divergences can be used to anticipate
// reversals.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:trix
//  https://www.investopedia.com/terms/t/trix.asp
type Trix struct {
	n    int64
	ema1 *Ema
	ema2 *Ema
	ema3 *Ema
	prev float64
	sz   int64
}

func NewTrix(n int64) *Trix {
	k := 2.0 / float64(n+1)
	return &Trix{
		n:    n,
		ema1: NewEma(n, k),
		ema2: NewEma(n, k),
		ema3: NewEma(n, k),
		prev: 0,
		sz:   0,
	}
}

func (t *Trix) Update(v float64) float64 {
	t.sz++

	v = t.ema1.Update(v)
	if t.sz < t.n {
		return 0
	}

	v = t.ema2.Update(v)
	if t.sz < 2*t.n-1 {
		return 0
	}

	v = t.ema3.Update(v)
	roc := float64(0)
	if !almostZero(t.prev) {
		roc = (v - t.prev) / t.prev
	}
	t.prev = v
	if t.sz < 3*t.n-2 {
		return 0
	}

	return roc * 100.0
}

func (t *Trix) InitPeriod() int64 {
	return 3*t.n - 3
}

func (t *Trix) Valid() bool {
	return t.sz > t.InitPeriod()
}

// TRIX is a momentum oscillator that displays the percent rate of change of
// a triple exponentially smoothed moving average. It was developed in the
// early 1980's by Jack Hutson, an editor for Technical Analysis of Stocks
// and Commodities magazine. With its triple smoothing, TRIX is designed to
// filter out insignificant price movements. Chartists can use TRIX to
// generate signals similar to MACD. A signal line can be applied to look
// for signal line crossovers. A directional bias can be determined with the
// absolute level. Bullish and bearish divergences can be used to anticipate
// reversals.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:trix
//  https://www.investopedia.com/terms/t/trix.asp
func TrixArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	t := NewTrix(n)
	for i, v := range in {
		out[i] = t.Update(v)
	}

	return out
}
