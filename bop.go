package tart

// Balance of Power (BOP) is an oscillator that measures the
// strength of buying and selling pressure. Introduced by Igor
// Levshin in the August 2001 issue of Technical Analysis of
// Stocks & Commodities magazine, this indicator compares the
// power of buyers to push prices to higher extremes with the
// power of sellers to move prices to lower extremes. When the
// indicator is in positive territory, the bulls are in charge;
// and sellers dominate when the indicator is negative. A reading
// near the zero line indicates a balance between the two and
// can mean a trend reversal.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:balance_of_power
//  https://www.investopedia.com/terms/b/bop.asp
type Bop struct {
}

func NewBop() *Bop {
	return &Bop{}
}

func (b *Bop) Update(o, h, l, c float64) float64 {
	d := h - l
	if almostZero(d) {
		return 0
	}
	return (c - o) / d
}

func (b *Bop) InitPeriod() int64 {
	return 0
}

func (b *Bop) Valid() bool {
	return true
}

// Balance of Power (BOP) is an oscillator that measures the
// strength of buying and selling pressure. Introduced by Igor
// Levshin in the August 2001 issue of Technical Analysis of
// Stocks & Commodities magazine, this indicator compares the
// power of buyers to push prices to higher extremes with the
// power of sellers to move prices to lower extremes. When the
// indicator is in positive territory, the bulls are in charge;
// and sellers dominate when the indicator is negative. A reading
// near the zero line indicates a balance between the two and
// can mean a trend reversal.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:balance_of_power
//  https://www.investopedia.com/terms/b/bop.asp
func BopArr(o, h, l, c []float64) []float64 {
	out := make([]float64, len(o))

	b := NewBop()
	for i := 0; i < len(o); i++ {
		out[i] = b.Update(o[i], h[i], l[i], c[i])
	}

	return out
}
