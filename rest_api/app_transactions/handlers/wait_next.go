package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	mySSE "github.com/ej-you/HamstersShaver/rest_api/ton_api_sse"
	"github.com/ej-you/HamstersShaver/rest_api/app_transactions/serializers"
	
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// эндпоинт ожидания окончания транзакции
// @Title Wait the end of transaction
// @Description Wait the end of any next transaction
// @Success 200 object serializers.WaitNextOut "Transaction hash"
// @Tag transactions
// @Route /transactions/wait-next [get]
func WaitNext(ctx echo.Context) error {

	// ожидание окончания транзакции
	transHash, err := mySSE.SubscribeToNextTransaction(constants.WaitNextTransactionTimeout)
	if err != nil {
		settings.ErrorLog.Println(err)
		return err
	}

	return ctx.JSON(http.StatusOK, serializers.WaitNextOut{Hash: transHash})
}
