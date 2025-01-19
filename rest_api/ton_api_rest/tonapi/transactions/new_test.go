package transactions

import (
	"context"
	"fmt"
	"testing"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// hash транзакции
const hash = "99821a8101a7d25e76811e01e97d606d5caa5f62419867245b8fd1f5f362590b"


func TestOld1(t *testing.T) {
	// создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), constants.GetTransInfoContextTimeout*2)
	defer cancel()

	// получение клиента для tonapi-go
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", constants.TonapiClientTimeout*2)
	if err != nil {
		panic(err)
	}

	// получение информации о транзе
	params1 := tonapi.GetBlockchainTransactionParams{TransactionID: hash}
	rawTransInfo, err := tonapiClient.GetBlockchainTransaction(ctx, params1)
	if err != nil {
		panic(err)
	}

	t.Log("rawTransInfo.Hash:", rawTransInfo.Hash)
	t.Log("rawTransInfo.Lt:", rawTransInfo.Lt)
	t.Log()

	// получение блока, содержащего данную транзакцию
	params2 := tonapi.GetBlockchainBlockParams{BlockID: rawTransInfo.Block}
	block, err := tonapiClient.GetBlockchainBlock(ctx, params2)

	t.Log("block.WorkchainID", block.WorkchainID)
	t.Log("block.Shard", block.Shard)
	t.Log("block.Seqno", block.Seqno)
	t.Log("block.RootHash", block.RootHash)
	t.Log("block.FileHash", block.FileHash)

	fmt.Printf(`
TRANS_LT = %d
WORKCHAIN = %d
SHARD_STR = %q
SEQNO = %d
ROOT_HASH = %q
FILE_HASH = %q
	%s`, rawTransInfo.Lt, block.WorkchainID, block.Shard, block.Seqno, block.RootHash, block.FileHash, "\n")
}


func unwrapTrace(trace *tonapi.Trace) {
	fmt.Println("\ntrace.Transaction:", trace.Transaction)

	if trace.Children != nil {
		unwrapTrace(&trace.Children[0])
	}
}
/*
CONFIG_PATH=../../../settings/config/ go test -v -run TestNew
*/
func TestNew(t *testing.T) {
	// получение клиента для tonapi-go
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", constants.TonapiClientTimeout)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	streamingAPI := tonapi.NewStreamingAPI(tonapi.WithStreamingToken(settings.TonApiToken))

	errChan := make(chan error)
	defer close(errChan)
	// resultChan := make(chan string)
	// defer close(resultChan)

	go func() {
		err := streamingAPI.SubscribeToTraces(ctx,
			[]string{settings.GetJsonWallet().Hash, "UQA4mfrV45OEIuTyJKDQe41FX1X0XD8IPJ9UYb7Tpu3gK6kO"},
			func(data tonapi.TraceEventData) {
				// cancel()
				t.Log("data:", data)

				trace, err := tonapiClient.GetTrace(ctx, tonapi.GetTraceParams{TraceID: data.Hash})
				if err != nil {
					errChan <- err
				}
				unwrapTrace(trace)

				// resultChan <- data.TxHash
			},
		)

		if err != nil {
			errChan <- err
		}
	}()

	select {
		// // успешное завершение
		// case transHash := <-resultChan:
		// 	t.Log("transHash:", transHash)
		// ошибка в горутине
		case err := <-errChan:
			cancel()
			t.Fatal("error:", err)
	}
}
