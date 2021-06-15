package tart

import (
	"testing"
)

func TestMacdExt(t *testing.T) {
	macd, signal, hist := MacdExtArr(testClose, SMA, 12, SMA, 26, SMA, 9)
	compare(t, "result, signal, hist = talib.MACDEXT(testClose, 12, talib.MA_Type.SMA, 26, talib.MA_Type.SMA, 9, talib.MA_Type.SMA)", macd)
	compare(t, "macd, result, hist = talib.MACDEXT(testClose, 12, talib.MA_Type.SMA, 26, talib.MA_Type.SMA, 9, talib.MA_Type.SMA)", signal)
	compare(t, "macd, signal, result = talib.MACDEXT(testClose, 12, talib.MA_Type.SMA, 26, talib.MA_Type.SMA, 9, talib.MA_Type.SMA)", hist)
}
