package tart

// The Money Flow Index (MFI) is an oscillator that uses
// both price and volume to measure buying and selling
// pressure. Created by Gene Quong and Avrum Soudack,
// MFI is also known as volume-weighted RSI. MFI starts
// with the typical price for each period. Money flow is
// positive when the typical price rises (buying pressure)
// and negative when the typical price declines (selling pressure).
// A ratio of positive and negative money flow is then plugged
// into an RSI formula to create an oscillator that moves
// between zero and one hundred. As a momentum oscillator
// tied to volume, MFI is best suited to identify reversals
// and price extremes with a variety of signals.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:money_flow_index_mfi
//  https://www.investopedia.com/terms/m/mfi.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/MFI
type Mfi struct {
	n        int64
	positive *Sum
	negative *Sum
	prevTp   float64
	sz       int64
}

func NewMfi(n int64) *Mfi {
	return &Mfi{
		n:        n,
		positive: NewSum(n),
		negative: NewSum(n),
		prevTp:   0,
		sz:       0,
	}
}

func (m *Mfi) Update(h, l, c, v float64) float64 {
	m.sz++

	tp := (h + l + c) / 3.0
	prevTp := m.prevTp
	m.prevTp = tp

	if m.sz == 1 {
		return 0
	}

	var pSum, nSum float64
	if tp > prevTp {
		pSum = m.positive.Update(tp * v)
		nSum = m.negative.Update(0)
	} else {
		pSum = m.positive.Update(0)
		nSum = m.negative.Update(tp * v)
	}

	if m.sz <= m.n {
		return 0
	}

	sum := pSum + nSum
	return pSum / sum * 100.0
}

func (m *Mfi) InitPeriod() int64 {
	return m.n
}

func (m *Mfi) Valid() bool {
	return m.sz > m.InitPeriod()
}

// The Money Flow Index (MFI) is an oscillator that uses
// both price and volume to measure buying and selling
// pressure. Created by Gene Quong and Avrum Soudack,
// MFI is also known as volume-weighted RSI. MFI starts
// with the typical price for each period. Money flow is
// positive when the typical price rises (buying pressure)
// and negative when the typical price declines (selling pressure).
// A ratio of positive and negative money flow is then plugged
// into an RSI formula to create an oscillator that moves
// between zero and one hundred. As a momentum oscillator
// tied to volume, MFI is best suited to identify reversals
// and price extremes with a variety of signals.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:money_flow_index_mfi
//  https://www.investopedia.com/terms/m/mfi.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/MFI
func MfiArr(h, l, c, v []float64, n int64) []float64 {
	out := make([]float64, len(c))

	m := NewMfi(n)
	for i := 0; i < len(c); i++ {
		out[i] = m.Update(h[i], l[i], c[i], v[i])
	}

	return out
}
