package handlers

import (
	"context"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	myTonapiTransactions "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/transactions"
	"github.com/ej-you/HamstersShaver/rest_api/app_transactions/serializers"
	
	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
)


// эндпоинт получения информации о прошедшей транзакции по её хэшу
//	@Summary		Transaction info
//	@Description	Get transaction info by given its hash and action (buy OR cell)
//	@Router			/transactions/info [get]
//	@ID				transactions-info
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			Query		query	serializers.InfoIn	true	"InfoIn struct params"
//	@Success		200		{object}	myTonapiTransactions.TransactionInfoWithStatusOK
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

	// создание контекста с таймаутом в 5 секунд
	tonApiContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// формирование структуры для ответа с таймаутом в 3 секунды
	dataOut, err = myTonapiTransactions.GetTransactionInfoWithStatusOKByHash(tonApiContext, dataIn.TransactionHash, dataIn.Action, 3*time.Second)
	if err != nil {
		return echo.NewHTTPError(500, map[string]string{"transactions": err.Error()})
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
