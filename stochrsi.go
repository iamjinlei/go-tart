package tart

// Developed by Tushar Chande and Stanley Kroll, StochRSI is an oscillator that
// measures the level of RSI relative to its high-low range over a set time period.
// StochRsi applies the Stochastics formula to RSI values, rather than price values,
// making it an indicator of an indicator. The result is an oscillator that
// fluctuates between 0 and 1. In their 1994 book, The New Technical Trader, Chande
// and Kroll explain that RSI can oscillate between 80 and 20 for extended periods
// without reaching extreme levels. Notice that 80 and 20 are used for overbought
// and oversold instead of the more traditional 70 and 30. Traders looking to enter
// a stock based on an overbought or oversold reading in RSI might find themselves
// continuously on the sidelines. Chande and Kroll developed StochRSI to increase
// sensitivity and generate more overbought/oversold signals.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:stochrsi
//  https://www.investopedia.com/terms/s/stochrsi.asp
type StochRsi struct {
	n     int64
	kN    int64
	dN    int64
	util  int64
	rsi   *Rsi
	stoch *StochFast
	sz    int64
}

func NewStochRsi(n int64, kN int64, dt MaType, dN int64) *StochRsi {
	return &StochRsi{
		n:     n,
		kN:    kN,
		dN:    dN,
		util:  n + kN + dN - 1,
		rsi:   NewRsi(n),
		stoch: NewStochFast(kN, dt, dN),
		sz:    0,
	}
}

func (s *StochRsi) Update(v float64) (float64, float64) {
	s.sz++

	rsi := s.rsi.Update(v)

	if s.sz <= s.n {
		return 0, 0
	}
	k, d := s.stoch.Update(rsi, rsi, rsi)
	if s.sz < s.util {
		return 0, 0
	}

	return k, d
}

func (s *StochRsi) InitPeriod() int64 {
	return s.util - 1
}

func (s *StochRsi) Valid() bool {
	return s.sz > s.InitPeriod()
}

// Developed by Tushar Chande and Stanley Kroll, StochRSI is an oscillator that
// measures the level of RSI relative to its high-low range over a set time period.
// StochRsi applies the Stochastics formula to RSI values, rather than price values,
// making it an indicator of an indicator. The result is an oscillator that
// fluctuates between 0 and 1. In their 1994 book, The New Technical Trader, Chande
// and Kroll explain that RSI can oscillate between 80 and 20 for extended periods
// without reaching extreme levels. Notice that 80 and 20 are used for overbought
// and oversold instead of the more traditional 70 and 30. Traders looking to enter
// a stock based on an overbought or oversold reading in RSI might find themselves
// continuously on the sidelines. Chande and Kroll developed StochRSI to increase
// sensitivity and generate more overbought/oversold signals.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:stochrsi
//  https://www.investopedia.com/terms/s/stochrsi.asp
func StochRsiArr(in []float64, n, kN int64, dt MaType, dN int64) ([]float64, []float64) {
	k := make([]float64, len(in))
	d := make([]float64, len(in))

	s := NewStochRsi(n, kN, dt, dN)
	for i, v := range in {
		k[i], d[i] = s.Update(v)
	}

	return k, d
}
