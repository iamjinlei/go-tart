package tart

import (
	"testing"
)

func TestKama(t *testing.T) {
	compare(t, "result = talib.KAMA(testClose, 10)", KamaArr(testClose, 10))
}
