package tart

import (
	"testing"
)

func TestEma(t *testing.T) {
	compare(t, "result = talib.EMA(testClose, 5)", EmaArr(testClose, 5))
	compare(t, "result = talib.EMA(testClose, 20)", EmaArr(testClose, 20))
	compare(t, "result = talib.EMA(testClose, 50)", EmaArr(testClose, 50))
	compare(t, "result = talib.EMA(testClose, 100)", EmaArr(testClose, 100))
}
