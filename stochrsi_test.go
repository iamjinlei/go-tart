package tart

import (
	"testing"
)

func TestStochRsi(t *testing.T) {
	k, d := StochRsiArr(testClose, 14, 5, SMA, 2)
	compare(t, "result, d = talib.STOCHRSI(testClose, 14, 5, 2, talib.MA_Type.SMA)", k)
	compare(t, "k, result = talib.STOCHRSI(testClose, 14, 5, 2, talib.MA_Type.SMA)", d)
}
