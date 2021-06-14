package tart

import (
	"testing"
)

func TestAroon(t *testing.T) {
	dn, up := AroonArr(testHigh, testLow, 14)
	compare(t, "result, up = talib.AROON(testHigh, testLow, 14)", dn)
	compare(t, "dn, result = talib.AROON(testHigh, testLow, 14)", up)
}
