package tart

// The Rate-of-Change (ROC) indicator, which is also referred to as simply
// Momentum, is a pure momentum oscillator that measures the percent change
// in price from one period to the next. The ROC calculation compares the
// current price with the price “n” periods ago. The plot forms an oscillator
// that fluctuates above and below the zero line as the Rate-of-Change moves
// from positive to negative. As a momentum oscillator, ROC signals include
// centerline crossovers, divergences and overbought-oversold readings.
// Divergences fail to foreshadow reversals more often than not, so this
// article will forgo a detailed discussion on them. Even though centerline
// crossovers are prone to whipsaw, especially short-term, these crossovers
// can be used to identify the overall trend. Identifying overbought or
// oversold extremes comes naturally to the Rate-of-Change oscillator.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:rate_of_change_roc_and_momentum
//  https://www.investopedia.com/terms/p/pricerateofchange.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/roc
type Roc struct {
	n    int64
	hist *CBuf
	sz   int64
}

func NewRoc(n int64) *Roc {
	return &Roc{
		n:    n,
		hist: NewCBuf(n),
		sz:   0,
	}
}

func (r *Roc) Update(v float64) float64 {
	r.sz++

	old := r.hist.Append(v)

	if r.sz <= r.n {
		return 0
	}

	if almostZero(old) {
		return 0
	}
	return (v - old) / old * 100.0
}

func (r *Roc) InitPeriod() int64 {
	return r.n
}

func (r *Roc) Valid() bool {
	return r.sz > r.InitPeriod()
}

// The Rate-of-Change (ROC) indicator, which is also referred to as simply
// Momentum, is a pure momentum oscillator that measures the percent change
// in price from one period to the next. The ROC calculation compares the
// current price with the price “n” periods ago. The plot forms an oscillator
// that fluctuates above and below the zero line as the Rate-of-Change moves
// from positive to negative. As a momentum oscillator, ROC signals include
// centerline crossovers, divergences and overbought-oversold readings.
// Divergences fail to foreshadow reversals more often than not, so this
// article will forgo a detailed discussion on them. Even though centerline
// crossovers are prone to whipsaw, especially short-term, these crossovers
// can be used to identify the overall trend. Identifying overbought or
// oversold extremes comes naturally to the Rate-of-Change oscillator.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:rate_of_change_roc_and_momentum
//  https://www.investopedia.com/terms/p/pricerateofchange.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/roc
func RocArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	r := NewRoc(n)
	for i, v := range in {
		out[i] = r.Update(v)
	}

	return out
}
