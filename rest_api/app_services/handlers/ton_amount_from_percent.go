package handlers

import (
	"context"
	"fmt"
	"math"
	"net/http"

	echo "github.com/labstack/echo/v4"

	myTonapiAccount "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/account"
	myTonapiServices "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/services"
	"github.com/ej-you/HamstersShaver/rest_api/app_services/serializers"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// эндпоинт получения кол-ва TON по проценту от их баланса
// @Title Get TON amount from percent of its balance
// @Description Get TON amount from percent of its balance (in string format and not floored)
// @Param Percent query int true "процент от баланса TON" "100"
// @Success 200 object serializers.TonAmountFromPercentOut "TonAmountFromPercent value"
// @Tag services
// @Route /services/ton-amount-from-percent [get]
func TonAmountFromPercent(ctx echo.Context) error {
	var err error
	var dataIn serializers.TonAmountFromPercentIn
	var dataOut serializers.TonAmountFromPercentOut

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
		settings.ErrorLog.Println(fmt.Errorf("get TON amount from percent: %w", err))
		return err
	}
	// создание контекста с таймаутом
	getBalanceTONContext, cancel := context.WithTimeout(context.Background(), constants.GetBalanceTONContextTimeout)
	defer cancel()
	// получение информации о TON у аккаунта
	tonInfo, err := myTonapiAccount.GetBalanceTON(getBalanceTONContext, tonapiClient)
	if err != nil {
		settings.ErrorLog.Println(fmt.Errorf("get TON amount from percent: %w", err))
		return err
	}

	tonBalanceFloat64 := float64(tonInfo.Balance) / math.Pow10(tonInfo.Decimals)
	// проверка на наличие хотя бы constants.GasAmountFloat64 TON для газа
	if tonBalanceFloat64 <= constants.GasAmountFloat64 {
		errText := fmt.Sprintf("TON balance must be greater than %v (gas amount)", constants.GasAmountFloat64)
		return coreErrors.New(
			fmt.Errorf("get TON amount from percent: %s", errText),
			errText,
			"rest_api",
			400,
		)
	}

	// получение части от баланса TON в соответствии с процентом
	tonPercentAmount := tonBalanceFloat64 / 100 * float64(dataIn.Percent)

	// если кол-во монет по проценту больше чем TonBalance с вычетом газа, то уменьшаем это кол-во до предельного значения
	if tonPercentAmount > (tonBalanceFloat64 - constants.GasAmountFloat64) {
		tonPercentAmount = tonBalanceFloat64 - constants.GasAmountFloat64
	}

	// формирование структуры с кол-вом монет (перевод в строку без округления)
	dataOut = serializers.TonAmountFromPercentOut{
		TonAmount: myTonapiServices.StringJettonAmountFromFloat64(tonPercentAmount, tonInfo.Decimals),
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
