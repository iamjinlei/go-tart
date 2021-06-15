package tart

import (
	"testing"
)

func TestMacd(t *testing.T) {
	macd, signal, hist := MacdArr(testClose, 12, 26, 9)
	compare(t, "result, signal, hist = talib.MACD(testClose, 12, 26, 9)", macd)
	compare(t, "macd, result, hist = talib.MACD(testClose, 12, 26, 9)", signal)
	compare(t, "macd, signal, result = talib.MACD(testClose, 12, 26, 9)", hist)
}
