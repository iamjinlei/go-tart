package tart

import (
	"testing"
)

func TestPpo(t *testing.T) {
	compare(t, "result = talib.PPO(testClose, 12, 26, talib.MA_Type.SMA)", PpoArr(testClose, SMA, 12, 26))
	compare(t, "result = talib.PPO(testClose, 26, 12, talib.MA_Type.SMA)", PpoArr(testClose, SMA, 26, 12))
}
