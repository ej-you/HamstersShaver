package handlers

import (
	"context"
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"

	myTonapiAccount "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/account"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// эндпоинт получения информации о всех монетах на аккаунте
// @Title Get jettons balances on account
// @Description Get all non-null jettons balances on account and other info about jettons
// @Success 200 array []myTonapiAccount.AccountJetton "AccountJettons list JSON"
// @Tag account
// @Route /account/get-jettons [get]
func GetJettons(ctx echo.Context) error {
	var err error
	var dataOut []myTonapiAccount.AccountJetton

	// создание API клиента TON для tonapi-go
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", constants.TonapiClientTimeout)
	if err != nil {
		settings.ErrorLog.Println(fmt.Errorf("get account jettons using tonapi: %w", err))
		return coreErrors.AssertAPIError(err).GetHTTPError()
	}

	// создание контекста с таймаутом
	getBalanceJettonsContext, cancel := context.WithTimeout(context.Background(), constants.GetBalanceJettonsContextTimeout)
	defer cancel()

	// получение монет аккаунта
	dataOut, err = myTonapiAccount.GetBalanceJettons(getBalanceJettonsContext, tonapiClient)
	if err != nil {
		settings.ErrorLog.Println(err)
		return coreErrors.AssertAPIError(err).GetHTTPError()
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
