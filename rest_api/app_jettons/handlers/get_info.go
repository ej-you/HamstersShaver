package handlers

import (
	"time"
	"net/http"

	echo "github.com/labstack/echo/v4"

	myStonfiJettons "github.com/Danil-114195722/HamstersShaver/rest_api/ton_api_rest/stonfi/jettons"
	JettonsErrors "github.com/Danil-114195722/HamstersShaver/rest_api/app_jettons/errors"
	"github.com/Danil-114195722/HamstersShaver/rest_api/app_jettons/serializers"
)


// эндпоинт получения информации о монете
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
