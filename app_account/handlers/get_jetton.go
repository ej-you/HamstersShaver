package handlers

import (
	"time"
	"context"
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"

	myTonapiAccount "github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/account"
	
	AccountErrors "github.com/Danil-114195722/HamstersShaver/app_account/errors"
	"github.com/Danil-114195722/HamstersShaver/app_account/serializers"

	coreErrors "github.com/Danil-114195722/HamstersShaver/core/errors"
	"github.com/Danil-114195722/HamstersShaver/settings"
)


// эндпоинт получения информации о монете аккаунта по её адресу
func GetJetton(ctx echo.Context) error {
	var err error
	var dataIn serializers.GetJettonIn
	var dataOut myTonapiAccount.AccountJetton

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = dataIn.Validate(); err != nil {
		return err
	}

	// создание API клиента TON для tonapi-go с таймаутом в 3 секунд
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
