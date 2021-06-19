package tart

// The Aroon Oscillator is the difference between Aroon-Up
// and Aroon-Down. These two indicators are usually plotted
// together for easy comparison, but chartists can also view
// the difference between these two indicators with the Aroon
// Oscillator. This indicator fluctuates between -100 and +100
// with zero as the middle line. An upward trend bias is present
// when the oscillator is positive, while a downward trend bias
// exists when the oscillator is negative. Chartists can also
// expand the bull-bear threshold to identify stronger signals.
// See our ChartSchool article for more details on Aroon-Up and
// Aroon-Down.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:aroon_oscillator
//  https://www.investopedia.com/terms/a/aroonoscillator.asp
type AroonOsc struct {
	n     int64
	aroon *Aroon
}

func NewAroonOsc(n int64) *AroonOsc {
	return &AroonOsc{
		n:     n,
		aroon: NewAroon(n),
	}
}

func (a *AroonOsc) Update(h, l float64) float64 {
	dn, up := a.aroon.Update(h, l)
	return up - dn
}

func (a *AroonOsc) InitPeriod() int64 {
	return a.aroon.InitPeriod()
}

func (a *AroonOsc) Valid() bool {
	return a.aroon.Valid()
}

// The Aroon Oscillator is the difference between Aroon-Up
// and Aroon-Down. These two indicators are usually plotted
// together for easy comparison, but chartists can also view
// the difference between these two indicators with the Aroon
// Oscillator. This indicator fluctuates between -100 and +100
// with zero as the middle line. An upward trend bias is present
// when the oscillator is positive, while a downward trend bias
// exists when the oscillator is negative. Chartists can also
// expand the bull-bear threshold to identify stronger signals.
// See our ChartSchool article for more details on Aroon-Up and
// Aroon-Down.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:aroon_oscillator
//  https://www.investopedia.com/terms/a/aroonoscillator.asp
func AroonOscArr(h, l []float64, n int64) []float64 {
	out := make([]float64, len(h))

	a := NewAroonOsc(n)
	for i := 0; i < len(h); i++ {
		out[i] = a.Update(h[i], l[i])
	}

	return out
}
