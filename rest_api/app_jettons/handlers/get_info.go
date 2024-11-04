package handlers

import (
	"time"
	"net/http"

	echo "github.com/labstack/echo/v4"

	myStonfiJettons "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/stonfi/jettons"
	JettonsErrors "github.com/ej-you/HamstersShaver/rest_api/app_jettons/errors"
	"github.com/ej-you/HamstersShaver/rest_api/app_jettons/serializers"
)


// эндпоинт получения информации о монете
//	@Summary		Get jetton info [NOT WORK IN SWAGGER]
//	@Description	Get jetton info from Stonfi API by it master address
//	@Router			/jettons/get-info [get]
//	@ID				get-info
//	@Tags			jettons
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			JSON		body	serializers.GetInfoIn	true	"GetInfoIn struct params"
//	@Success		200		{object}	myStonfiJettons.JettonParams
func GetInfo(ctx echo.Context) error {
	var err error
	var dataIn serializers.GetInfoIn
	var dataOut myStonfiJettons.JettonParams

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = dataIn.Validate(); err != nil {
		return err
	}

	// формирование структуры для ответа с таймаутом в 5 секунд
	dataOut, err = myStonfiJettons.GetJettonInfoByAddressWithTimeout(dataIn.MasterAddress, 5*time.Second)
	if err != nil {
		if err.Error() == "Jetton was not found" {
			return JettonsErrors.InvalidJettonAddressError
		}
		return echo.NewHTTPError(500, map[string]string{"jettons": err.Error()})
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
