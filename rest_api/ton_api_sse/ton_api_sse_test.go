package ton_api_sse

import (
	"time"

	"testing"
	"github.com/stretchr/testify/assert"
)


// обязательно указать эту переменную окружения перед командой запуска тестов (go test)
// CONFIG_PATH=../settings/config/


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

// subscribe_to_transaction.go
func TestSubscribeToTransaction(t *testing.T) {
	startTime := time.Now()

	t.Logf("Test subscribe to transaction via Server Sent Events")
	{
		transHash, err := SubscribeToTransaction(5*time.Minute)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tGot finished transaction hash: %s", successMarker, transHash)
		}
	}
	logExecTime(t, &startTime)
}
