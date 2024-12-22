package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	myTonapiAccount "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/account"
	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// эндпоинт получения информации о TON на аккаунте
// @Title Get TON balance on account
// @Description Get TON balance on account and other info about TON
// @Success 200 object myTonapiAccount.TonJetton "Account TON info"
// @Tag account
// @Route /account/get-ton [get]
func GetTon(ctx echo.Context) error {
	var err error
	var dataOut myTonapiAccount.TonJetton

	// создание API клиента TON для tonapi-go с таймаутом в 3 секунд
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", 3*time.Second)
	if err != nil {
		settings.ErrorLog.Println(fmt.Errorf("get account ton balance using tonapi: %w", err))
		return coreErrors.AssertAPIError(err).GetHTTPError()
	}

	// создание контекста с таймаутом в 5 секунд
	tonApiContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получение баланса TON аккаунта
	dataOut, err = myTonapiAccount.GetBalanceTON(tonApiContext, tonapiClient)
	if err != nil {
		settings.ErrorLog.Println(fmt.Errorf("get account ton balance using tonapi: %w", err))
		return coreErrors.AssertAPIError(err).GetHTTPError()
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
