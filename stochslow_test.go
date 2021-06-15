package tart

import (
	"testing"
)

func TestStochSlow(t *testing.T) {
	k, d := StochSlowArr(testHigh, testLow, testClose, 5, SMA, 3, SMA, 3)
	compare(t, "result, d = talib.STOCH(testHigh, testLow, testClose, 5, 3, talib.MA_Type.SMA, 3, talib.MA_Type.SMA)", k)
	compare(t, "k, result = talib.STOCH(testHigh, testLow, testClose, 5, 3, talib.MA_Type.SMA, 3, talib.MA_Type.SMA)", d)

	k, d = StochSlowArr(testHigh, testLow, testClose, 12, SMA, 3, SMA, 3)
	compare(t, "result, d = talib.STOCH(testHigh, testLow, testClose, 12, 3, talib.MA_Type.SMA, 3, talib.MA_Type.SMA)", k)
	compare(t, "k, result = talib.STOCH(testHigh, testLow, testClose, 12, 3, talib.MA_Type.SMA, 3, talib.MA_Type.SMA)", d)

	k, d = StochSlowArr(testHigh, testLow, testClose, 12, SMA, 3, SMA, 15)
	compare(t, "result, d = talib.STOCH(testHigh, testLow, testClose, 12, 3, talib.MA_Type.SMA, 15, talib.MA_Type.SMA)", k)
	compare(t, "k, result = talib.STOCH(testHigh, testLow, testClose, 12, 3, talib.MA_Type.SMA, 15, talib.MA_Type.SMA)", d)
}
