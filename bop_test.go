package tart

import (
	"testing"
)

func TestBop(t *testing.T) {
	compare(t, "result = talib.BOP(testOpen, testHigh, testLow, testClose)", BopArr(testOpen, testHigh, testLow, testClose))
}
