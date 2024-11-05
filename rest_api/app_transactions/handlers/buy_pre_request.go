package handlers

import (
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	JettonsErrors "github.com/ej-you/HamstersShaver/rest_api/app_jettons/errors"
	
	myTongoTransactions "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tongo/transactions"
	"github.com/ej-you/HamstersShaver/rest_api/app_transactions/serializers"
)


// эндпоинт получения информации о последующей транзакции покупки монет
//	@Summary		Buy pre-request
//	@Description	Get pre-request info about buy transaction
//	@Router			/transactions/buy/pre-request [get]
//	@ID				buy-pre-request
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			Query		query	serializers.BuyPreRequestIn	true	"BuyPreRequestIn struct params"
//	@Success		200		{object}	myTongoTransactions.PreRequestBuyJetton
func BuyPreRequest(ctx echo.Context) error {
	var err error
	var dataIn serializers.BuyPreRequestIn
	var dataOut myTongoTransactions.PreRequestBuyJetton

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = dataIn.Validate(); err != nil {
		return err
	}

	// формирование структуры для ответа с таймаутом в 3 секунд
	dataOut, err = myTongoTransactions.GetPreRequestBuyJetton(dataIn.JettonCA, dataIn.Amount, dataIn.Slippage, 3*time.Second)
	if err != nil {
		if err.Error() == "Jetton was not found" {
			return JettonsErrors.InvalidJettonAddressError
		}
		return echo.NewHTTPError(500, map[string]string{"transactions": err.Error()})
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
