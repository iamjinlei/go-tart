package tart

type Macd struct {
	*MacdExt
}

func NewMacd(fastN, slowN, signalN int64) *Macd {
	return &Macd{
		MacdExt: NewMacdExt(EMA, fastN, EMA, slowN, EMA, signalN),
	}
}

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
