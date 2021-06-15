package tart

type Macd struct {
	fastN   int64
	slowN   int64
	signalN int64
	n       int64
	fast    *Ema
	slow    *Ema
	signal  *Ema
	sz      int64
}

func NewMacd(fastN, slowN, signalN int64) *Macd {
	if slowN < fastN {
		fastN, slowN = slowN, fastN
	}
	return &Macd{
		fastN:   fastN,
		slowN:   slowN,
		signalN: signalN,
		fast:    NewEma(fastN, 2.0/float64(fastN+1)),
		slow:    NewEma(slowN, 2.0/float64(slowN+1)),
		signal:  NewEma(signalN, 2.0/float64(signalN+1)),
		n:       slowN - 1 + signalN,
		sz:      0,
	}
}

func (m *Macd) Update(v float64) (float64, float64, float64) {
	slow := m.slow.Update(v)
	m.sz++
	if m.sz <= m.slowN-m.fastN {
		// align the first valid result for fast & slow EMA
		return 0, 0, 0
	}

	fast := m.fast.Update(v)
	macd := fast - slow

	if m.sz < m.slowN {
		// wait until fast and slow EMAs are valid before populating signal EMA
		return 0, 0, 0
	}

	sig := m.signal.Update(macd)

	if m.sz < m.n {
		return 0, 0, 0
	}

	return macd, sig, macd - sig
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
