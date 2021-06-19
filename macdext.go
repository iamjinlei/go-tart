package tart

// Refer to MACD.
// This is a general version of MACD with moving average types
// for fast, slow, and signal lines as paremters.
type MacdExt struct {
	fastN   int64
	slowN   int64
	signalN int64
	n       int64
	fast    *Ma
	slow    *Ma
	signal  *Ma
	sz      int64
}

func NewMacdExt(fastT MaType, fastN int64, slowT MaType, slowN int64, signalT MaType, signalN int64) *MacdExt {
	if slowN < fastN {
		fastN, slowN = slowN, fastN
	}
	return &MacdExt{
		fastN:   fastN,
		slowN:   slowN,
		signalN: signalN,
		n:       slowN - 1 + signalN,
		fast:    NewMa(fastT, fastN),
		slow:    NewMa(slowT, slowN),
		signal:  NewMa(signalT, signalN),
		sz:      0,
	}
}

func (m *MacdExt) Update(v float64) (float64, float64, float64) {
	m.sz++

	slow := m.slow.Update(v)

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

func (m *MacdExt) InitPeriod() int64 {
	return m.n - 1
}

func (m *MacdExt) Valid() bool {
	return m.sz > m.InitPeriod()
}

// Refer to MACD.
// This is a general version of MACD with moving average types
// for fast, slow, and signal lines as paremters.
func MacdExtArr(in []float64, fastT MaType, fastN int64, slowT MaType, slowN int64, signalT MaType, signalN int64) ([]float64, []float64, []float64) {
	macd := make([]float64, len(in))
	signal := make([]float64, len(in))
	hist := make([]float64, len(in))

	m := NewMacdExt(fastT, fastN, slowT, slowN, signalT, signalN)
	for i, v := range in {
		macd[i], signal[i], hist[i] = m.Update(v)
	}

	return macd, signal, hist
}
