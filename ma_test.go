package tart

import (
	"testing"
)

func TestMa(t *testing.T) {
	compare(t, "result = talib.MA(testClose, 20, talib.MA_Type.DEMA)", MaArr(DEMA, testClose, 20))
}
