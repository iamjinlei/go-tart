package tart

import (
	"testing"
)

func TestMax(t *testing.T) {
	maxIdx, max := MaxArr(testClose, 10)
	compare(t, "result = talib.MAX(testClose, 10)", max)
	compare(t, "result = talib.MAXINDEX(testClose, 10)", maxIdx)
}
