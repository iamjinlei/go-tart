package tart

import (
	"testing"
)

func TestCci(t *testing.T) {
	compare(t, "result = talib.CCI(testHigh, testLow, testClose, 14)", CciArr(testHigh, testLow, testClose, 14))
}
