package tart

type Trima struct {
	n          int64
	factor     float64
	leftWidth  int64
	rightWidth int64
	sz         int64
	hist       []float64
	leftSum    float64
	rightSum   float64
	sum        float64
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
		n:          n,
		factor:     factor,
		leftWidth:  n - half,
		rightWidth: half,
		sz:         0,
		hist:       make([]float64, n),
		leftSum:    0,
		rightSum:   0,
		sum:        0,
	}
}

func (t *Trima) Update(v float64) float64 {
	leftStart := t.sz % t.n
	leftEnd := (leftStart + t.leftWidth - 1) % t.n
	rightStart := (leftEnd + 1) % t.n
	oldLeftSum := t.leftSum
	oldRightSum := t.rightSum
	t.leftSum += t.hist[rightStart] - t.hist[leftStart]
	t.rightSum += v - t.hist[rightStart]
	t.hist[leftStart] = v
	t.sz++

	if t.sz <= t.leftWidth {
		t.sum += float64(t.sz) * v
		return 0
	} else if t.sz <= t.n {
		t.sum += float64(t.n-t.sz+1) * v
		if t.sz < t.n {
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

func TrimaArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	t := NewTrima(n)
	for i, v := range in {
		out[i] = t.Update(v)
	}

	return out
}
