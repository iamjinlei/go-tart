package tart

import (
	"testing"
)

func TestWma(t *testing.T) {
	compare(t, "result = talib.WMA(testClose, 10)", WmaArr(testClose, 10))
}
