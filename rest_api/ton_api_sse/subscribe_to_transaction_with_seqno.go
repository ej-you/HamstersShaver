package ton_api_sse

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	tonapi "github.com/tonkeeper/tonapi-go"

	myTonapiTransactions "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/transactions"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// кол-во попыток получения информации о транзакции при получении timeout ошибки
const getTransInfoAttemps = 3 


// ожидание покупки/продажи монеты (полное завершение транзакции с определённым seqno)
// JettonNotify - 0x7362d09c (успех продажи, неудача покупки)
// Excess -	0xd53276db (успех покупки, неудача продажи)
func SubscribeToTransactionWithSeqno(timeout time.Duration, seqno int, action string) (myTonapiTransactions.TransactionInfoWithStatusOK, error) {
	var transInfo myTonapiTransactions.TransactionInfoWithStatusOK
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	streamingAPI := tonapi.NewStreamingAPI(tonapi.WithStreamingToken(settings.TonApiToken))

	errChan := make(chan error)
	defer close(errChan)
	resultChan := make(chan myTonapiTransactions.TransactionInfoWithStatusOK)
	defer close(resultChan)
	var wg sync.WaitGroup

	fmt.Println("Prepared to subscribe!")

	go func() {
		err := streamingAPI.SubscribeToTransactions(ctx,
			[]string{settings.JsonWallet.Hash},
			// получение транзакций только с этими операциями
			[]string{"0x7362d09c", "0xd53276db"},
			func(data tonapi.TransactionEventData) {
				fmt.Println("Caught new transaction!")
				wg.Add(1)
				go nextTransactionHandler(ctx, data.TxHash, seqno, action, &wg, resultChan, errChan)
			},
		)
		if err != nil {
			errChan <- fmt.Errorf("wait transaction with seqno: %v: %w", err, SseError("failed to subscribe to transactions using SSE"))
		}
		fmt.Println("Unsubscribe...")
	}()

	fmt.Println("Wait gorutines...")

	var err error
	select {
		// успешное завершение
		case transInfo = <-resultChan:
			fmt.Println("Successfully wait!")
			cancel()
			fmt.Println("Wait all gorutines to exit...")
			wg.Wait()
			fmt.Println("Exit from func...")
			return transInfo, nil
		// ошибка в горутине
		case gorutineErr := <-errChan:
			err = fmt.Errorf("wait transaction with seqno: %v: %w", gorutineErr, SseError("failed to wait transaction via SSE"))
		// если прошло время timeout, а данные не получены
		case <-time.After(timeout):
			err = fmt.Errorf("wait transaction with seqno: %w", SseError("failed to wait transaction via SSE: timeout error"))
	}
	fmt.Println("Got error in waiting!")
	cancel()
	wg.Wait()
	return transInfo, err
}


// обработчик каждой транзакции, принятой от SSE
func nextTransactionHandler(ctx context.Context, hash string, seqno int, action string, wg *sync.WaitGroup, resultChan chan<- myTonapiTransactions.TransactionInfoWithStatusOK, errChan chan<- error) {
	defer (*wg).Done()

	var err error
	var transInfo myTonapiTransactions.TransactionInfoWithStatusOK

	// создание контекста с таймаутом
	getTransInfoContext, cancel := context.WithTimeout(ctx, constants.GetTransInfoContextTimeout)
	defer cancel()

	fmt.Println("Get info about transaction", hash)

	canceledErr := context.Canceled
	// делаем getTransInfoAttemps попыток получения информации о транзакции
	for i := 0; i < getTransInfoAttemps; i++ {
		fmt.Println("Try to get info about transaction", hash)
		transInfo, err = myTonapiTransactions.GetTransactionInfoWithStatusOKByHash(getTransInfoContext, hash, action)

		if err == nil { // NOT err
			break
		}

		// если контекст был отменён через родительский контекст
		if errors.Is(err, canceledErr) {
			fmt.Printf("Already found info about transaction (not %s)\n", hash)

			return
		// если произошла неизвестная ошибка (не timeout)
		} else if !coreErrors.AssertAPIError(err).IsTimeout() {
			fmt.Println("Error while getting info about transaction", hash)
			errChan <- err
			return
		}
	}
	// если попытки не помогли, и осталась timeout ошибка
	if err != nil {
		fmt.Println("Timeout error while getting info about transaction", hash)
		errChan <- err
		return
	}

	fmt.Println("Got info about transaction", hash)
	fmt.Println("transInfo:", transInfo)
	fmt.Println("SeqnoBeforeTransaction:", seqno)

	// if transInfo.SeqnoBeforeTransaction == seqno {
	// 	fmt.Println("Successfully got info about transaction", hash)
	// 	resultChan <- transInfo
	// }
}
