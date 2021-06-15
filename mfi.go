package tart

type Mfi struct {
	n        int64
	positive *Sum
	negative *Sum
	prevTp   float64
	sz       int64
}

func NewMfi(n int64) *Mfi {
	return &Mfi{
		n:        n,
		positive: NewSum(n),
		negative: NewSum(n),
		prevTp:   0,
		sz:       0,
	}
}

func (m *Mfi) Update(h, l, c, v float64) float64 {
	tp := (h + l + c) / 3.0
	prevTp := m.prevTp
	m.prevTp = tp
	m.sz++

	if m.sz == 1 {
		return 0
	}

	var pSum, nSum float64
	if tp > prevTp {
		pSum = m.positive.Update(tp * v)
		nSum = m.negative.Update(0)
	} else {
		pSum = m.positive.Update(0)
		nSum = m.negative.Update(tp * v)
	}

	if m.sz <= m.n {
		return 0
	}

	sum := pSum + nSum
	return pSum / sum * 100.0
}

func MfiArr(h, l, c, v []float64, n int64) []float64 {
	out := make([]float64, len(c))

	m := NewMfi(n)
	for i := 0; i < len(c); i++ {
		out[i] = m.Update(h[i], l[i], c[i], v[i])
	}

	return out
}
