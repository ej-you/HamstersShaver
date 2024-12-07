package ton_api_sse

import (
	"context"
	"errors"
	"fmt"
	"time"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// ожидание покупки монеты (полное завершение транзакции)
// Excess -	0xd53276db
func SubscribeToBuyJettonsTransaction(timeout time.Duration) error {
	return subscribeToTransaction([]string{"0xd53276db"}, timeout)
}

// ожидание продажи монеты (полное завершение транзакции)
// JettonNotify - 0x7362d09c
func SubscribeToCellJettonsTransaction(timeout time.Duration) error {
	return subscribeToTransaction([]string{"0x7362d09c"}, timeout)
}


// ожидание транзакции на аккаунте с переданными операциями
func subscribeToTransaction(operations []string, timeout time.Duration) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	streamingAPI := tonapi.NewStreamingAPI(tonapi.WithStreamingToken(settings.TonApiToken))

	errChan := make(chan error)
	defer close(errChan)

	go func() {
		err := streamingAPI.SubscribeToTransactions(ctx,
			[]string{settings.JsonWallet.Hash},
			// получение транзакций только с этими операциями
			operations,
			func(data tonapi.TransactionEventData) {
				fmt.Printf("New tx with hash: %v\n", data.TxHash)
				fmt.Printf("New lt: %v\n", data.Lt)
				cancel()
			},
		)

		if err != nil {
			errChan <- err
		}
	}()

	select {
		// успешное завершение
		case <-ctx.Done():
			return nil
		// ошибка в горутине
		case err := <-errChan:
			cancel()
			return errors.New("Failed to get transaction info via SSE: " + err.Error())
		// если прошло время timeout, а данные не получены
		case <-time.After(timeout):
			cancel()
			return errors.New("Failed to get transaction info via SSE: timeout error")
	}
}
