package tart

// Min value of the selected period.
type Min struct {
	n      int64
	hist   *CBuf
	min    float64
	minIdx int64
}

func NewMin(n int64) *Min {
	return &Min{
		n:      n,
		hist:   NewCBuf(n),
		min:    0,
		minIdx: 0,
	}
}

func (m *Min) Update(v float64) (int64, float64) {
	m.hist.Append(v)

	if m.hist.Size() < m.n {
		return 0, 0
	}

	if m.hist.Size() == m.n || m.minIdx == m.hist.NewestIndex() {
		m.minIdx, m.min = m.hist.Min()
	} else if m.min >= v {
		// conforming to TA-Lib which updates minIdx on equality
		m.min = v
		m.minIdx = m.hist.NewestIndex()
	}

	return m.hist.IndexToSeq(m.minIdx), m.min
}

func (m *Min) InitPeriod() int64 {
	return m.n - 1
}

func (m *Min) Valid() bool {
	return m.hist.Size() > m.InitPeriod()
}

// Min value of the selected period.
func MinArr(in []float64, n int64) ([]int64, []float64) {
	outIdx := make([]int64, len(in))
	out := make([]float64, len(in))

	m := NewMin(n)
	for i, v := range in {
		outIdx[i], out[i] = m.Update(v)
	}

	return outIdx, out
}
