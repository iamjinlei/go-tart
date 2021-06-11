package tart

import (
	"testing"
)

func TestDema(t *testing.T) {
	compare(t, "result = talib.DEMA(testClose, 10)", DemaArr(testClose, 10))
}
