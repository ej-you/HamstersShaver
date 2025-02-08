package jettons

import (
	"time"

	"testing"
	"github.com/stretchr/testify/assert"
)


// маркеры успеха и неудачи
const (
    successMarker = "\u2713"
    failedMarker  = "\u2717"
)


func logExecTime(t *testing.T, startTime *time.Time) {
	endTime := time.Now()
	t.Logf("\t\tExec time: %v", endTime.Sub(*startTime))
	*startTime = endTime
}


// jetton_info.go
func TestGetJettonInfoByAddressWithTimeout(t *testing.T) {
	startTime := time.Now()

	// валидный адрес монеты (в base64 формате)
	t.Logf("Test getting jetton info with VALID jetton masterAddress in base64")
	{
		// In
		var masterAddress string = "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
		var timeout time.Duration = 5*time.Second
		t.Logf("\t\tInput: masterAddress=%q | timeout=%v", masterAddress, timeout)

		// Out
		gotJettonParams, err := GetJettonInfoByAddressWithTimeout(masterAddress, timeout)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tOutput (got jetton info): %v", successMarker, gotJettonParams)
		}
	}
	logExecTime(t, &startTime)

	// валидный адрес монеты (в HEX формате)
	t.Logf("Test getting jetton info with VALID jetton masterAddress in HEX")
	{
		// In
		var masterAddress string = "0:afc49cb8786f21c87045b19ede78fc6b46c51048513f8e9a6d44060199c1bf0c"
		var timeout time.Duration = 5*time.Second
		t.Logf("\t\tInput: masterAddress=%q | timeout=%v", masterAddress, timeout)

		// Out
		gotJettonParams, err := GetJettonInfoByAddressWithTimeout(masterAddress, timeout)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tOutput (got jetton info): %v", successMarker, gotJettonParams)
		}
	}
	logExecTime(t, &startTime)

	// невалидный адрес монеты
	t.Logf("Test getting jetton info with INVALID jetton masterAddress")
	{
		// In
		var masterAddress string = "EQC47093oX5Xhb0xuk2lCr2RhS8rj"
		var timeout time.Duration = 5*time.Second
		t.Logf("\t\tInput: masterAddress=%q | timeout=%v", masterAddress, timeout)

		// Out
		gotJettonParams, err := GetJettonInfoByAddressWithTimeout(masterAddress, timeout)
		if assert.Error(t, err) {
			if assert.Equalf(t, err.Error(), "Jetton was not found", "\t%s\tFailed: %s", failedMarker, err.Error()) {
				t.Logf("\t%s\tOutput: %s", successMarker, err.Error())
			}
		} else {
			t.Errorf("\t%s\tFailed. Output (got jetton info): %v", failedMarker, gotJettonParams)
		}
	}
	logExecTime(t, &startTime)
}
