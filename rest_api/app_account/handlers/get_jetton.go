package handlers

import (
	"time"
	"context"
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"

	myTonapiAccount "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/account"
	
	AccountErrors "github.com/ej-you/HamstersShaver/rest_api/app_account/errors"
	"github.com/ej-you/HamstersShaver/rest_api/app_account/serializers"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
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

	// создание API клиента TON для tonapi-go с таймаутом в 3 секунды
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", 3*time.Second)
	if err != nil {
		return coreErrors.GetTonapiClientError
	}
	// создание контекста с таймаутом в 5 секунд
	tonApiContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// формирование структуры для ответа
	dataOut, err = myTonapiAccount.GetAccountJetton(tonApiContext, tonapiClient, dataIn.MasterAddress)
	if err != nil {
		// если такой монеты нет у данного аккаунта
		if strings.HasPrefix(err.Error(), "Failed to get account jetton info: decode response: error: code 404: {Error:account") {
			return AccountErrors.AccountHasNotJettonError
		// если был дан неверный адрес
		} else if strings.HasPrefix(err.Error(), "Failed to get account jetton info: decode response: error: code 4") {
			return AccountErrors.InvalidJettonAddressError
		}
		return echo.NewHTTPError(500, map[string]string{"account": err.Error()})
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
