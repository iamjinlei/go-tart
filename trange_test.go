package tart

import (
	"testing"
)

func TestTRange(t *testing.T) {
	compare(t, "result = talib.TRANGE(testHigh, testLow, testClose)", TRangeArr(testHigh, testLow, testClose))
}
