package tart

import (
	"math"
)

type Dev struct {
	n    int64
	hist *cBuf
	sum  float64
}

func NewDev(n int64) *Dev {
	return &Dev{
		n:    n,
		hist: newCBuf(n),
		sum:  0,
	}
}

func (d *Dev) Update(v float64) float64 {
	old := d.hist.append(v)
	d.sum += v - old

	if d.hist.size() < d.n {
		return 0
	}

	mean := d.sum / float64(d.n)
	sum := float64(0)
	d.hist.iter(func(v float64) {
		sum += math.Abs(v - mean)
	})

	return sum / float64(d.n)
}

func VarDev(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	d := NewDev(n)
	for i, v := range in {
		out[i] = d.Update(v)
	}

	return out
}
