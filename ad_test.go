package tart

import (
	"testing"
)

func TestAd(t *testing.T) {
	// NOTE: examined ta-lib c code with the same implementation. But the calculation leads to different precision
	compareWithError(t, "result = talib.AD(testHigh, testLow, testClose, testVolume)", AdArr(testHigh, testLow, testClose, testVolume), 1e-2)
}
