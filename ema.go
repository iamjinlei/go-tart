package tart

// Exponential moving averages (EMAs) reduce the lag by
// applying more weight to recent prices. The weighting
// applied to the most recent price depends on the number
// of periods in the moving average. EMAs differ from
// simple moving averages in that a given day's EMA
// calculation depends on the EMA calculations for all
// the days prior to that day. You need far more than 10
// days of data to calculate a reasonably accurate 10-day EMA.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:moving_averages
type Ema struct {
	n  int64
	k1 float64
	k2 float64
	sz int64
	ma float64
}

func NewEma(n int64, k float64) *Ema {
	return &Ema{
		n:  n,
		k1: k,
		k2: 1 - k,
		sz: 0,
		ma: 0,
	}
}

func (e *Ema) Update(v float64) float64 {
	e.sz++

	if e.sz <= e.n {
		e.ma += v / float64(e.n)
		if e.sz < e.n {
			return 0
		}
	} else {
		e.ma = v*e.k1 + e.ma*e.k2
	}

	return e.ma
}

func (e *Ema) InitPeriod() int64 {
	return e.n - 1
}

func (e *Ema) Valid() bool {
	return e.sz > e.InitPeriod()
}

// Exponential moving averages (EMAs) reduce the lag by
// applying more weight to recent prices. The weighting
// applied to the most recent price depends on the number
// of periods in the moving average. EMAs differ from
// simple moving averages in that a given day's EMA
// calculation depends on the EMA calculations for all
// the days prior to that day. You need far more than 10
// days of data to calculate a reasonably accurate 10-day EMA.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:moving_averages
func EmaArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	k := 2.0 / float64(n+1)
	e := NewEma(n, k)
	for i, v := range in {
		out[i] = e.Update(v)
	}

	return out
}
