package tart

type Diff struct {
	n    int64
	hist *cBuf
	sz   int64
}

func NewDiff(n int64) *Diff {
	return &Diff{
		n:    n,
		hist: newCBuf(n),
		sz:   0,
	}
}

func (d *Diff) Update(v float64) float64 {
	old := d.hist.append(v)

	d.sz++
	if d.sz <= d.n {
		return 0
	}

	return v - old
}

func DiffArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	d := NewDiff(n)
	for i, v := range in {
		out[i] = d.Update(v)
	}

	return out
}
