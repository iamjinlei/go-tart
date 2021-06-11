package tart

import (
	"testing"
)

func TestTrima(t *testing.T) {
	compare(t, "result = talib.TRIMA(testClose, 10)", TrimaArr(testClose, 10))
	compare(t, "result = talib.TRIMA(testClose, 11)", TrimaArr(testClose, 11))
}
