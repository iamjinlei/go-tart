package tart

// Developed by John Bollinger, Bollinger Bands are volatility
// bands placed above and below a moving average. Volatility
// is based on the standard deviation, which changes as volatility
// increases and decreases. The bands automatically widen when
// volatility increases and contract when volatility decreases.
// Their dynamic nature allows them to be used on different
// securities with the standard settings. Bollinger Bands can be
// used to identify M-Tops and W-Bottoms or to determine the
// strength of the trend. Signals based on the distance between
// the upper and lower band, including the popular Bollinger Band
// Squeeze, are identified using the related Bollinger BandWidth
// indicator.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:bollinger_bands
//  https://www.investopedia.com/terms/b/bollingerbands.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/bollinger-bands
type BBands struct {
	initPeriod int64
	ma         *Ma
	stdDev     *StdDev
	upNStdDev  float64
	dnNStdDev  float64
	sz         int64
}

func NewBBands(t MaType, n int64, upNStdDev, dnNStdDev float64) *BBands {
	ma := NewMa(t, n)
	stdDev := NewStdDev(n)
	a := ma.InitPeriod()
	b := stdDev.InitPeriod()
	if a < b {
		a = b
	}
	return &BBands{
		initPeriod: a,
		ma:         ma,
		stdDev:     stdDev,
		upNStdDev:  upNStdDev,
		dnNStdDev:  dnNStdDev,
		sz:         0,
	}
}

// upper, middle, lower
func (b *BBands) Update(v float64) (float64, float64, float64) {
	b.sz++

	m := b.ma.Update(v)
	stddev := b.stdDev.Update(v)

	return m + b.upNStdDev*stddev, m, m - b.upNStdDev*stddev
}

func (b *BBands) InitPeriod() int64 {
	return b.initPeriod
}

func (b *BBands) Valid() bool {
	return b.sz > b.initPeriod
}

// Developed by John Bollinger, Bollinger Bands are volatility
// bands placed above and below a moving average. Volatility
// is based on the standard deviation, which changes as volatility
// increases and decreases. The bands automatically widen when
// volatility increases and contract when volatility decreases.
// Their dynamic nature allows them to be used on different
// securities with the standard settings. Bollinger Bands can be
// used to identify M-Tops and W-Bottoms or to determine the
// strength of the trend. Signals based on the distance between
// the upper and lower band, including the popular Bollinger Band
// Squeeze, are identified using the related Bollinger BandWidth
// indicator.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:bollinger_bands
//  https://www.investopedia.com/terms/b/bollingerbands.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/bollinger-bands
func BBandsArr(t MaType, in []float64, n int64, upNStdDev, dnNStdDev float64) ([]float64, []float64, []float64) {
	m := make([]float64, len(in))
	u := make([]float64, len(in))
	l := make([]float64, len(in))

	b := NewBBands(t, n, upNStdDev, dnNStdDev)
	for i, v := range in {
		u[i], m[i], l[i] = b.Update(v)
	}

	return u, m, l
}
