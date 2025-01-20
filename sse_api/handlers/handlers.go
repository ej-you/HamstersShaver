package handlers

import (
	"context"
	"fmt"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/sse_api/settings"
)


func SubscribeToAccountTraces(ctx echo.Context) error {
	clientID := uuid.New().String()
	settings.InfoLog.Printf("New client %s connected\n", clientID)

	respWriter := ctx.Response()
	respWriter.Header().Set("Content-Type", "text/event-stream")
	respWriter.Header().Set("Cache-Control", "no-cache")
	respWriter.Header().Set("Connection", "keep-alive")

	// канал для отслеживания отключения юзера
	clientDisconnected := ctx.Request().Context().Done()

	// контекст и канал для ошибок для горутины с подпиской на trace аккаунта
	subscribeCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errChan := make(chan error)
	defer close(errChan)

	streamingAPI := tonapi.NewStreamingAPI(tonapi.WithStreamingToken(settings.TonApiToken))
	// запуск подписки на trace аккаунта
	go func() {
		// trading bot wallet for tests: UQA4mfrV45OEIuTyJKDQe41FX1X0XD8IPJ9UYb7Tpu3gK6kO
		err := streamingAPI.SubscribeToTraces(subscribeCtx,
			[]string{settings.WalletHash, "UQA4mfrV45OEIuTyJKDQe41FX1X0XD8IPJ9UYb7Tpu3gK6kO"},
			func(data tonapi.TraceEventData) {
				settings.InfoLog.Printf("Got txHash %s for client %s\n", data.Hash, clientID)

				// добавление хэша полученной транзы в буфер ответа
				_, err := fmt.Fprintf(respWriter, "data: %s\n\n", data.Hash)
				if err != nil {
					errChan <- err
				}
				// errChan <- fmt.Errorf("test error")
				// отправка данных из буфера клиенту без ожидания окончания запроса
				respWriter.Flush()
			},
		)
		if err != nil {
			errChan <- err
		}
	}()

	select {
		// отключение клиента
		case <-clientDisconnected:
			settings.InfoLog.Printf("Client %s disconnected\n", clientID)
			cancel()
			return nil
		// ошибка в горутине
		case err := <-errChan:
			cancel()
			return err
	}
}
