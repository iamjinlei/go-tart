package tart

// Developed by Gerald Appel in the late seventies, the
// Moving Average Convergence/Divergence oscillator (MACD)
// is one of the simplest and most effective momentum indicators
// available. The MACD turns two trend-following indicators,
// moving averages, into a momentum oscillator by subtracting
// the longer moving average from the shorter one. As a result,
// the MACD offers the best of both worlds: trend following and
// momentum. The MACD fluctuates above and below the zero line
// as the moving averages converge, cross and diverge. Traders
// can look for signal line crossovers, centerline crossovers
// and divergences to generate signals. Because the MACD is
// unbounded, it is not particularly useful for identifying
// overbought and oversold levels.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:moving_average_convergence_divergence_macd
//  https://www.investopedia.com/terms/m/macd.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/macd
type Macd struct {
	*MacdExt
}

func NewMacd(fastN, slowN, signalN int64) *Macd {
	return &Macd{
		MacdExt: NewMacdExt(EMA, fastN, EMA, slowN, EMA, signalN),
	}
}

// Developed by Gerald Appel in the late seventies, the
// Moving Average Convergence/Divergence oscillator (MACD)
// is one of the simplest and most effective momentum indicators
// available. The MACD turns two trend-following indicators,
// moving averages, into a momentum oscillator by subtracting
// the longer moving average from the shorter one. As a result,
// the MACD offers the best of both worlds: trend following and
// momentum. The MACD fluctuates above and below the zero line
// as the moving averages converge, cross and diverge. Traders
// can look for signal line crossovers, centerline crossovers
// and divergences to generate signals. Because the MACD is
// unbounded, it is not particularly useful for identifying
// overbought and oversold levels.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:moving_average_convergence_divergence_macd
//  https://www.investopedia.com/terms/m/macd.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/macd
func MacdArr(in []float64, fastN, slowN, signalN int64) ([]float64, []float64, []float64) {
	macd := make([]float64, len(in))
	signal := make([]float64, len(in))
	hist := make([]float64, len(in))

	m := NewMacd(fastN, slowN, signalN)
	for i, v := range in {
		macd[i], signal[i], hist[i] = m.Update(v)
	}

	return macd, signal, hist
}
