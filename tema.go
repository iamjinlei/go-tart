package tart

// The Triple Exponential Moving Average (TEMA) reduces the lag of traditional
// EMAs, making it more responsive and better-suited for short-term trading.
// Shortly after developing the Double Exponential Moving Average (DEMA) in 1994,
// Patrick Mulloy took the concept a step further and created the Triple
// Exponential Moving Average (TEMA). Like its predecessor DEMA, the TEMA overlay
// uses the lag difference between different EMAs to adjust a traditional EMA.
// However, TEMA's formula uses a triple-smoothed EMA in addition to the single-
// and double-smoothed EMAs employed in the formula for DEMA. The offset created
// using these three EMAs produces a moving average that stays even closer to the
// price bars than DEMA.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:tema
//  https://www.investopedia.com/terms/t/triple-exponential-moving-average.asp
type Tema struct {
	n    int64
	sz   int64
	ema1 *Ema
	ema2 *Ema
	ema3 *Ema
}

func NewTema(n int64, k float64) *Tema {
	return &Tema{
		n:    n,
		sz:   0,
		ema1: NewEma(n, k),
		ema2: NewEma(n, k),
		ema3: NewEma(n, k),
	}
}

func (t *Tema) Update(v float64) float64 {
	t.sz++

	e1 := t.ema1.Update(v)

	if t.sz > t.n-1 {
		e2 := t.ema2.Update(e1)

		if t.sz > 2*t.n-2 {
			e3 := t.ema3.Update(e2)

			if t.sz > t.n*3-3 {
				return e3 + 3.0*(e1-e2)
			}
		}
	}

	return 0
}

func (t *Tema) InitPeriod() int64 {
	return t.n*3 - 3
}

func (t *Tema) Valid() bool {
	return t.sz > t.InitPeriod()
}

// The Triple Exponential Moving Average (TEMA) reduces the lag of traditional
// EMAs, making it more responsive and better-suited for short-term trading.
// Shortly after developing the Double Exponential Moving Average (DEMA) in 1994,
// Patrick Mulloy took the concept a step further and created the Triple
// Exponential Moving Average (TEMA). Like its predecessor DEMA, the TEMA overlay
// uses the lag difference between different EMAs to adjust a traditional EMA.
// However, TEMA's formula uses a triple-smoothed EMA in addition to the single-
// and double-smoothed EMAs employed in the formula for DEMA. The offset created
// using these three EMAs produces a moving average that stays even closer to the
// price bars than DEMA.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:tema
//  https://www.investopedia.com/terms/t/triple-exponential-moving-average.asp
func TemaArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	k := 2.0 / float64(n+1)
	t := NewTema(n, k)
	for i, v := range in {
		out[i] = t.Update(v)
	}

	return out
}
