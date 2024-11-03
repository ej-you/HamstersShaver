package handlers

import (
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	JettonsErrors "github.com/Danil-114195722/HamstersShaver/app_jettons/errors"
	
	myTongoTransactions "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/transactions"
	"github.com/Danil-114195722/HamstersShaver/app_transactions/serializers"
)


// эндпоинт получения информации о последующей транзакции продажи монет
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
