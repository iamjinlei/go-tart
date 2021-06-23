package tart

// A Weighted Moving Average puts more weight on recent data and less on past
// data. This is done by multiplying each bar’s price by a weighting factor.
// Because of its unique calculation, WMA will follow prices more closely
// than a corresponding Simple Moving Average.
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/wma
type Wma struct {
	n    int64
	d    float64
	hist *CBuf
	sum  float64
	wsum float64
}

func NewWma(n int64) *Wma {
	return &Wma{
		n:    n,
		d:    float64(n*(n+1)) / 2,
		hist: NewCBuf(n),
		sum:  0,
		wsum: 0,
	}
}

func (w *Wma) Update(v float64) float64 {
	if w.n == 1 {
		return v
	}

	old := w.hist.Append(v)
	w.sum += v - old

	sz := w.hist.Size()
	if sz < w.n {
		w.wsum += v * float64(sz)
		return 0
	} else {
		w.wsum += v * float64(w.n)
	}

	ret := w.wsum / w.d
	w.wsum -= w.sum

	return ret
}

func (w *Wma) InitPeriod() int64 {
	return w.n - 1
}

func (w *Wma) Valid() bool {
	return w.hist.Size() > w.InitPeriod()
}

// A Weighted Moving Average puts more weight on recent data and less on past
// data. This is done by multiplying each bar’s price by a weighting factor.
// Because of its unique calculation, WMA will follow prices more closely
// than a corresponding Simple Moving Average.
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/wma
func WmaArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	w := NewWma(n)
	for i, v := range in {
		out[i] = w.Update(v)
	}

	return out
}
