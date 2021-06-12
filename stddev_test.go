package tart

import (
	"testing"
)

func TestStdDev(t *testing.T) {
	compare(t, "result = talib.STDDEV(testClose, 10, 1.0)", StdDevArr(testClose, 10))
}
