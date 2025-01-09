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

	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// эндпоинт получения кол-ва монет по проценту от их баланса
// @Title Get jettons amount from percent of its balance
// @Description Get jettons amount from percent of its balance (in string format and not floored)
// @Param Percent query int true "процент от баланса монеты" "50"
// @Param MasterAddress query string true "мастер-адрес монеты (jetton_master)" "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
// @Success 200 object serializers.JettonAmountFromPercentOut "JettonAmountFromPercent value"
// @Tag services
// @Route /services/jetton-amount-from-percent [get]
func JettonAmountFromPercent(ctx echo.Context) error {
	var err error
	var dataIn serializers.JettonAmountFromPercentIn
	var dataOut serializers.JettonAmountFromPercentOut

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
		settings.ErrorLog.Println(fmt.Errorf("get jetton amount from percent: %w", err))
		return err
	}
	// создание контекста с таймаутом
	getAccountJettonContext, cancel := context.WithTimeout(context.Background(), constants.GetAccountJettonContextTimeout)
	defer cancel()
	// получение информации о монете аккаунта
	jettonInfo, err := myTonapiAccount.GetAccountJetton(getAccountJettonContext, tonapiClient, dataIn.MasterAddress)
	if err != nil {
		settings.ErrorLog.Println(fmt.Errorf("get jetton amount from percent: %w", err))
		return err
	}

	// получение части от баланса монеты в соответствии с процентом
	jettonBalanceFloat64 := float64(jettonInfo.Balance) / math.Pow10(jettonInfo.Decimals)
	jettonPercentAmount := jettonBalanceFloat64 / 100 * float64(dataIn.Percent)

	// формирование структуры с кол-вом монет (перевод в строку без округления)
	dataOut = serializers.JettonAmountFromPercentOut{
		JettonAmount: myTonapiServices.StringJettonAmountFromFloat64(jettonPercentAmount, jettonInfo.Decimals),
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
