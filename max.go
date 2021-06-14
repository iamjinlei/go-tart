package tart

type Max struct {
	n      int64
	hist   *cBuf
	max    float64
	maxIdx int64
}

func NewMax(n int64) *Max {
	return &Max{
		n:      n,
		hist:   newCBuf(n),
		max:    0,
		maxIdx: 0,
	}
}

func (m *Max) Update(v float64) (int64, float64) {
	m.hist.append(v)

	if m.hist.size() < m.n {
		return 0, 0
	}

	if m.hist.size() == m.n || m.maxIdx == m.hist.newestIndex() {
		m.maxIdx, m.max = m.hist.max()
	} else if m.max <= v {
		// conforming to TA-Lib which updates maxIdx on equality
		m.max = v
		m.maxIdx = m.hist.newestIndex()
	}

	return m.hist.indexToSeq(m.maxIdx), m.max
}

func MaxArr(in []float64, n int64) ([]int64, []float64) {
	outIdx := make([]int64, len(in))
	out := make([]float64, len(in))

	m := NewMax(n)
	for i, v := range in {
		outIdx[i], out[i] = m.Update(v)
	}

	return outIdx, out
}
