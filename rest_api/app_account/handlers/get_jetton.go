package handlers

import (
	"context"
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"

	myTonapiAccount "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/account"
	
	"github.com/ej-you/HamstersShaver/rest_api/app_account/serializers"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// эндпоинт получения информации о монете аккаунта по её адресу
// @Title Get jetton balance on account
// @Description Get jetton balance on account and other info about jetton by it master address
// @Param MasterAddress query string true "мастер-адрес монеты (jetton_master)" "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
// @Success 200 object myTonapiAccount.AccountJetton "AccountJetton JSON"
// @Tag account
// @Route /account/get-jetton [get]
func GetJetton(ctx echo.Context) error {
	var err error
	var dataIn serializers.GetJettonIn
	var dataOut myTonapiAccount.AccountJetton

	// парсинг query-параметров
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.Validate(&dataIn); err != nil {
		return err
	}

	// создание API клиента TON для tonapi-go
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", constants.TonapiClientTimeout)
	if err != nil {
		settings.ErrorLog.Println(fmt.Errorf("get account jetton using tonapi: %w", err))
		return coreErrors.AssertAPIError(err).GetHTTPError()
	}

	// создание контекста с таймаутом
	getAccountJettonContext, cancel := context.WithTimeout(context.Background(), constants.GetAccountJettonContextTimeout)
	defer cancel()

	// получение информации о монете аккаунта
	dataOut, err = myTonapiAccount.GetAccountJetton(getAccountJettonContext, tonapiClient, dataIn.MasterAddress)
	if err != nil {
		settings.ErrorLog.Println(err)
		return coreErrors.AssertAPIError(err).GetHTTPError()
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
