package tart

import (
	"testing"
)

func TestRsi(t *testing.T) {
	compare(t, "result = talib.RSI(testClose, 10)", RsiArr(testClose, 10))
}
