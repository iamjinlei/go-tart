package tart

import (
	"testing"
)

func TestRoc(t *testing.T) {
	compare(t, "result = talib.ROC(testClose, 10)", RocArr(testClose, 10))
}
