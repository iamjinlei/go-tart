package tart

import (
	"testing"
)

func TestNatr(t *testing.T) {
	compare(t, "result = talib.NATR(testHigh, testLow, testClose, 14)", NatrArr(testHigh, testLow, testClose, 14))
}
