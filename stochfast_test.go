package tart

import (
	"testing"
)

func TestStochFast(t *testing.T) {
	k, d := StochFastArr(testHigh, testLow, testClose, 5, SMA, 3)
	compare(t, "result, d = talib.STOCHF(testHigh, testLow, testClose, 5, 3, talib.MA_Type.SMA)", k)
	compare(t, "k, result = talib.STOCHF(testHigh, testLow, testClose, 5, 3, talib.MA_Type.SMA)", d)
}
