package tart

import (
	"testing"
)

func TestBBands(t *testing.T) {
	upper, mid, lower := BBandsArr(SMA, testClose, 5, 2.0, 2.0)
	compare(t, "result, mid, lower = talib.BBANDS(testClose, 5, 2.0, 2.0)", upper)
	compare(t, "upper, result, lower = talib.BBANDS(testClose, 5, 2.0, 2.0)", mid)
	compare(t, "upper, mid, result = talib.BBANDS(testClose, 5, 2.0, 2.0)", lower)
}
