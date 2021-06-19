package tart

// Developed by Donald Lambert and featured in Commodities
// magazine in 1980, the Commodity Channel Index (CCI) is
// a versatile indicator that can be used to identify a new
// trend or warn of extreme conditions. Lambert originally
// developed CCI to identify cyclical turns in commodities,
// but the indicator can be successfully applied to indices,
// ETFs, stocks and other securities. In general, CCI measures
// the current price level relative to an average price level
// over a given period of time. CCI is relatively high when
// prices are far above their average, but is relatively low
// when prices are far below their average. In this manner,
// CCI can be used to identify overbought and oversold levels.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:commodity_channel_index_cci
//  https://www.investopedia.com/terms/c/commoditychannelindex.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/cci
type Cci struct {
	n          int64
	initPeriod int64
	avg        *Sma
	dev        *Dev
	sz         int64
}

func NewCci(n int64) *Cci {
	avg := NewSma(n)
	dev := NewDev(n)
	a := avg.InitPeriod()
	b := dev.InitPeriod()
	if a < b {
		a = b
	}
	return &Cci{
		n:          n,
		initPeriod: a,
		avg:        avg,
		dev:        dev,
		sz:         0,
	}
}

func (d *Cci) Update(h, l, c float64) float64 {
	d.sz++

	m := (h + l + c) / 3.0
	avg := d.avg.Update(m)
	dev := d.dev.Update(m)

	if almostZero(dev) {
		return 0
	}

	return (m - avg) / (0.015 * dev)
}

func (d *Cci) InitPeriod() int64 {
	return d.initPeriod
}

func (d *Cci) Valid() bool {
	return d.sz > d.initPeriod
}

// Developed by Donald Lambert and featured in Commodities
// magazine in 1980, the Commodity Channel Index (CCI) is
// a versatile indicator that can be used to identify a new
// trend or warn of extreme conditions. Lambert originally
// developed CCI to identify cyclical turns in commodities,
// but the indicator can be successfully applied to indices,
// ETFs, stocks and other securities. In general, CCI measures
// the current price level relative to an average price level
// over a given period of time. CCI is relatively high when
// prices are far above their average, but is relatively low
// when prices are far below their average. In this manner,
// CCI can be used to identify overbought and oversold levels.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:commodity_channel_index_cci
//  https://www.investopedia.com/terms/c/commoditychannelindex.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/cci
func CciArr(h, l, c []float64, n int64) []float64 {
	out := make([]float64, len(h))

	d := NewCci(n)
	for i := 0; i < len(h); i++ {
		out[i] = d.Update(h[i], l[i], c[i])
	}

	return out
}
