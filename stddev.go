package tart

import (
	"math"
)

type StdDev struct {
	v *Var
}

func NewStdDev(n int64) *StdDev {
	return &StdDev{
		v: NewVar(n),
	}
}

func (s *StdDev) Update(v float64) float64 {
	v = s.v.Update(v)
	return math.Sqrt(v)
}

func StdDevArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	s := NewStdDev(n)
	for i, v := range in {
		out[i] = s.Update(v)
	}

	return out
}
