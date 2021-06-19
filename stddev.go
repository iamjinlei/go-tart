package tart

import (
	"math"
)

// Standard deviation is a statistical term that measures the amount of
// variability or dispersion around an average. Standard deviation is also
// a measure of volatility. Generally speaking, dispersion is the difference
// between the actual value and the average value. The larger this dispersion
// or variability is, the higher the standard deviation. The smaller this
// dispersion or variability is, the lower the standard deviation. Chartists
// can use the standard deviation to measure expected risk and determine
// the significance of certain price movements.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:standard_deviation_volatility
//  https://www.investopedia.com/terms/s/standarddeviation.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/standard-deviation
type StdDev struct {
	v *Var
}

func NewStdDev(n int64) *StdDev {
	return &StdDev{
		v: NewVar(n),
	}
}

func (s *StdDev) Update(v float64) float64 {
	v = s.v.Update(v)
	return math.Sqrt(v)
}

func (s *StdDev) InitPeriod() int64 {
	return s.v.InitPeriod()
}

func (s *StdDev) Valid() bool {
	return s.v.Valid()
}

// Standard deviation is a statistical term that measures the amount of
// variability or dispersion around an average. Standard deviation is also
// a measure of volatility. Generally speaking, dispersion is the difference
// between the actual value and the average value. The larger this dispersion
// or variability is, the higher the standard deviation. The smaller this
// dispersion or variability is, the lower the standard deviation. Chartists
// can use the standard deviation to measure expected risk and determine
// the significance of certain price movements.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:standard_deviation_volatility
//  https://www.investopedia.com/terms/s/standarddeviation.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/standard-deviation
func StdDevArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	s := NewStdDev(n)
	for i, v := range in {
		out[i] = s.Update(v)
	}

	return out
}
