package tart

// The triangular moving average (TMA) is a technical indicator that is similar
// to other moving averages. The TMA shows the average (or mean) price of an
// asset over a specified number of data points—usually a number of price bars.
// However, the triangular moving average differs in that it is double
// smoothed—which also means averaged twice.
//  https://www.thebalance.com/triangular-moving-average-tma-description-and-uses-1031203
//  https://www.fidelity.com/viewpoints/active-investor/moving-averages
type Trima struct {
	n         int64
	factor    float64
	leftWidth int64
	hist      *CBuf
	leftSum   float64
	rightSum  float64
	sum       float64
}

func NewTrima(n int64) *Trima {
	half := n >> 1
	var factor float64
	if n%2 == 1 {
		factor = 1.0 / float64((half+1)*(half+1))
	} else {
		factor = 1.0 / float64(half*(half+1))
	}

	return &Trima{
		n:         n,
		factor:    factor,
		leftWidth: n - half,
		hist:      NewCBuf(n),
		leftSum:   0,
		rightSum:  0,
		sum:       0,
	}
}

func (t *Trima) Update(v float64) float64 {
	oldLeftSum := t.leftSum
	oldRightSum := t.rightSum
	old := t.hist.Append(v)
	mid := t.hist.NthOldest(t.leftWidth - 1)
	t.leftSum += mid - old
	t.rightSum += v - mid

	sz := t.hist.Size()
	if sz <= t.leftWidth {
		t.sum += float64(sz) * v
		return 0
	} else if sz <= t.n {
		t.sum += float64(t.n-sz+1) * v
		if sz < t.n {
			return 0
		}
	} else {
		if t.n%2 == 1 {
			t.sum += -oldLeftSum + oldRightSum + v
		} else {
			t.sum += -oldLeftSum + t.rightSum
		}
	}

	return t.sum * t.factor
}

func (t *Trima) InitPeriod() int64 {
	return t.n - 1
}

func (t *Trima) Valid() bool {
	return t.hist.Size() > t.InitPeriod()
}

// The triangular moving average (TMA) is a technical indicator that is similar
// to other moving averages. The TMA shows the average (or mean) price of an
// asset over a specified number of data points—usually a number of price bars.
// However, the triangular moving average differs in that it is double
// smoothed—which also means averaged twice.
//  https://www.thebalance.com/triangular-moving-average-tma-description-and-uses-1031203
//  https://www.fidelity.com/viewpoints/active-investor/moving-averages
func TrimaArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	t := NewTrima(n)
	for i, v := range in {
		out[i] = t.Update(v)
	}

	return out
}
