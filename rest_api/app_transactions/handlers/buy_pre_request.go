package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	myTongoTransactions "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tongo/transactions"
	"github.com/ej-you/HamstersShaver/rest_api/app_transactions/serializers"
	
	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// эндпоинт получения информации о последующей транзакции покупки монет
// @Title Buy pre-request
// @Description Get pre-request info about buy transaction
// @Param JettonCA query string true "мастер-адрес покупаемой монеты (jetton_master)" "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
// @Param Amount query float64 true "кол-во используемых TON для покупки в формате, удобном для человека" "0.1"
// @Param Slippage query int true "процент проскальзывания" "20"
// @Success 200 object myTongoTransactions.PreRequestBuyJetton "PreRequestBuyJetton JSON"
// @Tag transactions
// @Route /transactions/buy/pre-request [get]
func BuyPreRequest(ctx echo.Context) error {
	var err error
	var dataIn serializers.BuyPreRequestIn
	var dataOut myTongoTransactions.PreRequestBuyJetton

	// парсинг query-параметров
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.Validate(&dataIn); err != nil {
		return err
	}

	// получение примерных данных о будующей транзакции
	dataOut, err = myTongoTransactions.GetPreRequestBuyJetton(dataIn.JettonCA, dataIn.Amount, dataIn.Slippage)
	if err != nil {
		settings.ErrorLog.Println(err)
		return coreErrors.AssertAPIError(err).GetHTTPError()
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
