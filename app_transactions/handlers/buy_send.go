package handlers

import (
	"context"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	JettonsErrors "github.com/Danil-114195722/HamstersShaver/app_jettons/errors"
	
	myTongoTransactions "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/transactions"
	"github.com/Danil-114195722/HamstersShaver/app_transactions/serializers"
)


// эндпоинт отправки транзакции на покупку
func BuySend(ctx echo.Context) error {
	var err error
	var dataIn serializers.BuySendIn

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = dataIn.Validate(); err != nil {
		return err
	}

	// создание контекста с таймаутом в 5 секунд
	tonApiContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// отправка транзакции на покупку с таймаутом в 5 секунд
	err = myTongoTransactions.BuyJetton(tonApiContext, 10*time.Second, dataIn.JettonCA, dataIn.Amount, dataIn.Slippage)
	if err != nil {
		if err.Error() == "Jetton was not found" {
			return JettonsErrors.InvalidJettonAddressError
		}
		return echo.NewHTTPError(500, map[string]string{"transactions": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, "ok")
}
