package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	myTonapiServices "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/services"
	"github.com/ej-you/HamstersShaver/rest_api/app_services/serializers"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
)


// эндпоинт получения информации о монете аккаунта по её адресу
// @Title Convert raw balance into beauty balance
// @Description Convert raw balance (in int64 forman) into beauty balance (rounded float in string format)
// @Param RawBalance query int64 true "баланс монеты в int64 формате" "326166742480"
// @Param Decimals query int true "decimals монеты" "9"
// @Success 200 object serializers.BeautyBalanceOut "BeautyBalance value"
// @Tag services
// @Route /services/beauty-balance [get]
func BeautyBalance(ctx echo.Context) error {
	var err error
	var dataIn serializers.BeautyBalanceIn
	var dataOut serializers.BeautyBalanceOut

	// парсинг query-параметров
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.Validate(&dataIn); err != nil {
		return err
	}

	// формирование структуры для ответа
	dataOut = serializers.BeautyBalanceOut{
		BeautyBalance: myTonapiServices.JettonBalanceFormat(dataIn.RawBalance, dataIn.Decimals),
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
