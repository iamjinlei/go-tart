package tart

import (
	"testing"
)

func TestObv(t *testing.T) {
	compare(t, "result = talib.OBV(testClose, testVolume)", ObvArr(testClose, testVolume))
}
