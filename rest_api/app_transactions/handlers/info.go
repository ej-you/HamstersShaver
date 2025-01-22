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


// эндпоинт получения информации о прошедшей транзакции по хэшу её первой операции
// @Title Transaction info
// @Description Get transaction info by given its hash (hash of first operation) and action (buy OR cell)
// @Param TransactionHash query string true "хэш транзакции" "4f8ff3378e1d4cc80488750fda3bcc6b730b71b69429d9c44a775b377bdc66a4"
// @Param Action query string true "действие с монетами в транзакции (покупка/продажа)" "cell"
// @Success 200 object myTonapiTransactions.TransactionInfo "TransactionInfo JSON"
// @Tag transactions
// @Route /transactions/info [get]
func Info(ctx echo.Context) error {
	var err error
	var dataIn serializers.InfoIn
	var dataOut myTonapiTransactions.TransactionInfo

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
	dataOut, err = myTonapiTransactions.GetTransactionInfoWithStatusOK(getTransInfoContext, dataIn.TransactionHash, dataIn.Action)
	if err != nil {
		settings.ErrorLog.Println(err)
		return err
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
