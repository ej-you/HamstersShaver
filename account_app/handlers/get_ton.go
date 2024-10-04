package handlers

import (
	"time"
	"context"
	"net/http"

	echo "github.com/labstack/echo/v4"

	myTonapiAccount "github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/account"
	"github.com/Danil-114195722/HamstersShaver/settings"
)


// эндпоинт получения информации о TON на аккаунте
func GetTon(ctx echo.Context) error {
	var err error
	var dataOut myTonapiAccount.TonJetton

	// создание API клиента TON для tonapi-go
	tonapiClient, err := settings.GetTonClientTonapi("mainnet")
	if err != nil {
		return err
	}

	// создание контекста с таймаутом в 5 секунд
	tonApiContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// формирование структуры для ответа
	dataOut, err = myTonapiAccount.GetBalanceTON(tonApiContext, tonapiClient)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
