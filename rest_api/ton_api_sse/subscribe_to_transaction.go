package ton_api_sse

import (
	"context"
	"errors"
	"time"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// ожидание покупки/продажи монеты (полное завершение транзакции)
// JettonNotify - 0x7362d09c (успех продажи, неудача покупки)
// Excess -	0xd53276db (успех покупки, неудача продажи)
func SubscribeToTransaction(timeout time.Duration) (string, error) {
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
			return "", errors.New("Failed to wait transaction via SSE: " + err.Error())
		// если прошло время timeout, а данные не получены
		case <-time.After(timeout):
			cancel()
			return "", errors.New("Failed to wait transaction via SSE: timeout error")
	}
}
