package tart

import (
	"testing"
)

func TestUltOsc(t *testing.T) {
	compare(t, "result = talib.ULTOSC(testHigh, testLow, testClose, 7, 14, 28)", UltOscArr(testHigh, testLow, testClose, 7, 14, 28))
}
