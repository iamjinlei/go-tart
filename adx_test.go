package tart

import (
	"testing"
)

func TestAdx(t *testing.T) {
	compare(t, "result = talib.ADX(testHigh, testLow, testClose, 14)", AdxArr(testHigh, testLow, testClose, 14))
}
