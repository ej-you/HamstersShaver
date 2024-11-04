package handlers

import (
	"context"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	JettonsErrors "github.com/ej-you/HamstersShaver/rest_api/app_jettons/errors"
	
	myTongoTransactions "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tongo/transactions"
	"github.com/ej-you/HamstersShaver/rest_api/app_transactions/serializers"
)


// эндпоинт отправки транзакции на покупку
//	@Summary		Buy send
//	@Description	Send transaction to buy jettons using TON
//	@Router			/transactions/buy/send [post]
//	@ID				buy-send
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			JSON		body	serializers.BuySendIn	true	"BuySendIn struct params"
//	@Success		201		{object}	serializers.BuySendOut
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

	return ctx.JSON(http.StatusCreated, map[string]bool{"success": true})
}
