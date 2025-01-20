package handlers

import (
	"context"
	"net/http"

	echo "github.com/labstack/echo/v4"

	myTonapiTransactions "github.com/ej-you/HamstersShaver/rest_api/ton_api/tonapi/transactions"
	"github.com/ej-you/HamstersShaver/rest_api/app_transactions/serializers"
	
	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// эндпоинт получения информации о прошедшей транзакции по её хэшу
// @Title Transaction info
// @Description Get transaction info by given its hash and action (buy OR cell)
// @Param TransactionHash query string true "хэш транзакции" "29a301e4d2a05713f4eab6c8f0daa3c58eed15d1d41678068cd50fe46ca7f6a5"
// @Param Action query string true "действие с монетами в транзакции (покупка/продажа)" "cell"
// @Success 200 object myTonapiTransactions.TransactionInfoWithStatusOK "TransactionInfoWithStatusOK JSON"
// @Tag transactions
// @Route /transactions/info [get]
func Info(ctx echo.Context) error {
	var err error
	var dataIn serializers.InfoIn
	var dataOut myTonapiTransactions.TransactionInfoWithStatusOK

	// парсинг query-параметров
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.Validate(&dataIn); err != nil {
		return err
	}

	// создание контекста с таймаутом
	getTransInfoContext, cancel := context.WithTimeout(context.Background(), constants.GetTransInfoContextTimeout)
	defer cancel()

	// получение информации о транзакции
	dataOut, err = myTonapiTransactions.GetTransactionInfoWithStatusOKByHash(getTransInfoContext, dataIn.TransactionHash, dataIn.Action)
	if err != nil {
		settings.ErrorLog.Println(err)
		return err
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
