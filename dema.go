package tart

// The Double Exponential Moving Average (DEMA) reduces the lag
// of traditional EMAs, making it more responsive and better-suited
// for short-term traders. DEMA was developed by Patrick Mulloy,
// and introduced in the January 1994 issue of Technical Analysis
// of Stocks & Commodities magazine. The overlay uses the lag
// difference between a single-smoothed EMA and a double-smoothed
// EMA to offset the single-smoothed EMA. This offset produces a
// moving average that remains smooth, but stays closer to the
// price bars than either the single- or double-smoothed
// traditional EMA.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:dema
//  https://www.investopedia.com/terms/d/double-exponential-moving-average.asp
type Dema struct {
	n    int64
	ema1 *Ema
	ema2 *Ema
	sz   int64
}

func NewDema(n int64, k float64) *Dema {
	return &Dema{
		n:    n,
		ema1: NewEma(n, k),
		ema2: NewEma(n, k),
		sz:   0,
	}
}

func (d *Dema) Update(v float64) float64 {
	d.sz++

	e1 := d.ema1.Update(v)

	if d.sz > d.n-1 {
		e2 := d.ema2.Update(e1)
		if d.sz > d.n*2-2 {
			return 2.0*e1 - e2
		}
	}

	return 0
}

func (d *Dema) InitPeriod() int64 {
	return d.n*2 - 2
}

func (d *Dema) Valid() bool {
	return d.sz > d.InitPeriod()
}

// The Double Exponential Moving Average (DEMA) reduces the lag
// of traditional EMAs, making it more responsive and better-suited
// for short-term traders. DEMA was developed by Patrick Mulloy,
// and introduced in the January 1994 issue of Technical Analysis
// of Stocks & Commodities magazine. The overlay uses the lag
// difference between a single-smoothed EMA and a double-smoothed
// EMA to offset the single-smoothed EMA. This offset produces a
// moving average that remains smooth, but stays closer to the
// price bars than either the single- or double-smoothed
// traditional EMA.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:dema
//  https://www.investopedia.com/terms/d/double-exponential-moving-average.asp
func DemaArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	k := 2.0 / float64(n+1)
	d := NewDema(n, k)
	for i, v := range in {
		out[i] = d.Update(v)
	}

	return out
}
