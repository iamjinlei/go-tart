package tart

// circular buffer maintaining fixed sized history
type cBuf struct {
	n      int64
	hist   []float64
	oldest int64
	newest int64
	sz     int64
}

func newCBuf(n int64) *cBuf {
	return &cBuf{
		n:      n,
		hist:   make([]float64, n),
		oldest: 0,
		newest: n - 1,
	}
}

func (c *cBuf) append(v float64) float64 {
	old := c.hist[c.oldest]
	c.hist[c.oldest] = v
	c.newest = c.oldest
	c.oldest = (c.oldest + 1) % c.n
	c.sz++
	return old
}

func (c *cBuf) size() int64 {
	return c.sz
}

func (c *cBuf) indexToSeq(idx int64) int64 {
	if idx < c.oldest {
		return c.sz - (c.oldest - idx)
	} else {
		return c.sz - c.n + (idx - c.oldest)
	}
}

func (c *cBuf) newestIndex() int64 {
	return c.newest
}

func (c *cBuf) oldestIndex() int64 {
	return c.oldest
}

// nthNewest(0) = newest
// nthNewest(1) = 2nd newest
func (c *cBuf) nthNewest(offset int64) float64 {
	return c.hist[(c.newest+c.n-offset)%c.n]
}

// nthOldest(0) = oldest
// nthOldest(1) = 2nd oldest
func (c *cBuf) nthOldest(offset int64) float64 {
	return c.hist[(c.oldest+offset)%c.n]
}

func (c *cBuf) min() (int64, float64) {
	min := c.hist[0]
	minIdx := int64(0)
	for i := 1; i < len(c.hist); i++ {
		if c.hist[i] < min {
			min = c.hist[i]
			minIdx = int64(i)
		}
	}
	return minIdx, min
}

func (c *cBuf) max() (int64, float64) {
	max := c.hist[0]
	maxIdx := int64(0)
	for i := 1; i < len(c.hist); i++ {
		if c.hist[i] > max {
			max = c.hist[i]
			maxIdx = int64(i)
		}
	}
	return maxIdx, max
}

func (c *cBuf) iter(fn func(v float64)) {
	idx := c.oldest
	for i := int64(0); i < c.n; i++ {
		fn(c.hist[idx])
		idx = (idx + 1) % c.n
	}
}
