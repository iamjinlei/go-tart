package tart

// A simple moving average is formed by computing the average price of a security
// over a specific number of periods. Most moving averages are based on closing
// prices; for example, a 5-day simple moving average is the five-day sum of closing
// prices divided by five. As its name implies, a moving average is an average that
// moves. Old data is dropped as new data becomes available, causing the average
// to move along the time scale. The example below shows a 5-day moving average
// evolving over three days.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:moving_averages
//  https://www.investopedia.com/terms/s/sma.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/sma
type Sma struct {
	n    int64
	hist *CBuf
	sz   int64
	sum  float64
}

func NewSma(n int64) *Sma {
	return &Sma{
		n:    n,
		hist: NewCBuf(n),
		sz:   0,
		sum:  0,
	}
}

func (s *Sma) Update(v float64) float64 {
	s.sz++

	old := s.hist.Append(v)
	s.sum += v - old

	if s.sz < s.n {
		return 0
	}

	return s.sum / float64(s.n)
}

func (s *Sma) InitPeriod() int64 {
	return s.n - 1
}

func (s *Sma) Valid() bool {
	return s.sz > s.InitPeriod()
}

// A simple moving average is formed by computing the average price of a security
// over a specific number of periods. Most moving averages are based on closing
// prices; for example, a 5-day simple moving average is the five-day sum of closing
// prices divided by five. As its name implies, a moving average is an average that
// moves. Old data is dropped as new data becomes available, causing the average
// to move along the time scale. The example below shows a 5-day moving average
// evolving over three days.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:moving_averages
//  https://www.investopedia.com/terms/s/sma.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/sma
func SmaArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	s := NewSma(n)
	for i, v := range in {
		out[i] = s.Update(v)
	}

	return out
}
