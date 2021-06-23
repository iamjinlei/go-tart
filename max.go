package tart

// Max value of the selected period.
type Max struct {
	n      int64
	hist   *CBuf
	max    float64
	maxIdx int64
}

func NewMax(n int64) *Max {
	return &Max{
		n:      n,
		hist:   NewCBuf(n),
		max:    0,
		maxIdx: 0,
	}
}

func (m *Max) Update(v float64) (int64, float64) {
	m.hist.Append(v)

	if m.hist.Size() < m.n {
		return 0, 0
	}

	if m.hist.Size() == m.n || m.maxIdx == m.hist.NewestIndex() {
		m.maxIdx, m.max = m.hist.Max()
	} else if m.max <= v {
		// conforming to TA-Lib which updates maxIdx on equality
		m.max = v
		m.maxIdx = m.hist.NewestIndex()
	}

	return m.hist.IndexToSeq(m.maxIdx), m.max
}

func (m *Max) InitPeriod() int64 {
	return m.n - 1
}

func (m *Max) Valid() bool {
	return m.hist.Size() > m.InitPeriod()
}

// Max value of the selected period.
func MaxArr(in []float64, n int64) ([]int64, []float64) {
	outIdx := make([]int64, len(in))
	out := make([]float64, len(in))

	m := NewMax(n)
	for i, v := range in {
		outIdx[i], out[i] = m.Update(v)
	}

	return outIdx, out
}
