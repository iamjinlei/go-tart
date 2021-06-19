package tart

// Developed by J. Welles Wilder, the Average True Range (ATR)
// is an indicator that measures volatility. As with most of
// his indicators, Wilder designed ATR with commodities and
// daily prices in mind. Commodities are frequently more volatile
// than stocks. They were are often subject to gaps and limit moves,
// which occur when a commodity opens up or down its maximum
// allowed move for the session. A volatility formula based only
// on the high-low range would fail to capture volatility from
// gap or limit moves. Wilder created Average True Range to
// capture this “missing” volatility. It is important to remember
// that ATR does not provide an indication of price direction,
// just volatility.
//
// Wilder features ATR in his 1978 book, New Concepts in Technical
// Trading Systems. This book also includes the Parabolic SAR,
// RSI and the Directional Movement Concept (ADX). Despite being
// developed before the computer age, Wilder's indicators have
// stood the test of time and remain extremely popular.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:average_true_range_atr
//  https://www.investopedia.com/terms/a/atr.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/atr
type Atr struct {
	n   int64
	tr  *TRange
	ema *Ema
	sz  int64
}

func NewAtr(n int64) *Atr {
	return &Atr{
		n:   n,
		tr:  NewTRange(),
		ema: NewEma(n, 1.0/float64(n)),
		sz:  0,
	}
}

func (a *Atr) Update(h, l, c float64) float64 {
	a.sz++

	tr := a.tr.Update(h, l, c)

	if a.sz == 1 {
		return 0
	}

	return a.ema.Update(tr)
}

func (a *Atr) InitPeriod() int64 {
	return 1
}

func (a *Atr) Valid() bool {
	return a.sz > 1
}

// Developed by J. Welles Wilder, the Average True Range (ATR)
// is an indicator that measures volatility. As with most of
// his indicators, Wilder designed ATR with commodities and
// daily prices in mind. Commodities are frequently more volatile
// than stocks. They were are often subject to gaps and limit moves,
// which occur when a commodity opens up or down its maximum
// allowed move for the session. A volatility formula based only
// on the high-low range would fail to capture volatility from
// gap or limit moves. Wilder created Average True Range to
// capture this “missing” volatility. It is important to remember
// that ATR does not provide an indication of price direction,
// just volatility.
//
// Wilder features ATR in his 1978 book, New Concepts in Technical
// Trading Systems. This book also includes the Parabolic SAR,
// RSI and the Directional Movement Concept (ADX). Despite being
// developed before the computer age, Wilder's indicators have
// stood the test of time and remain extremely popular.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:average_true_range_atr
//  https://www.investopedia.com/terms/a/atr.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/atr
func AtrArr(h, l, c []float64, n int64) []float64 {
	out := make([]float64, len(c))

	a := NewAtr(n)
	for i := 0; i < len(c); i++ {
		out[i] = a.Update(h[i], l[i], c[i])
	}

	return out
}
