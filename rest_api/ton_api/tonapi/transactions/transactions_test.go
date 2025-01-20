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

// info_by_hash.go
func TestGetTransactionInfoByHash(t *testing.T) {
	startTime := time.Now()

	t.Logf("Test getting info about success cell jettons transaction by its hash")
	{
		// In
		// hash := "79c2a5559c671e1ea56f6e345eeb88ef9f689d65a27c709be457f4bc4fa1e7a7"
		hash := "a8ec992c341230a885f9adfe6598eb307660c306a5f00cf0d302c72e7d966389"
		action := "cell"

		// Out
		transInfo, err := GetTransactionInfoWithStatusOKByHash(tonApiContext, hash, action)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tGot cell transaction info: %v", successMarker, transInfo)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test getting info about failed (slippage) cell jettons transaction by its hash")
	{
		// In
		hash := "32fec73c3305a18420cb6568ce56ce960b04b23ccbe07da0cdfb33cac0506c54"
		action := "cell"

		// Out
		transInfo, err := GetTransactionInfoWithStatusOKByHash(tonApiContext, hash, action)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tGot cell transaction info: %v", successMarker, transInfo)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test getting info about success buy jettons transaction by its hash")
	{
		// In
		hash := "f77c04ca40caf0b606ee1dc4dbd35e578faffdb262b24df99606ee77a35de077"
		action := "buy"

		// Out
		transInfo, err := GetTransactionInfoWithStatusOKByHash(tonApiContext, hash, action)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tGot buy transaction info: %v", successMarker, transInfo)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test getting info about failed (slippage) buy jettons transaction by its hash")
	{
		// In
		hash := "eca5765b54e5a1bc0bf1f4010732ac085edf179c7da19a2b60223cd15a19e862"
		action := "buy"

		// Out
		transInfo, err := GetTransactionInfoWithStatusOKByHash(tonApiContext, hash, action)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tGot buy transaction info: %v", successMarker, transInfo)
		}
	}
	logExecTime(t, &startTime)
}
