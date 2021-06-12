package tart

type Min struct {
	n      int64
	sz     int64
	hist   []float64
	min    float64
	minIdx int64
}

func NewMin(n int64) *Min {
	return &Min{
		n:      n,
		sz:     0,
		hist:   make([]float64, n),
		min:    0,
		minIdx: 0,
	}
}

func (m *Min) arrIdx(histIdx int64) int64 {
	// m.sz % m.n points to oldest value with array index m.sz - n
	idx := m.sz % m.n

	if histIdx < idx {
		return m.sz - (idx - histIdx)
	} else {
		return m.sz - m.n + (histIdx - idx)
	}
}

func (m *Min) Update(v float64) (int64, float64) {
	idx := m.sz % m.n
	m.hist[idx] = v

	m.sz++
	if m.sz < m.n {
		if m.sz == 1 || m.min > v {
			m.min = v
			m.minIdx = idx
		}
		return 0, 0
	}

	if m.min >= v {
		// conforming to TA-Lib which updates minIdx on equality
		m.min = v
		m.minIdx = idx
	} else if m.minIdx == idx {
		// previous min overwritten, recalulate
		m.min = m.hist[0]
		m.minIdx = 0
		for i := int64(1); i < m.n; i++ {
			if m.min > m.hist[i] {
				m.min = m.hist[i]
				m.minIdx = i
			}
		}
	}

	return m.arrIdx(m.minIdx), m.min
}

func MinArr(in []float64, n int64) ([]int64, []float64) {
	outIdx := make([]int64, len(in))
	out := make([]float64, len(in))

	m := NewMin(n)
	for i, v := range in {
		outIdx[i], out[i] = m.Update(v)
	}

	return outIdx, out
}
