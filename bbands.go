package tart

type BBands struct {
	ma        *Ma
	stdDev    *StdDev
	upNStdDev float64
	dnNStdDev float64
}

func NewBBands(t MaType, n int64, upNStdDev, dnNStdDev float64) *BBands {
	return &BBands{
		ma:        NewMa(t, n),
		stdDev:    NewStdDev(n),
		upNStdDev: upNStdDev,
		dnNStdDev: dnNStdDev,
	}
}

// upper, middle, lower
func (b *BBands) Update(v float64) (float64, float64, float64) {
	m := b.ma.Update(v)
	stddev := b.stdDev.Update(v)

	return m + b.upNStdDev*stddev, m, m - b.upNStdDev*stddev
}

func BBandsArr(t MaType, in []float64, n int64, upNStdDev, dnNStdDev float64) ([]float64, []float64, []float64) {
	m := make([]float64, len(in))
	u := make([]float64, len(in))
	l := make([]float64, len(in))

	b := NewBBands(t, n, upNStdDev, dnNStdDev)
	for i, v := range in {
		u[i], m[i], l[i] = b.Update(v)
	}

	return u, m, l
}
