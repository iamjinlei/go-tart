package tart

// The term variance refers to a statistical measurement of the spread between
// numbers in a data set. More specifically, variance measures how far each
// number in the set is from the mean and thus from every other number in the
// set. Variance is often depicted by this symbol: σ2. It is used by both
// analysts and traders to determine volatility and market security. The square
// root of the variance is the standard deviation (σ), which helps determine
// the consistency of an investment’s returns over a period of time.
//  https://www.investopedia.com/terms/v/variance.asp
type Var struct {
	n    int64
	hist *cBuf
	sum  float64
}

func NewVar(n int64) *Var {
	return &Var{
		n:    n,
		hist: newCBuf(n),
		sum:  0,
	}
}

func (r *Var) Update(v float64) float64 {
	old := r.hist.append(v)
	r.sum += v - old

	if r.hist.size() < r.n {
		return 0
	}

	mean := r.sum / float64(r.n)
	sum := float64(0)
	r.hist.iter(func(v float64) {
		diff := (v - mean)
		sum += diff * diff
	})

	return sum / float64(r.n)
}

// The term variance refers to a statistical measurement of the spread between
// numbers in a data set. More specifically, variance measures how far each
// number in the set is from the mean and thus from every other number in the
// set. Variance is often depicted by this symbol: σ2. It is used by both
// analysts and traders to determine volatility and market security. The square
// root of the variance is the standard deviation (σ), which helps determine
// the consistency of an investment’s returns over a period of time.
//  https://www.investopedia.com/terms/v/variance.asp
func VarArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	s := NewVar(n)
	for i, v := range in {
		out[i] = s.Update(v)
	}

	return out
}
