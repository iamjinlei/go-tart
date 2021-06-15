package tart

import (
	"testing"
)

func TestMfi(t *testing.T) {
	compare(t, "result = talib.MFI(testHigh, testLow, testClose, testVolume, 14)", MfiArr(testHigh, testLow, testClose, testVolume, 14))
}
