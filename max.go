package tart

type Max struct {
	n      int64
	sz     int64
	hist   []float64
	max    float64
	maxIdx int64
}

func NewMax(n int64) *Max {
	return &Max{
		n:      n,
		sz:     0,
		hist:   make([]float64, n),
		max:    0,
		maxIdx: 0,
	}
}

func (m *Max) arrIdx(histIdx int64) int64 {
	// m.sz % m.n points to oldest value with array index m.sz - n
	idx := m.sz % m.n

	if histIdx < idx {
		return m.sz - (idx - histIdx)
	} else {
		return m.sz - m.n + (histIdx - idx)
	}
}

func (m *Max) Update(v float64) (int64, float64) {
	idx := m.sz % m.n
	m.hist[idx] = v

	m.sz++
	if m.sz < m.n {
		if m.sz == 0 || m.max < v {
			m.max = v
			m.maxIdx = idx
		}
		return 0, 0
	}

	if m.max <= v {
		// conforming to TA-Lib which updates maxIdx on equality
		m.max = v
		m.maxIdx = idx
	} else if m.maxIdx == idx {
		// previous max overwritten, recalulate
		m.max = m.hist[0]
		m.maxIdx = 0
		for i := int64(1); i < m.n; i++ {
			if m.max < m.hist[i] {
				m.max = m.hist[i]
				m.maxIdx = i
			}
		}
	}

	return m.arrIdx(m.maxIdx), m.max
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
