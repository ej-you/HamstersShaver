package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	myTongoTransactions "github.com/ej-you/HamstersShaver/rest_api/ton_api/tongo/transactions"
	"github.com/ej-you/HamstersShaver/rest_api/app_transactions/serializers"
	
	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// эндпоинт получения информации о последующей транзакции продажи монет
// @Title Cell pre-request
// @Description Get pre-request info about cell transaction
// @Param JettonCA query string true "мастер-адрес продаваемой монеты (jetton_master)" "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
// @Param Amount query float64 true "кол-во используемых монет на продажу в формате, удобном для человека" "200"
// @Param Slippage query int true "процент проскальзывания" "20"
// @Success 200 object myTongoTransactions.PreRequestCellJetton "PreRequestCellJetton JSON"
// @Tag transactions
// @Route /transactions/cell/pre-request [get]
func CellPreRequest(ctx echo.Context) error {
	var err error
	var dataIn serializers.CellPreRequestIn
	var dataOut myTongoTransactions.PreRequestCellJetton

	// парсинг query-параметров
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.Validate(&dataIn); err != nil {
		return err
	}

	// получение примерных данных о будующей транзакции
	dataOut, err = myTongoTransactions.GetPreRequestCellJetton(dataIn.JettonCA, dataIn.Amount, dataIn.Slippage)
	if err != nil {
		settings.ErrorLog.Println(err)
		return err
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
