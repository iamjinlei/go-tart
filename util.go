package tart

func almostZero(v float64) bool {
	return v > -0.00000000000001 && v < 0.00000000000001
}

// math.Max() does some extra we don'tc care (overhead)
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
