package tart

type Aroon struct {
	n   int64
	min *Min
	max *Max
	sz  int64
}

// Developed by Tushar Chande in 1995, Aroon is an indicator
// system that determines whether a stock is trending or not
// and how strong the trend is. “Aroon” means “Dawn's Early Light”
// in Sanskrit. Chande chose this name because the indicators
// are designed to reveal the beginning of a new trend. The
// Aroon indicators measure the number of periods since price
// recorded an x-day high or low. There are two separate
// indicators: Aroon-Up and Aroon-Down. A 25-day Aroon-Up
// measures the number of days since a 25-day high. A 25-day
// Aroon-Down measures the number of days since a 25-day low.
// In this sense, the Aroon indicators are quite different
// from typical momentum oscillators, which focus on price
// relative to time. Aroon is unique because it focuses on
// time relative to price. Chartists can use the Aroon
// indicators to spot emerging trends, identify consolidations,
// define correction periods and anticipate reversals.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:aroon
//  https://www.investopedia.com/terms/a/aroon.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/aroon-indicator
func NewAroon(n int64) *Aroon {
	return &Aroon{
		n:   n,
		min: NewMin(n + 1),
		max: NewMax(n + 1),
		sz:  0,
	}
}

func (a *Aroon) aroonValue(idx int64) float64 {
	return float64(a.n-(a.sz-1-idx)) / float64(a.n) * 100.0
}

func (a *Aroon) Update(h, l float64) (float64, float64) {
	a.sz++

	minIdx, _ := a.min.Update(l)
	maxIdx, _ := a.max.Update(h)

	if a.sz <= a.n {
		return 0, 0
	}

	return a.aroonValue(minIdx), a.aroonValue(maxIdx)
}

func (a *Aroon) InitPeriod() int64 {
	return a.n
}

func (a *Aroon) Valid() bool {
	return a.sz > a.InitPeriod()
}

// Developed by Tushar Chande in 1995, Aroon is an indicator
// system that determines whether a stock is trending or not
// and how strong the trend is. “Aroon” means “Dawn's Early Light”
// in Sanskrit. Chande chose this name because the indicators
// are designed to reveal the beginning of a new trend. The
// Aroon indicators measure the number of periods since price
// recorded an x-day high or low. There are two separate
// indicators: Aroon-Up and Aroon-Down. A 25-day Aroon-Up
// measures the number of days since a 25-day high. A 25-day
// Aroon-Down measures the number of days since a 25-day low.
// In this sense, the Aroon indicators are quite different
// from typical momentum oscillators, which focus on price
// relative to time. Aroon is unique because it focuses on
// time relative to price. Chartists can use the Aroon
// indicators to spot emerging trends, identify consolidations,
// define correction periods and anticipate reversals.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:aroon
//  https://www.investopedia.com/terms/a/aroon.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/aroon-indicator
func AroonArr(h, l []float64, n int64) ([]float64, []float64) {
	dn := make([]float64, len(h))
	up := make([]float64, len(h))

	a := NewAroon(n)
	for i := 0; i < len(h); i++ {
		dn[i], up[i] = a.Update(h[i], l[i])
	}

	return dn, up
}
