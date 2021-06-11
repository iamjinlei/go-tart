package tart

import (
	"testing"
)

func TestTema(t *testing.T) {
	compare(t, "result = talib.TEMA(testClose, 10)", TemaArr(testClose, 10))
}
