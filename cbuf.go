package tart

// circular buffer maintaining fixed sized history
type cBuf struct {
	n      int64
	hist   []float64
	oldest int64
	newest int64
}

func newCBuf(n int64) *cBuf {
	return &cBuf{
		n:      n,
		hist:   make([]float64, n),
		oldest: 0,
		newest: n - 1,
	}
}

func (c *cBuf) append(v float64) {
	c.hist[c.oldest] = v
	c.newest = c.oldest
	c.oldest = (c.oldest + 1) % c.n
}

// nthNewest(0) = newest
// nthNewest(1) = 2nd newest
func (c *cBuf) nthNewest(offset int) float64 {
	return c.hist[(c.newest+c.n-int64(offset))%c.n]
}

// nthOldest(0) = oldest
// nthOldest(1) = 2nd oldest
func (c *cBuf) nthOldest(offset int) float64 {
	return c.hist[(c.oldest+int64(offset))%c.n]
}

func (c *cBuf) min() float64 {
	min := c.hist[0]
	for i := 1; i < len(c.hist); i++ {
		if c.hist[i] < min {
			min = c.hist[i]
		}
	}
	return min
}

func (c *cBuf) max() float64 {
	max := c.hist[0]
	for i := 1; i < len(c.hist); i++ {
		if c.hist[i] > max {
			max = c.hist[i]
		}
	}
	return max
}
