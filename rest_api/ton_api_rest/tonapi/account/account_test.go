package account

import (
	"context"
	"os"
	"time"
	tonapi "github.com/tonkeeper/tonapi-go"

	myTongoWallet "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tongo/wallet"
	"github.com/ej-you/HamstersShaver/rest_api/settings"

	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


// обязательно указать эту переменную окружения перед командой запуска тестов (go test)
// CONFIG_PATH=../../../settings/config/


// маркеры успеха и неудачи
const (
    successMarker = "\u2713"
    failedMarker  = "\u2717"
)

var (
	tonapiClient *tonapi.Client
	tonApiContext context.Context
)


func logExecTime(t *testing.T, startTime *time.Time) {
	endTime := time.Now()
	t.Logf("\t\tExec time: %v", endTime.Sub(*startTime))
	*startTime = endTime
}


func TestMain(m *testing.M) {
	var err error
	var cancel context.CancelFunc

	// создание API клиента TON для tonapi-go с таймаутом в 3 секунды
	tonapiClient, err = settings.GetTonClientTonapiWithTimeout("mainnet", 3*time.Second)
	if err != nil {
		panic(err)
	}
	// создание контекста с таймаутом в 5 секунд
	tonApiContext, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exitCode := m.Run()
	os.Exit(exitCode)
}


// account_jetton.go
func TestGetAccountJetton(t *testing.T) {
	startTime := time.Now()

	// валидный адрес монеты (в base64 формате)
	t.Logf("Test getting account jetton info with VALID jetton masterAddress in base64")
	{
		// In
		var masterAddress string = "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
		t.Logf("\t\tInput: masterAddress=%q", masterAddress)

		// Out
		gotAccountJetton, err := GetAccountJetton(tonApiContext, tonapiClient, masterAddress)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tOutput (got account jetton info): %v", successMarker, gotAccountJetton)
		}
	}
	logExecTime(t, &startTime)

	// валидный адрес монеты (в HEX формате)
	t.Logf("Test getting jetton info with VALID jetton masterAddress in HEX")
	{
		// In
		var masterAddress string = "0:afc49cb8786f21c87045b19ede78fc6b46c51048513f8e9a6d44060199c1bf0c"
		t.Logf("\t\tInput: masterAddress=%q", masterAddress)

		// Out
		gotAccountJetton, err := GetAccountJetton(tonApiContext, tonapiClient, masterAddress)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tOutput (got account jetton info): %v", successMarker, gotAccountJetton)
		}
	}
	logExecTime(t, &startTime)

	// невалидный адрес монеты
	t.Logf("Test getting jetton info with INVALID jetton masterAddress")
	{
		// In
		var masterAddress string = "0:afc49cb8786f21c87045b19ede78fc6b46"
		t.Logf("\t\tInput: masterAddress=%q", masterAddress)

		// Out
		gotAccountJetton, err := GetAccountJetton(tonApiContext, tonapiClient, masterAddress)
		if assert.Error(t, err) {
			if assert.Containsf(t, err.Error(), "Failed to get account jetton info: decode response: error: code 4", "\t%s\tFailed: %s", failedMarker, err.Error()) {
				t.Logf("\t%s\tOutput: %s", successMarker, err.Error())
			}
		} else {
			t.Errorf("\t%s\tFailed. Output (got account jetton info): %v", failedMarker, gotAccountJetton)
		}
	}
	logExecTime(t, &startTime)

	// валидный адрес монеты, которой нет на этом аккаунте
	t.Logf("Test getting jetton info with VALID jetton masterAddress but account HAS NOT this jetton")
	{
		// In
		var masterAddress string = "EQAQXlWJvGbbFfE8F3oS8s87lIgdovS455IsWFaRdmJetTon"
		t.Logf("\t\tInput: masterAddress=%q", masterAddress)

		// Out
		gotAccountJetton, err := GetAccountJetton(tonApiContext, tonapiClient, masterAddress)
		if assert.Error(t, err) {
			if assert.Containsf(t, err.Error(), "Failed to get account jetton info: decode response: error: code 404: {Error:account", "\t%s\tFailed: %s", failedMarker, err.Error()) {
				t.Logf("\t%s\tOutput: %s", successMarker, err.Error())
			}
		} else {
			t.Errorf("\t%s\tFailed. Output (got account jetton info): %v", failedMarker, gotAccountJetton)
		}
	}
	logExecTime(t, &startTime)
}


// account_jettons.go
func TestGetBalanceJettons(t *testing.T) {
	startTime := time.Now()

	t.Logf("Test getting all account jettons info")
	{
		// In
		t.Logf("\t\tInput: -")

		// Out
		gotBalanceJettons, err := GetBalanceJettons(tonApiContext, tonapiClient)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tOutput (got all account jettons info): %v", successMarker, gotBalanceJettons)
		}
	}
	logExecTime(t, &startTime)
}


// account_seqno.go
func TestGetAccountSeqno(t *testing.T) {
	startTime := time.Now()

	// создание API клиента TON для tongo с таймаутом в 3 секунд
	tongoClient, err := settings.GetTonClientTongoWithTimeout("mainnet", 3*time.Second)
	require.NoErrorf(t, err, "\t%s\tFailed to get tongoClient: %v", failedMarker, err)

	// получение данных о кошельке через tongo
	realWallet, err := myTongoWallet.GetWallet(tongoClient)
	require.NoErrorf(t, err, "\t%s\tFailed to get wallet: %v", failedMarker, err)

	t.Logf("Test getting account seqno")
	{
		// In
		t.Logf("\t\tInput: -")

		// Out
		gotAccountSeqno, err := GetAccountSeqno(tonApiContext, tonapiClient, realWallet)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tOutput (got account seqno): %v", successMarker, gotAccountSeqno)
		}
	}
	logExecTime(t, &startTime)
}


// account_ton.go
func TestGetBalanceTON(t *testing.T) {
	var startTime, endTime time.Time

	startTime = time.Now()

	t.Logf("Test getting account TON info")
	{
		// In
		t.Logf("\t\tInput: -")

		// Out
		gotAccountTON, err := GetBalanceTON(tonApiContext, tonapiClient)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tOutput (got account TON info): %v", successMarker, gotAccountTON)
		}
	}

	endTime = time.Now()
	t.Logf("\t\tExec time: %v", endTime.Sub(startTime))
	startTime = endTime
}
