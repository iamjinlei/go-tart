package tart

import (
	"testing"
)

func TestWillR(t *testing.T) {
	compare(t, "result = talib.WILLR(testHigh, testLow, testClose, 9)", WillRArr(testHigh, testLow, testClose, 9))
}
