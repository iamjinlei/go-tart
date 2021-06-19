package tart

// The Average Directional Index (ADX), Minus Directional
// Indicator (-DI) and Plus Directional Indicator (+DI)
// represent a group of directional movement indicators
// that form a trading system developed by Welles Wilder.
// Although Wilder designed his Directional Movement System
// with commodities and daily prices in mind, these indicators
// can also be applied to stocks.
//
// Positive and negative directional movement form the backbone
// of the Directional Movement System. Wilder determined
// directional movement by comparing the difference between
// two consecutive lows with the difference between their
// respective highs.
//
// The Plus Directional Indicator (+DI) and Minus Directional
// Indicator (-DI) are derived from smoothed averages of
// these differences and measure trend direction over time.
// These two indicators are often collectively referred to
// as the Directional Movement Indicator (DMI).
//
// The Average Directional Index (ADX) is in turn derived
// from the smoothed averages of the difference between +DI
// and -DI; it measures the strength of the trend
// (regardless of direction) over time.
//
// Using these three indicators together, chartists can
// determine both the direction and strength of the trend.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:average_directional_index_adx
//  https://www.investopedia.com/articles/trading/07/adx-trend-indicator.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/adx
type Adx struct {
	n   int64
	dx  *Dx
	adx *Ema
	sz  int64
}

func NewAdx(n int64) *Adx {
	return &Adx{
		n:   n,
		dx:  NewDx(n),
		adx: NewEma(n, 1.0/float64(n)),
		sz:  0,
	}
}

func (a *Adx) Update(h, l, c float64) float64 {
	a.sz++

	dx := a.dx.Update(h, l, c)

	if a.sz <= a.n {
		return 0
	}

	return a.adx.Update(dx)
}

func (a *Adx) InitPeriod() int64 {
	return a.n
}

func (a *Adx) Valid() bool {
	return a.sz > a.InitPeriod()
}

// The Average Directional Index (ADX), Minus Directional
// Indicator (-DI) and Plus Directional Indicator (+DI)
// represent a group of directional movement indicators
// that form a trading system developed by Welles Wilder.
// Although Wilder designed his Directional Movement System
// with commodities and daily prices in mind, these indicators
// can also be applied to stocks.
//
// Positive and negative directional movement form the backbone
// of the Directional Movement System. Wilder determined
// directional movement by comparing the difference between
// two consecutive lows with the difference between their
// respective highs.
//
// The Plus Directional Indicator (+DI) and Minus Directional
// Indicator (-DI) are derived from smoothed averages of
// these differences and measure trend direction over time.
// These two indicators are often collectively referred to
// as the Directional Movement Indicator (DMI).
//
// The Average Directional Index (ADX) is in turn derived
// from the smoothed averages of the difference between +DI
// and -DI; it measures the strength of the trend
// (regardless of direction) over time.
//
// Using these three indicators together, chartists can
// determine both the direction and strength of the trend.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:average_directional_index_adx
//  https://www.investopedia.com/articles/trading/07/adx-trend-indicator.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/adx
func AdxArr(h, l, c []float64, n int64) []float64 {
	out := make([]float64, len(c))

	a := NewAdx(n)
	for i := 0; i < len(c); i++ {
		out[i] = a.Update(h[i], l[i], c[i])
	}

	return out
}
