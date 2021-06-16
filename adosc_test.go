package tart

import (
	"testing"
)

func TestAdOsc(t *testing.T) {
	// NOTE: the error difference is introduced by Ad
	compareWithError(t, "result = talib.ADOSC(testHigh, testLow, testClose, testVolume, 3, 10)", AdOscArr(testHigh, testLow, testClose, testVolume, 3, 10), 1e-3)
}
