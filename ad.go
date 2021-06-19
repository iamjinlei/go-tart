package tart

// Developed by Marc Chaikin, the Accumulation Distribution
// Line is a volume-based indicator designed to measure the
// cumulative flow of money into and out of a security.
// Chaikin originally referred to the indicator as the
// Cumulative Money Flow Line. As with cumulative indicators,
// the Accumulation Distribution Line is a running total of
// each period's Money Flow Volume. First, a multiplier is
// calculated based on the relationship of the close to the
// high-low range. Second, the Money Flow Multiplier is
// multiplied by the period's volume to come up with a
// Money Flow Volume. A running total of the Money Flow
// Volume forms the Accumulation Distribution Line.
// Chartists can use this indicator to affirm a security's
// underlying trend or anticipate reversals when the
// indicator diverges from the security price.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:accumulation_distribution_line
//  https://www.investopedia.com/terms/a/accumulationdistribution.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/accumulation-distribution
type Ad struct {
	ad float64
}

func NewAd() *Ad {
	return &Ad{
		ad: 0,
	}
}

func (a *Ad) Update(h, l, c, v float64) float64 {
	h2l := h - l
	if h2l > 0.0 {
		a.ad += (((c - l) - (h - c)) / h2l) * v
	}
	return a.ad
}

func (a *Ad) InitPeriod() int64 {
	return 0
}

func (a *Ad) Valid() bool {
	return true
}

// Developed by Marc Chaikin, the Accumulation Distribution
// Line is a volume-based indicator designed to measure the
// cumulative flow of money into and out of a security.
// Chaikin originally referred to the indicator as the
// Cumulative Money Flow Line. As with cumulative indicators,
// the Accumulation Distribution Line is a running total of
// each period's Money Flow Volume. First, a multiplier is
// calculated based on the relationship of the close to the
// high-low range. Second, the Money Flow Multiplier is
// multiplied by the period's volume to come up with a
// Money Flow Volume. A running total of the Money Flow
// Volume forms the Accumulation Distribution Line.
// Chartists can use this indicator to affirm a security's
// underlying trend or anticipate reversals when the
// indicator diverges from the security price.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:accumulation_distribution_line
//  https://www.investopedia.com/terms/a/accumulationdistribution.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/accumulation-distribution
func AdArr(h, l, c, v []float64) []float64 {
	out := make([]float64, len(c))

	a := NewAd()
	for i := 0; i < len(c); i++ {
		out[i] = a.Update(h[i], l[i], c[i], v[i])
	}
	return out
}
