package tart

import (
	"testing"
)

func TestTrix(t *testing.T) {
	compare(t, "result = talib.TRIX(testClose, 5)", TrixArr(testClose, 5))
	compare(t, "result = talib.TRIX(testClose, 30)", TrixArr(testClose, 30))
}
