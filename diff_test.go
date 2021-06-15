package tart

import (
	"testing"
)

func TestDiff(t *testing.T) {
	compare(t, "result = talib.MOM(testClose, 10)", DiffArr(testClose, 10))
}
