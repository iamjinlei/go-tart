package tart

import (
	"testing"
)

func TestAdxR(t *testing.T) {
	compare(t, "result = talib.ADXR(testHigh, testLow, testClose, 5)", AdxRArr(testHigh, testLow, testClose, 5))
}
