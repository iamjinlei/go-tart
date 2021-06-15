package tart

type Cci struct {
	n   int64
	avg *Sma
	dev *Dev
}

func NewCci(n int64) *Cci {
	return &Cci{
		n:   n,
		avg: NewSma(n),
		dev: NewDev(n),
	}
}

func (d *Cci) Update(h, l, c float64) float64 {
	m := (h + l + c) / 3.0
	avg := d.avg.Update(m)
	dev := d.dev.Update(m)

	if almostZero(dev) {
		return 0
	}

	return (m - avg) / (0.015 * dev)
}

func CciArr(h, l, c []float64, n int64) []float64 {
	out := make([]float64, len(h))

	d := NewCci(n)
	for i := 0; i < len(h); i++ {
		out[i] = d.Update(h[i], l[i], c[i])
	}

	return out
}
