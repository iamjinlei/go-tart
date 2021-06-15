package tart

import (
	"testing"
)

func TestCmo(t *testing.T) {
	compare(t, "result = talib.CMO(testClose, 14)", CmoArr(testClose, 14))
}
