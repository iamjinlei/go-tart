package tart

// circular buffer maintaining fixed sized history
type CBuf struct {
	n      int64
	hist   []float64
	oldest int64
	newest int64
	sz     int64
}

func NewCBuf(n int64) *CBuf {
	return &CBuf{
		n:      n,
		hist:   make([]float64, n),
		oldest: 0,
		newest: n - 1,
	}
}

// Append latest value and return oldest one
func (c *CBuf) Append(v float64) float64 {
	old := c.hist[c.oldest]
	c.hist[c.oldest] = v
	c.newest = c.oldest
	c.oldest = (c.oldest + 1) % c.n
	c.sz++
	return old
}

// Number of values appended
func (c *CBuf) Size() int64 {
	return c.sz
}

// From circular buf position to total sequence index
func (c *CBuf) IndexToSeq(idx int64) int64 {
	if idx < c.oldest {
		return c.sz - (c.oldest - idx)
	} else {
		return c.sz - c.n + (idx - c.oldest)
	}
}

// Index of the latest value
func (c *CBuf) NewestIndex() int64 {
	return c.newest
}

// Index of the oldest value
func (c *CBuf) OldestIndex() int64 {
	return c.oldest
}

// nthNewest(0) = newest
// nthNewest(1) = 2nd newest
func (c *CBuf) NthNewest(offset int64) float64 {
	return c.hist[(c.newest+c.n-offset)%c.n]
}

// nthOldest(0) = oldest
// nthOldest(1) = 2nd oldest
func (c *CBuf) NthOldest(offset int64) float64 {
	return c.hist[(c.oldest+offset)%c.n]
}

// Min value in buf
func (c *CBuf) Min() (int64, float64) {
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

// Max value in buf
func (c *CBuf) Max() (int64, float64) {
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

// Iterate through buf elements and call function for each
func (c *CBuf) Iter(fn func(v float64)) {
	idx := c.oldest
	for i := int64(0); i < c.n; i++ {
		fn(c.hist[idx])
		idx = (idx + 1) % c.n
	}
}
