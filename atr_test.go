package tart

import (
	"testing"
)

func TestAtr(t *testing.T) {
	compare(t, "result = talib.ATR(testHigh, testLow, testClose, 14)", AtrArr(testHigh, testLow, testClose, 14))
}
