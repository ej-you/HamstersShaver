package transactions

import (
	"context"
	"os"
	"time"

	"testing"
	"github.com/stretchr/testify/assert"
)


// маркеры успеха и неудачи
const (
    successMarker = "\u2713"
    failedMarker  = "\u2717"
)

var tonApiContext context.Context


func logExecTime(t *testing.T, startTime *time.Time) {
	endTime := time.Now()
	t.Logf("\t\tExec time: %v", endTime.Sub(*startTime))
	*startTime = endTime
}

func TestMain(m *testing.M) {
	var cancel context.CancelFunc

	// создание контекста с таймаутом в 5 секунд
	tonApiContext, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	exitCode := m.Run()
	os.Exit(exitCode)
}

// info.go
func TestGetTransactionInfoWithStatusOK(t *testing.T) {
	startTime := time.Now()

	t.Logf("Test getting info about SUCCESS cell jettons transaction by its hash")
	{
		// In
		// hash := "a8ec992c341230a885f9adfe6598eb307660c306a5f00cf0d302c72e7d966389"
		hash := "4f8ff3378e1d4cc80488750fda3bcc6b730b71b69429d9c44a775b377bdc66a4"
		action := "cell"

		// Out
		transInfo, err := GetTransactionInfoWithStatusOK(tonApiContext, hash, action)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tGot cell transaction info: %v", successMarker, transInfo)
		}
		assert.Equal(t, transInfo.StatusOK, true, "The StatusOK should be true")
	}
	logExecTime(t, &startTime)

	t.Logf("Test getting info about FAILED (slippage) cell jettons transaction by its hash")
	{
		// In
		// hash := "32fec73c3305a18420cb6568ce56ce960b04b23ccbe07da0cdfb33cac0506c54"
		hash := "396ae1f2fd595598d9bffb09506d5fd80b12fc5704ce9e6ab0d9782b68da24f7"
		action := "cell"

		// Out
		transInfo, err := GetTransactionInfoWithStatusOK(tonApiContext, hash, action)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tGot cell transaction info: %v", successMarker, transInfo)
		}
		assert.Equal(t, transInfo.StatusOK, false, "The StatusOK should be false")
	}
	logExecTime(t, &startTime)

	t.Logf("Test getting info about SUCCESS buy jettons transaction by its hash")
	{
		// In
		// hash := "f77c04ca40caf0b606ee1dc4dbd35e578faffdb262b24df99606ee77a35de077"
		hash := "a556193068a7777adc0c6b0ea0feae9878add5eec87dd955d39d83444838cb8e"
		action := "buy"

		// Out
		transInfo, err := GetTransactionInfoWithStatusOK(tonApiContext, hash, action)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tGot buy transaction info: %v", successMarker, transInfo)
		}
		assert.Equal(t, transInfo.StatusOK, true, "The StatusOK should be true")
	}
	logExecTime(t, &startTime)

	t.Logf("Test getting info about FAILED (slippage) buy jettons transaction by its hash")
	{
		// In
		// hash := "eca5765b54e5a1bc0bf1f4010732ac085edf179c7da19a2b60223cd15a19e862"
		hash := "caf8d9e8d888395bd1c530a301be6918c2f93fb1aa8ee69eab7fa485a17c624f"
		action := "buy"

		// Out
		transInfo, err := GetTransactionInfoWithStatusOK(tonApiContext, hash, action)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tGot buy transaction info: %v", successMarker, transInfo)
		}
		assert.Equal(t, transInfo.StatusOK, false, "The StatusOK should be false")
	}
	logExecTime(t, &startTime)
}
