package ton_api_sse

import (
	"context"
	"fmt"
	"time"

	tonapi "github.com/tonkeeper/tonapi-go"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// ожидание покупки/продажи монеты (полное завершение любой следующей транзакции)
// JettonNotify - 0x7362d09c (успех продажи, неудача покупки)
// Excess -	0xd53276db (успех покупки, неудача продажи)
func SubscribeToNextTransaction(timeout time.Duration) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	streamingAPI := tonapi.NewStreamingAPI(tonapi.WithStreamingToken(settings.TonApiToken))

	errChan := make(chan error)
	defer close(errChan)
	resultChan := make(chan string)
	defer close(resultChan)

	go func() {
		err := streamingAPI.SubscribeToTransactions(ctx,
			[]string{settings.JsonWallet.Hash},
			// получение транзакций только с этими операциями
			[]string{"0x7362d09c", "0xd53276db"},
			func(data tonapi.TransactionEventData) {
				cancel()
				resultChan <- data.TxHash
			},
		)

		if err != nil {
			errChan <- err
		}
	}()

	select {
		// успешное завершение
		case transHash := <-resultChan:
			return transHash, nil
		// ошибка в горутине
		case err := <-errChan:
			cancel()
			return "", coreErrors.New(
				fmt.Errorf("wait next transaction: %w", err),
				"failed to wait transaction via SSE",
				"sse",
				500,
			)
		// если прошло время timeout, а данные не получены
		case <-time.After(timeout):
			cancel()
			return "", coreErrors.NewTimeout(
				fmt.Errorf("wait next transaction: timeout error"),
				"wait transaction via SSE: timeout error",
				"sse",
				500,
			)
	}
}
