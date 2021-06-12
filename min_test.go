package tart

import (
	"testing"
)

func TestMin(t *testing.T) {
	minIdx, min := MinArr(testClose, 10)
	compare(t, "result = talib.MIN(testClose, 10)", min)
	compare(t, "result = talib.MININDEX(testClose, 10)", minIdx)
}
