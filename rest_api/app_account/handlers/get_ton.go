package handlers

import (
	"context"
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"

	myTonapiAccount "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/account"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
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

	// создание API клиента TON для tonapi-go
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", constants.TonapiClientTimeout)
	if err != nil {
		settings.ErrorLog.Println(fmt.Errorf("get account ton balance using tonapi: %w", err))
		return err
	}

	// создание контекста с таймаутом
	getBalanceTONContext, cancel := context.WithTimeout(context.Background(), constants.GetBalanceTONContextTimeout)
	defer cancel()

	// получение баланса TON аккаунта
	dataOut, err = myTonapiAccount.GetBalanceTON(getBalanceTONContext, tonapiClient)
	if err != nil {
		settings.ErrorLog.Println(err)
		return err
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
