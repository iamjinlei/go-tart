package tart

// This is also known as Momentum (MOM).
// The Momentum (MOM) indicator compares the current price with
// the previous price from a selected number of periods ago.
// This indicator is similar to the “Rate of Change” indicator,
// but the MOM does not normalize the price, so different
// instruments can have different indicator values based on
// their point values.
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

// This is also known as Momentum (MOM).
// The Momentum (MOM) indicator compares the current price with
// the previous price from a selected number of periods ago.
// This indicator is similar to the “Rate of Change” indicator,
// but the MOM does not normalize the price, so different
// instruments can have different indicator values based on
// their point values.
func DiffArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	d := NewDiff(n)
	for i, v := range in {
		out[i] = d.Update(v)
	}

	return out
}
