package handlers

import (
	echo "github.com/labstack/echo/v4"

	myStonfiJettons "github.com/ej-you/HamstersShaver/rest_api/ton_api/stonfi/jettons"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// структура входных данных для получения информации о монете по её адресу
type GetInfoIn struct {
	// мастер-адрес монеты (jetton_master)
	MasterAddress string `query:"masterAddress" json:"masterAddress" validate:"required"`
}


// эндпоинт получения информации о монете
// @Title Get jetton info
// @Description Get jetton info from Stonfi API by it master address
// @Param MasterAddress query string true "мастер-адрес монеты (jetton_master)" "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
// @Success 200 object myStonfiJettons.JettonParams "JettonParams JSON"
// @Tag jettons
// @Route /jettons/get-info [get]
func GetInfo(ctx echo.Context) error {
	var err error
	var dataIn GetInfoIn
	var dataOut myStonfiJettons.JettonParams

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Struct(&dataIn); err != nil {
		return err
	}

	// получение информации о монете
	dataOut, err = myStonfiJettons.GetJettonInfoByAddressWithTimeout(dataIn.MasterAddress, constants.GetJettonInfoByAddressTimeout)
	if err != nil {
		settings.ErrorLog.Println(err)
		return err
	}

	return ctx.JSON(200, dataOut)
}
