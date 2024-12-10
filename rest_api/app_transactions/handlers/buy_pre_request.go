package handlers

import (
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	JettonsErrors "github.com/ej-you/HamstersShaver/rest_api/app_jettons/errors"
	myTongoTransactions "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tongo/transactions"
	"github.com/ej-you/HamstersShaver/rest_api/app_transactions/serializers"
	
	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
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
