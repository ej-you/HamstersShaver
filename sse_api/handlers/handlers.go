package handlers

import (
	"context"
	"fmt"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"

	customErrors "github.com/ej-you/HamstersShaver/sse_api/errors"
	"github.com/ej-you/HamstersShaver/sse_api/settings"
)


// рекомендованная задержка перед переподключением в миллисекундах для клиента
const messageRetry = 1000


func SubscribeToAccountTraces(ctx echo.Context) error {
	clientID := uuid.New().String()
	settings.InfoLog.Printf("New client %s connected\n", clientID)

	respWriter := ctx.Response()
	respWriter.Header().Set("Content-Type", "text/event-stream")
	respWriter.Header().Set("Cache-Control", "no-cache")
	respWriter.Header().Set("Connection", "keep-alive")

	// канал для отслеживания отключения юзера
	clientDisconnected := ctx.Request().Context().Done()

	// контекст для горутины с подпиской на trace аккаунта
	subscribeCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// канал для ошибок для горутины с подпиской на trace аккаунта
	errChan := make(chan error)

	// запуск подписки на trace аккаунта
	go func() {
		defer close(errChan)

		// строка SSE сообщения
		var message string

		// trading bot wallet for tests: UQA4mfrV45OEIuTyJKDQe41FX1X0XD8IPJ9UYb7Tpu3gK6kO
		streamingAPI := tonapi.NewStreamingAPI(tonapi.WithStreamingToken(settings.TonApiToken))
		err := streamingAPI.SubscribeToTraces(subscribeCtx,
			[]string{settings.WalletHash},
			func(data tonapi.TraceEventData) {
				settings.InfoLog.Printf("Got txHash %s for client %s\n", data.Hash, clientID)
				// создание сообщения с хэшем полученной транзы
				message = fmt.Sprintf("event: trace\ndata: %s\nid: %s\nretry: %d\n\n", data.Hash, uuid.New().String(), messageRetry)

				// добавление нового сообщения в буфер ответа
				_, err := fmt.Fprint(respWriter, message)
				if err != nil {
					errChan <- customErrors.NewSseErrorf("trying to add message to response buffer: %v", err)
					return
				}

				// отправка данных из буфера клиенту без ожидания окончания запроса
				respWriter.Flush()
			},
		)
		if err != nil {
			errChan <- customErrors.NewSseErrorf("trying to subscribe to account traces: %v", err)
		}
	}()

	select {
		// отключение клиента
		case <-clientDisconnected:
			settings.InfoLog.Printf("Client %s disconnected\n", clientID)
			return nil
		// ошибка в горутине
		case err := <-errChan:
			return err
	}
}
