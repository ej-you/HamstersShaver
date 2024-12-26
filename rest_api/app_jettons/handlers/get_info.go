package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	myStonfiJettons "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/stonfi/jettons"
	"github.com/ej-you/HamstersShaver/rest_api/app_jettons/serializers"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// эндпоинт получения информации о монете
// @Title Get jetton info
// @Description Get jetton info from Stonfi API by it master address
// @Param MasterAddress query string true "мастер-адрес монеты (jetton_master)" "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
// @Success 200 object myStonfiJettons.JettonParams "JettonParams JSON"
// @Tag jettons
// @Route /jettons/get-info [get]
func GetInfo(ctx echo.Context) error {
	var err error
	var dataIn serializers.GetInfoIn
	var dataOut myStonfiJettons.JettonParams

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.Validate(&dataIn); err != nil {
		return err
	}

	// получение информации о монете
	dataOut, err = myStonfiJettons.GetJettonInfoByAddressWithTimeout(dataIn.MasterAddress, constants.GetJettonInfoByAddressTimeout)
	if err != nil {
		settings.ErrorLog.Println(err)
		return coreErrors.AssertAPIError(err).GetHTTPError()
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
