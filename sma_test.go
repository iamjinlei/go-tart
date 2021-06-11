package tart

import (
	"testing"
)

func TestSma(t *testing.T) {
	compare(t, "result = talib.SMA(testClose, 20)", SmaArr(testClose, 20))
}
