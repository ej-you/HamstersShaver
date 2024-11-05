package handlers

import (
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	JettonsErrors "github.com/ej-you/HamstersShaver/rest_api/app_jettons/errors"
	
	myTongoTransactions "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tongo/transactions"
	"github.com/ej-you/HamstersShaver/rest_api/app_transactions/serializers"
)


// эндпоинт получения информации о последующей транзакции продажи монет
//	@Summary		Cell pre-request
//	@Description	Get pre-request info about cell transaction
//	@Router			/transactions/cell/pre-request [get]
//	@ID				cell-pre-request
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			Query		query	serializers.CellPreRequestIn	true	"CellPreRequestIn struct params"
//	@Success		200		{object}	myTongoTransactions.PreRequestCellJetton
func CellPreRequest(ctx echo.Context) error {
	var err error
	var dataIn serializers.CellPreRequestIn
	var dataOut myTongoTransactions.PreRequestCellJetton

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = dataIn.Validate(); err != nil {
		return err
	}

	// формирование структуры для ответа с таймаутом в 3 секунд
	dataOut, err = myTongoTransactions.GetPreRequestCellJetton(dataIn.JettonCA, dataIn.Amount, dataIn.Slippage, 3*time.Second)
	if err != nil {
		if err.Error() == "Jetton was not found" {
			return JettonsErrors.InvalidJettonAddressError
		}
		return echo.NewHTTPError(500, map[string]string{"transactions": err.Error()})
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
