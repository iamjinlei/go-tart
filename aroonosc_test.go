package tart

import (
	"testing"
)

func TestAroonOsc(t *testing.T) {
	compare(t, "result = talib.AROONOSC(testHigh, testLow, 14)", AroonOscArr(testHigh, testLow, 14))
}
