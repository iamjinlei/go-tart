package tart

type MaType int

// TODO: support Mama & T3
const (
	SMA MaType = iota
	EMA
	WMA
	DEMA
	TEMA
	TRIMA
	KAMA
)

// Convenient wrapper for different moving average types
type Ma struct {
	mu maUpdater
}

func NewMa(t MaType, n int64) *Ma {
	k := 2.0 / float64(n+1)
	var mu maUpdater
	switch t {
	case SMA:
		mu = NewSma(n)
	case EMA:
		mu = NewEma(n, k)
	case WMA:
		mu = NewWma(n)
	case DEMA:
		mu = NewDema(n, k)
	case TEMA:
		mu = NewTema(n, k)
	case TRIMA:
		mu = NewTrima(n)
	case KAMA:
		mu = NewKama(n)
	default:
		return nil
	}

	return &Ma{
		mu: mu,
	}
}

func (m *Ma) Update(v float64) float64 {
	return m.mu.Update(v)
}

func (m *Ma) InitPeriod() int64 {
	return m.mu.InitPeriod()
}

func (m *Ma) Valid() bool {
	return m.mu.Valid()
}

// Convenient wrapper for different moving average types
func MaArr(t MaType, in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	m := NewMa(t, n)
	for i, v := range in {
		out[i] = m.Update(v)
	}

	return out
}

type maUpdater interface {
	Update(v float64) float64
	InitPeriod() int64
	Valid() bool
}
