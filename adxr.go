package tart

// Average Directional Movement Index Rating (ADXR) is a simple
// average of today’s ADX value and the ADX from N periods ago.
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/dmi
type AdxR struct {
	n    int64
	adx  *Adx
	hist *cBuf
}

func NewAdxR(n int64) *AdxR {
	return &AdxR{
		n:    n,
		adx:  NewAdx(n),
		hist: newCBuf(n - 1),
	}
}

func (a *AdxR) Update(h, l, c float64) float64 {
	v := a.adx.Update(h, l, c)
	old := a.hist.append(v)

	if a.hist.size() <= 3*a.n-2 {
		return 0
	}

	return (old + v) / 2.0
}

// Average Directional Movement Index Rating (ADXR) is a simple
// average of today’s ADX value and the ADX from N periods ago.
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/dmi
func AdxRArr(h, l, c []float64, n int64) []float64 {
	out := make([]float64, len(c))

	a := NewAdxR(n)
	for i := 0; i < len(c); i++ {
		out[i] = a.Update(h[i], l[i], c[i])
	}

	return out
}
