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
func TestSubscribeToCellJettonsTransaction(t *testing.T) {
	startTime := time.Now()

	t.Logf("Test subscribe to cell jettons transactions via Server Sent Events")
	{
		err := SubscribeToCellJettonsTransaction(5*time.Minute) // 5*time.Second
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tNotified about finished cell transaction", successMarker)
		}
	}
	logExecTime(t, &startTime)
}

// subscribe_to_transaction.go
func TestSubscribeToBuyJettonsTransaction(t *testing.T) {
	startTime := time.Now()

	t.Logf("Test subscribe to buy jettons transactions via Server Sent Events")
	{
		err := SubscribeToBuyJettonsTransaction(5*time.Minute) // 5*time.Second
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tNotified about finished buy transaction", successMarker)
		}
	}
	logExecTime(t, &startTime)
}
