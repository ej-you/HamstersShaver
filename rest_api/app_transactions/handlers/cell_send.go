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


// эндпоинт отправки транзакции на продажу
// @Title Cell send
// @Description Send transaction to cell jettons to TON
// @Param CellSendIn body serializers.CellSendIn true "Cтруктура входных данных для отправки транзакции на продажу"
// @Success 201 object serializers.CellSendOut "Transaction was sent successfully"
// @Tag transactions
// @Route /transactions/cell/send [post]
func CellSend(ctx echo.Context) error {
	var err error
	var dataIn serializers.CellSendIn

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.Validate(&dataIn); err != nil {
		return err
	}

	// создание контекста с таймаутом
	sendCellJettonContext, cancel := context.WithTimeout(context.Background(), constants.SendCellJettonContextTimeout)
	defer cancel()

	// отправка транзакции на продажу
	err = myTongoTransactions.CellJetton(sendCellJettonContext, dataIn.JettonCA, dataIn.Amount, dataIn.Slippage)
	if err != nil {
		settings.ErrorLog.Println(err)
		return err
	}

	return ctx.JSON(http.StatusCreated, serializers.CellSendOut{Success: true})
}
