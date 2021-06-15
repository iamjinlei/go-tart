package tart

import (
	"testing"
)

func TestDx(t *testing.T) {
	compare(t, "result = talib.DX(testHigh, testLow, testClose, 14)", DxArr(testHigh, testLow, testClose, 14))
}
