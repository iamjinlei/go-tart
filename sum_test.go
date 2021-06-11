package tart

import (
	"testing"
)

func TestSum(t *testing.T) {
	compare(t, "result = talib.SUM(testClose, 10)", SumArr(testClose, 10))
}
