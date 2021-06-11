package tart

import (
	"testing"
)

func TestVar(t *testing.T) {
	compare(t, "result = talib.VAR(testClose, 10)", VarArr(testClose, 10))
}
