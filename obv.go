package tart

// On Balance Volume (OBV) measures buying and selling
// pressure as a cumulative indicator, adding volume on
// up days and subtracting it on down days. OBV was
// developed by Joe Granville and introduced in his 1963
// book Granville's New Key to Stock Market Profits.
// It was one of the first indicators to measure positive
// and negative volume flow. Chartists can look for
// divergences between OBV and price to predict price
// movements or use OBV to confirm price trends.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:on_balance_volume_obv
//  https://www.investopedia.com/terms/o/onbalancevolume.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/obv
type Obv struct {
	prev float64
	obv  float64
	sz   int64
}

func NewObv() *Obv {
	return &Obv{
		prev: 0,
		obv:  0,
		sz:   0,
	}
}

func (o *Obv) Update(c, v float64) float64 {
	o.sz++

	prev := o.prev
	o.prev = c
	if o.sz == 1 {
		o.obv = v
		return o.obv
	}

	if c > prev {
		o.obv += v
	} else if c < prev {
		o.obv -= v
	}

	return o.obv
}

func (o *Obv) InitPeriod() int64 {
	return 0
}

func (o *Obv) Valid() bool {
	return true
}

// On Balance Volume (OBV) measures buying and selling
// pressure as a cumulative indicator, adding volume on
// up days and subtracting it on down days. OBV was
// developed by Joe Granville and introduced in his 1963
// book Granville's New Key to Stock Market Profits.
// It was one of the first indicators to measure positive
// and negative volume flow. Chartists can look for
// divergences between OBV and price to predict price
// movements or use OBV to confirm price trends.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:on_balance_volume_obv
//  https://www.investopedia.com/terms/o/onbalancevolume.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/obv
func ObvArr(c, v []float64) []float64 {
	out := make([]float64, len(c))

	o := NewObv()
	for i := 0; i < len(c); i++ {
		out[i] = o.Update(c[i], v[i])
	}

	return out
}
