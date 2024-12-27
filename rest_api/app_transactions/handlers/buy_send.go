package handlers

import (
	"context"
	"net/http"

	echo "github.com/labstack/echo/v4"

	myTongoTransactions "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tongo/transactions"
	"github.com/ej-you/HamstersShaver/rest_api/app_transactions/serializers"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// эндпоинт отправки транзакции на покупку
// @Title Buy send
// @Description Send transaction to buy jettons using TON
// @Param BuySendIn body serializers.BuySendIn true "Cтруктура входных данных для отправки транзакции на покупку"
// @Success 201 object serializers.BuySendOut "Transaction was sent successfully"
// @Tag transactions
// @Route /transactions/buy/send [post]
func BuySend(ctx echo.Context) error {
	var err error
	var dataIn serializers.BuySendIn

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.Validate(&dataIn); err != nil {
		return err
	}

	// создание контекста с таймаутом
	sendBuyJettonContext, cancel := context.WithTimeout(context.Background(), constants.SendBuyJettonContextTimeout)
	defer cancel()

	// отправка транзакции на покупку
	err = myTongoTransactions.BuyJetton(sendBuyJettonContext, dataIn.JettonCA, dataIn.Amount, dataIn.Slippage)
	if err != nil {
		settings.ErrorLog.Println(err)
		return err
	}

	return ctx.JSON(http.StatusCreated, serializers.BuySendOut{Success: true})
}
