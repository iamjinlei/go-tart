package tart

import (
	"testing"
)

func TestMa(t *testing.T) {
	compare(t, "result = talib.MA(testClose, 10, talib.MA_Type.DEMA)", MaArr(DEMA, testClose, 20))
}
