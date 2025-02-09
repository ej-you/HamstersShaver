package transactions

import (
	"context"
	"os"
	"time"

	"testing"
	"github.com/stretchr/testify/assert"
)


// опционально указать эту переменную окружения для запуска тестов на отправку транзакций покупки и продажи токенов
// TEST_SEND_TRANS=""


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


// buy_jettons.go
func TestGetPreRequestBuyJetton(t *testing.T) {
	startTime := time.Now()

	// валидные данные
	t.Logf("Test getting buy pre-request info with VALID given parameters")
	{
		// In
		var masterAddress string = "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
		var tonAmount float64 = 0.1
		var slippage int = 20
		var timeout time.Duration = 5*time.Second
		t.Logf("\t\tInput: masterAddress=%q | tonAmount=%f | slippage=%d | timeout=%v", masterAddress, tonAmount, slippage, timeout)

		// Out
		buyPreRequestInfo, err := GetPreRequestBuyJetton(masterAddress, tonAmount, slippage, timeout)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tOutput (got buy pre-request info): %v", successMarker, buyPreRequestInfo)
		}
	}
	logExecTime(t, &startTime)
}


// cell_jettons.go
func TestGetPreRequestCellJetton(t *testing.T) {
	startTime := time.Now()

	// валидные данные
	t.Logf("Test getting buy pre-request info with VALID given parameters")
	{
		// In
		var masterAddress string = "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG_DOGS"
		var jettonAmount float64 = 2000
		var slippage int = 20
		var timeout time.Duration = 5*time.Second
		t.Logf("\t\tInput: masterAddress=%q | jettonAmount=%f | slippage=%d | timeout=%v", masterAddress, jettonAmount, slippage, timeout)

		// Out
		cellPreRequestInfo, err := GetPreRequestCellJetton(masterAddress, jettonAmount, slippage, timeout)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tOutput (got cell pre-request info): %v", successMarker, cellPreRequestInfo)
		}
	}
	logExecTime(t, &startTime)
}


// buy_jettons.go (костыль, запустится с переменной окружения TEST_SEND_TRANS="")
func TestBuyJetton(t *testing.T) {
	_, sendTransEnvExists := os.LookupEnv("TEST_SEND_TRANS")
	if !sendTransEnvExists {
	    t.Skip("Skipping test in not verbose mode (use env TEST_SEND_TRANS=\"\" for sending transactions)")
    }

	// создание контекста с таймаутом в 5 секунд
	buyJettonContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	startTime := time.Now()
	// валидные данные
	t.Logf("Send buy jetton transaction with VALID given parameters")
	{
		// In
		var masterAddress string = "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
		var tonAmount float64 = 0.1
		var slippage int = 20
		var timeout time.Duration = 10*time.Second
		t.Logf("\t\tInput: masterAddress=%q | tonAmount=%f | slippage=%d | timeout=%v", masterAddress, tonAmount, slippage, timeout)

		// Out
		err := BuyJetton(buyJettonContext, timeout, masterAddress, tonAmount, slippage)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tSuccessfully!", successMarker)
		}
	}
	logExecTime(t, &startTime)
}


// cell_jettons.go (костыль, запустится с переменной окружения TEST_SEND_TRANS="")
func TestCellJetton(t *testing.T) {
	_, sendTransEnvExists := os.LookupEnv("TEST_SEND_TRANS")
	if !sendTransEnvExists {
	    t.Skip("Skipping test in not verbose mode (use env TEST_SEND_TRANS=\"\" for sending transactions)")
    }

	// создание контекста с таймаутом в 5 секунд
	cellJettonContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	startTime := time.Now()
	// валидные данные
	t.Logf("Send cell jetton transaction with VALID given parameters")
	{
		// In
		var masterAddress string = "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG_DOGS"
		var jettonAmount float64 = 2000
		var slippage int = 20
		var timeout time.Duration = 10*time.Second
		t.Logf("\t\tInput: masterAddress=%q | jettonAmount=%f | slippage=%d | timeout=%v", masterAddress, jettonAmount, slippage, timeout)

		// Out
		err := CellJetton(cellJettonContext, timeout, masterAddress, jettonAmount, slippage)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tSuccessfully!", successMarker)
		}
	}
	logExecTime(t, &startTime)
}
