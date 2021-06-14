package tart

import (
	"testing"
)

func TestApo(t *testing.T) {
	compare(t, "result = talib.APO(testClose, 12, 26, talib.MA_Type.SMA)", ApoArr(SMA, testClose, 12, 26))
	compare(t, "result = talib.APO(testClose, 26, 12, talib.MA_Type.SMA)", ApoArr(SMA, testClose, 26, 12))
}
